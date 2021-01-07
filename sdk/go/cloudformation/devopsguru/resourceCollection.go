// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devopsguru

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-devopsguru-resourcecollection.html
type ResourceCollection struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ResourceCollectionAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResourceCollectionPropertiesOutput `pulumi:"properties"`
}

// NewResourceCollection registers a new resource with the given unique name, arguments, and options.
func NewResourceCollection(ctx *pulumi.Context,
	name string, args *ResourceCollectionArgs, opts ...pulumi.ResourceOption) (*ResourceCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ResourceCollection
	err := ctx.RegisterResource("cloudformation:DevOpsGuru:ResourceCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceCollection gets an existing ResourceCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceCollectionState, opts ...pulumi.ResourceOption) (*ResourceCollection, error) {
	var resource ResourceCollection
	err := ctx.ReadResource("cloudformation:DevOpsGuru:ResourceCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceCollection resources.
type resourceCollectionState struct {
	// The attributes associated with the resource
	Attributes *ResourceCollectionAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ResourceCollectionProperties `pulumi:"properties"`
}

type ResourceCollectionState struct {
	// The attributes associated with the resource
	Attributes ResourceCollectionAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResourceCollectionPropertiesPtrInput
}

func (ResourceCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceCollectionState)(nil)).Elem()
}

type resourceCollectionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResourceCollectionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ResourceCollection resource.
type ResourceCollectionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResourceCollectionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ResourceCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceCollectionArgs)(nil)).Elem()
}

type ResourceCollectionInput interface {
	pulumi.Input

	ToResourceCollectionOutput() ResourceCollectionOutput
	ToResourceCollectionOutputWithContext(ctx context.Context) ResourceCollectionOutput
}

func (*ResourceCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollection)(nil))
}

func (i *ResourceCollection) ToResourceCollectionOutput() ResourceCollectionOutput {
	return i.ToResourceCollectionOutputWithContext(context.Background())
}

func (i *ResourceCollection) ToResourceCollectionOutputWithContext(ctx context.Context) ResourceCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceCollectionOutput)
}

type ResourceCollectionOutput struct {
	*pulumi.OutputState
}

func (ResourceCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceCollection)(nil))
}

func (o ResourceCollectionOutput) ToResourceCollectionOutput() ResourceCollectionOutput {
	return o
}

func (o ResourceCollectionOutput) ToResourceCollectionOutputWithContext(ctx context.Context) ResourceCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceCollectionOutput{})
}