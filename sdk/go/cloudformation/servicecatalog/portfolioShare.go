// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-portfolioshare.html
type PortfolioShare struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes PortfolioShareAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties PortfolioSharePropertiesOutput `pulumi:"properties"`
}

// NewPortfolioShare registers a new resource with the given unique name, arguments, and options.
func NewPortfolioShare(ctx *pulumi.Context,
	name string, args *PortfolioShareArgs, opts ...pulumi.ResourceOption) (*PortfolioShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource PortfolioShare
	err := ctx.RegisterResource("cloudformation:ServiceCatalog:PortfolioShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPortfolioShare gets an existing PortfolioShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPortfolioShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PortfolioShareState, opts ...pulumi.ResourceOption) (*PortfolioShare, error) {
	var resource PortfolioShare
	err := ctx.ReadResource("cloudformation:ServiceCatalog:PortfolioShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PortfolioShare resources.
type portfolioShareState struct {
	// The attributes associated with the resource
	Attributes *PortfolioShareAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *PortfolioShareProperties `pulumi:"properties"`
}

type PortfolioShareState struct {
	// The attributes associated with the resource
	Attributes PortfolioShareAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties PortfolioSharePropertiesPtrInput
}

func (PortfolioShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioShareState)(nil)).Elem()
}

type portfolioShareArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties PortfolioShareProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a PortfolioShare resource.
type PortfolioShareArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties PortfolioSharePropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (PortfolioShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*portfolioShareArgs)(nil)).Elem()
}

type PortfolioShareInput interface {
	pulumi.Input

	ToPortfolioShareOutput() PortfolioShareOutput
	ToPortfolioShareOutputWithContext(ctx context.Context) PortfolioShareOutput
}

func (*PortfolioShare) ElementType() reflect.Type {
	return reflect.TypeOf((*PortfolioShare)(nil))
}

func (i *PortfolioShare) ToPortfolioShareOutput() PortfolioShareOutput {
	return i.ToPortfolioShareOutputWithContext(context.Background())
}

func (i *PortfolioShare) ToPortfolioShareOutputWithContext(ctx context.Context) PortfolioShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PortfolioShareOutput)
}

type PortfolioShareOutput struct {
	*pulumi.OutputState
}

func (PortfolioShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PortfolioShare)(nil))
}

func (o PortfolioShareOutput) ToPortfolioShareOutput() PortfolioShareOutput {
	return o
}

func (o PortfolioShareOutput) ToPortfolioShareOutputWithContext(ctx context.Context) PortfolioShareOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PortfolioShareOutput{})
}
