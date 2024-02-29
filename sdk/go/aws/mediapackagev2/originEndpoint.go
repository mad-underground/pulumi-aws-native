// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackagev2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// <p>Represents an origin endpoint that is associated with a channel, offering a dynamically repackaged version of its content through various streaming media protocols. The content can be efficiently disseminated to end-users via a Content Delivery Network (CDN), like Amazon CloudFront.</p>
type OriginEndpoint struct {
	pulumi.CustomResourceState

	// <p>The Amazon Resource Name (ARN) associated with the resource.</p>
	Arn              pulumi.StringOutput                  `pulumi:"arn"`
	ChannelGroupName pulumi.StringOutput                  `pulumi:"channelGroupName"`
	ChannelName      pulumi.StringOutput                  `pulumi:"channelName"`
	ContainerType    OriginEndpointContainerTypePtrOutput `pulumi:"containerType"`
	// <p>The date and time the origin endpoint was created.</p>
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// <p>An HTTP live streaming (HLS) manifest configuration.</p>
	HlsManifests OriginEndpointHlsManifestConfigurationArrayOutput `pulumi:"hlsManifests"`
	// <p>A low-latency HLS manifest configuration.</p>
	LowLatencyHlsManifests OriginEndpointLowLatencyHlsManifestConfigurationArrayOutput `pulumi:"lowLatencyHlsManifests"`
	// <p>The date and time the origin endpoint was modified.</p>
	ModifiedAt         pulumi.StringOutput            `pulumi:"modifiedAt"`
	OriginEndpointName pulumi.StringOutput            `pulumi:"originEndpointName"`
	Segment            OriginEndpointSegmentPtrOutput `pulumi:"segment"`
	// <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
	StartoverWindowSeconds pulumi.IntPtrOutput `pulumi:"startoverWindowSeconds"`
	Tags                   aws.TagArrayOutput  `pulumi:"tags"`
}

// NewOriginEndpoint registers a new resource with the given unique name, arguments, and options.
func NewOriginEndpoint(ctx *pulumi.Context,
	name string, args *OriginEndpointArgs, opts ...pulumi.ResourceOption) (*OriginEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChannelGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ChannelGroupName'")
	}
	if args.ChannelName == nil {
		return nil, errors.New("invalid value for required argument 'ChannelName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"channelGroupName",
		"channelName",
		"originEndpointName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OriginEndpoint
	err := ctx.RegisterResource("aws-native:mediapackagev2:OriginEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOriginEndpoint gets an existing OriginEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginEndpointState, opts ...pulumi.ResourceOption) (*OriginEndpoint, error) {
	var resource OriginEndpoint
	err := ctx.ReadResource("aws-native:mediapackagev2:OriginEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OriginEndpoint resources.
type originEndpointState struct {
}

type OriginEndpointState struct {
}

func (OriginEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*originEndpointState)(nil)).Elem()
}

type originEndpointArgs struct {
	ChannelGroupName string                       `pulumi:"channelGroupName"`
	ChannelName      string                       `pulumi:"channelName"`
	ContainerType    *OriginEndpointContainerType `pulumi:"containerType"`
	// <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
	Description *string `pulumi:"description"`
	// <p>An HTTP live streaming (HLS) manifest configuration.</p>
	HlsManifests []OriginEndpointHlsManifestConfiguration `pulumi:"hlsManifests"`
	// <p>A low-latency HLS manifest configuration.</p>
	LowLatencyHlsManifests []OriginEndpointLowLatencyHlsManifestConfiguration `pulumi:"lowLatencyHlsManifests"`
	OriginEndpointName     *string                                            `pulumi:"originEndpointName"`
	Segment                *OriginEndpointSegment                             `pulumi:"segment"`
	// <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
	StartoverWindowSeconds *int      `pulumi:"startoverWindowSeconds"`
	Tags                   []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a OriginEndpoint resource.
type OriginEndpointArgs struct {
	ChannelGroupName pulumi.StringInput
	ChannelName      pulumi.StringInput
	ContainerType    OriginEndpointContainerTypePtrInput
	// <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
	Description pulumi.StringPtrInput
	// <p>An HTTP live streaming (HLS) manifest configuration.</p>
	HlsManifests OriginEndpointHlsManifestConfigurationArrayInput
	// <p>A low-latency HLS manifest configuration.</p>
	LowLatencyHlsManifests OriginEndpointLowLatencyHlsManifestConfigurationArrayInput
	OriginEndpointName     pulumi.StringPtrInput
	Segment                OriginEndpointSegmentPtrInput
	// <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
	StartoverWindowSeconds pulumi.IntPtrInput
	Tags                   aws.TagArrayInput
}

func (OriginEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originEndpointArgs)(nil)).Elem()
}

type OriginEndpointInput interface {
	pulumi.Input

	ToOriginEndpointOutput() OriginEndpointOutput
	ToOriginEndpointOutputWithContext(ctx context.Context) OriginEndpointOutput
}

func (*OriginEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginEndpoint)(nil)).Elem()
}

