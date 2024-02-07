// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3express

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::S3Express::DirectoryBucket.
type DirectoryBucket struct {
	pulumi.CustomResourceState

	// Returns the Amazon Resource Name (ARN) of the specified bucket.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	BucketName pulumi.StringPtrOutput `pulumi:"bucketName"`
	// Specifies the number of Availability Zone that's used for redundancy for the bucket.
	DataRedundancy DirectoryBucketDataRedundancyOutput `pulumi:"dataRedundancy"`
	// Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
	LocationName pulumi.StringOutput `pulumi:"locationName"`
}

// NewDirectoryBucket registers a new resource with the given unique name, arguments, and options.
func NewDirectoryBucket(ctx *pulumi.Context,
	name string, args *DirectoryBucketArgs, opts ...pulumi.ResourceOption) (*DirectoryBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataRedundancy == nil {
		return nil, errors.New("invalid value for required argument 'DataRedundancy'")
	}
	if args.LocationName == nil {
		return nil, errors.New("invalid value for required argument 'LocationName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucketName",
		"dataRedundancy",
		"locationName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DirectoryBucket
	err := ctx.RegisterResource("aws-native:s3express:DirectoryBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDirectoryBucket gets an existing DirectoryBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDirectoryBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DirectoryBucketState, opts ...pulumi.ResourceOption) (*DirectoryBucket, error) {
	var resource DirectoryBucket
	err := ctx.ReadResource("aws-native:s3express:DirectoryBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DirectoryBucket resources.
type directoryBucketState struct {
}

type DirectoryBucketState struct {
}

func (DirectoryBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryBucketState)(nil)).Elem()
}

type directoryBucketArgs struct {
	// Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	BucketName *string `pulumi:"bucketName"`
	// Specifies the number of Availability Zone that's used for redundancy for the bucket.
	DataRedundancy DirectoryBucketDataRedundancy `pulumi:"dataRedundancy"`
	// Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
	LocationName string `pulumi:"locationName"`
}

// The set of arguments for constructing a DirectoryBucket resource.
type DirectoryBucketArgs struct {
	// Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
	BucketName pulumi.StringPtrInput
	// Specifies the number of Availability Zone that's used for redundancy for the bucket.
	DataRedundancy DirectoryBucketDataRedundancyInput
	// Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
	LocationName pulumi.StringInput
}

func (DirectoryBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*directoryBucketArgs)(nil)).Elem()
}

type DirectoryBucketInput interface {
	pulumi.Input

	ToDirectoryBucketOutput() DirectoryBucketOutput
	ToDirectoryBucketOutputWithContext(ctx context.Context) DirectoryBucketOutput
}

func (*DirectoryBucket) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryBucket)(nil)).Elem()
}

func (i *DirectoryBucket) ToDirectoryBucketOutput() DirectoryBucketOutput {
	return i.ToDirectoryBucketOutputWithContext(context.Background())
}

func (i *DirectoryBucket) ToDirectoryBucketOutputWithContext(ctx context.Context) DirectoryBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectoryBucketOutput)
}

type DirectoryBucketOutput struct{ *pulumi.OutputState }

func (DirectoryBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectoryBucket)(nil)).Elem()
}

func (o DirectoryBucketOutput) ToDirectoryBucketOutput() DirectoryBucketOutput {
	return o
}

func (o DirectoryBucketOutput) ToDirectoryBucketOutputWithContext(ctx context.Context) DirectoryBucketOutput {
	return o
}

// Returns the Amazon Resource Name (ARN) of the specified bucket.
func (o DirectoryBucketOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
func (o DirectoryBucketOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.StringPtrOutput { return v.BucketName }).(pulumi.StringPtrOutput)
}

// Specifies the number of Availability Zone that's used for redundancy for the bucket.
func (o DirectoryBucketOutput) DataRedundancy() DirectoryBucketDataRedundancyOutput {
	return o.ApplyT(func(v *DirectoryBucket) DirectoryBucketDataRedundancyOutput { return v.DataRedundancy }).(DirectoryBucketDataRedundancyOutput)
}

// Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
func (o DirectoryBucketOutput) LocationName() pulumi.StringOutput {
	return o.ApplyT(func(v *DirectoryBucket) pulumi.StringOutput { return v.LocationName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DirectoryBucketInput)(nil)).Elem(), &DirectoryBucket{})
	pulumi.RegisterOutputType(DirectoryBucketOutput{})
}
