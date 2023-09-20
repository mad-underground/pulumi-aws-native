// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package entityresolution

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

type MatchingWorkflowResolutionTechniquesResolutionType string

const (
	MatchingWorkflowResolutionTechniquesResolutionTypeRuleMatching = MatchingWorkflowResolutionTechniquesResolutionType("RULE_MATCHING")
	MatchingWorkflowResolutionTechniquesResolutionTypeMlMatching   = MatchingWorkflowResolutionTechniquesResolutionType("ML_MATCHING")
)

func (MatchingWorkflowResolutionTechniquesResolutionType) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchingWorkflowResolutionTechniquesResolutionType)(nil)).Elem()
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToMatchingWorkflowResolutionTechniquesResolutionTypeOutput() MatchingWorkflowResolutionTechniquesResolutionTypeOutput {
	return pulumi.ToOutput(e).(MatchingWorkflowResolutionTechniquesResolutionTypeOutput)
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToMatchingWorkflowResolutionTechniquesResolutionTypeOutputWithContext(ctx context.Context) MatchingWorkflowResolutionTechniquesResolutionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MatchingWorkflowResolutionTechniquesResolutionTypeOutput)
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutput() MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return e.ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(context.Background())
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(ctx context.Context) MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return MatchingWorkflowResolutionTechniquesResolutionType(e).ToMatchingWorkflowResolutionTechniquesResolutionTypeOutputWithContext(ctx).ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(ctx)
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MatchingWorkflowResolutionTechniquesResolutionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MatchingWorkflowResolutionTechniquesResolutionTypeOutput struct{ *pulumi.OutputState }

func (MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchingWorkflowResolutionTechniquesResolutionType)(nil)).Elem()
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToMatchingWorkflowResolutionTechniquesResolutionTypeOutput() MatchingWorkflowResolutionTechniquesResolutionTypeOutput {
	return o
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToMatchingWorkflowResolutionTechniquesResolutionTypeOutputWithContext(ctx context.Context) MatchingWorkflowResolutionTechniquesResolutionTypeOutput {
	return o
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutput() MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return o.ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(context.Background())
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(ctx context.Context) MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MatchingWorkflowResolutionTechniquesResolutionType) *MatchingWorkflowResolutionTechniquesResolutionType {
		return &v
	}).(MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput)
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToOutput(ctx context.Context) pulumix.Output[MatchingWorkflowResolutionTechniquesResolutionType] {
	return pulumix.Output[MatchingWorkflowResolutionTechniquesResolutionType]{
		OutputState: o.OutputState,
	}
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MatchingWorkflowResolutionTechniquesResolutionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MatchingWorkflowResolutionTechniquesResolutionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput struct{ *pulumi.OutputState }

func (MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MatchingWorkflowResolutionTechniquesResolutionType)(nil)).Elem()
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutput() MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return o
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(ctx context.Context) MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return o
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*MatchingWorkflowResolutionTechniquesResolutionType] {
	return pulumix.Output[*MatchingWorkflowResolutionTechniquesResolutionType]{
		OutputState: o.OutputState,
	}
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput) Elem() MatchingWorkflowResolutionTechniquesResolutionTypeOutput {
	return o.ApplyT(func(v *MatchingWorkflowResolutionTechniquesResolutionType) MatchingWorkflowResolutionTechniquesResolutionType {
		if v != nil {
			return *v
		}
		var ret MatchingWorkflowResolutionTechniquesResolutionType
		return ret
	}).(MatchingWorkflowResolutionTechniquesResolutionTypeOutput)
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MatchingWorkflowResolutionTechniquesResolutionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MatchingWorkflowResolutionTechniquesResolutionTypeInput is an input type that accepts MatchingWorkflowResolutionTechniquesResolutionTypeArgs and MatchingWorkflowResolutionTechniquesResolutionTypeOutput values.
// You can construct a concrete instance of `MatchingWorkflowResolutionTechniquesResolutionTypeInput` via:
//
//	MatchingWorkflowResolutionTechniquesResolutionTypeArgs{...}
type MatchingWorkflowResolutionTechniquesResolutionTypeInput interface {
	pulumi.Input

	ToMatchingWorkflowResolutionTechniquesResolutionTypeOutput() MatchingWorkflowResolutionTechniquesResolutionTypeOutput
	ToMatchingWorkflowResolutionTechniquesResolutionTypeOutputWithContext(context.Context) MatchingWorkflowResolutionTechniquesResolutionTypeOutput
}

var matchingWorkflowResolutionTechniquesResolutionTypePtrType = reflect.TypeOf((**MatchingWorkflowResolutionTechniquesResolutionType)(nil)).Elem()

type MatchingWorkflowResolutionTechniquesResolutionTypePtrInput interface {
	pulumi.Input

	ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutput() MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput
	ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(context.Context) MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput
}

type matchingWorkflowResolutionTechniquesResolutionTypePtr string

func MatchingWorkflowResolutionTechniquesResolutionTypePtr(v string) MatchingWorkflowResolutionTechniquesResolutionTypePtrInput {
	return (*matchingWorkflowResolutionTechniquesResolutionTypePtr)(&v)
}

func (*matchingWorkflowResolutionTechniquesResolutionTypePtr) ElementType() reflect.Type {
	return matchingWorkflowResolutionTechniquesResolutionTypePtrType
}

func (in *matchingWorkflowResolutionTechniquesResolutionTypePtr) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutput() MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return pulumi.ToOutput(in).(MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput)
}

