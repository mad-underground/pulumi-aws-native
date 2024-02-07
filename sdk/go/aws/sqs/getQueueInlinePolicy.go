// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Schema for SQS QueueInlinePolicy
func LookupQueueInlinePolicy(ctx *pulumi.Context, args *LookupQueueInlinePolicyArgs, opts ...pulumi.InvokeOption) (*LookupQueueInlinePolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQueueInlinePolicyResult
	err := ctx.Invoke("aws-native:sqs:getQueueInlinePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueInlinePolicyArgs struct {
	// The URL of the SQS queue.
	Queue string `pulumi:"queue"`
}

type LookupQueueInlinePolicyResult struct {
	// A policy document that contains permissions to add to the specified SQS queue
	PolicyDocument interface{} `pulumi:"policyDocument"`
}

func LookupQueueInlinePolicyOutput(ctx *pulumi.Context, args LookupQueueInlinePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupQueueInlinePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueInlinePolicyResult, error) {
			args := v.(LookupQueueInlinePolicyArgs)
			r, err := LookupQueueInlinePolicy(ctx, &args, opts...)
			var s LookupQueueInlinePolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueInlinePolicyResultOutput)
}

type LookupQueueInlinePolicyOutputArgs struct {
	// The URL of the SQS queue.
	Queue pulumi.StringInput `pulumi:"queue"`
}

func (LookupQueueInlinePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueInlinePolicyArgs)(nil)).Elem()
}

type LookupQueueInlinePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupQueueInlinePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueInlinePolicyResult)(nil)).Elem()
}

func (o LookupQueueInlinePolicyResultOutput) ToLookupQueueInlinePolicyResultOutput() LookupQueueInlinePolicyResultOutput {
	return o
}

func (o LookupQueueInlinePolicyResultOutput) ToLookupQueueInlinePolicyResultOutputWithContext(ctx context.Context) LookupQueueInlinePolicyResultOutput {
	return o
}

// A policy document that contains permissions to add to the specified SQS queue
func (o LookupQueueInlinePolicyResultOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupQueueInlinePolicyResult) interface{} { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueInlinePolicyResultOutput{})
}
