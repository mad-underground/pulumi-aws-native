// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediatailor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Definition of AWS::MediaTailor::Channel Resource Type
func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupChannelResult
	err := ctx.Invoke("aws-native:mediatailor:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	ChannelName string `pulumi:"channelName"`
}

type LookupChannelResult struct {
	// <p>The ARN of the channel.</p>
	Arn              *string                            `pulumi:"arn"`
	FillerSlate      *ChannelSlateSource                `pulumi:"fillerSlate"`
	LogConfiguration *ChannelLogConfigurationForChannel `pulumi:"logConfiguration"`
	PlaybackMode     *ChannelPlaybackMode               `pulumi:"playbackMode"`
	// The tags to assign to the channel.
	Tags []ChannelTag `pulumi:"tags"`
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelResult, error) {
			args := v.(LookupChannelArgs)
			r, err := LookupChannel(ctx, &args, opts...)
			var s LookupChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	ChannelName pulumi.StringInput `pulumi:"channelName"`
}

func (LookupChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelArgs)(nil)).Elem()
}

type LookupChannelResultOutput struct{ *pulumi.OutputState }

func (LookupChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelResult)(nil)).Elem()
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutput() LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutputWithContext(ctx context.Context) LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupChannelResult] {
	return pulumix.Output[LookupChannelResult]{
		OutputState: o.OutputState,
	}
}

// <p>The ARN of the channel.</p>
func (o LookupChannelResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) FillerSlate() ChannelSlateSourcePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelSlateSource { return v.FillerSlate }).(ChannelSlateSourcePtrOutput)
}

func (o LookupChannelResultOutput) LogConfiguration() ChannelLogConfigurationForChannelPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelLogConfigurationForChannel { return v.LogConfiguration }).(ChannelLogConfigurationForChannelPtrOutput)
}

func (o LookupChannelResultOutput) PlaybackMode() ChannelPlaybackModePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelPlaybackMode { return v.PlaybackMode }).(ChannelPlaybackModePtrOutput)
}

// The tags to assign to the channel.
func (o LookupChannelResultOutput) Tags() ChannelTagArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []ChannelTag { return v.Tags }).(ChannelTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}