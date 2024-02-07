// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a fleet provisioning template.
func LookupProvisioningTemplate(ctx *pulumi.Context, args *LookupProvisioningTemplateArgs, opts ...pulumi.InvokeOption) (*LookupProvisioningTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProvisioningTemplateResult
	err := ctx.Invoke("aws-native:iot:getProvisioningTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProvisioningTemplateArgs struct {
	TemplateName string `pulumi:"templateName"`
}

type LookupProvisioningTemplateResult struct {
	Description         *string                               `pulumi:"description"`
	Enabled             *bool                                 `pulumi:"enabled"`
	PreProvisioningHook *ProvisioningTemplateProvisioningHook `pulumi:"preProvisioningHook"`
	ProvisioningRoleArn *string                               `pulumi:"provisioningRoleArn"`
	Tags                []ProvisioningTemplateTag             `pulumi:"tags"`
	TemplateArn         *string                               `pulumi:"templateArn"`
	TemplateBody        *string                               `pulumi:"templateBody"`
}

func LookupProvisioningTemplateOutput(ctx *pulumi.Context, args LookupProvisioningTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupProvisioningTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProvisioningTemplateResult, error) {
			args := v.(LookupProvisioningTemplateArgs)
			r, err := LookupProvisioningTemplate(ctx, &args, opts...)
			var s LookupProvisioningTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProvisioningTemplateResultOutput)
}

type LookupProvisioningTemplateOutputArgs struct {
	TemplateName pulumi.StringInput `pulumi:"templateName"`
}

func (LookupProvisioningTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProvisioningTemplateArgs)(nil)).Elem()
}

type LookupProvisioningTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupProvisioningTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProvisioningTemplateResult)(nil)).Elem()
}

func (o LookupProvisioningTemplateResultOutput) ToLookupProvisioningTemplateResultOutput() LookupProvisioningTemplateResultOutput {
	return o
}

func (o LookupProvisioningTemplateResultOutput) ToLookupProvisioningTemplateResultOutputWithContext(ctx context.Context) LookupProvisioningTemplateResultOutput {
	return o
}

func (o LookupProvisioningTemplateResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProvisioningTemplateResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupProvisioningTemplateResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupProvisioningTemplateResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupProvisioningTemplateResultOutput) PreProvisioningHook() ProvisioningTemplateProvisioningHookPtrOutput {
	return o.ApplyT(func(v LookupProvisioningTemplateResult) *ProvisioningTemplateProvisioningHook {
		return v.PreProvisioningHook
	}).(ProvisioningTemplateProvisioningHookPtrOutput)
}

func (o LookupProvisioningTemplateResultOutput) ProvisioningRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProvisioningTemplateResult) *string { return v.ProvisioningRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupProvisioningTemplateResultOutput) Tags() ProvisioningTemplateTagArrayOutput {
	return o.ApplyT(func(v LookupProvisioningTemplateResult) []ProvisioningTemplateTag { return v.Tags }).(ProvisioningTemplateTagArrayOutput)
}

func (o LookupProvisioningTemplateResultOutput) TemplateArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProvisioningTemplateResult) *string { return v.TemplateArn }).(pulumi.StringPtrOutput)
}

func (o LookupProvisioningTemplateResultOutput) TemplateBody() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProvisioningTemplateResult) *string { return v.TemplateBody }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProvisioningTemplateResultOutput{})
}
