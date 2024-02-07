// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opensearchserverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Amazon OpenSearchServerless collection resource
type Collection struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the collection.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The endpoint for the collection.
	CollectionEndpoint pulumi.StringOutput `pulumi:"collectionEndpoint"`
	// The OpenSearch Dashboards endpoint for the collection.
	DashboardEndpoint pulumi.StringOutput `pulumi:"dashboardEndpoint"`
	// The description of the collection
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the collection.
	//
	// The name must meet the following criteria:
	// Unique to your account and AWS Region
	// Starts with a lowercase letter
	// Contains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)
	// Contains between 3 and 32 characters
	Name            pulumi.StringOutput                `pulumi:"name"`
	StandbyReplicas CollectionStandbyReplicasPtrOutput `pulumi:"standbyReplicas"`
	// List of tags to be added to the resource
	Tags CollectionTagArrayOutput `pulumi:"tags"`
	Type CollectionTypePtrOutput  `pulumi:"type"`
}

// NewCollection registers a new resource with the given unique name, arguments, and options.
func NewCollection(ctx *pulumi.Context,
	name string, args *CollectionArgs, opts ...pulumi.ResourceOption) (*Collection, error) {
	if args == nil {
		args = &CollectionArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
		"tags[*]",
		"type",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Collection
	err := ctx.RegisterResource("aws-native:opensearchserverless:Collection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCollection gets an existing Collection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CollectionState, opts ...pulumi.ResourceOption) (*Collection, error) {
	var resource Collection
	err := ctx.ReadResource("aws-native:opensearchserverless:Collection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Collection resources.
type collectionState struct {
}

type CollectionState struct {
}

func (CollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*collectionState)(nil)).Elem()
}

type collectionArgs struct {
	// The description of the collection
	Description *string `pulumi:"description"`
	// The name of the collection.
	//
	// The name must meet the following criteria:
	// Unique to your account and AWS Region
	// Starts with a lowercase letter
	// Contains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)
	// Contains between 3 and 32 characters
	Name            *string                    `pulumi:"name"`
	StandbyReplicas *CollectionStandbyReplicas `pulumi:"standbyReplicas"`
	// List of tags to be added to the resource
	Tags []CollectionTag `pulumi:"tags"`
	Type *CollectionType `pulumi:"type"`
}

// The set of arguments for constructing a Collection resource.
type CollectionArgs struct {
	// The description of the collection
	Description pulumi.StringPtrInput
	// The name of the collection.
	//
	// The name must meet the following criteria:
	// Unique to your account and AWS Region
	// Starts with a lowercase letter
	// Contains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)
	// Contains between 3 and 32 characters
	Name            pulumi.StringPtrInput
	StandbyReplicas CollectionStandbyReplicasPtrInput
	// List of tags to be added to the resource
	Tags CollectionTagArrayInput
	Type CollectionTypePtrInput
}

func (CollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*collectionArgs)(nil)).Elem()
}

type CollectionInput interface {
	pulumi.Input

	ToCollectionOutput() CollectionOutput
	ToCollectionOutputWithContext(ctx context.Context) CollectionOutput
}

func (*Collection) ElementType() reflect.Type {
	return reflect.TypeOf((**Collection)(nil)).Elem()
}

func (i *Collection) ToCollectionOutput() CollectionOutput {
	return i.ToCollectionOutputWithContext(context.Background())
}

func (i *Collection) ToCollectionOutputWithContext(ctx context.Context) CollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectionOutput)
}

type CollectionOutput struct{ *pulumi.OutputState }

func (CollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Collection)(nil)).Elem()
}

func (o CollectionOutput) ToCollectionOutput() CollectionOutput {
	return o
}

func (o CollectionOutput) ToCollectionOutputWithContext(ctx context.Context) CollectionOutput {
	return o
}

// The Amazon Resource Name (ARN) of the collection.
func (o CollectionOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Collection) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The endpoint for the collection.
func (o CollectionOutput) CollectionEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Collection) pulumi.StringOutput { return v.CollectionEndpoint }).(pulumi.StringOutput)
}

// The OpenSearch Dashboards endpoint for the collection.
func (o CollectionOutput) DashboardEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Collection) pulumi.StringOutput { return v.DashboardEndpoint }).(pulumi.StringOutput)
}

// The description of the collection
func (o CollectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Collection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the collection.
//
// The name must meet the following criteria:
// Unique to your account and AWS Region
// Starts with a lowercase letter
// Contains only lowercase letters a-z, the numbers 0-9 and the hyphen (-)
// Contains between 3 and 32 characters
func (o CollectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Collection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CollectionOutput) StandbyReplicas() CollectionStandbyReplicasPtrOutput {
	return o.ApplyT(func(v *Collection) CollectionStandbyReplicasPtrOutput { return v.StandbyReplicas }).(CollectionStandbyReplicasPtrOutput)
}

// List of tags to be added to the resource
func (o CollectionOutput) Tags() CollectionTagArrayOutput {
	return o.ApplyT(func(v *Collection) CollectionTagArrayOutput { return v.Tags }).(CollectionTagArrayOutput)
}

func (o CollectionOutput) Type() CollectionTypePtrOutput {
	return o.ApplyT(func(v *Collection) CollectionTypePtrOutput { return v.Type }).(CollectionTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CollectionInput)(nil)).Elem(), &Collection{})
	pulumi.RegisterOutputType(CollectionOutput{})
}
