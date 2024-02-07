// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconnect

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaConnect::FlowSource
type FlowSource struct {
	pulumi.CustomResourceState

	// The type of encryption that is used on the content ingested from this source.
	Decryption FlowSourceEncryptionPtrOutput `pulumi:"decryption"`
	// A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
	EntitlementArn pulumi.StringPtrOutput `pulumi:"entitlementArn"`
	// The ARN of the flow.
	FlowArn pulumi.StringPtrOutput `pulumi:"flowArn"`
	// The source configuration for cloud flows receiving a stream from a bridge.
	GatewayBridgeSource FlowSourceGatewayBridgeSourcePtrOutput `pulumi:"gatewayBridgeSource"`
	// The IP address that the flow will be listening on for incoming content.
	IngestIp pulumi.StringOutput `pulumi:"ingestIp"`
	// The port that the flow will be listening on for incoming content.
	IngestPort pulumi.IntPtrOutput `pulumi:"ingestPort"`
	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate pulumi.IntPtrOutput `pulumi:"maxBitrate"`
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	MaxLatency pulumi.IntPtrOutput `pulumi:"maxLatency"`
	// The minimum latency in milliseconds.
	MinLatency pulumi.IntPtrOutput `pulumi:"minLatency"`
	// The name of the source.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protocol that is used by the source.
	Protocol FlowSourceProtocolPtrOutput `pulumi:"protocol"`
	// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
	SenderControlPort pulumi.IntPtrOutput `pulumi:"senderControlPort"`
	// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
	SenderIpAddress pulumi.StringPtrOutput `pulumi:"senderIpAddress"`
	// The ARN of the source.
	SourceArn pulumi.StringOutput `pulumi:"sourceArn"`
	// The port that the flow will be listening on for incoming content.(ReadOnly)
	SourceIngestPort pulumi.StringOutput `pulumi:"sourceIngestPort"`
	// Source IP or domain name for SRT-caller protocol.
	SourceListenerAddress pulumi.StringPtrOutput `pulumi:"sourceListenerAddress"`
	// Source port for SRT-caller protocol.
	SourceListenerPort pulumi.IntPtrOutput `pulumi:"sourceListenerPort"`
	// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
	StreamId pulumi.StringPtrOutput `pulumi:"streamId"`
	// The name of the VPC Interface this Source is configured with.
	VpcInterfaceName pulumi.StringPtrOutput `pulumi:"vpcInterfaceName"`
	// The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr pulumi.StringPtrOutput `pulumi:"whitelistCidr"`
}

// NewFlowSource registers a new resource with the given unique name, arguments, and options.
func NewFlowSource(ctx *pulumi.Context,
	name string, args *FlowSourceArgs, opts ...pulumi.ResourceOption) (*FlowSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FlowSource
	err := ctx.RegisterResource("aws-native:mediaconnect:FlowSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlowSource gets an existing FlowSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlowSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowSourceState, opts ...pulumi.ResourceOption) (*FlowSource, error) {
	var resource FlowSource
	err := ctx.ReadResource("aws-native:mediaconnect:FlowSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlowSource resources.
type flowSourceState struct {
}

type FlowSourceState struct {
}

func (FlowSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSourceState)(nil)).Elem()
}

type flowSourceArgs struct {
	// The type of encryption that is used on the content ingested from this source.
	Decryption *FlowSourceEncryption `pulumi:"decryption"`
	// A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
	Description string `pulumi:"description"`
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
	EntitlementArn *string `pulumi:"entitlementArn"`
	// The ARN of the flow.
	FlowArn *string `pulumi:"flowArn"`
	// The source configuration for cloud flows receiving a stream from a bridge.
	GatewayBridgeSource *FlowSourceGatewayBridgeSource `pulumi:"gatewayBridgeSource"`
	// The port that the flow will be listening on for incoming content.
	IngestPort *int `pulumi:"ingestPort"`
	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *int `pulumi:"maxBitrate"`
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	MaxLatency *int `pulumi:"maxLatency"`
	// The minimum latency in milliseconds.
	MinLatency *int `pulumi:"minLatency"`
	// The name of the source.
	Name *string `pulumi:"name"`
	// The protocol that is used by the source.
	Protocol *FlowSourceProtocol `pulumi:"protocol"`
	// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
	SenderControlPort *int `pulumi:"senderControlPort"`
	// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
	SenderIpAddress *string `pulumi:"senderIpAddress"`
	// Source IP or domain name for SRT-caller protocol.
	SourceListenerAddress *string `pulumi:"sourceListenerAddress"`
	// Source port for SRT-caller protocol.
	SourceListenerPort *int `pulumi:"sourceListenerPort"`
	// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
	StreamId *string `pulumi:"streamId"`
	// The name of the VPC Interface this Source is configured with.
	VpcInterfaceName *string `pulumi:"vpcInterfaceName"`
	// The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr *string `pulumi:"whitelistCidr"`
}

// The set of arguments for constructing a FlowSource resource.
type FlowSourceArgs struct {
	// The type of encryption that is used on the content ingested from this source.
	Decryption FlowSourceEncryptionPtrInput
	// A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
	Description pulumi.StringInput
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
	EntitlementArn pulumi.StringPtrInput
	// The ARN of the flow.
	FlowArn pulumi.StringPtrInput
	// The source configuration for cloud flows receiving a stream from a bridge.
	GatewayBridgeSource FlowSourceGatewayBridgeSourcePtrInput
	// The port that the flow will be listening on for incoming content.
	IngestPort pulumi.IntPtrInput
	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate pulumi.IntPtrInput
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	MaxLatency pulumi.IntPtrInput
	// The minimum latency in milliseconds.
	MinLatency pulumi.IntPtrInput
	// The name of the source.
	Name pulumi.StringPtrInput
	// The protocol that is used by the source.
	Protocol FlowSourceProtocolPtrInput
	// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
	SenderControlPort pulumi.IntPtrInput
	// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
	SenderIpAddress pulumi.StringPtrInput
	// Source IP or domain name for SRT-caller protocol.
	SourceListenerAddress pulumi.StringPtrInput
	// Source port for SRT-caller protocol.
	SourceListenerPort pulumi.IntPtrInput
	// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
	StreamId pulumi.StringPtrInput
	// The name of the VPC Interface this Source is configured with.
	VpcInterfaceName pulumi.StringPtrInput
	// The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
	WhitelistCidr pulumi.StringPtrInput
}

func (FlowSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowSourceArgs)(nil)).Elem()
}

