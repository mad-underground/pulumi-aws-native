// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::TrafficMirrorTarget
func LookupTrafficMirrorTarget(ctx *pulumi.Context, args *LookupTrafficMirrorTargetArgs, opts ...pulumi.InvokeOption) (*LookupTrafficMirrorTargetResult, error) {
	var rv LookupTrafficMirrorTargetResult
	err := ctx.Invoke("aws-native:ec2:getTrafficMirrorTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficMirrorTargetArgs struct {
	Id string `pulumi:"id"`
}

type LookupTrafficMirrorTargetResult struct {
	Id   *string                  `pulumi:"id"`
	Tags []TrafficMirrorTargetTag `pulumi:"tags"`
}

func LookupTrafficMirrorTargetOutput(ctx *pulumi.Context, args LookupTrafficMirrorTargetOutputArgs, opts ...pulumi.InvokeOption) LookupTrafficMirrorTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrafficMirrorTargetResult, error) {
			args := v.(LookupTrafficMirrorTargetArgs)
			r, err := LookupTrafficMirrorTarget(ctx, &args, opts...)
			return *r, err
		}).(LookupTrafficMirrorTargetResultOutput)
}

type LookupTrafficMirrorTargetOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTrafficMirrorTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficMirrorTargetArgs)(nil)).Elem()
}

type LookupTrafficMirrorTargetResultOutput struct{ *pulumi.OutputState }

func (LookupTrafficMirrorTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficMirrorTargetResult)(nil)).Elem()
}

func (o LookupTrafficMirrorTargetResultOutput) ToLookupTrafficMirrorTargetResultOutput() LookupTrafficMirrorTargetResultOutput {
	return o
}

func (o LookupTrafficMirrorTargetResultOutput) ToLookupTrafficMirrorTargetResultOutputWithContext(ctx context.Context) LookupTrafficMirrorTargetResultOutput {
	return o
}

func (o LookupTrafficMirrorTargetResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorTargetResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTrafficMirrorTargetResultOutput) Tags() TrafficMirrorTargetTagArrayOutput {
	return o.ApplyT(func(v LookupTrafficMirrorTargetResult) []TrafficMirrorTargetTag { return v.Tags }).(TrafficMirrorTargetTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrafficMirrorTargetResultOutput{})
}