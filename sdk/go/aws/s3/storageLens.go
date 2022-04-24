// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::S3::StorageLens resource is an Amazon S3 resource type that you can use to create Storage Lens configurations.
type StorageLens struct {
	pulumi.CustomResourceState

	StorageLensConfiguration StorageLensConfigurationOutput `pulumi:"storageLensConfiguration"`
	// A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.
	Tags StorageLensTagArrayOutput `pulumi:"tags"`
}

// NewStorageLens registers a new resource with the given unique name, arguments, and options.
func NewStorageLens(ctx *pulumi.Context,
	name string, args *StorageLensArgs, opts ...pulumi.ResourceOption) (*StorageLens, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageLensConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'StorageLensConfiguration'")
	}
	var resource StorageLens
	err := ctx.RegisterResource("aws-native:s3:StorageLens", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageLens gets an existing StorageLens resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageLens(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageLensState, opts ...pulumi.ResourceOption) (*StorageLens, error) {
	var resource StorageLens
	err := ctx.ReadResource("aws-native:s3:StorageLens", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageLens resources.
type storageLensState struct {
}

type StorageLensState struct {
}

func (StorageLensState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageLensState)(nil)).Elem()
}

type storageLensArgs struct {
	StorageLensConfiguration StorageLensConfiguration `pulumi:"storageLensConfiguration"`
	// A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.
	Tags []StorageLensTag `pulumi:"tags"`
}

// The set of arguments for constructing a StorageLens resource.
type StorageLensArgs struct {
	StorageLensConfiguration StorageLensConfigurationInput
	// A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.
	Tags StorageLensTagArrayInput
}

func (StorageLensArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageLensArgs)(nil)).Elem()
}

type StorageLensInput interface {
	pulumi.Input

	ToStorageLensOutput() StorageLensOutput
	ToStorageLensOutputWithContext(ctx context.Context) StorageLensOutput
}

func (*StorageLens) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageLens)(nil)).Elem()
}

func (i *StorageLens) ToStorageLensOutput() StorageLensOutput {
	return i.ToStorageLensOutputWithContext(context.Background())
}

func (i *StorageLens) ToStorageLensOutputWithContext(ctx context.Context) StorageLensOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageLensOutput)
}

type StorageLensOutput struct{ *pulumi.OutputState }

func (StorageLensOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageLens)(nil)).Elem()
}

func (o StorageLensOutput) ToStorageLensOutput() StorageLensOutput {
	return o
}

func (o StorageLensOutput) ToStorageLensOutputWithContext(ctx context.Context) StorageLensOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageLensInput)(nil)).Elem(), &StorageLens{})
	pulumi.RegisterOutputType(StorageLensOutput{})
}
