// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The “AWS::EFS::FileSystem“ resource creates a new, empty file system in EFSlong (EFS). You must create a mount target ([AWS::EFS::MountTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html)) to mount your EFS file system on an EC2 or other AWS cloud compute resource.
func LookupFileSystem(ctx *pulumi.Context, args *LookupFileSystemArgs, opts ...pulumi.InvokeOption) (*LookupFileSystemResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFileSystemResult
	err := ctx.Invoke("aws-native:efs:getFileSystem", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFileSystemArgs struct {
	FileSystemId string `pulumi:"fileSystemId"`
}

type LookupFileSystemResult struct {
	Arn *string `pulumi:"arn"`
	// Use the ``BackupPolicy`` to turn automatic backups on or off for the file system.
	BackupPolicy *FileSystemBackupPolicy `pulumi:"backupPolicy"`
	FileSystemId *string                 `pulumi:"fileSystemId"`
	// The ``FileSystemPolicy`` for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EFS::FileSystem` for more information about the expected schema for this property.
	FileSystemPolicy interface{} `pulumi:"fileSystemPolicy"`
	// Describes the protection on the file system.
	FileSystemProtection *FileSystemProtection `pulumi:"fileSystemProtection"`
	// Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a ``"Key":"Name","Value":"{value}"`` key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*.
	FileSystemTags []aws.Tag `pulumi:"fileSystemTags"`
	// An array of ``LifecyclePolicy`` objects that define the file system's ``LifecycleConfiguration`` object. A ``LifecycleConfiguration`` object informs Lifecycle management of the following:
	//   +  When to move files in the file system from primary storage to IA storage.
	//   + When to move files in the file system from primary storage or IA storage to Archive storage.
	//  +  When to move files that are in IA or Archive storage to primary storage.
	//
	//   EFS requires that each ``LifecyclePolicy`` object have only a single transition. This means that in a request body, ``LifecyclePolicies`` needs to be structured as an array of ``LifecyclePolicy`` objects, one object for each transition, ``TransitionToIA``, ``TransitionToArchive`` ``TransitionToPrimaryStorageClass``. See the example requests in the following section for more information.
	LifecyclePolicies []FileSystemLifecyclePolicy `pulumi:"lifecyclePolicies"`
	// The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ``ThroughputMode`` is set to ``provisioned``. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*.
	ProvisionedThroughputInMibps *float64 `pulumi:"provisionedThroughputInMibps"`
	// Describes the replication configuration for a specific file system.
	ReplicationConfiguration *FileSystemReplicationConfiguration `pulumi:"replicationConfiguration"`
	// Specifies the throughput mode for the file system. The mode can be ``bursting``, ``provisioned``, or ``elastic``. If you set ``ThroughputMode`` to ``provisioned``, you must also set a value for ``ProvisionedThroughputInMibps``. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*.
	//  Default is ``bursting``.
	ThroughputMode *string `pulumi:"throughputMode"`
}

func LookupFileSystemOutput(ctx *pulumi.Context, args LookupFileSystemOutputArgs, opts ...pulumi.InvokeOption) LookupFileSystemResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFileSystemResult, error) {
			args := v.(LookupFileSystemArgs)
			r, err := LookupFileSystem(ctx, &args, opts...)
			var s LookupFileSystemResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFileSystemResultOutput)
}

type LookupFileSystemOutputArgs struct {
	FileSystemId pulumi.StringInput `pulumi:"fileSystemId"`
}

func (LookupFileSystemOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileSystemArgs)(nil)).Elem()
}

type LookupFileSystemResultOutput struct{ *pulumi.OutputState }

func (LookupFileSystemResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFileSystemResult)(nil)).Elem()
}

func (o LookupFileSystemResultOutput) ToLookupFileSystemResultOutput() LookupFileSystemResultOutput {
	return o
}

func (o LookupFileSystemResultOutput) ToLookupFileSystemResultOutputWithContext(ctx context.Context) LookupFileSystemResultOutput {
	return o
}

func (o LookupFileSystemResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Use the “BackupPolicy“ to turn automatic backups on or off for the file system.
func (o LookupFileSystemResultOutput) BackupPolicy() FileSystemBackupPolicyPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *FileSystemBackupPolicy { return v.BackupPolicy }).(FileSystemBackupPolicyPtrOutput)
}

func (o LookupFileSystemResultOutput) FileSystemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *string { return v.FileSystemId }).(pulumi.StringPtrOutput)
}

// The “FileSystemPolicy“ for the EFS file system. A file system policy is an IAM resource policy used to control NFS access to an EFS file system. For more information, see [Using to control NFS access to Amazon EFS](https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html) in the *Amazon EFS User Guide*.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::EFS::FileSystem` for more information about the expected schema for this property.
func (o LookupFileSystemResultOutput) FileSystemPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupFileSystemResult) interface{} { return v.FileSystemPolicy }).(pulumi.AnyOutput)
}

// Describes the protection on the file system.
func (o LookupFileSystemResultOutput) FileSystemProtection() FileSystemProtectionPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *FileSystemProtection { return v.FileSystemProtection }).(FileSystemProtectionPtrOutput)
}

// Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a “"Key":"Name","Value":"{value}"“ key-value pair. Each key must be unique. For more information, see [Tagging resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *General Reference Guide*.
func (o LookupFileSystemResultOutput) FileSystemTags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupFileSystemResult) []aws.Tag { return v.FileSystemTags }).(aws.TagArrayOutput)
}

// An array of “LifecyclePolicy“ objects that define the file system's “LifecycleConfiguration“ object. A “LifecycleConfiguration“ object informs Lifecycle management of the following:
//
//   - When to move files in the file system from primary storage to IA storage.
//
//   - When to move files in the file system from primary storage or IA storage to Archive storage.
//
//   - When to move files that are in IA or Archive storage to primary storage.
//
//     EFS requires that each “LifecyclePolicy“ object have only a single transition. This means that in a request body, “LifecyclePolicies“ needs to be structured as an array of “LifecyclePolicy“ objects, one object for each transition, “TransitionToIA“, “TransitionToArchive“ “TransitionToPrimaryStorageClass“. See the example requests in the following section for more information.
func (o LookupFileSystemResultOutput) LifecyclePolicies() FileSystemLifecyclePolicyArrayOutput {
	return o.ApplyT(func(v LookupFileSystemResult) []FileSystemLifecyclePolicy { return v.LifecyclePolicies }).(FileSystemLifecyclePolicyArrayOutput)
}

// The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if “ThroughputMode“ is set to “provisioned“. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact SUP. For more information, see [Amazon EFS quotas that you can increase](https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the *Amazon EFS User Guide*.
func (o LookupFileSystemResultOutput) ProvisionedThroughputInMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *float64 { return v.ProvisionedThroughputInMibps }).(pulumi.Float64PtrOutput)
}

// Describes the replication configuration for a specific file system.
func (o LookupFileSystemResultOutput) ReplicationConfiguration() FileSystemReplicationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *FileSystemReplicationConfiguration { return v.ReplicationConfiguration }).(FileSystemReplicationConfigurationPtrOutput)
}

// Specifies the throughput mode for the file system. The mode can be “bursting“, “provisioned“, or “elastic“. If you set “ThroughputMode“ to “provisioned“, you must also set a value for “ProvisionedThroughputInMibps“. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see [Specifying throughput with provisioned mode](https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the *Amazon EFS User Guide*.
//
//	Default is ``bursting``.
func (o LookupFileSystemResultOutput) ThroughputMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFileSystemResult) *string { return v.ThroughputMode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFileSystemResultOutput{})
}
