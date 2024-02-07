// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::TrafficMirrorSession
func LookupTrafficMirrorSession(ctx *pulumi.Context, args *LookupTrafficMirrorSessionArgs, opts ...pulumi.InvokeOption) (*LookupTrafficMirrorSessionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTrafficMirrorSessionResult
	err := ctx.Invoke("aws-native:ec2:getTrafficMirrorSession", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrafficMirrorSessionArgs struct {
	Id string `pulumi:"id"`
}

type LookupTrafficMirrorSessionResult struct {
	Description           *string                   `pulumi:"description"`
	Id                    *string                   `pulumi:"id"`
	PacketLength          *int                      `pulumi:"packetLength"`
	SessionNumber         *int                      `pulumi:"sessionNumber"`
	Tags                  []TrafficMirrorSessionTag `pulumi:"tags"`
	TrafficMirrorFilterId *string                   `pulumi:"trafficMirrorFilterId"`
	TrafficMirrorTargetId *string                   `pulumi:"trafficMirrorTargetId"`
	VirtualNetworkId      *int                      `pulumi:"virtualNetworkId"`
}

func LookupTrafficMirrorSessionOutput(ctx *pulumi.Context, args LookupTrafficMirrorSessionOutputArgs, opts ...pulumi.InvokeOption) LookupTrafficMirrorSessionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrafficMirrorSessionResult, error) {
			args := v.(LookupTrafficMirrorSessionArgs)
			r, err := LookupTrafficMirrorSession(ctx, &args, opts...)
			var s LookupTrafficMirrorSessionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrafficMirrorSessionResultOutput)
}

type LookupTrafficMirrorSessionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupTrafficMirrorSessionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficMirrorSessionArgs)(nil)).Elem()
}

type LookupTrafficMirrorSessionResultOutput struct{ *pulumi.OutputState }

func (LookupTrafficMirrorSessionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrafficMirrorSessionResult)(nil)).Elem()
}

func (o LookupTrafficMirrorSessionResultOutput) ToLookupTrafficMirrorSessionResultOutput() LookupTrafficMirrorSessionResultOutput {
	return o
}

func (o LookupTrafficMirrorSessionResultOutput) ToLookupTrafficMirrorSessionResultOutputWithContext(ctx context.Context) LookupTrafficMirrorSessionResultOutput {
	return o
}

func (o LookupTrafficMirrorSessionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTrafficMirrorSessionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTrafficMirrorSessionResultOutput) PacketLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) *int { return v.PacketLength }).(pulumi.IntPtrOutput)
}

func (o LookupTrafficMirrorSessionResultOutput) SessionNumber() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) *int { return v.SessionNumber }).(pulumi.IntPtrOutput)
}

func (o LookupTrafficMirrorSessionResultOutput) Tags() TrafficMirrorSessionTagArrayOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) []TrafficMirrorSessionTag { return v.Tags }).(TrafficMirrorSessionTagArrayOutput)
}

func (o LookupTrafficMirrorSessionResultOutput) TrafficMirrorFilterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) *string { return v.TrafficMirrorFilterId }).(pulumi.StringPtrOutput)
}

func (o LookupTrafficMirrorSessionResultOutput) TrafficMirrorTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) *string { return v.TrafficMirrorTargetId }).(pulumi.StringPtrOutput)
}

func (o LookupTrafficMirrorSessionResultOutput) VirtualNetworkId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupTrafficMirrorSessionResult) *int { return v.VirtualNetworkId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrafficMirrorSessionResultOutput{})
}
