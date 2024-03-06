// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ElastiCache::CacheCluster
//
// Deprecated: CacheCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type CacheCluster struct {
	pulumi.CustomResourceState

	AutoMinorVersionUpgrade      pulumi.BoolPtrOutput                                   `pulumi:"autoMinorVersionUpgrade"`
	AwsId                        pulumi.StringOutput                                    `pulumi:"awsId"`
	AzMode                       pulumi.StringPtrOutput                                 `pulumi:"azMode"`
	CacheNodeType                pulumi.StringOutput                                    `pulumi:"cacheNodeType"`
	CacheParameterGroupName      pulumi.StringPtrOutput                                 `pulumi:"cacheParameterGroupName"`
	CacheSecurityGroupNames      pulumi.StringArrayOutput                               `pulumi:"cacheSecurityGroupNames"`
	CacheSubnetGroupName         pulumi.StringPtrOutput                                 `pulumi:"cacheSubnetGroupName"`
	ClusterName                  pulumi.StringPtrOutput                                 `pulumi:"clusterName"`
	ConfigurationEndpointAddress pulumi.StringPtrOutput                                 `pulumi:"configurationEndpointAddress"`
	ConfigurationEndpointPort    pulumi.StringPtrOutput                                 `pulumi:"configurationEndpointPort"`
	Engine                       pulumi.StringOutput                                    `pulumi:"engine"`
	EngineVersion                pulumi.StringPtrOutput                                 `pulumi:"engineVersion"`
	IpDiscovery                  pulumi.StringPtrOutput                                 `pulumi:"ipDiscovery"`
	LogDeliveryConfigurations    CacheClusterLogDeliveryConfigurationRequestArrayOutput `pulumi:"logDeliveryConfigurations"`
	NetworkType                  pulumi.StringPtrOutput                                 `pulumi:"networkType"`
	NotificationTopicArn         pulumi.StringPtrOutput                                 `pulumi:"notificationTopicArn"`
	NumCacheNodes                pulumi.IntOutput                                       `pulumi:"numCacheNodes"`
	Port                         pulumi.IntPtrOutput                                    `pulumi:"port"`
	PreferredAvailabilityZone    pulumi.StringPtrOutput                                 `pulumi:"preferredAvailabilityZone"`
	PreferredAvailabilityZones   pulumi.StringArrayOutput                               `pulumi:"preferredAvailabilityZones"`
	PreferredMaintenanceWindow   pulumi.StringPtrOutput                                 `pulumi:"preferredMaintenanceWindow"`
	RedisEndpointAddress         pulumi.StringPtrOutput                                 `pulumi:"redisEndpointAddress"`
	RedisEndpointPort            pulumi.StringPtrOutput                                 `pulumi:"redisEndpointPort"`
	SnapshotArns                 pulumi.StringArrayOutput                               `pulumi:"snapshotArns"`
	SnapshotName                 pulumi.StringPtrOutput                                 `pulumi:"snapshotName"`
	SnapshotRetentionLimit       pulumi.IntPtrOutput                                    `pulumi:"snapshotRetentionLimit"`
	SnapshotWindow               pulumi.StringPtrOutput                                 `pulumi:"snapshotWindow"`
	Tags                         aws.TagArrayOutput                                     `pulumi:"tags"`
	TransitEncryptionEnabled     pulumi.BoolPtrOutput                                   `pulumi:"transitEncryptionEnabled"`
	VpcSecurityGroupIds          pulumi.StringArrayOutput                               `pulumi:"vpcSecurityGroupIds"`
}

