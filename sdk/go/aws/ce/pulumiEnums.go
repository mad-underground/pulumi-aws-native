// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ce

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CostCategoryRuleVersion string

const (
	CostCategoryRuleVersionCostCategoryExpressionV1 = CostCategoryRuleVersion("CostCategoryExpression.v1")
)

func (CostCategoryRuleVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*CostCategoryRuleVersion)(nil)).Elem()
}

func (e CostCategoryRuleVersion) ToCostCategoryRuleVersionOutput() CostCategoryRuleVersionOutput {
	return pulumi.ToOutput(e).(CostCategoryRuleVersionOutput)
}

func (e CostCategoryRuleVersion) ToCostCategoryRuleVersionOutputWithContext(ctx context.Context) CostCategoryRuleVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CostCategoryRuleVersionOutput)
}

func (e CostCategoryRuleVersion) ToCostCategoryRuleVersionPtrOutput() CostCategoryRuleVersionPtrOutput {
	return e.ToCostCategoryRuleVersionPtrOutputWithContext(context.Background())
}

func (e CostCategoryRuleVersion) ToCostCategoryRuleVersionPtrOutputWithContext(ctx context.Context) CostCategoryRuleVersionPtrOutput {
	return CostCategoryRuleVersion(e).ToCostCategoryRuleVersionOutputWithContext(ctx).ToCostCategoryRuleVersionPtrOutputWithContext(ctx)
}

func (e CostCategoryRuleVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CostCategoryRuleVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CostCategoryRuleVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CostCategoryRuleVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CostCategoryRuleVersionOutput struct{ *pulumi.OutputState }

func (CostCategoryRuleVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostCategoryRuleVersion)(nil)).Elem()
}

func (o CostCategoryRuleVersionOutput) ToCostCategoryRuleVersionOutput() CostCategoryRuleVersionOutput {
	return o
}

func (o CostCategoryRuleVersionOutput) ToCostCategoryRuleVersionOutputWithContext(ctx context.Context) CostCategoryRuleVersionOutput {
	return o
}

func (o CostCategoryRuleVersionOutput) ToCostCategoryRuleVersionPtrOutput() CostCategoryRuleVersionPtrOutput {
	return o.ToCostCategoryRuleVersionPtrOutputWithContext(context.Background())
}

func (o CostCategoryRuleVersionOutput) ToCostCategoryRuleVersionPtrOutputWithContext(ctx context.Context) CostCategoryRuleVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CostCategoryRuleVersion) *CostCategoryRuleVersion {
		return &v
	}).(CostCategoryRuleVersionPtrOutput)
}

func (o CostCategoryRuleVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CostCategoryRuleVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CostCategoryRuleVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CostCategoryRuleVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CostCategoryRuleVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CostCategoryRuleVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CostCategoryRuleVersionPtrOutput struct{ *pulumi.OutputState }

func (CostCategoryRuleVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostCategoryRuleVersion)(nil)).Elem()
}

func (o CostCategoryRuleVersionPtrOutput) ToCostCategoryRuleVersionPtrOutput() CostCategoryRuleVersionPtrOutput {
	return o
}

func (o CostCategoryRuleVersionPtrOutput) ToCostCategoryRuleVersionPtrOutputWithContext(ctx context.Context) CostCategoryRuleVersionPtrOutput {
	return o
}

func (o CostCategoryRuleVersionPtrOutput) Elem() CostCategoryRuleVersionOutput {
	return o.ApplyT(func(v *CostCategoryRuleVersion) CostCategoryRuleVersion {
		if v != nil {
			return *v
		}
		var ret CostCategoryRuleVersion
		return ret
	}).(CostCategoryRuleVersionOutput)
}

func (o CostCategoryRuleVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CostCategoryRuleVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CostCategoryRuleVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CostCategoryRuleVersionInput is an input type that accepts CostCategoryRuleVersionArgs and CostCategoryRuleVersionOutput values.
// You can construct a concrete instance of `CostCategoryRuleVersionInput` via:
//
//          CostCategoryRuleVersionArgs{...}
type CostCategoryRuleVersionInput interface {
	pulumi.Input

	ToCostCategoryRuleVersionOutput() CostCategoryRuleVersionOutput
	ToCostCategoryRuleVersionOutputWithContext(context.Context) CostCategoryRuleVersionOutput
}

var costCategoryRuleVersionPtrType = reflect.TypeOf((**CostCategoryRuleVersion)(nil)).Elem()

type CostCategoryRuleVersionPtrInput interface {
	pulumi.Input

	ToCostCategoryRuleVersionPtrOutput() CostCategoryRuleVersionPtrOutput
	ToCostCategoryRuleVersionPtrOutputWithContext(context.Context) CostCategoryRuleVersionPtrOutput
}

type costCategoryRuleVersionPtr string

func CostCategoryRuleVersionPtr(v string) CostCategoryRuleVersionPtrInput {
	return (*costCategoryRuleVersionPtr)(&v)
}

func (*costCategoryRuleVersionPtr) ElementType() reflect.Type {
	return costCategoryRuleVersionPtrType
}

func (in *costCategoryRuleVersionPtr) ToCostCategoryRuleVersionPtrOutput() CostCategoryRuleVersionPtrOutput {
	return pulumi.ToOutput(in).(CostCategoryRuleVersionPtrOutput)
}

func (in *costCategoryRuleVersionPtr) ToCostCategoryRuleVersionPtrOutputWithContext(ctx context.Context) CostCategoryRuleVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CostCategoryRuleVersionPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CostCategoryRuleVersionOutput{})
	pulumi.RegisterOutputType(CostCategoryRuleVersionPtrOutput{})
}