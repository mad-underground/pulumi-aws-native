// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediaconnect

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaConnect::FlowSource
func LookupFlowSource(ctx *pulumi.Context, args *LookupFlowSourceArgs, opts ...pulumi.InvokeOption) (*LookupFlowSourceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFlowSourceResult
	err := ctx.Invoke("aws-native:mediaconnect:getFlowSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFlowSourceArgs struct {
	// The ARN of the source.
	SourceArn string `pulumi:"sourceArn"`
}

type LookupFlowSourceResult struct {
	// The type of encryption that is used on the content ingested from this source.
	Decryption *FlowSourceEncryption `pulumi:"decryption"`
	// A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
	Description *string `pulumi:"description"`
	// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
	EntitlementArn *string `pulumi:"entitlementArn"`
	// The ARN of the flow.
	FlowArn *string `pulumi:"flowArn"`
	// The source configuration for cloud flows receiving a stream from a bridge.
	GatewayBridgeSource *FlowSourceGatewayBridgeSource `pulumi:"gatewayBridgeSource"`
	// The IP address that the flow will be listening on for incoming content.
	IngestIp *string `pulumi:"ingestIp"`
	// The port that the flow will be listening on for incoming content.
	IngestPort *int `pulumi:"ingestPort"`
	// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
	MaxBitrate *int `pulumi:"maxBitrate"`
	// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
	MaxLatency *int `pulumi:"maxLatency"`
	// The minimum latency in milliseconds.
	MinLatency *int `pulumi:"minLatency"`
	// The protocol that is used by the source.
	Protocol *FlowSourceProtocol `pulumi:"protocol"`
	// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
	SenderControlPort *int `pulumi:"senderControlPort"`
	// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
	SenderIpAddress *string `pulumi:"senderIpAddress"`
	// The ARN of the source.
	SourceArn *string `pulumi:"sourceArn"`
	// The port that the flow will be listening on for incoming content.(ReadOnly)
	SourceIngestPort *string `pulumi:"sourceIngestPort"`
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

func LookupFlowSourceOutput(ctx *pulumi.Context, args LookupFlowSourceOutputArgs, opts ...pulumi.InvokeOption) LookupFlowSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFlowSourceResult, error) {
			args := v.(LookupFlowSourceArgs)
			r, err := LookupFlowSource(ctx, &args, opts...)
			var s LookupFlowSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFlowSourceResultOutput)
}

type LookupFlowSourceOutputArgs struct {
	// The ARN of the source.
	SourceArn pulumi.StringInput `pulumi:"sourceArn"`
}

func (LookupFlowSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowSourceArgs)(nil)).Elem()
}

type LookupFlowSourceResultOutput struct{ *pulumi.OutputState }

func (LookupFlowSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFlowSourceResult)(nil)).Elem()
}

func (o LookupFlowSourceResultOutput) ToLookupFlowSourceResultOutput() LookupFlowSourceResultOutput {
	return o
}

func (o LookupFlowSourceResultOutput) ToLookupFlowSourceResultOutputWithContext(ctx context.Context) LookupFlowSourceResultOutput {
	return o
}

// The type of encryption that is used on the content ingested from this source.
func (o LookupFlowSourceResultOutput) Decryption() FlowSourceEncryptionPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *FlowSourceEncryption { return v.Decryption }).(FlowSourceEncryptionPtrOutput)
}

// A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.
func (o LookupFlowSourceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.
func (o LookupFlowSourceResultOutput) EntitlementArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.EntitlementArn }).(pulumi.StringPtrOutput)
}

// The ARN of the flow.
func (o LookupFlowSourceResultOutput) FlowArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.FlowArn }).(pulumi.StringPtrOutput)
}

// The source configuration for cloud flows receiving a stream from a bridge.
func (o LookupFlowSourceResultOutput) GatewayBridgeSource() FlowSourceGatewayBridgeSourcePtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *FlowSourceGatewayBridgeSource { return v.GatewayBridgeSource }).(FlowSourceGatewayBridgeSourcePtrOutput)
}

// The IP address that the flow will be listening on for incoming content.
func (o LookupFlowSourceResultOutput) IngestIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.IngestIp }).(pulumi.StringPtrOutput)
}

// The port that the flow will be listening on for incoming content.
func (o LookupFlowSourceResultOutput) IngestPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *int { return v.IngestPort }).(pulumi.IntPtrOutput)
}

// The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.
func (o LookupFlowSourceResultOutput) MaxBitrate() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *int { return v.MaxBitrate }).(pulumi.IntPtrOutput)
}

// The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.
func (o LookupFlowSourceResultOutput) MaxLatency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *int { return v.MaxLatency }).(pulumi.IntPtrOutput)
}

// The minimum latency in milliseconds.
func (o LookupFlowSourceResultOutput) MinLatency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *int { return v.MinLatency }).(pulumi.IntPtrOutput)
}

// The protocol that is used by the source.
func (o LookupFlowSourceResultOutput) Protocol() FlowSourceProtocolPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *FlowSourceProtocol { return v.Protocol }).(FlowSourceProtocolPtrOutput)
}

// The port that the flow uses to send outbound requests to initiate connection with the sender for fujitsu-qos protocol.
func (o LookupFlowSourceResultOutput) SenderControlPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *int { return v.SenderControlPort }).(pulumi.IntPtrOutput)
}

// The IP address that the flow communicates with to initiate connection with the sender for fujitsu-qos protocol.
func (o LookupFlowSourceResultOutput) SenderIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.SenderIpAddress }).(pulumi.StringPtrOutput)
}

// The ARN of the source.
func (o LookupFlowSourceResultOutput) SourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.SourceArn }).(pulumi.StringPtrOutput)
}

// The port that the flow will be listening on for incoming content.(ReadOnly)
func (o LookupFlowSourceResultOutput) SourceIngestPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.SourceIngestPort }).(pulumi.StringPtrOutput)
}

// Source IP or domain name for SRT-caller protocol.
func (o LookupFlowSourceResultOutput) SourceListenerAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.SourceListenerAddress }).(pulumi.StringPtrOutput)
}

// Source port for SRT-caller protocol.
func (o LookupFlowSourceResultOutput) SourceListenerPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *int { return v.SourceListenerPort }).(pulumi.IntPtrOutput)
}

// The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.
func (o LookupFlowSourceResultOutput) StreamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.StreamId }).(pulumi.StringPtrOutput)
}

// The name of the VPC Interface this Source is configured with.
func (o LookupFlowSourceResultOutput) VpcInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.VpcInterfaceName }).(pulumi.StringPtrOutput)
}

// The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.
func (o LookupFlowSourceResultOutput) WhitelistCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFlowSourceResult) *string { return v.WhitelistCidr }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFlowSourceResultOutput{})
}
