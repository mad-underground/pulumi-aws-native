// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-organizationconformancepack.html
type OrganizationConformancePack struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes OrganizationConformancePackAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties OrganizationConformancePackPropertiesOutput `pulumi:"properties"`
}

// NewOrganizationConformancePack registers a new resource with the given unique name, arguments, and options.
func NewOrganizationConformancePack(ctx *pulumi.Context,
	name string, args *OrganizationConformancePackArgs, opts ...pulumi.ResourceOption) (*OrganizationConformancePack, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource OrganizationConformancePack
	err := ctx.RegisterResource("cloudformation:Configuration:OrganizationConformancePack", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationConformancePack gets an existing OrganizationConformancePack resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationConformancePack(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationConformancePackState, opts ...pulumi.ResourceOption) (*OrganizationConformancePack, error) {
	var resource OrganizationConformancePack
	err := ctx.ReadResource("cloudformation:Configuration:OrganizationConformancePack", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationConformancePack resources.
type organizationConformancePackState struct {
	// The attributes associated with the resource
	Attributes *OrganizationConformancePackAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *OrganizationConformancePackProperties `pulumi:"properties"`
}

type OrganizationConformancePackState struct {
	// The attributes associated with the resource
	Attributes OrganizationConformancePackAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties OrganizationConformancePackPropertiesPtrInput
}

func (OrganizationConformancePackState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationConformancePackState)(nil)).Elem()
}

type organizationConformancePackArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties OrganizationConformancePackProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a OrganizationConformancePack resource.
type OrganizationConformancePackArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties OrganizationConformancePackPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (OrganizationConformancePackArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationConformancePackArgs)(nil)).Elem()
}

type OrganizationConformancePackInput interface {
	pulumi.Input

	ToOrganizationConformancePackOutput() OrganizationConformancePackOutput
	ToOrganizationConformancePackOutputWithContext(ctx context.Context) OrganizationConformancePackOutput
}

func (*OrganizationConformancePack) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationConformancePack)(nil))
}

func (i *OrganizationConformancePack) ToOrganizationConformancePackOutput() OrganizationConformancePackOutput {
	return i.ToOrganizationConformancePackOutputWithContext(context.Background())
}

func (i *OrganizationConformancePack) ToOrganizationConformancePackOutputWithContext(ctx context.Context) OrganizationConformancePackOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationConformancePackOutput)
}

type OrganizationConformancePackOutput struct {
	*pulumi.OutputState
}

func (OrganizationConformancePackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationConformancePack)(nil))
}

func (o OrganizationConformancePackOutput) ToOrganizationConformancePackOutput() OrganizationConformancePackOutput {
	return o
}

func (o OrganizationConformancePackOutput) ToOrganizationConformancePackOutputWithContext(ctx context.Context) OrganizationConformancePackOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationConformancePackOutput{})
}
