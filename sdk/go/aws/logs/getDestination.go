// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and enables you to subscribe that resource to a stream of log events.
func LookupDestination(ctx *pulumi.Context, args *LookupDestinationArgs, opts ...pulumi.InvokeOption) (*LookupDestinationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDestinationResult
	err := ctx.Invoke("aws-native:logs:getDestination", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDestinationArgs struct {
	// The name of the destination resource
	DestinationName string `pulumi:"destinationName"`
}

type LookupDestinationResult struct {
	Arn *string `pulumi:"arn"`
	// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
	DestinationPolicy *string `pulumi:"destinationPolicy"`
	// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource
	RoleArn *string `pulumi:"roleArn"`
	// The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)
	TargetArn *string `pulumi:"targetArn"`
}

func LookupDestinationOutput(ctx *pulumi.Context, args LookupDestinationOutputArgs, opts ...pulumi.InvokeOption) LookupDestinationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDestinationResult, error) {
			args := v.(LookupDestinationArgs)
			r, err := LookupDestination(ctx, &args, opts...)
			var s LookupDestinationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDestinationResultOutput)
}

type LookupDestinationOutputArgs struct {
	// The name of the destination resource
	DestinationName pulumi.StringInput `pulumi:"destinationName"`
}

func (LookupDestinationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationArgs)(nil)).Elem()
}

type LookupDestinationResultOutput struct{ *pulumi.OutputState }

func (LookupDestinationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDestinationResult)(nil)).Elem()
}

func (o LookupDestinationResultOutput) ToLookupDestinationResultOutput() LookupDestinationResultOutput {
	return o
}

func (o LookupDestinationResultOutput) ToLookupDestinationResultOutputWithContext(ctx context.Context) LookupDestinationResultOutput {
	return o
}

func (o LookupDestinationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// An IAM policy document that governs which AWS accounts can create subscription filters against this destination.
func (o LookupDestinationResultOutput) DestinationPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.DestinationPolicy }).(pulumi.StringPtrOutput)
}

// The ARN of an IAM role that permits CloudWatch Logs to send data to the specified AWS resource
func (o LookupDestinationResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// The ARN of the physical target where the log events are delivered (for example, a Kinesis stream)
func (o LookupDestinationResultOutput) TargetArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDestinationResult) *string { return v.TargetArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDestinationResultOutput{})
}