type FlowSourceInput interface {
	pulumi.Input

	ToFlowSourceOutput() FlowSourceOutput
	ToFlowSourceOutputWithContext(ctx context.Context) FlowSourceOutput
}

func (*FlowSource) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSource)(nil)).Elem()
}

func (i *FlowSource) ToFlowSourceOutput() FlowSourceOutput {
	return i.ToFlowSourceOutputWithContext(context.Background())
}

func (i *FlowSource) ToFlowSourceOutputWithContext(ctx context.Context) FlowSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlowSourceOutput)
}

type FlowSourceOutput struct{ *pulumi.OutputState }

func (FlowSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlowSource)(nil)).Elem()
}

func (o FlowSourceOutput) ToFlowSourceOutput() FlowSourceOutput {
	return o
}

func (o FlowSourceOutput) ToFlowSourceOutputWithContext(ctx context.Context) FlowSourceOutput {
	return o
}

// The type of encryption that is used on the content ingested from this source.
func (o FlowSourceOutput) Decryption() FlowSourceEncryptionPtrOutput {
	return o.ApplyT(func(v *FlowSource) FlowSourceEncryptionPtrOutput { return v.Decryption }).(FlowSourceEncryptionPtrOutput)
}

// A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
func (o FlowSourceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
func (o FlowSourceOutput) EntitlementArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringPtrOutput { return v.EntitlementArn }).(pulumi.StringPtrOutput)
}

// The ARN of the flow.
func (o FlowSourceOutput) FlowArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringPtrOutput { return v.FlowArn }).(pulumi.StringPtrOutput)
}

// The source configuration for cloud flows receiving a stream from a bridge.
func (o FlowSourceOutput) GatewayBridgeSource() FlowSourceGatewayBridgeSourcePtrOutput {
	return o.ApplyT(func(v *FlowSource) FlowSourceGatewayBridgeSourcePtrOutput { return v.GatewayBridgeSource }).(FlowSourceGatewayBridgeSourcePtrOutput)
}

// The IP address that the flow will be listening on for incoming content.
func (o FlowSourceOutput) IngestIp() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringOutput { return v.IngestIp }).(pulumi.StringOutput)
}

// The port that the flow will be listening on for incoming content.
func (o FlowSourceOutput) IngestPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.IntPtrOutput { return v.IngestPort }).(pulumi.IntPtrOutput)
}

// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
func (o FlowSourceOutput) MaxBitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.IntPtrOutput { return v.MaxBitrate }).(pulumi.IntPtrOutput)
}

// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
func (o FlowSourceOutput) MaxLatency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.IntPtrOutput { return v.MaxLatency }).(pulumi.IntPtrOutput)
}

// The minimum latency in milliseconds.
func (o FlowSourceOutput) MinLatency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.IntPtrOutput { return v.MinLatency }).(pulumi.IntPtrOutput)
}

// The name of the source.
func (o FlowSourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protocol that is used by the source.
func (o FlowSourceOutput) Protocol() FlowSourceProtocolPtrOutput {
	return o.ApplyT(func(v *FlowSource) FlowSourceProtocolPtrOutput { return v.Protocol }).(FlowSourceProtocolPtrOutput)
}

// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
func (o FlowSourceOutput) SenderControlPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.IntPtrOutput { return v.SenderControlPort }).(pulumi.IntPtrOutput)
}

// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
func (o FlowSourceOutput) SenderIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringPtrOutput { return v.SenderIpAddress }).(pulumi.StringPtrOutput)
}

// The ARN of the source.
func (o FlowSourceOutput) SourceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringOutput { return v.SourceArn }).(pulumi.StringOutput)
}

// The port that the flow will be listening on for incoming content.(ReadOnly)
func (o FlowSourceOutput) SourceIngestPort() pulumi.StringOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringOutput { return v.SourceIngestPort }).(pulumi.StringOutput)
}

// Source IP or domain name for SRT-caller protocol.
func (o FlowSourceOutput) SourceListenerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringPtrOutput { return v.SourceListenerAddress }).(pulumi.StringPtrOutput)
}

// Source port for SRT-caller protocol.
func (o FlowSourceOutput) SourceListenerPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.IntPtrOutput { return v.SourceListenerPort }).(pulumi.IntPtrOutput)
}

// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
func (o FlowSourceOutput) StreamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringPtrOutput { return v.StreamId }).(pulumi.StringPtrOutput)
}

// The name of the VPC Interface this Source is configured with.
func (o FlowSourceOutput) VpcInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringPtrOutput { return v.VpcInterfaceName }).(pulumi.StringPtrOutput)
}

// The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
func (o FlowSourceOutput) WhitelistCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FlowSource) pulumi.StringPtrOutput { return v.WhitelistCidr }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FlowSourceInput)(nil)).Elem(), &FlowSource{})
	pulumi.RegisterOutputType(FlowSourceOutput{})
}