func (in *matchingWorkflowResolutionTechniquesResolutionTypePtr) ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(ctx context.Context) MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput)
}

func (in *matchingWorkflowResolutionTechniquesResolutionTypePtr) ToOutput(ctx context.Context) pulumix.Output[*MatchingWorkflowResolutionTechniquesResolutionType] {
	return pulumix.Output[*MatchingWorkflowResolutionTechniquesResolutionType]{
		OutputState: in.ToMatchingWorkflowResolutionTechniquesResolutionTypePtrOutputWithContext(ctx).OutputState,
	}
}

type MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel string

const (
	MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOneToOne   = MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel("ONE_TO_ONE")
	MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelManyToMany = MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel("MANY_TO_MANY")
)

func (MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel)(nil)).Elem()
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput {
	return pulumi.ToOutput(e).(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput)
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutputWithContext(ctx context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput)
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return e.ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(context.Background())
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(ctx context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel(e).ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutputWithContext(ctx).ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(ctx)
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput struct{ *pulumi.OutputState }

func (MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel)(nil)).Elem()
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput {
	return o
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutputWithContext(ctx context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput {
	return o
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return o.ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(context.Background())
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(ctx context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) *MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel {
		return &v
	}).(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput)
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToOutput(ctx context.Context) pulumix.Output[MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel] {
	return pulumix.Output[MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel]{
		OutputState: o.OutputState,
	}
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput struct{ *pulumi.OutputState }

func (MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel)(nil)).Elem()
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return o
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(ctx context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return o
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput) ToOutput(ctx context.Context) pulumix.Output[*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel] {
	return pulumix.Output[*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel]{
		OutputState: o.OutputState,
	}
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput) Elem() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput {
	return o.ApplyT(func(v *MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel {
		if v != nil {
			return *v
		}
		var ret MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel
		return ret
	}).(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput)
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelInput is an input type that accepts MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelArgs and MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput values.
// You can construct a concrete instance of `MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelInput` via:
//
//	MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelArgs{...}
type MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelInput interface {
	pulumi.Input

	ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput
	ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutputWithContext(context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput
}

var matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrType = reflect.TypeOf((**MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel)(nil)).Elem()

type MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrInput interface {
	pulumi.Input

	ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput
	ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput
}

type matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtr string

func MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtr(v string) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrInput {
	return (*matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtr)(&v)
}

func (*matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtr) ElementType() reflect.Type {
	return matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrType
}

func (in *matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtr) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput() MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return pulumi.ToOutput(in).(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput)
}

func (in *matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtr) ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(ctx context.Context) MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput)
}

func (in *matchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtr) ToOutput(ctx context.Context) pulumix.Output[*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel] {
	return pulumix.Output[*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel]{
		OutputState: in.ToMatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutputWithContext(ctx).OutputState,
	}
}

type SchemaMappingSchemaAttributeType string

const (
	SchemaMappingSchemaAttributeTypeName              = SchemaMappingSchemaAttributeType("NAME")
	SchemaMappingSchemaAttributeTypeNameFirst         = SchemaMappingSchemaAttributeType("NAME_FIRST")
	SchemaMappingSchemaAttributeTypeNameMiddle        = SchemaMappingSchemaAttributeType("NAME_MIDDLE")
	SchemaMappingSchemaAttributeTypeNameLast          = SchemaMappingSchemaAttributeType("NAME_LAST")
	SchemaMappingSchemaAttributeTypeAddress           = SchemaMappingSchemaAttributeType("ADDRESS")
	SchemaMappingSchemaAttributeTypeAddressStreet1    = SchemaMappingSchemaAttributeType("ADDRESS_STREET1")
	SchemaMappingSchemaAttributeTypeAddressStreet2    = SchemaMappingSchemaAttributeType("ADDRESS_STREET2")
	SchemaMappingSchemaAttributeTypeAddressStreet3    = SchemaMappingSchemaAttributeType("ADDRESS_STREET3")
	SchemaMappingSchemaAttributeTypeAddressCity       = SchemaMappingSchemaAttributeType("ADDRESS_CITY")
	SchemaMappingSchemaAttributeTypeAddressState      = SchemaMappingSchemaAttributeType("ADDRESS_STATE")
	SchemaMappingSchemaAttributeTypeAddressCountry    = SchemaMappingSchemaAttributeType("ADDRESS_COUNTRY")
	SchemaMappingSchemaAttributeTypeAddressPostalcode = SchemaMappingSchemaAttributeType("ADDRESS_POSTALCODE")
	SchemaMappingSchemaAttributeTypePhone             = SchemaMappingSchemaAttributeType("PHONE")
	SchemaMappingSchemaAttributeTypePhoneNumber       = SchemaMappingSchemaAttributeType("PHONE_NUMBER")
	SchemaMappingSchemaAttributeTypePhoneCountrycode  = SchemaMappingSchemaAttributeType("PHONE_COUNTRYCODE")
	SchemaMappingSchemaAttributeTypeEmailAddress      = SchemaMappingSchemaAttributeType("EMAIL_ADDRESS")
	SchemaMappingSchemaAttributeTypeUniqueId          = SchemaMappingSchemaAttributeType("UNIQUE_ID")
	SchemaMappingSchemaAttributeTypeDate              = SchemaMappingSchemaAttributeType("DATE")
	SchemaMappingSchemaAttributeTypeString            = SchemaMappingSchemaAttributeType("STRING")
)

func (SchemaMappingSchemaAttributeType) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaMappingSchemaAttributeType)(nil)).Elem()
}

