// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAFRegional::IPSet
func LookupIPSet(ctx *pulumi.Context, args *LookupIPSetArgs, opts ...pulumi.InvokeOption) (*LookupIPSetResult, error) {
	var rv LookupIPSetResult
	err := ctx.Invoke("aws-native:wafregional:getIPSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPSetArgs struct {
	Id string `pulumi:"id"`
}

type LookupIPSetResult struct {
	IPSetDescriptors []IPSetDescriptor `pulumi:"iPSetDescriptors"`
	Id               *string           `pulumi:"id"`
}

func LookupIPSetOutput(ctx *pulumi.Context, args LookupIPSetOutputArgs, opts ...pulumi.InvokeOption) LookupIPSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPSetResult, error) {
			args := v.(LookupIPSetArgs)
			r, err := LookupIPSet(ctx, &args, opts...)
			return *r, err
		}).(LookupIPSetResultOutput)
}

type LookupIPSetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIPSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPSetArgs)(nil)).Elem()
}

type LookupIPSetResultOutput struct{ *pulumi.OutputState }

func (LookupIPSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPSetResult)(nil)).Elem()
}

func (o LookupIPSetResultOutput) ToLookupIPSetResultOutput() LookupIPSetResultOutput {
	return o
}

func (o LookupIPSetResultOutput) ToLookupIPSetResultOutputWithContext(ctx context.Context) LookupIPSetResultOutput {
	return o
}

func (o LookupIPSetResultOutput) IPSetDescriptors() IPSetDescriptorArrayOutput {
	return o.ApplyT(func(v LookupIPSetResult) []IPSetDescriptor { return v.IPSetDescriptors }).(IPSetDescriptorArrayOutput)
}

func (o LookupIPSetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPSetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPSetResultOutput{})
}