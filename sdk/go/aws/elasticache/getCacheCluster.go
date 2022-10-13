// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElastiCache::CacheCluster
func LookupCacheCluster(ctx *pulumi.Context, args *LookupCacheClusterArgs, opts ...pulumi.InvokeOption) (*LookupCacheClusterResult, error) {
	var rv LookupCacheClusterResult
	err := ctx.Invoke("aws-native:elasticache:getCacheCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCacheClusterArgs struct {
	Id string `pulumi:"id"`
}

type LookupCacheClusterResult struct {
	AZMode                       *string                                       `pulumi:"aZMode"`
	AutoMinorVersionUpgrade      *bool                                         `pulumi:"autoMinorVersionUpgrade"`
	CacheNodeType                *string                                       `pulumi:"cacheNodeType"`
	CacheParameterGroupName      *string                                       `pulumi:"cacheParameterGroupName"`
	CacheSecurityGroupNames      []string                                      `pulumi:"cacheSecurityGroupNames"`
	ConfigurationEndpointAddress *string                                       `pulumi:"configurationEndpointAddress"`
	ConfigurationEndpointPort    *string                                       `pulumi:"configurationEndpointPort"`
	EngineVersion                *string                                       `pulumi:"engineVersion"`
	Id                           *string                                       `pulumi:"id"`
	LogDeliveryConfigurations    []CacheClusterLogDeliveryConfigurationRequest `pulumi:"logDeliveryConfigurations"`
	NotificationTopicArn         *string                                       `pulumi:"notificationTopicArn"`
	NumCacheNodes                *int                                          `pulumi:"numCacheNodes"`
	PreferredAvailabilityZone    *string                                       `pulumi:"preferredAvailabilityZone"`
	PreferredAvailabilityZones   []string                                      `pulumi:"preferredAvailabilityZones"`
	PreferredMaintenanceWindow   *string                                       `pulumi:"preferredMaintenanceWindow"`
	RedisEndpointAddress         *string                                       `pulumi:"redisEndpointAddress"`
	RedisEndpointPort            *string                                       `pulumi:"redisEndpointPort"`
	SnapshotRetentionLimit       *int                                          `pulumi:"snapshotRetentionLimit"`
	SnapshotWindow               *string                                       `pulumi:"snapshotWindow"`
	Tags                         []CacheClusterTag                             `pulumi:"tags"`
	TransitEncryptionEnabled     *bool                                         `pulumi:"transitEncryptionEnabled"`
	VpcSecurityGroupIds          []string                                      `pulumi:"vpcSecurityGroupIds"`
}

func LookupCacheClusterOutput(ctx *pulumi.Context, args LookupCacheClusterOutputArgs, opts ...pulumi.InvokeOption) LookupCacheClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCacheClusterResult, error) {
			args := v.(LookupCacheClusterArgs)
			r, err := LookupCacheCluster(ctx, &args, opts...)
			var s LookupCacheClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCacheClusterResultOutput)
}

type LookupCacheClusterOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupCacheClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheClusterArgs)(nil)).Elem()
}

type LookupCacheClusterResultOutput struct{ *pulumi.OutputState }

func (LookupCacheClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheClusterResult)(nil)).Elem()
}

func (o LookupCacheClusterResultOutput) ToLookupCacheClusterResultOutput() LookupCacheClusterResultOutput {
	return o
}

func (o LookupCacheClusterResultOutput) ToLookupCacheClusterResultOutputWithContext(ctx context.Context) LookupCacheClusterResultOutput {
	return o
}

func (o LookupCacheClusterResultOutput) AZMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.AZMode }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *bool { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupCacheClusterResultOutput) CacheNodeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.CacheNodeType }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) CacheParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.CacheParameterGroupName }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) CacheSecurityGroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) []string { return v.CacheSecurityGroupNames }).(pulumi.StringArrayOutput)
}

func (o LookupCacheClusterResultOutput) ConfigurationEndpointAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.ConfigurationEndpointAddress }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) ConfigurationEndpointPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.ConfigurationEndpointPort }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) LogDeliveryConfigurations() CacheClusterLogDeliveryConfigurationRequestArrayOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) []CacheClusterLogDeliveryConfigurationRequest {
		return v.LogDeliveryConfigurations
	}).(CacheClusterLogDeliveryConfigurationRequestArrayOutput)
}

func (o LookupCacheClusterResultOutput) NotificationTopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.NotificationTopicArn }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) NumCacheNodes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *int { return v.NumCacheNodes }).(pulumi.IntPtrOutput)
}

func (o LookupCacheClusterResultOutput) PreferredAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.PreferredAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) PreferredAvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) []string { return v.PreferredAvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o LookupCacheClusterResultOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) RedisEndpointAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.RedisEndpointAddress }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) RedisEndpointPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.RedisEndpointPort }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) SnapshotRetentionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *int { return v.SnapshotRetentionLimit }).(pulumi.IntPtrOutput)
}

func (o LookupCacheClusterResultOutput) SnapshotWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *string { return v.SnapshotWindow }).(pulumi.StringPtrOutput)
}

func (o LookupCacheClusterResultOutput) Tags() CacheClusterTagArrayOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) []CacheClusterTag { return v.Tags }).(CacheClusterTagArrayOutput)
}

func (o LookupCacheClusterResultOutput) TransitEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) *bool { return v.TransitEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupCacheClusterResultOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCacheClusterResult) []string { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCacheClusterResultOutput{})
}
