// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package medialive

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaLive::Multiplexprogram
type Multiplexprogram struct {
	pulumi.CustomResourceState

	// The MediaLive channel associated with the program.
	ChannelId pulumi.StringPtrOutput `pulumi:"channelId"`
	// The ID of the multiplex that the program belongs to.
	MultiplexId pulumi.StringPtrOutput `pulumi:"multiplexId"`
	// The settings for this multiplex program.
	MultiplexProgramSettings MultiplexprogramMultiplexProgramSettingsPtrOutput `pulumi:"multiplexProgramSettings"`
	// The packet identifier map for this multiplex program.
	PacketIdentifiersMap MultiplexprogramMultiplexProgramPacketIdentifiersMapPtrOutput `pulumi:"packetIdentifiersMap"`
	// Contains information about the current sources for the specified program in the specified multiplex. Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time.
	PipelineDetails MultiplexprogramMultiplexProgramPipelineDetailArrayOutput `pulumi:"pipelineDetails"`
	// The settings for this multiplex program.
	PreferredChannelPipeline MultiplexprogramPreferredChannelPipelinePtrOutput `pulumi:"preferredChannelPipeline"`
	// The name of the multiplex program.
	ProgramName pulumi.StringPtrOutput `pulumi:"programName"`
}

// NewMultiplexprogram registers a new resource with the given unique name, arguments, and options.
func NewMultiplexprogram(ctx *pulumi.Context,
	name string, args *MultiplexprogramArgs, opts ...pulumi.ResourceOption) (*Multiplexprogram, error) {
	if args == nil {
		args = &MultiplexprogramArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"multiplexId",
		"programName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Multiplexprogram
	err := ctx.RegisterResource("aws-native:medialive:Multiplexprogram", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMultiplexprogram gets an existing Multiplexprogram resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMultiplexprogram(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultiplexprogramState, opts ...pulumi.ResourceOption) (*Multiplexprogram, error) {
	var resource Multiplexprogram
	err := ctx.ReadResource("aws-native:medialive:Multiplexprogram", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Multiplexprogram resources.
type multiplexprogramState struct {
}

type MultiplexprogramState struct {
}

func (MultiplexprogramState) ElementType() reflect.Type {
	return reflect.TypeOf((*multiplexprogramState)(nil)).Elem()
}

type multiplexprogramArgs struct {
	// The MediaLive channel associated with the program.
	ChannelId *string `pulumi:"channelId"`
	// The ID of the multiplex that the program belongs to.
	MultiplexId *string `pulumi:"multiplexId"`
	// The settings for this multiplex program.
	MultiplexProgramSettings *MultiplexprogramMultiplexProgramSettings `pulumi:"multiplexProgramSettings"`
	// The packet identifier map for this multiplex program.
	PacketIdentifiersMap *MultiplexprogramMultiplexProgramPacketIdentifiersMap `pulumi:"packetIdentifiersMap"`
	// Contains information about the current sources for the specified program in the specified multiplex. Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time.
	PipelineDetails []MultiplexprogramMultiplexProgramPipelineDetail `pulumi:"pipelineDetails"`
	// The settings for this multiplex program.
	PreferredChannelPipeline *MultiplexprogramPreferredChannelPipeline `pulumi:"preferredChannelPipeline"`
	// The name of the multiplex program.
	ProgramName *string `pulumi:"programName"`
}

// The set of arguments for constructing a Multiplexprogram resource.
type MultiplexprogramArgs struct {
	// The MediaLive channel associated with the program.
	ChannelId pulumi.StringPtrInput
	// The ID of the multiplex that the program belongs to.
	MultiplexId pulumi.StringPtrInput
	// The settings for this multiplex program.
	MultiplexProgramSettings MultiplexprogramMultiplexProgramSettingsPtrInput
	// The packet identifier map for this multiplex program.
	PacketIdentifiersMap MultiplexprogramMultiplexProgramPacketIdentifiersMapPtrInput
	// Contains information about the current sources for the specified program in the specified multiplex. Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time.
	PipelineDetails MultiplexprogramMultiplexProgramPipelineDetailArrayInput
	// The settings for this multiplex program.
	PreferredChannelPipeline MultiplexprogramPreferredChannelPipelinePtrInput
	// The name of the multiplex program.
	ProgramName pulumi.StringPtrInput
}

func (MultiplexprogramArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multiplexprogramArgs)(nil)).Elem()
}

type MultiplexprogramInput interface {
	pulumi.Input

	ToMultiplexprogramOutput() MultiplexprogramOutput
	ToMultiplexprogramOutputWithContext(ctx context.Context) MultiplexprogramOutput
}

func (*Multiplexprogram) ElementType() reflect.Type {
	return reflect.TypeOf((**Multiplexprogram)(nil)).Elem()
}

func (i *Multiplexprogram) ToMultiplexprogramOutput() MultiplexprogramOutput {
	return i.ToMultiplexprogramOutputWithContext(context.Background())
}

func (i *Multiplexprogram) ToMultiplexprogramOutputWithContext(ctx context.Context) MultiplexprogramOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiplexprogramOutput)
}

type MultiplexprogramOutput struct{ *pulumi.OutputState }

func (MultiplexprogramOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Multiplexprogram)(nil)).Elem()
}

