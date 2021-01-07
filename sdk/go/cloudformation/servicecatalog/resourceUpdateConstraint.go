// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-resourceupdateconstraint.html
type ResourceUpdateConstraint struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ResourceUpdateConstraintAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResourceUpdateConstraintPropertiesOutput `pulumi:"properties"`
}

// NewResourceUpdateConstraint registers a new resource with the given unique name, arguments, and options.
func NewResourceUpdateConstraint(ctx *pulumi.Context,
	name string, args *ResourceUpdateConstraintArgs, opts ...pulumi.ResourceOption) (*ResourceUpdateConstraint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ResourceUpdateConstraint
	err := ctx.RegisterResource("cloudformation:ServiceCatalog:ResourceUpdateConstraint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceUpdateConstraint gets an existing ResourceUpdateConstraint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceUpdateConstraint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceUpdateConstraintState, opts ...pulumi.ResourceOption) (*ResourceUpdateConstraint, error) {
	var resource ResourceUpdateConstraint
	err := ctx.ReadResource("cloudformation:ServiceCatalog:ResourceUpdateConstraint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceUpdateConstraint resources.
type resourceUpdateConstraintState struct {
	// The attributes associated with the resource
	Attributes *ResourceUpdateConstraintAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ResourceUpdateConstraintProperties `pulumi:"properties"`
}

type ResourceUpdateConstraintState struct {
	// The attributes associated with the resource
	Attributes ResourceUpdateConstraintAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResourceUpdateConstraintPropertiesPtrInput
}

func (ResourceUpdateConstraintState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceUpdateConstraintState)(nil)).Elem()
}

type resourceUpdateConstraintArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResourceUpdateConstraintProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ResourceUpdateConstraint resource.
type ResourceUpdateConstraintArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResourceUpdateConstraintPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ResourceUpdateConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceUpdateConstraintArgs)(nil)).Elem()
}

type ResourceUpdateConstraintInput interface {
	pulumi.Input

	ToResourceUpdateConstraintOutput() ResourceUpdateConstraintOutput
	ToResourceUpdateConstraintOutputWithContext(ctx context.Context) ResourceUpdateConstraintOutput
}

func (*ResourceUpdateConstraint) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceUpdateConstraint)(nil))
}

func (i *ResourceUpdateConstraint) ToResourceUpdateConstraintOutput() ResourceUpdateConstraintOutput {
	return i.ToResourceUpdateConstraintOutputWithContext(context.Background())
}

func (i *ResourceUpdateConstraint) ToResourceUpdateConstraintOutputWithContext(ctx context.Context) ResourceUpdateConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceUpdateConstraintOutput)
}

type ResourceUpdateConstraintOutput struct {
	*pulumi.OutputState
}

func (ResourceUpdateConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceUpdateConstraint)(nil))
}

func (o ResourceUpdateConstraintOutput) ToResourceUpdateConstraintOutput() ResourceUpdateConstraintOutput {
	return o
}

func (o ResourceUpdateConstraintOutput) ToResourceUpdateConstraintOutputWithContext(ctx context.Context) ResourceUpdateConstraintOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceUpdateConstraintOutput{})
}
