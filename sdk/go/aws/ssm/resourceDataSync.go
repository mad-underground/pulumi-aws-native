// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html
type ResourceDataSync struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketname
	BucketName pulumi.StringPtrOutput `pulumi:"bucketName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketprefix
	BucketPrefix pulumi.StringPtrOutput `pulumi:"bucketPrefix"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketregion
	BucketRegion pulumi.StringPtrOutput `pulumi:"bucketRegion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-kmskeyarn
	KMSKeyArn pulumi.StringPtrOutput `pulumi:"kMSKeyArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-s3destination
	S3Destination ResourceDataSyncS3DestinationPtrOutput `pulumi:"s3Destination"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncformat
	SyncFormat pulumi.StringPtrOutput `pulumi:"syncFormat"`
	SyncName   pulumi.StringOutput    `pulumi:"syncName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncsource
	SyncSource ResourceDataSyncSyncSourcePtrOutput `pulumi:"syncSource"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-synctype
	SyncType pulumi.StringPtrOutput `pulumi:"syncType"`
}

// NewResourceDataSync registers a new resource with the given unique name, arguments, and options.
func NewResourceDataSync(ctx *pulumi.Context,
	name string, args *ResourceDataSyncArgs, opts ...pulumi.ResourceOption) (*ResourceDataSync, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SyncName == nil {
		return nil, errors.New("invalid value for required argument 'SyncName'")
	}
	var resource ResourceDataSync
	err := ctx.RegisterResource("aws-native:ssm:ResourceDataSync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceDataSync gets an existing ResourceDataSync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceDataSync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceDataSyncState, opts ...pulumi.ResourceOption) (*ResourceDataSync, error) {
	var resource ResourceDataSync
	err := ctx.ReadResource("aws-native:ssm:ResourceDataSync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceDataSync resources.
type resourceDataSyncState struct {
}

type ResourceDataSyncState struct {
}

func (ResourceDataSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDataSyncState)(nil)).Elem()
}

type resourceDataSyncArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketname
	BucketName *string `pulumi:"bucketName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketprefix
	BucketPrefix *string `pulumi:"bucketPrefix"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketregion
	BucketRegion *string `pulumi:"bucketRegion"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-kmskeyarn
	KMSKeyArn *string `pulumi:"kMSKeyArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-s3destination
	S3Destination *ResourceDataSyncS3Destination `pulumi:"s3Destination"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncformat
	SyncFormat *string `pulumi:"syncFormat"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncname
	SyncName string `pulumi:"syncName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncsource
	SyncSource *ResourceDataSyncSyncSource `pulumi:"syncSource"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-synctype
	SyncType *string `pulumi:"syncType"`
}

// The set of arguments for constructing a ResourceDataSync resource.
type ResourceDataSyncArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketname
	BucketName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketprefix
	BucketPrefix pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-bucketregion
	BucketRegion pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-kmskeyarn
	KMSKeyArn pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-s3destination
	S3Destination ResourceDataSyncS3DestinationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncformat
	SyncFormat pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncname
	SyncName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-syncsource
	SyncSource ResourceDataSyncSyncSourcePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html#cfn-ssm-resourcedatasync-synctype
	SyncType pulumi.StringPtrInput
}

func (ResourceDataSyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDataSyncArgs)(nil)).Elem()
}

type ResourceDataSyncInput interface {
	pulumi.Input

	ToResourceDataSyncOutput() ResourceDataSyncOutput
	ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput
}

func (*ResourceDataSync) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDataSync)(nil))
}

func (i *ResourceDataSync) ToResourceDataSyncOutput() ResourceDataSyncOutput {
	return i.ToResourceDataSyncOutputWithContext(context.Background())
}

func (i *ResourceDataSync) ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncOutput)
}

type ResourceDataSyncOutput struct{ *pulumi.OutputState }

func (ResourceDataSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDataSync)(nil))
}

func (o ResourceDataSyncOutput) ToResourceDataSyncOutput() ResourceDataSyncOutput {
	return o
}

func (o ResourceDataSyncOutput) ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceDataSyncOutput{})
}