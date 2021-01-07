// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip-association.html
type EIPAssociation struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes EIPAssociationAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties EIPAssociationPropertiesOutput `pulumi:"properties"`
}

// NewEIPAssociation registers a new resource with the given unique name, arguments, and options.
func NewEIPAssociation(ctx *pulumi.Context,
	name string, args *EIPAssociationArgs, opts ...pulumi.ResourceOption) (*EIPAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource EIPAssociation
	err := ctx.RegisterResource("cloudformation:EC2:EIPAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEIPAssociation gets an existing EIPAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEIPAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EIPAssociationState, opts ...pulumi.ResourceOption) (*EIPAssociation, error) {
	var resource EIPAssociation
	err := ctx.ReadResource("cloudformation:EC2:EIPAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EIPAssociation resources.
type eipassociationState struct {
	// The attributes associated with the resource
	Attributes *EIPAssociationAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *EIPAssociationProperties `pulumi:"properties"`
}

type EIPAssociationState struct {
	// The attributes associated with the resource
	Attributes EIPAssociationAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties EIPAssociationPropertiesPtrInput
}

func (EIPAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipassociationState)(nil)).Elem()
}

type eipassociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties EIPAssociationProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a EIPAssociation resource.
type EIPAssociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties EIPAssociationPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (EIPAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipassociationArgs)(nil)).Elem()
}

type EIPAssociationInput interface {
	pulumi.Input

	ToEIPAssociationOutput() EIPAssociationOutput
	ToEIPAssociationOutputWithContext(ctx context.Context) EIPAssociationOutput
}

func (*EIPAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*EIPAssociation)(nil))
}

func (i *EIPAssociation) ToEIPAssociationOutput() EIPAssociationOutput {
	return i.ToEIPAssociationOutputWithContext(context.Background())
}

func (i *EIPAssociation) ToEIPAssociationOutputWithContext(ctx context.Context) EIPAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EIPAssociationOutput)
}

type EIPAssociationOutput struct {
	*pulumi.OutputState
}

func (EIPAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EIPAssociation)(nil))
}

func (o EIPAssociationOutput) ToEIPAssociationOutput() EIPAssociationOutput {
	return o
}

func (o EIPAssociationOutput) ToEIPAssociationOutputWithContext(ctx context.Context) EIPAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EIPAssociationOutput{})
}
