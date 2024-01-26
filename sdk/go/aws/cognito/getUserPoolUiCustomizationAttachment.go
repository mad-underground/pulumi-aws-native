// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
func LookupUserPoolUiCustomizationAttachment(ctx *pulumi.Context, args *LookupUserPoolUiCustomizationAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolUiCustomizationAttachmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserPoolUiCustomizationAttachmentResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolUiCustomizationAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolUiCustomizationAttachmentArgs struct {
	ClientId   string `pulumi:"clientId"`
	UserPoolId string `pulumi:"userPoolId"`
}

type LookupUserPoolUiCustomizationAttachmentResult struct {
	Css *string `pulumi:"css"`
}

func LookupUserPoolUiCustomizationAttachmentOutput(ctx *pulumi.Context, args LookupUserPoolUiCustomizationAttachmentOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolUiCustomizationAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolUiCustomizationAttachmentResult, error) {
			args := v.(LookupUserPoolUiCustomizationAttachmentArgs)
			r, err := LookupUserPoolUiCustomizationAttachment(ctx, &args, opts...)
			var s LookupUserPoolUiCustomizationAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolUiCustomizationAttachmentResultOutput)
}

type LookupUserPoolUiCustomizationAttachmentOutputArgs struct {
	ClientId   pulumi.StringInput `pulumi:"clientId"`
	UserPoolId pulumi.StringInput `pulumi:"userPoolId"`
}

func (LookupUserPoolUiCustomizationAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolUiCustomizationAttachmentArgs)(nil)).Elem()
}

type LookupUserPoolUiCustomizationAttachmentResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolUiCustomizationAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolUiCustomizationAttachmentResult)(nil)).Elem()
}

func (o LookupUserPoolUiCustomizationAttachmentResultOutput) ToLookupUserPoolUiCustomizationAttachmentResultOutput() LookupUserPoolUiCustomizationAttachmentResultOutput {
	return o
}

func (o LookupUserPoolUiCustomizationAttachmentResultOutput) ToLookupUserPoolUiCustomizationAttachmentResultOutputWithContext(ctx context.Context) LookupUserPoolUiCustomizationAttachmentResultOutput {
	return o
}

func (o LookupUserPoolUiCustomizationAttachmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserPoolUiCustomizationAttachmentResult] {
	return pulumix.Output[LookupUserPoolUiCustomizationAttachmentResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupUserPoolUiCustomizationAttachmentResultOutput) Css() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolUiCustomizationAttachmentResult) *string { return v.Css }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolUiCustomizationAttachmentResultOutput{})
}
