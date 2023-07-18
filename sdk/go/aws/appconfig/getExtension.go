// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppConfig::Extension
func LookupExtension(ctx *pulumi.Context, args *LookupExtensionArgs, opts ...pulumi.InvokeOption) (*LookupExtensionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupExtensionResult
	err := ctx.Invoke("aws-native:appconfig:getExtension", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExtensionArgs struct {
	Id string `pulumi:"id"`
}

type LookupExtensionResult struct {
	Actions interface{} `pulumi:"actions"`
	Arn     *string     `pulumi:"arn"`
	// Description of the extension.
	Description   *string     `pulumi:"description"`
	Id            *string     `pulumi:"id"`
	Parameters    interface{} `pulumi:"parameters"`
	VersionNumber *int        `pulumi:"versionNumber"`
}

func LookupExtensionOutput(ctx *pulumi.Context, args LookupExtensionOutputArgs, opts ...pulumi.InvokeOption) LookupExtensionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExtensionResult, error) {
			args := v.(LookupExtensionArgs)
			r, err := LookupExtension(ctx, &args, opts...)
			var s LookupExtensionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExtensionResultOutput)
}

type LookupExtensionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupExtensionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionArgs)(nil)).Elem()
}

type LookupExtensionResultOutput struct{ *pulumi.OutputState }

func (LookupExtensionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExtensionResult)(nil)).Elem()
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutput() LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) ToLookupExtensionResultOutputWithContext(ctx context.Context) LookupExtensionResultOutput {
	return o
}

func (o LookupExtensionResultOutput) Actions() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupExtensionResult) interface{} { return v.Actions }).(pulumi.AnyOutput)
}

func (o LookupExtensionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Description of the extension.
func (o LookupExtensionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExtensionResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupExtensionResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupExtensionResultOutput) VersionNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExtensionResult) *int { return v.VersionNumber }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExtensionResultOutput{})
}
