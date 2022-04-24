// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iottwinmaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::IoTTwinMaker::ComponentType
func LookupComponentType(ctx *pulumi.Context, args *LookupComponentTypeArgs, opts ...pulumi.InvokeOption) (*LookupComponentTypeResult, error) {
	var rv LookupComponentTypeResult
	err := ctx.Invoke("aws-native:iottwinmaker:getComponentType", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupComponentTypeArgs struct {
	// The ID of the component type.
	ComponentTypeId string `pulumi:"componentTypeId"`
	// The ID of the workspace that contains the component type.
	WorkspaceId string `pulumi:"workspaceId"`
}

type LookupComponentTypeResult struct {
	// The ARN of the component type.
	Arn *string `pulumi:"arn"`
	// The date and time when the component type was created.
	CreationDateTime *string `pulumi:"creationDateTime"`
	// The description of the component type.
	Description *string `pulumi:"description"`
	// Specifies the parent component type to extend.
	ExtendsFrom []string `pulumi:"extendsFrom"`
	// a Map of functions in the component type. Each function's key must be unique to this map.
	Functions interface{} `pulumi:"functions"`
	// A Boolean value that specifies whether the component type is abstract.
	IsAbstract *bool `pulumi:"isAbstract"`
	// A Boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run.
	IsSchemaInitialized *bool `pulumi:"isSchemaInitialized"`
	// A Boolean value that specifies whether an entity can have more than one component of this type.
	IsSingleton *bool `pulumi:"isSingleton"`
	// An map of the property definitions in the component type. Each property definition's key must be unique to this map.
	PropertyDefinitions interface{} `pulumi:"propertyDefinitions"`
	// The current status of the component type.
	Status *ComponentTypeStatus `pulumi:"status"`
	// A map of key-value pairs to associate with a resource.
	Tags interface{} `pulumi:"tags"`
	// The last date and time when the component type was updated.
	UpdateDateTime *string `pulumi:"updateDateTime"`
}

func LookupComponentTypeOutput(ctx *pulumi.Context, args LookupComponentTypeOutputArgs, opts ...pulumi.InvokeOption) LookupComponentTypeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComponentTypeResult, error) {
			args := v.(LookupComponentTypeArgs)
			r, err := LookupComponentType(ctx, &args, opts...)
			var s LookupComponentTypeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComponentTypeResultOutput)
}

type LookupComponentTypeOutputArgs struct {
	// The ID of the component type.
	ComponentTypeId pulumi.StringInput `pulumi:"componentTypeId"`
	// The ID of the workspace that contains the component type.
	WorkspaceId pulumi.StringInput `pulumi:"workspaceId"`
}

func (LookupComponentTypeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentTypeArgs)(nil)).Elem()
}

type LookupComponentTypeResultOutput struct{ *pulumi.OutputState }

func (LookupComponentTypeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComponentTypeResult)(nil)).Elem()
}

func (o LookupComponentTypeResultOutput) ToLookupComponentTypeResultOutput() LookupComponentTypeResultOutput {
	return o
}

func (o LookupComponentTypeResultOutput) ToLookupComponentTypeResultOutputWithContext(ctx context.Context) LookupComponentTypeResultOutput {
	return o
}

// The ARN of the component type.
func (o LookupComponentTypeResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The date and time when the component type was created.
func (o LookupComponentTypeResultOutput) CreationDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *string { return v.CreationDateTime }).(pulumi.StringPtrOutput)
}

// The description of the component type.
func (o LookupComponentTypeResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the parent component type to extend.
func (o LookupComponentTypeResultOutput) ExtendsFrom() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) []string { return v.ExtendsFrom }).(pulumi.StringArrayOutput)
}

// a Map of functions in the component type. Each function's key must be unique to this map.
func (o LookupComponentTypeResultOutput) Functions() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) interface{} { return v.Functions }).(pulumi.AnyOutput)
}

// A Boolean value that specifies whether the component type is abstract.
func (o LookupComponentTypeResultOutput) IsAbstract() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *bool { return v.IsAbstract }).(pulumi.BoolPtrOutput)
}

// A Boolean value that specifies whether the component type has a schema initializer and that the schema initializer has run.
func (o LookupComponentTypeResultOutput) IsSchemaInitialized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *bool { return v.IsSchemaInitialized }).(pulumi.BoolPtrOutput)
}

// A Boolean value that specifies whether an entity can have more than one component of this type.
func (o LookupComponentTypeResultOutput) IsSingleton() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *bool { return v.IsSingleton }).(pulumi.BoolPtrOutput)
}

// An map of the property definitions in the component type. Each property definition's key must be unique to this map.
func (o LookupComponentTypeResultOutput) PropertyDefinitions() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) interface{} { return v.PropertyDefinitions }).(pulumi.AnyOutput)
}

// The current status of the component type.
func (o LookupComponentTypeResultOutput) Status() ComponentTypeStatusPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *ComponentTypeStatus { return v.Status }).(ComponentTypeStatusPtrOutput)
}

// A map of key-value pairs to associate with a resource.
func (o LookupComponentTypeResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

// The last date and time when the component type was updated.
func (o LookupComponentTypeResultOutput) UpdateDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComponentTypeResult) *string { return v.UpdateDateTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComponentTypeResultOutput{})
}
