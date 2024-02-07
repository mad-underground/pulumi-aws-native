// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package verifiedpermissions

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IdentitySourceOpenIdIssuer string

const (
	IdentitySourceOpenIdIssuerCognito = IdentitySourceOpenIdIssuer("COGNITO")
)

type IdentitySourceOpenIdIssuerOutput struct{ *pulumi.OutputState }

func (IdentitySourceOpenIdIssuerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentitySourceOpenIdIssuer)(nil)).Elem()
}

func (o IdentitySourceOpenIdIssuerOutput) ToIdentitySourceOpenIdIssuerOutput() IdentitySourceOpenIdIssuerOutput {
	return o
}

func (o IdentitySourceOpenIdIssuerOutput) ToIdentitySourceOpenIdIssuerOutputWithContext(ctx context.Context) IdentitySourceOpenIdIssuerOutput {
	return o
}

func (o IdentitySourceOpenIdIssuerOutput) ToIdentitySourceOpenIdIssuerPtrOutput() IdentitySourceOpenIdIssuerPtrOutput {
	return o.ToIdentitySourceOpenIdIssuerPtrOutputWithContext(context.Background())
}

func (o IdentitySourceOpenIdIssuerOutput) ToIdentitySourceOpenIdIssuerPtrOutputWithContext(ctx context.Context) IdentitySourceOpenIdIssuerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentitySourceOpenIdIssuer) *IdentitySourceOpenIdIssuer {
		return &v
	}).(IdentitySourceOpenIdIssuerPtrOutput)
}

func (o IdentitySourceOpenIdIssuerOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentitySourceOpenIdIssuerOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentitySourceOpenIdIssuer) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentitySourceOpenIdIssuerOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentitySourceOpenIdIssuerOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentitySourceOpenIdIssuer) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentitySourceOpenIdIssuerPtrOutput struct{ *pulumi.OutputState }

func (IdentitySourceOpenIdIssuerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentitySourceOpenIdIssuer)(nil)).Elem()
}

func (o IdentitySourceOpenIdIssuerPtrOutput) ToIdentitySourceOpenIdIssuerPtrOutput() IdentitySourceOpenIdIssuerPtrOutput {
	return o
}

func (o IdentitySourceOpenIdIssuerPtrOutput) ToIdentitySourceOpenIdIssuerPtrOutputWithContext(ctx context.Context) IdentitySourceOpenIdIssuerPtrOutput {
	return o
}

func (o IdentitySourceOpenIdIssuerPtrOutput) Elem() IdentitySourceOpenIdIssuerOutput {
	return o.ApplyT(func(v *IdentitySourceOpenIdIssuer) IdentitySourceOpenIdIssuer {
		if v != nil {
			return *v
		}
		var ret IdentitySourceOpenIdIssuer
		return ret
	}).(IdentitySourceOpenIdIssuerOutput)
}

func (o IdentitySourceOpenIdIssuerPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentitySourceOpenIdIssuerPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentitySourceOpenIdIssuer) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyStoreValidationMode string

const (
	PolicyStoreValidationModeOff    = PolicyStoreValidationMode("OFF")
	PolicyStoreValidationModeStrict = PolicyStoreValidationMode("STRICT")
)

func (PolicyStoreValidationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStoreValidationMode)(nil)).Elem()
}

func (e PolicyStoreValidationMode) ToPolicyStoreValidationModeOutput() PolicyStoreValidationModeOutput {
	return pulumi.ToOutput(e).(PolicyStoreValidationModeOutput)
}

func (e PolicyStoreValidationMode) ToPolicyStoreValidationModeOutputWithContext(ctx context.Context) PolicyStoreValidationModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyStoreValidationModeOutput)
}

func (e PolicyStoreValidationMode) ToPolicyStoreValidationModePtrOutput() PolicyStoreValidationModePtrOutput {
	return e.ToPolicyStoreValidationModePtrOutputWithContext(context.Background())
}

func (e PolicyStoreValidationMode) ToPolicyStoreValidationModePtrOutputWithContext(ctx context.Context) PolicyStoreValidationModePtrOutput {
	return PolicyStoreValidationMode(e).ToPolicyStoreValidationModeOutputWithContext(ctx).ToPolicyStoreValidationModePtrOutputWithContext(ctx)
}

func (e PolicyStoreValidationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyStoreValidationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyStoreValidationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyStoreValidationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyStoreValidationModeOutput struct{ *pulumi.OutputState }

func (PolicyStoreValidationModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyStoreValidationMode)(nil)).Elem()
}

func (o PolicyStoreValidationModeOutput) ToPolicyStoreValidationModeOutput() PolicyStoreValidationModeOutput {
	return o
}

func (o PolicyStoreValidationModeOutput) ToPolicyStoreValidationModeOutputWithContext(ctx context.Context) PolicyStoreValidationModeOutput {
	return o
}