func (i *OriginEndpoint) ToOriginEndpointOutput() OriginEndpointOutput {
	return i.ToOriginEndpointOutputWithContext(context.Background())
}

func (i *OriginEndpoint) ToOriginEndpointOutputWithContext(ctx context.Context) OriginEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginEndpointOutput)
}

type OriginEndpointOutput struct{ *pulumi.OutputState }

func (OriginEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OriginEndpoint)(nil)).Elem()
}

func (o OriginEndpointOutput) ToOriginEndpointOutput() OriginEndpointOutput {
	return o
}

func (o OriginEndpointOutput) ToOriginEndpointOutputWithContext(ctx context.Context) OriginEndpointOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) associated with the resource.</p>
func (o OriginEndpointOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o OriginEndpointOutput) ChannelGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.StringOutput { return v.ChannelGroupName }).(pulumi.StringOutput)
}

func (o OriginEndpointOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.StringOutput { return v.ChannelName }).(pulumi.StringOutput)
}

func (o OriginEndpointOutput) ContainerType() OriginEndpointContainerTypePtrOutput {
	return o.ApplyT(func(v *OriginEndpoint) OriginEndpointContainerTypePtrOutput { return v.ContainerType }).(OriginEndpointContainerTypePtrOutput)
}

// <p>The date and time the origin endpoint was created.</p>
func (o OriginEndpointOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// <p>Enter any descriptive text that helps you to identify the origin endpoint.</p>
func (o OriginEndpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// <p>An HTTP live streaming (HLS) manifest configuration.</p>
func (o OriginEndpointOutput) HlsManifests() OriginEndpointHlsManifestConfigurationArrayOutput {
	return o.ApplyT(func(v *OriginEndpoint) OriginEndpointHlsManifestConfigurationArrayOutput { return v.HlsManifests }).(OriginEndpointHlsManifestConfigurationArrayOutput)
}

// <p>A low-latency HLS manifest configuration.</p>
func (o OriginEndpointOutput) LowLatencyHlsManifests() OriginEndpointLowLatencyHlsManifestConfigurationArrayOutput {
	return o.ApplyT(func(v *OriginEndpoint) OriginEndpointLowLatencyHlsManifestConfigurationArrayOutput {
		return v.LowLatencyHlsManifests
	}).(OriginEndpointLowLatencyHlsManifestConfigurationArrayOutput)
}

// <p>The date and time the origin endpoint was modified.</p>
func (o OriginEndpointOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.StringOutput { return v.ModifiedAt }).(pulumi.StringOutput)
}

func (o OriginEndpointOutput) OriginEndpointName() pulumi.StringOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.StringOutput { return v.OriginEndpointName }).(pulumi.StringOutput)
}

func (o OriginEndpointOutput) Segment() OriginEndpointSegmentPtrOutput {
	return o.ApplyT(func(v *OriginEndpoint) OriginEndpointSegmentPtrOutput { return v.Segment }).(OriginEndpointSegmentPtrOutput)
}

// <p>The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing. Viewers can start-over or catch-up on content that falls within the window. The maximum startover window is 1,209,600 seconds (14 days).</p>
func (o OriginEndpointOutput) StartoverWindowSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OriginEndpoint) pulumi.IntPtrOutput { return v.StartoverWindowSeconds }).(pulumi.IntPtrOutput)
}

func (o OriginEndpointOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *OriginEndpoint) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OriginEndpointInput)(nil)).Elem(), &OriginEndpoint{})
	pulumi.RegisterOutputType(OriginEndpointOutput{})
}