func (o MultiplexprogramOutput) ToMultiplexprogramOutput() MultiplexprogramOutput {
	return o
}

func (o MultiplexprogramOutput) ToMultiplexprogramOutputWithContext(ctx context.Context) MultiplexprogramOutput {
	return o
}

// The MediaLive channel associated with the program.
func (o MultiplexprogramOutput) ChannelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multiplexprogram) pulumi.StringPtrOutput { return v.ChannelId }).(pulumi.StringPtrOutput)
}

// The ID of the multiplex that the program belongs to.
func (o MultiplexprogramOutput) MultiplexId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multiplexprogram) pulumi.StringPtrOutput { return v.MultiplexId }).(pulumi.StringPtrOutput)
}

// The settings for this multiplex program.
func (o MultiplexprogramOutput) MultiplexProgramSettings() MultiplexprogramMultiplexProgramSettingsPtrOutput {
	return o.ApplyT(func(v *Multiplexprogram) MultiplexprogramMultiplexProgramSettingsPtrOutput {
		return v.MultiplexProgramSettings
	}).(MultiplexprogramMultiplexProgramSettingsPtrOutput)
}

// The packet identifier map for this multiplex program.
func (o MultiplexprogramOutput) PacketIdentifiersMap() MultiplexprogramMultiplexProgramPacketIdentifiersMapPtrOutput {
	return o.ApplyT(func(v *Multiplexprogram) MultiplexprogramMultiplexProgramPacketIdentifiersMapPtrOutput {
		return v.PacketIdentifiersMap
	}).(MultiplexprogramMultiplexProgramPacketIdentifiersMapPtrOutput)
}

// Contains information about the current sources for the specified program in the specified multiplex. Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time.
func (o MultiplexprogramOutput) PipelineDetails() MultiplexprogramMultiplexProgramPipelineDetailArrayOutput {
	return o.ApplyT(func(v *Multiplexprogram) MultiplexprogramMultiplexProgramPipelineDetailArrayOutput {
		return v.PipelineDetails
	}).(MultiplexprogramMultiplexProgramPipelineDetailArrayOutput)
}

// The settings for this multiplex program.
func (o MultiplexprogramOutput) PreferredChannelPipeline() MultiplexprogramPreferredChannelPipelinePtrOutput {
	return o.ApplyT(func(v *Multiplexprogram) MultiplexprogramPreferredChannelPipelinePtrOutput {
		return v.PreferredChannelPipeline
	}).(MultiplexprogramPreferredChannelPipelinePtrOutput)
}

// The name of the multiplex program.
func (o MultiplexprogramOutput) ProgramName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Multiplexprogram) pulumi.StringPtrOutput { return v.ProgramName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MultiplexprogramInput)(nil)).Elem(), &Multiplexprogram{})
	pulumi.RegisterOutputType(MultiplexprogramOutput{})
}
