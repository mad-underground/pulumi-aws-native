// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFormation::WaitConditionHandle
func LookupWaitConditionHandle(ctx *pulumi.Context, args *LookupWaitConditionHandleArgs, opts ...pulumi.InvokeOption) (*LookupWaitConditionHandleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWaitConditionHandleResult
	err := ctx.Invoke("aws-native:cloudformation:getWaitConditionHandle", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWaitConditionHandleArgs struct {
	Id string `pulumi:"id"`
}

type LookupWaitConditionHandleResult struct {
	Id *string `pulumi:"id"`
}

func LookupWaitConditionHandleOutput(ctx *pulumi.Context, args LookupWaitConditionHandleOutputArgs, opts ...pulumi.InvokeOption) LookupWaitConditionHandleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWaitConditionHandleResult, error) {
			args := v.(LookupWaitConditionHandleArgs)
			r, err := LookupWaitConditionHandle(ctx, &args, opts...)
			var s LookupWaitConditionHandleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWaitConditionHandleResultOutput)
}

type LookupWaitConditionHandleOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWaitConditionHandleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWaitConditionHandleArgs)(nil)).Elem()
}

type LookupWaitConditionHandleResultOutput struct{ *pulumi.OutputState }

func (LookupWaitConditionHandleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWaitConditionHandleResult)(nil)).Elem()
}

func (o LookupWaitConditionHandleResultOutput) ToLookupWaitConditionHandleResultOutput() LookupWaitConditionHandleResultOutput {
	return o
}

func (o LookupWaitConditionHandleResultOutput) ToLookupWaitConditionHandleResultOutputWithContext(ctx context.Context) LookupWaitConditionHandleResultOutput {
	return o
}

func (o LookupWaitConditionHandleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWaitConditionHandleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWaitConditionHandleResultOutput{})
}