func (o PolicyStoreValidationModeOutput) ToPolicyStoreValidationModePtrOutput() PolicyStoreValidationModePtrOutput {
	return o.ToPolicyStoreValidationModePtrOutputWithContext(context.Background())
}

func (o PolicyStoreValidationModeOutput) ToPolicyStoreValidationModePtrOutputWithContext(ctx context.Context) PolicyStoreValidationModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyStoreValidationMode) *PolicyStoreValidationMode {
		return &v
	}).(PolicyStoreValidationModePtrOutput)
}

func (o PolicyStoreValidationModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyStoreValidationModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyStoreValidationMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyStoreValidationModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyStoreValidationModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyStoreValidationMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyStoreValidationModePtrOutput struct{ *pulumi.OutputState }

func (PolicyStoreValidationModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyStoreValidationMode)(nil)).Elem()
}

func (o PolicyStoreValidationModePtrOutput) ToPolicyStoreValidationModePtrOutput() PolicyStoreValidationModePtrOutput {
	return o
}

func (o PolicyStoreValidationModePtrOutput) ToPolicyStoreValidationModePtrOutputWithContext(ctx context.Context) PolicyStoreValidationModePtrOutput {
	return o
}

func (o PolicyStoreValidationModePtrOutput) Elem() PolicyStoreValidationModeOutput {
	return o.ApplyT(func(v *PolicyStoreValidationMode) PolicyStoreValidationMode {
		if v != nil {
			return *v
		}
		var ret PolicyStoreValidationMode
		return ret
	}).(PolicyStoreValidationModeOutput)
}

func (o PolicyStoreValidationModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyStoreValidationModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyStoreValidationMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// PolicyStoreValidationModeInput is an input type that accepts values of the PolicyStoreValidationMode enum
// A concrete instance of `PolicyStoreValidationModeInput` can be one of the following:
//
//	PolicyStoreValidationModeOff
//	PolicyStoreValidationModeStrict
type PolicyStoreValidationModeInput interface {
	pulumi.Input

	ToPolicyStoreValidationModeOutput() PolicyStoreValidationModeOutput
	ToPolicyStoreValidationModeOutputWithContext(context.Context) PolicyStoreValidationModeOutput
}

var policyStoreValidationModePtrType = reflect.TypeOf((**PolicyStoreValidationMode)(nil)).Elem()

type PolicyStoreValidationModePtrInput interface {
	pulumi.Input

	ToPolicyStoreValidationModePtrOutput() PolicyStoreValidationModePtrOutput
	ToPolicyStoreValidationModePtrOutputWithContext(context.Context) PolicyStoreValidationModePtrOutput
}

type policyStoreValidationModePtr string

func PolicyStoreValidationModePtr(v string) PolicyStoreValidationModePtrInput {
	return (*policyStoreValidationModePtr)(&v)
}

func (*policyStoreValidationModePtr) ElementType() reflect.Type {
	return policyStoreValidationModePtrType
}

func (in *policyStoreValidationModePtr) ToPolicyStoreValidationModePtrOutput() PolicyStoreValidationModePtrOutput {
	return pulumi.ToOutput(in).(PolicyStoreValidationModePtrOutput)
}

func (in *policyStoreValidationModePtr) ToPolicyStoreValidationModePtrOutputWithContext(ctx context.Context) PolicyStoreValidationModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyStoreValidationModePtrOutput)
}

type PolicyType string

const (
	PolicyTypeStatic         = PolicyType("STATIC")
	PolicyTypeTemplateLinked = PolicyType("TEMPLATE_LINKED")
)

type PolicyTypeOutput struct{ *pulumi.OutputState }

func (PolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (o PolicyTypeOutput) ToPolicyTypeOutput() PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyType) *PolicyType {
		return &v
	}).(PolicyTypePtrOutput)
}

func (o PolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyTypePtrOutput struct{ *pulumi.OutputState }

func (PolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyType)(nil)).Elem()
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) Elem() PolicyTypeOutput {
	return o.ApplyT(func(v *PolicyType) PolicyType {
		if v != nil {
			return *v
		}
		var ret PolicyType
		return ret
	}).(PolicyTypeOutput)
}

func (o PolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStoreValidationModeInput)(nil)).Elem(), PolicyStoreValidationMode("OFF"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyStoreValidationModePtrInput)(nil)).Elem(), PolicyStoreValidationMode("OFF"))
	pulumi.RegisterOutputType(IdentitySourceOpenIdIssuerOutput{})
	pulumi.RegisterOutputType(IdentitySourceOpenIdIssuerPtrOutput{})
	pulumi.RegisterOutputType(PolicyStoreValidationModeOutput{})
	pulumi.RegisterOutputType(PolicyStoreValidationModePtrOutput{})
	pulumi.RegisterOutputType(PolicyTypeOutput{})
	pulumi.RegisterOutputType(PolicyTypePtrOutput{})
}
