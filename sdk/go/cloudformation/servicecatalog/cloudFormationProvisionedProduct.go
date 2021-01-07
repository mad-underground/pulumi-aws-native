// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-cloudformationprovisionedproduct.html
type CloudFormationProvisionedProduct struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes CloudFormationProvisionedProductAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties CloudFormationProvisionedProductPropertiesOutput `pulumi:"properties"`
}

// NewCloudFormationProvisionedProduct registers a new resource with the given unique name, arguments, and options.
func NewCloudFormationProvisionedProduct(ctx *pulumi.Context,
	name string, args *CloudFormationProvisionedProductArgs, opts ...pulumi.ResourceOption) (*CloudFormationProvisionedProduct, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource CloudFormationProvisionedProduct
	err := ctx.RegisterResource("cloudformation:ServiceCatalog:CloudFormationProvisionedProduct", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudFormationProvisionedProduct gets an existing CloudFormationProvisionedProduct resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudFormationProvisionedProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudFormationProvisionedProductState, opts ...pulumi.ResourceOption) (*CloudFormationProvisionedProduct, error) {
	var resource CloudFormationProvisionedProduct
	err := ctx.ReadResource("cloudformation:ServiceCatalog:CloudFormationProvisionedProduct", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudFormationProvisionedProduct resources.
type cloudFormationProvisionedProductState struct {
	// The attributes associated with the resource
	Attributes *CloudFormationProvisionedProductAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *CloudFormationProvisionedProductProperties `pulumi:"properties"`
}

type CloudFormationProvisionedProductState struct {
	// The attributes associated with the resource
	Attributes CloudFormationProvisionedProductAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties CloudFormationProvisionedProductPropertiesPtrInput
}

func (CloudFormationProvisionedProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudFormationProvisionedProductState)(nil)).Elem()
}

type cloudFormationProvisionedProductArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties CloudFormationProvisionedProductProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a CloudFormationProvisionedProduct resource.
type CloudFormationProvisionedProductArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties CloudFormationProvisionedProductPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (CloudFormationProvisionedProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudFormationProvisionedProductArgs)(nil)).Elem()
}

type CloudFormationProvisionedProductInput interface {
	pulumi.Input

	ToCloudFormationProvisionedProductOutput() CloudFormationProvisionedProductOutput
	ToCloudFormationProvisionedProductOutputWithContext(ctx context.Context) CloudFormationProvisionedProductOutput
}

func (*CloudFormationProvisionedProduct) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudFormationProvisionedProduct)(nil))
}

func (i *CloudFormationProvisionedProduct) ToCloudFormationProvisionedProductOutput() CloudFormationProvisionedProductOutput {
	return i.ToCloudFormationProvisionedProductOutputWithContext(context.Background())
}

func (i *CloudFormationProvisionedProduct) ToCloudFormationProvisionedProductOutputWithContext(ctx context.Context) CloudFormationProvisionedProductOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudFormationProvisionedProductOutput)
}

type CloudFormationProvisionedProductOutput struct {
	*pulumi.OutputState
}

func (CloudFormationProvisionedProductOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudFormationProvisionedProduct)(nil))
}

func (o CloudFormationProvisionedProductOutput) ToCloudFormationProvisionedProductOutput() CloudFormationProvisionedProductOutput {
	return o
}

func (o CloudFormationProvisionedProductOutput) ToCloudFormationProvisionedProductOutputWithContext(ctx context.Context) CloudFormationProvisionedProductOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CloudFormationProvisionedProductOutput{})
}
