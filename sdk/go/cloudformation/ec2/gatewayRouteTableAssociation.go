// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-gatewayroutetableassociation.html
type GatewayRouteTableAssociation struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes GatewayRouteTableAssociationAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GatewayRouteTableAssociationPropertiesOutput `pulumi:"properties"`
}

// NewGatewayRouteTableAssociation registers a new resource with the given unique name, arguments, and options.
func NewGatewayRouteTableAssociation(ctx *pulumi.Context,
	name string, args *GatewayRouteTableAssociationArgs, opts ...pulumi.ResourceOption) (*GatewayRouteTableAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource GatewayRouteTableAssociation
	err := ctx.RegisterResource("cloudformation:EC2:GatewayRouteTableAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayRouteTableAssociation gets an existing GatewayRouteTableAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayRouteTableAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayRouteTableAssociationState, opts ...pulumi.ResourceOption) (*GatewayRouteTableAssociation, error) {
	var resource GatewayRouteTableAssociation
	err := ctx.ReadResource("cloudformation:EC2:GatewayRouteTableAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayRouteTableAssociation resources.
type gatewayRouteTableAssociationState struct {
	// The attributes associated with the resource
	Attributes *GatewayRouteTableAssociationAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *GatewayRouteTableAssociationProperties `pulumi:"properties"`
}

type GatewayRouteTableAssociationState struct {
	// The attributes associated with the resource
	Attributes GatewayRouteTableAssociationAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GatewayRouteTableAssociationPropertiesPtrInput
}

func (GatewayRouteTableAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteTableAssociationState)(nil)).Elem()
}

type gatewayRouteTableAssociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GatewayRouteTableAssociationProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a GatewayRouteTableAssociation resource.
type GatewayRouteTableAssociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GatewayRouteTableAssociationPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (GatewayRouteTableAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteTableAssociationArgs)(nil)).Elem()
}

type GatewayRouteTableAssociationInput interface {
	pulumi.Input

	ToGatewayRouteTableAssociationOutput() GatewayRouteTableAssociationOutput
	ToGatewayRouteTableAssociationOutputWithContext(ctx context.Context) GatewayRouteTableAssociationOutput
}

func (*GatewayRouteTableAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteTableAssociation)(nil))
}

func (i *GatewayRouteTableAssociation) ToGatewayRouteTableAssociationOutput() GatewayRouteTableAssociationOutput {
	return i.ToGatewayRouteTableAssociationOutputWithContext(context.Background())
}

func (i *GatewayRouteTableAssociation) ToGatewayRouteTableAssociationOutputWithContext(ctx context.Context) GatewayRouteTableAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteTableAssociationOutput)
}

type GatewayRouteTableAssociationOutput struct {
	*pulumi.OutputState
}

func (GatewayRouteTableAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRouteTableAssociation)(nil))
}

func (o GatewayRouteTableAssociationOutput) ToGatewayRouteTableAssociationOutput() GatewayRouteTableAssociationOutput {
	return o
}

func (o GatewayRouteTableAssociationOutput) ToGatewayRouteTableAssociationOutputWithContext(ctx context.Context) GatewayRouteTableAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayRouteTableAssociationOutput{})
}