// NewCacheCluster registers a new resource with the given unique name, arguments, and options.
func NewCacheCluster(ctx *pulumi.Context,
	name string, args *CacheClusterArgs, opts ...pulumi.ResourceOption) (*CacheCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CacheNodeType == nil {
		return nil, errors.New("invalid value for required argument 'CacheNodeType'")
	}
	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.NumCacheNodes == nil {
		return nil, errors.New("invalid value for required argument 'NumCacheNodes'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"cacheSubnetGroupName",
		"clusterName",
		"engine",
		"networkType",
		"port",
		"snapshotArns[*]",
		"snapshotName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CacheCluster
	err := ctx.RegisterResource("aws-native:elasticache:CacheCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCacheCluster gets an existing CacheCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCacheCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CacheClusterState, opts ...pulumi.ResourceOption) (*CacheCluster, error) {
	var resource CacheCluster
	err := ctx.ReadResource("aws-native:elasticache:CacheCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CacheCluster resources.
type cacheClusterState struct {
}

type CacheClusterState struct {
}

func (CacheClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheClusterState)(nil)).Elem()
}

type cacheClusterArgs struct {
	AutoMinorVersionUpgrade      *bool                                         `pulumi:"autoMinorVersionUpgrade"`
	AzMode                       *string                                       `pulumi:"azMode"`
	CacheNodeType                string                                        `pulumi:"cacheNodeType"`
	CacheParameterGroupName      *string                                       `pulumi:"cacheParameterGroupName"`
	CacheSecurityGroupNames      []string                                      `pulumi:"cacheSecurityGroupNames"`
	CacheSubnetGroupName         *string                                       `pulumi:"cacheSubnetGroupName"`
	ClusterName                  *string                                       `pulumi:"clusterName"`
	ConfigurationEndpointAddress *string                                       `pulumi:"configurationEndpointAddress"`
	ConfigurationEndpointPort    *string                                       `pulumi:"configurationEndpointPort"`
	Engine                       string                                        `pulumi:"engine"`
	EngineVersion                *string                                       `pulumi:"engineVersion"`
	IpDiscovery                  *string                                       `pulumi:"ipDiscovery"`
	LogDeliveryConfigurations    []CacheClusterLogDeliveryConfigurationRequest `pulumi:"logDeliveryConfigurations"`
	NetworkType                  *string                                       `pulumi:"networkType"`
	NotificationTopicArn         *string                                       `pulumi:"notificationTopicArn"`
	NumCacheNodes                int                                           `pulumi:"numCacheNodes"`
	Port                         *int                                          `pulumi:"port"`
	PreferredAvailabilityZone    *string                                       `pulumi:"preferredAvailabilityZone"`
	PreferredAvailabilityZones   []string                                      `pulumi:"preferredAvailabilityZones"`
	PreferredMaintenanceWindow   *string                                       `pulumi:"preferredMaintenanceWindow"`
	RedisEndpointAddress         *string                                       `pulumi:"redisEndpointAddress"`
	RedisEndpointPort            *string                                       `pulumi:"redisEndpointPort"`
	SnapshotArns                 []string                                      `pulumi:"snapshotArns"`
	SnapshotName                 *string                                       `pulumi:"snapshotName"`
	SnapshotRetentionLimit       *int                                          `pulumi:"snapshotRetentionLimit"`
	SnapshotWindow               *string                                       `pulumi:"snapshotWindow"`
	Tags                         []aws.Tag                                     `pulumi:"tags"`
	TransitEncryptionEnabled     *bool                                         `pulumi:"transitEncryptionEnabled"`
	VpcSecurityGroupIds          []string                                      `pulumi:"vpcSecurityGroupIds"`
}

// The set of arguments for constructing a CacheCluster resource.
type CacheClusterArgs struct {
	AutoMinorVersionUpgrade      pulumi.BoolPtrInput
	AzMode                       pulumi.StringPtrInput
	CacheNodeType                pulumi.StringInput
	CacheParameterGroupName      pulumi.StringPtrInput
	CacheSecurityGroupNames      pulumi.StringArrayInput
	CacheSubnetGroupName         pulumi.StringPtrInput
	ClusterName                  pulumi.StringPtrInput
	ConfigurationEndpointAddress pulumi.StringPtrInput
	ConfigurationEndpointPort    pulumi.StringPtrInput
	Engine                       pulumi.StringInput
	EngineVersion                pulumi.StringPtrInput
	IpDiscovery                  pulumi.StringPtrInput
	LogDeliveryConfigurations    CacheClusterLogDeliveryConfigurationRequestArrayInput
	NetworkType                  pulumi.StringPtrInput
	NotificationTopicArn         pulumi.StringPtrInput
	NumCacheNodes                pulumi.IntInput
	Port                         pulumi.IntPtrInput
	PreferredAvailabilityZone    pulumi.StringPtrInput
	PreferredAvailabilityZones   pulumi.StringArrayInput
	PreferredMaintenanceWindow   pulumi.StringPtrInput
	RedisEndpointAddress         pulumi.StringPtrInput
	RedisEndpointPort            pulumi.StringPtrInput
	SnapshotArns                 pulumi.StringArrayInput
	SnapshotName                 pulumi.StringPtrInput
	SnapshotRetentionLimit       pulumi.IntPtrInput
	SnapshotWindow               pulumi.StringPtrInput
	Tags                         aws.TagArrayInput
	TransitEncryptionEnabled     pulumi.BoolPtrInput
	VpcSecurityGroupIds          pulumi.StringArrayInput
}

func (CacheClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cacheClusterArgs)(nil)).Elem()
}

type CacheClusterInput interface {
	pulumi.Input

	ToCacheClusterOutput() CacheClusterOutput
	ToCacheClusterOutputWithContext(ctx context.Context) CacheClusterOutput
}

func (*CacheCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheCluster)(nil)).Elem()
}