func (e SchemaMappingSchemaAttributeType) ToSchemaMappingSchemaAttributeTypeOutput() SchemaMappingSchemaAttributeTypeOutput {
	return pulumi.ToOutput(e).(SchemaMappingSchemaAttributeTypeOutput)
}

func (e SchemaMappingSchemaAttributeType) ToSchemaMappingSchemaAttributeTypeOutputWithContext(ctx context.Context) SchemaMappingSchemaAttributeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SchemaMappingSchemaAttributeTypeOutput)
}

func (e SchemaMappingSchemaAttributeType) ToSchemaMappingSchemaAttributeTypePtrOutput() SchemaMappingSchemaAttributeTypePtrOutput {
	return e.ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(context.Background())
}

func (e SchemaMappingSchemaAttributeType) ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(ctx context.Context) SchemaMappingSchemaAttributeTypePtrOutput {
	return SchemaMappingSchemaAttributeType(e).ToSchemaMappingSchemaAttributeTypeOutputWithContext(ctx).ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(ctx)
}

func (e SchemaMappingSchemaAttributeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SchemaMappingSchemaAttributeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SchemaMappingSchemaAttributeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SchemaMappingSchemaAttributeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SchemaMappingSchemaAttributeTypeOutput struct{ *pulumi.OutputState }

func (SchemaMappingSchemaAttributeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchemaMappingSchemaAttributeType)(nil)).Elem()
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToSchemaMappingSchemaAttributeTypeOutput() SchemaMappingSchemaAttributeTypeOutput {
	return o
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToSchemaMappingSchemaAttributeTypeOutputWithContext(ctx context.Context) SchemaMappingSchemaAttributeTypeOutput {
	return o
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToSchemaMappingSchemaAttributeTypePtrOutput() SchemaMappingSchemaAttributeTypePtrOutput {
	return o.ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(context.Background())
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(ctx context.Context) SchemaMappingSchemaAttributeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SchemaMappingSchemaAttributeType) *SchemaMappingSchemaAttributeType {
		return &v
	}).(SchemaMappingSchemaAttributeTypePtrOutput)
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToOutput(ctx context.Context) pulumix.Output[SchemaMappingSchemaAttributeType] {
	return pulumix.Output[SchemaMappingSchemaAttributeType]{
		OutputState: o.OutputState,
	}
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SchemaMappingSchemaAttributeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SchemaMappingSchemaAttributeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SchemaMappingSchemaAttributeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SchemaMappingSchemaAttributeTypePtrOutput struct{ *pulumi.OutputState }

func (SchemaMappingSchemaAttributeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SchemaMappingSchemaAttributeType)(nil)).Elem()
}

func (o SchemaMappingSchemaAttributeTypePtrOutput) ToSchemaMappingSchemaAttributeTypePtrOutput() SchemaMappingSchemaAttributeTypePtrOutput {
	return o
}

func (o SchemaMappingSchemaAttributeTypePtrOutput) ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(ctx context.Context) SchemaMappingSchemaAttributeTypePtrOutput {
	return o
}

func (o SchemaMappingSchemaAttributeTypePtrOutput) ToOutput(ctx context.Context) pulumix.Output[*SchemaMappingSchemaAttributeType] {
	return pulumix.Output[*SchemaMappingSchemaAttributeType]{
		OutputState: o.OutputState,
	}
}

