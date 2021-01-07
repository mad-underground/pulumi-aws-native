// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-carriergateway.html
type CarrierGateway struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes CarrierGatewayAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties CarrierGatewayPropertiesOutput `pulumi:"properties"`
}

// NewCarrierGateway registers a new resource with the given unique name, arguments, and options.
func NewCarrierGateway(ctx *pulumi.Context,
	name string, args *CarrierGatewayArgs, opts ...pulumi.ResourceOption) (*CarrierGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource CarrierGateway
	err := ctx.RegisterResource("cloudformation:EC2:CarrierGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCarrierGateway gets an existing CarrierGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCarrierGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CarrierGatewayState, opts ...pulumi.ResourceOption) (*CarrierGateway, error) {
	var resource CarrierGateway
	err := ctx.ReadResource("cloudformation:EC2:CarrierGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CarrierGateway resources.
type carrierGatewayState struct {
	// The attributes associated with the resource
	Attributes *CarrierGatewayAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *CarrierGatewayProperties `pulumi:"properties"`
}

type CarrierGatewayState struct {
	// The attributes associated with the resource
	Attributes CarrierGatewayAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties CarrierGatewayPropertiesPtrInput
}

func (CarrierGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*carrierGatewayState)(nil)).Elem()
}

type carrierGatewayArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties CarrierGatewayProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a CarrierGateway resource.
type CarrierGatewayArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties CarrierGatewayPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (CarrierGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*carrierGatewayArgs)(nil)).Elem()
}

type CarrierGatewayInput interface {
	pulumi.Input

	ToCarrierGatewayOutput() CarrierGatewayOutput
	ToCarrierGatewayOutputWithContext(ctx context.Context) CarrierGatewayOutput
}

func (*CarrierGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*CarrierGateway)(nil))
}

func (i *CarrierGateway) ToCarrierGatewayOutput() CarrierGatewayOutput {
	return i.ToCarrierGatewayOutputWithContext(context.Background())
}

func (i *CarrierGateway) ToCarrierGatewayOutputWithContext(ctx context.Context) CarrierGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CarrierGatewayOutput)
}

type CarrierGatewayOutput struct {
	*pulumi.OutputState
}

func (CarrierGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CarrierGateway)(nil))
}

func (o CarrierGatewayOutput) ToCarrierGatewayOutput() CarrierGatewayOutput {
	return o
}

func (o CarrierGatewayOutput) ToCarrierGatewayOutputWithContext(ctx context.Context) CarrierGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CarrierGatewayOutput{})
}