func (i *CacheCluster) ToCacheClusterOutput() CacheClusterOutput {
	return i.ToCacheClusterOutputWithContext(context.Background())
}

func (i *CacheCluster) ToCacheClusterOutputWithContext(ctx context.Context) CacheClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CacheClusterOutput)
}

type CacheClusterOutput struct{ *pulumi.OutputState }

func (CacheClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CacheCluster)(nil)).Elem()
}

func (o CacheClusterOutput) ToCacheClusterOutput() CacheClusterOutput {
	return o
}

func (o CacheClusterOutput) ToCacheClusterOutputWithContext(ctx context.Context) CacheClusterOutput {
	return o
}

func (o CacheClusterOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.BoolPtrOutput { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

func (o CacheClusterOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o CacheClusterOutput) AzMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.AzMode }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) CacheNodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringOutput { return v.CacheNodeType }).(pulumi.StringOutput)
}

func (o CacheClusterOutput) CacheParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.CacheParameterGroupName }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) CacheSecurityGroupNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringArrayOutput { return v.CacheSecurityGroupNames }).(pulumi.StringArrayOutput)
}

func (o CacheClusterOutput) CacheSubnetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.CacheSubnetGroupName }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.ClusterName }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) ConfigurationEndpointAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.ConfigurationEndpointAddress }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) ConfigurationEndpointPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.ConfigurationEndpointPort }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

func (o CacheClusterOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) IpDiscovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.IpDiscovery }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) LogDeliveryConfigurations() CacheClusterLogDeliveryConfigurationRequestArrayOutput {
	return o.ApplyT(func(v *CacheCluster) CacheClusterLogDeliveryConfigurationRequestArrayOutput {
		return v.LogDeliveryConfigurations
	}).(CacheClusterLogDeliveryConfigurationRequestArrayOutput)
}

func (o CacheClusterOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) NotificationTopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.NotificationTopicArn }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) NumCacheNodes() pulumi.IntOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.IntOutput { return v.NumCacheNodes }).(pulumi.IntOutput)
}

func (o CacheClusterOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

func (o CacheClusterOutput) PreferredAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.PreferredAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) PreferredAvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringArrayOutput { return v.PreferredAvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o CacheClusterOutput) PreferredMaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.PreferredMaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) RedisEndpointAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.RedisEndpointAddress }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) RedisEndpointPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.RedisEndpointPort }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) SnapshotArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringArrayOutput { return v.SnapshotArns }).(pulumi.StringArrayOutput)
}

func (o CacheClusterOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) SnapshotRetentionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.IntPtrOutput { return v.SnapshotRetentionLimit }).(pulumi.IntPtrOutput)
}

func (o CacheClusterOutput) SnapshotWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringPtrOutput { return v.SnapshotWindow }).(pulumi.StringPtrOutput)
}

func (o CacheClusterOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *CacheCluster) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o CacheClusterOutput) TransitEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.BoolPtrOutput { return v.TransitEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

func (o CacheClusterOutput) VpcSecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CacheCluster) pulumi.StringArrayOutput { return v.VpcSecurityGroupIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CacheClusterInput)(nil)).Elem(), &CacheCluster{})
	pulumi.RegisterOutputType(CacheClusterOutput{})
}
