// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package memorydb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::MemoryDB::Cluster resource creates an Amazon MemoryDB Cluster.
type Cluster struct {
	pulumi.CustomResourceState

	// The name of the Access Control List to associate with the cluster.
	ACLName pulumi.StringOutput `pulumi:"aCLName"`
	// The Amazon Resource Name (ARN) of the cluster.
	ARN pulumi.StringOutput `pulumi:"aRN"`
	// A flag that enables automatic minor version upgrade when set to true.
	//
	// You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
	AutoMinorVersionUpgrade pulumi.BoolPtrOutput `pulumi:"autoMinorVersionUpgrade"`
	// The cluster endpoint.
	ClusterEndpoint ClusterEndpointPtrOutput `pulumi:"clusterEndpoint"`
	// The name of the cluster. This value must be unique as it also serves as the cluster identifier.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Enables data tiering. Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.
	DataTiering ClusterDataTieringStatusPtrOutput `pulumi:"dataTiering"`
	// An optional description of the cluster.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Redis engine version used by the cluster.
	EngineVersion pulumi.StringPtrOutput `pulumi:"engineVersion"`
	// The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
	FinalSnapshotName pulumi.StringPtrOutput `pulumi:"finalSnapshotName"`
	// The ID of the KMS key used to encrypt the cluster.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	MaintenanceWindow pulumi.StringPtrOutput `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes in the cluster.
	NodeType pulumi.StringOutput `pulumi:"nodeType"`
	// The number of replicas to apply to each shard. The limit is 5.
	NumReplicasPerShard pulumi.IntPtrOutput `pulumi:"numReplicasPerShard"`
	// The number of shards the cluster will contain.
	NumShards pulumi.IntPtrOutput `pulumi:"numShards"`
	// The name of the parameter group associated with the cluster.
	ParameterGroupName pulumi.StringPtrOutput `pulumi:"parameterGroupName"`
	// The status of the parameter group used by the cluster.
	ParameterGroupStatus pulumi.StringOutput `pulumi:"parameterGroupStatus"`
	// The port number on which each member of the cluster accepts connections.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// One or more Amazon VPC security groups associated with this cluster.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.
	SnapshotArns pulumi.StringArrayOutput `pulumi:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.
	SnapshotName pulumi.StringPtrOutput `pulumi:"snapshotName"`
	// The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
	SnapshotRetentionLimit pulumi.IntPtrOutput `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
	SnapshotWindow pulumi.StringPtrOutput `pulumi:"snapshotWindow"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	SnsTopicArn pulumi.StringPtrOutput `pulumi:"snsTopicArn"`
	// The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
	SnsTopicStatus pulumi.StringPtrOutput `pulumi:"snsTopicStatus"`
	// The status of the cluster. For example, Available, Updating, Creating.
	Status pulumi.StringOutput `pulumi:"status"`
	// The name of the subnet group to be used for the cluster.
	SubnetGroupName pulumi.StringPtrOutput `pulumi:"subnetGroupName"`
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	TLSEnabled pulumi.BoolPtrOutput `pulumi:"tLSEnabled"`
	// An array of key-value pairs to apply to this cluster.
	Tags ClusterTagArrayOutput `pulumi:"tags"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ACLName == nil {
		return nil, errors.New("invalid value for required argument 'ACLName'")
	}
	if args.NodeType == nil {
		return nil, errors.New("invalid value for required argument 'NodeType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("aws-native:memorydb:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("aws-native:memorydb:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// The name of the Access Control List to associate with the cluster.
	ACLName string `pulumi:"aCLName"`
	// A flag that enables automatic minor version upgrade when set to true.
	//
	// You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
	AutoMinorVersionUpgrade *bool `pulumi:"autoMinorVersionUpgrade"`
	// The cluster endpoint.
	ClusterEndpoint *ClusterEndpoint `pulumi:"clusterEndpoint"`
	// The name of the cluster. This value must be unique as it also serves as the cluster identifier.
	ClusterName *string `pulumi:"clusterName"`
	// Enables data tiering. Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.
	DataTiering *ClusterDataTieringStatus `pulumi:"dataTiering"`
	// An optional description of the cluster.
	Description *string `pulumi:"description"`
	// The Redis engine version used by the cluster.
	EngineVersion *string `pulumi:"engineVersion"`
	// The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
	FinalSnapshotName *string `pulumi:"finalSnapshotName"`
	// The ID of the KMS key used to encrypt the cluster.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	// The compute and memory capacity of the nodes in the cluster.
	NodeType string `pulumi:"nodeType"`
	// The number of replicas to apply to each shard. The limit is 5.
	NumReplicasPerShard *int `pulumi:"numReplicasPerShard"`
	// The number of shards the cluster will contain.
	NumShards *int `pulumi:"numShards"`
	// The name of the parameter group associated with the cluster.
	ParameterGroupName *string `pulumi:"parameterGroupName"`
	// The port number on which each member of the cluster accepts connections.
	Port *int `pulumi:"port"`
	// One or more Amazon VPC security groups associated with this cluster.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.
	SnapshotArns []string `pulumi:"snapshotArns"`
	// The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.
	SnapshotName *string `pulumi:"snapshotName"`
	// The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
	SnapshotRetentionLimit *int `pulumi:"snapshotRetentionLimit"`
	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
	SnapshotWindow *string `pulumi:"snapshotWindow"`
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	SnsTopicArn *string `pulumi:"snsTopicArn"`
	// The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
	SnsTopicStatus *string `pulumi:"snsTopicStatus"`
	// The name of the subnet group to be used for the cluster.
	SubnetGroupName *string `pulumi:"subnetGroupName"`
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	TLSEnabled *bool `pulumi:"tLSEnabled"`
	// An array of key-value pairs to apply to this cluster.
	Tags []ClusterTag `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// The name of the Access Control List to associate with the cluster.
	ACLName pulumi.StringInput
	// A flag that enables automatic minor version upgrade when set to true.
	//
	// You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
	AutoMinorVersionUpgrade pulumi.BoolPtrInput
	// The cluster endpoint.
	ClusterEndpoint ClusterEndpointPtrInput
	// The name of the cluster. This value must be unique as it also serves as the cluster identifier.
	ClusterName pulumi.StringPtrInput
	// Enables data tiering. Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.
	DataTiering ClusterDataTieringStatusPtrInput
	// An optional description of the cluster.
	Description pulumi.StringPtrInput
	// The Redis engine version used by the cluster.
	EngineVersion pulumi.StringPtrInput
	// The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
	FinalSnapshotName pulumi.StringPtrInput
	// The ID of the KMS key used to encrypt the cluster.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
	MaintenanceWindow pulumi.StringPtrInput
	// The compute and memory capacity of the nodes in the cluster.
	NodeType pulumi.StringInput
	// The number of replicas to apply to each shard. The limit is 5.
	NumReplicasPerShard pulumi.IntPtrInput
	// The number of shards the cluster will contain.
	NumShards pulumi.IntPtrInput
	// The name of the parameter group associated with the cluster.
	ParameterGroupName pulumi.StringPtrInput
	// The port number on which each member of the cluster accepts connections.
	Port pulumi.IntPtrInput
	// One or more Amazon VPC security groups associated with this cluster.
	SecurityGroupIds pulumi.StringArrayInput
	// A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.
	SnapshotArns pulumi.StringArrayInput
	// The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.
	SnapshotName pulumi.StringPtrInput
	// The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
	SnapshotRetentionLimit pulumi.IntPtrInput
	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
	SnapshotWindow pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
	SnsTopicArn pulumi.StringPtrInput
	// The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
	SnsTopicStatus pulumi.StringPtrInput
	// The name of the subnet group to be used for the cluster.
	SubnetGroupName pulumi.StringPtrInput
	// A flag that enables in-transit encryption when set to true.
	//
	// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
	TLSEnabled pulumi.BoolPtrInput
	// An array of key-value pairs to apply to this cluster.
	Tags ClusterTagArrayInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// The name of the Access Control List to associate with the cluster.
func (o ClusterOutput) ACLName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ACLName }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the cluster.
func (o ClusterOutput) ARN() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ARN }).(pulumi.StringOutput)
}

// A flag that enables automatic minor version upgrade when set to true.
//
// You cannot modify the value of AutoMinorVersionUpgrade after the cluster is created. To enable AutoMinorVersionUpgrade on a cluster you must set AutoMinorVersionUpgrade to true when you create a cluster.
func (o ClusterOutput) AutoMinorVersionUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.AutoMinorVersionUpgrade }).(pulumi.BoolPtrOutput)
}

// The cluster endpoint.
func (o ClusterOutput) ClusterEndpoint() ClusterEndpointPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterEndpointPtrOutput { return v.ClusterEndpoint }).(ClusterEndpointPtrOutput)
}

// The name of the cluster. This value must be unique as it also serves as the cluster identifier.
func (o ClusterOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Enables data tiering. Data tiering is only supported for clusters using the r6gd node type. This parameter must be set when using r6gd nodes.
func (o ClusterOutput) DataTiering() ClusterDataTieringStatusPtrOutput {
	return o.ApplyT(func(v *Cluster) ClusterDataTieringStatusPtrOutput { return v.DataTiering }).(ClusterDataTieringStatusPtrOutput)
}

// An optional description of the cluster.
func (o ClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Redis engine version used by the cluster.
func (o ClusterOutput) EngineVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.EngineVersion }).(pulumi.StringPtrOutput)
}

// The user-supplied name of a final cluster snapshot. This is the unique name that identifies the snapshot. MemoryDB creates the snapshot, and then deletes the cluster immediately afterward.
func (o ClusterOutput) FinalSnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.FinalSnapshotName }).(pulumi.StringPtrOutput)
}

// The ID of the KMS key used to encrypt the cluster.
func (o ClusterOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.
func (o ClusterOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

// The compute and memory capacity of the nodes in the cluster.
func (o ClusterOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.NodeType }).(pulumi.StringOutput)
}

// The number of replicas to apply to each shard. The limit is 5.
func (o ClusterOutput) NumReplicasPerShard() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.NumReplicasPerShard }).(pulumi.IntPtrOutput)
}

// The number of shards the cluster will contain.
func (o ClusterOutput) NumShards() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.NumShards }).(pulumi.IntPtrOutput)
}

// The name of the parameter group associated with the cluster.
func (o ClusterOutput) ParameterGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.ParameterGroupName }).(pulumi.StringPtrOutput)
}

// The status of the parameter group used by the cluster.
func (o ClusterOutput) ParameterGroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ParameterGroupStatus }).(pulumi.StringOutput)
}

// The port number on which each member of the cluster accepts connections.
func (o ClusterOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// One or more Amazon VPC security groups associated with this cluster.
func (o ClusterOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new cluster. The Amazon S3 object name in the ARN cannot contain any commas.
func (o ClusterOutput) SnapshotArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringArrayOutput { return v.SnapshotArns }).(pulumi.StringArrayOutput)
}

// The name of a snapshot from which to restore data into the new cluster. The snapshot status changes to restoring while the new cluster is being created.
func (o ClusterOutput) SnapshotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SnapshotName }).(pulumi.StringPtrOutput)
}

// The number of days for which MemoryDB retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.
func (o ClusterOutput) SnapshotRetentionLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.IntPtrOutput { return v.SnapshotRetentionLimit }).(pulumi.IntPtrOutput)
}

// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your cluster.
func (o ClusterOutput) SnapshotWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SnapshotWindow }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.
func (o ClusterOutput) SnsTopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SnsTopicArn }).(pulumi.StringPtrOutput)
}

// The status of the Amazon SNS notification topic. Notifications are sent only if the status is enabled.
func (o ClusterOutput) SnsTopicStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SnsTopicStatus }).(pulumi.StringPtrOutput)
}

// The status of the cluster. For example, Available, Updating, Creating.
func (o ClusterOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The name of the subnet group to be used for the cluster.
func (o ClusterOutput) SubnetGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.SubnetGroupName }).(pulumi.StringPtrOutput)
}

// A flag that enables in-transit encryption when set to true.
//
// You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.
func (o ClusterOutput) TLSEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolPtrOutput { return v.TLSEnabled }).(pulumi.BoolPtrOutput)
}

// An array of key-value pairs to apply to this cluster.
func (o ClusterOutput) Tags() ClusterTagArrayOutput {
	return o.ApplyT(func(v *Cluster) ClusterTagArrayOutput { return v.Tags }).(ClusterTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterOutputType(ClusterOutput{})
}
