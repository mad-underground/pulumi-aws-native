// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amazonmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AmazonMQ::Broker
func LookupBroker(ctx *pulumi.Context, args *LookupBrokerArgs, opts ...pulumi.InvokeOption) (*LookupBrokerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBrokerResult
	err := ctx.Invoke("aws-native:amazonmq:getBroker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBrokerArgs struct {
	Id string `pulumi:"id"`
}

type LookupBrokerResult struct {
	AmqpEndpoints                   []string                  `pulumi:"amqpEndpoints"`
	Arn                             *string                   `pulumi:"arn"`
	AutoMinorVersionUpgrade         *bool                     `pulumi:"autoMinorVersionUpgrade"`
	Configuration                   *BrokerConfigurationId    `pulumi:"configuration"`
	ConfigurationId                 *string                   `pulumi:"configurationId"`
	ConfigurationRevision           *int                      `pulumi:"configurationRevision"`
	DataReplicationMode             *string                   `pulumi:"dataReplicationMode"`
	DataReplicationPrimaryBrokerArn *string                   `pulumi:"dataReplicationPrimaryBrokerArn"`
	EngineVersion                   *string                   `pulumi:"engineVersion"`
	HostInstanceType                *string                   `pulumi:"hostInstanceType"`
	Id                              *string                   `pulumi:"id"`
	IpAddresses                     []string                  `pulumi:"ipAddresses"`
	LdapServerMetadata              *BrokerLdapServerMetadata `pulumi:"ldapServerMetadata"`
	Logs                            *BrokerLogList            `pulumi:"logs"`
	MaintenanceWindowStartTime      *BrokerMaintenanceWindow  `pulumi:"maintenanceWindowStartTime"`
	MqttEndpoints                   []string                  `pulumi:"mqttEndpoints"`
	OpenWireEndpoints               []string                  `pulumi:"openWireEndpoints"`
	SecurityGroups                  []string                  `pulumi:"securityGroups"`
	StompEndpoints                  []string                  `pulumi:"stompEndpoints"`
	Tags                            []BrokerTagsEntry         `pulumi:"tags"`
	Users                           []BrokerUser              `pulumi:"users"`
	WssEndpoints                    []string                  `pulumi:"wssEndpoints"`
}

func LookupBrokerOutput(ctx *pulumi.Context, args LookupBrokerOutputArgs, opts ...pulumi.InvokeOption) LookupBrokerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBrokerResult, error) {
			args := v.(LookupBrokerArgs)
			r, err := LookupBroker(ctx, &args, opts...)
			var s LookupBrokerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBrokerResultOutput)
}

type LookupBrokerOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupBrokerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBrokerArgs)(nil)).Elem()
}

type LookupBrokerResultOutput struct{ *pulumi.OutputState }

func (LookupBrokerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBrokerResult)(nil)).Elem()
}

func (o LookupBrokerResultOutput) ToLookupBrokerResultOutput() LookupBrokerResultOutput {
	return o
}

func (o LookupBrokerResultOutput) ToLookupBrokerResultOutputWithContext(ctx context.Context) LookupBrokerResultOutput {
	return o
}

func (o LookupBrokerResultOutput) AmqpEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []string { return v.AmqpEndpoints }).(pulumi.StringArrayOutput)
}

func (o LookupBrokerResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupBrokerResultOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *bool { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupBrokerResultOutput) Configuration() BrokerConfigurationIdPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *BrokerConfigurationId { return v.Configuration }).(BrokerConfigurationIdPtrOutput)
}

func (o LookupBrokerResultOutput) ConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *string { return v.ConfigurationId }).(pulumi.StringPtrOutput)
}

func (o LookupBrokerResultOutput) ConfigurationRevision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *int { return v.ConfigurationRevision }).(pulumi.IntPtrOutput)
}

func (o LookupBrokerResultOutput) DataReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *string { return v.DataReplicationMode }).(pulumi.StringPtrOutput)
}

func (o LookupBrokerResultOutput) DataReplicationPrimaryBrokerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *string { return v.DataReplicationPrimaryBrokerArn }).(pulumi.StringPtrOutput)
}

func (o LookupBrokerResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o LookupBrokerResultOutput) HostInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *string { return v.HostInstanceType }).(pulumi.StringPtrOutput)
}

func (o LookupBrokerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupBrokerResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupBrokerResultOutput) LdapServerMetadata() BrokerLdapServerMetadataPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *BrokerLdapServerMetadata { return v.LdapServerMetadata }).(BrokerLdapServerMetadataPtrOutput)
}

func (o LookupBrokerResultOutput) Logs() BrokerLogListPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *BrokerLogList { return v.Logs }).(BrokerLogListPtrOutput)
}

func (o LookupBrokerResultOutput) MaintenanceWindowStartTime() BrokerMaintenanceWindowPtrOutput {
	return o.ApplyT(func(v LookupBrokerResult) *BrokerMaintenanceWindow { return v.MaintenanceWindowStartTime }).(BrokerMaintenanceWindowPtrOutput)
}

func (o LookupBrokerResultOutput) MqttEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []string { return v.MqttEndpoints }).(pulumi.StringArrayOutput)
}

func (o LookupBrokerResultOutput) OpenWireEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []string { return v.OpenWireEndpoints }).(pulumi.StringArrayOutput)
}

func (o LookupBrokerResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o LookupBrokerResultOutput) StompEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []string { return v.StompEndpoints }).(pulumi.StringArrayOutput)
}

func (o LookupBrokerResultOutput) Tags() BrokerTagsEntryArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []BrokerTagsEntry { return v.Tags }).(BrokerTagsEntryArrayOutput)
}

func (o LookupBrokerResultOutput) Users() BrokerUserArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []BrokerUser { return v.Users }).(BrokerUserArrayOutput)
}

func (o LookupBrokerResultOutput) WssEndpoints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupBrokerResult) []string { return v.WssEndpoints }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBrokerResultOutput{})
}