func (o SchemaMappingSchemaAttributeTypePtrOutput) Elem() SchemaMappingSchemaAttributeTypeOutput {
	return o.ApplyT(func(v *SchemaMappingSchemaAttributeType) SchemaMappingSchemaAttributeType {
		if v != nil {
			return *v
		}
		var ret SchemaMappingSchemaAttributeType
		return ret
	}).(SchemaMappingSchemaAttributeTypeOutput)
}

func (o SchemaMappingSchemaAttributeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SchemaMappingSchemaAttributeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SchemaMappingSchemaAttributeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SchemaMappingSchemaAttributeTypeInput is an input type that accepts SchemaMappingSchemaAttributeTypeArgs and SchemaMappingSchemaAttributeTypeOutput values.
// You can construct a concrete instance of `SchemaMappingSchemaAttributeTypeInput` via:
//
//	SchemaMappingSchemaAttributeTypeArgs{...}
type SchemaMappingSchemaAttributeTypeInput interface {
	pulumi.Input

	ToSchemaMappingSchemaAttributeTypeOutput() SchemaMappingSchemaAttributeTypeOutput
	ToSchemaMappingSchemaAttributeTypeOutputWithContext(context.Context) SchemaMappingSchemaAttributeTypeOutput
}

var schemaMappingSchemaAttributeTypePtrType = reflect.TypeOf((**SchemaMappingSchemaAttributeType)(nil)).Elem()

type SchemaMappingSchemaAttributeTypePtrInput interface {
	pulumi.Input

	ToSchemaMappingSchemaAttributeTypePtrOutput() SchemaMappingSchemaAttributeTypePtrOutput
	ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(context.Context) SchemaMappingSchemaAttributeTypePtrOutput
}

type schemaMappingSchemaAttributeTypePtr string

func SchemaMappingSchemaAttributeTypePtr(v string) SchemaMappingSchemaAttributeTypePtrInput {
	return (*schemaMappingSchemaAttributeTypePtr)(&v)
}

func (*schemaMappingSchemaAttributeTypePtr) ElementType() reflect.Type {
	return schemaMappingSchemaAttributeTypePtrType
}

func (in *schemaMappingSchemaAttributeTypePtr) ToSchemaMappingSchemaAttributeTypePtrOutput() SchemaMappingSchemaAttributeTypePtrOutput {
	return pulumi.ToOutput(in).(SchemaMappingSchemaAttributeTypePtrOutput)
}

func (in *schemaMappingSchemaAttributeTypePtr) ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(ctx context.Context) SchemaMappingSchemaAttributeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SchemaMappingSchemaAttributeTypePtrOutput)
}

func (in *schemaMappingSchemaAttributeTypePtr) ToOutput(ctx context.Context) pulumix.Output[*SchemaMappingSchemaAttributeType] {
	return pulumix.Output[*SchemaMappingSchemaAttributeType]{
		OutputState: in.ToSchemaMappingSchemaAttributeTypePtrOutputWithContext(ctx).OutputState,
	}
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MatchingWorkflowResolutionTechniquesResolutionTypeInput)(nil)).Elem(), MatchingWorkflowResolutionTechniquesResolutionType("RULE_MATCHING"))
	pulumi.RegisterInputType(reflect.TypeOf((*MatchingWorkflowResolutionTechniquesResolutionTypePtrInput)(nil)).Elem(), MatchingWorkflowResolutionTechniquesResolutionType("RULE_MATCHING"))
	pulumi.RegisterInputType(reflect.TypeOf((*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelInput)(nil)).Elem(), MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel("ONE_TO_ONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrInput)(nil)).Elem(), MatchingWorkflowRuleBasedPropertiesAttributeMatchingModel("ONE_TO_ONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaMappingSchemaAttributeTypeInput)(nil)).Elem(), SchemaMappingSchemaAttributeType("NAME"))
	pulumi.RegisterInputType(reflect.TypeOf((*SchemaMappingSchemaAttributeTypePtrInput)(nil)).Elem(), SchemaMappingSchemaAttributeType("NAME"))
	pulumi.RegisterOutputType(MatchingWorkflowResolutionTechniquesResolutionTypeOutput{})
	pulumi.RegisterOutputType(MatchingWorkflowResolutionTechniquesResolutionTypePtrOutput{})
	pulumi.RegisterOutputType(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelOutput{})
	pulumi.RegisterOutputType(MatchingWorkflowRuleBasedPropertiesAttributeMatchingModelPtrOutput{})
	pulumi.RegisterOutputType(SchemaMappingSchemaAttributeTypeOutput{})
	pulumi.RegisterOutputType(SchemaMappingSchemaAttributeTypePtrOutput{})
}