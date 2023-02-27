// Copyright 2016-2021, Pulumi Corporation.

package schema

import (
	"reflect"
	"strings"

	"github.com/mattbaird/jsonpatch"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// SdkToCfn converts Pulumi-SDK-shaped state to CloudFormation-shaped payload. In particular, SDK properties
// are lowerCamelCase, while CloudFormation is usually (but not always) PascalCase.
func SdkToCfn(res *CloudAPIResource, types map[string]CloudAPIType, properties map[string]interface{}) map[string]interface{} {
	converter := sdkToCfnConverter{res, types}
	return converter.sdkToCfn(properties)
}

// DiffToPatch converts a Pulumi object diff to a CloudFormation-shaped patch operation slice. Update/add/delete operations are
// mapped to corresponding patch terms, and SDK properties are translated to respective CFN names.
func DiffToPatch(res *CloudAPIResource, types map[string]CloudAPIType, diff *resource.ObjectDiff) ([]jsonpatch.JsonPatchOperation, error) {
	if diff == nil {
		return []jsonpatch.JsonPatchOperation{}, nil
	}
	converter := sdkToCfnConverter{res, types}
	return converter.diffToPatch(diff), nil
}

type sdkToCfnConverter struct {
	spec  *CloudAPIResource
	types map[string]CloudAPIType
}

func (c *sdkToCfnConverter) sdkToCfn(properties map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, prop := range c.spec.Inputs {
		if v, ok := properties[k]; ok {
			result[ToCfnName(k)] = c.sdkTypedValueToCfn(&prop.TypeSpec, v)
		}
	}
	for k, attr := range c.spec.Outputs {
		if v, ok := properties[k]; ok {
			result[ToCfnName(k)] = c.sdkTypedValueToCfn(&attr.TypeSpec, v)
		}
	}
	return result
}

func (c *sdkToCfnConverter) sdkTypedValueToCfn(spec *pschema.TypeSpec, v interface{}) interface{} {
	if spec.Ref != "" {
		if spec.Ref == "pulumi.json#/Any" {
			return v
		}

		typName := strings.TrimPrefix(spec.Ref, "#/types/")
		return c.sdkObjectValueToCfn(typName, v)
	}

	if spec.OneOf != nil {
		for _, item := range spec.OneOf {
			converted := c.sdkTypedValueToCfn(&item, v)
			// If the conversion resulted in a different value, return it.
			// We should probably fail if we can't convert rather than considering change a success.
			// This is likely to happen if one of the types is any.
			if !reflect.DeepEqual(converted, v) {
				return converted
			}
		}
	}

	switch spec.Type {
	case "array":
		array := v.([]interface{})
		vs := make([]interface{}, len(array))
		for i, item := range array {
			vs[i] = c.sdkTypedValueToCfn(spec.Items, item)
		}
		return vs
	case "object":
		sourceMap := v.(map[string]interface{})
		vs := map[string]interface{}{}
		for n, item := range sourceMap {
			vs[n] = c.sdkTypedValueToCfn(spec.AdditionalProperties, item)
		}
		return vs
	default:
		return v
	}
}

func (c *sdkToCfnConverter) sdkObjectValueToCfn(typeName string, value interface{}) interface{} {
	properties, ok := value.(map[string]interface{})
	if !ok {
		return value
	}

	spec := c.types[typeName]
	result := map[string]interface{}{}
	for k, prop := range spec.Properties {
		if v, ok := properties[k]; ok {
			result[ToCfnName(k)] = c.sdkTypedValueToCfn(&prop.TypeSpec, v)
		}
	}
	return result
}

func (c *sdkToCfnConverter) diffToPatch(diff *resource.ObjectDiff) []jsonpatch.JsonPatchOperation {
	var ops []jsonpatch.JsonPatchOperation
	for sdkName, prop := range c.spec.Inputs {
		cfnName := ToCfnName(sdkName)
		key := resource.PropertyKey(sdkName)
		if v, ok := diff.Updates[key]; ok {
			// Check if properties are write-only, and use an "add" if so. This is because the old values of write only
			// properties don't show up when applying the diff within CC, so we need to do an "add".
			if c.isPropWriteOnly(sdkName) {
				op := c.valueToPatch("add", cfnName, prop, v.New)
				ops = append(ops, op)
			} else {
				op := c.valueToPatch("replace", cfnName, prop, v.New)
				ops = append(ops, op)
			}
		}
		if v, ok := diff.Adds[key]; ok {
			op := c.valueToPatch("add", cfnName, prop, v)
			ops = append(ops, op)
		}
		if _, ok := diff.Deletes[key]; ok {
			op := jsonpatch.NewPatch("remove", "/"+cfnName, nil)
			ops = append(ops, op)
		}
	}
	return ops
}

func (c *sdkToCfnConverter) valueToPatch(opName, propName string, prop pschema.PropertySpec, value resource.PropertyValue) jsonpatch.JsonPatchOperation {
	op := jsonpatch.NewPatch(opName, "/"+propName, nil)
	switch {
	case value.IsNumber() && prop.Type == "integer":
		i := int32(value.NumberValue())
		op.Value = i
	case value.IsNumber():
		op.Value = value.NumberValue()
	case value.IsBool():
		op.Value = value.BoolValue()
	case value.IsString():
		op.Value = value.StringValue()
	default:
		sdkObj := value.MapRepl(nil, nil)
		cfnObj := c.sdkTypedValueToCfn(&prop.TypeSpec, sdkObj)
		op.Value = cfnObj
	}
	return op
}

func (c *sdkToCfnConverter) isPropWriteOnly(cfnName string) bool {
	for _, createOnlyProp := range c.spec.WriteOnly {
		if createOnlyProp == cfnName {
			return true
		}
	}
	return false
}

// CfnToSdk converts CloudFormation-shaped payload to Pulumi-SDK-shaped state. In particular, SDK properties
// are lowerCamelCase, while CloudFormation is usually (but not always) PascalCase.
func CfnToSdk(properties map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for n, v := range properties {
		result[ToSdkName(n)] = cfnValueToSdk(v)
	}
	return result
}

func cfnValueToSdk(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}
		return CfnToSdk(valueMap)
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(value)
		result := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			result[i] = cfnValueToSdk(s.Index(i).Interface())
		}
		return result
	default:
		return value
	}
}
