// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-findingsfilter.html
type FindingsFilter struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes FindingsFilterAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties FindingsFilterPropertiesOutput `pulumi:"properties"`
}

// NewFindingsFilter registers a new resource with the given unique name, arguments, and options.
func NewFindingsFilter(ctx *pulumi.Context,
	name string, args *FindingsFilterArgs, opts ...pulumi.ResourceOption) (*FindingsFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource FindingsFilter
	err := ctx.RegisterResource("cloudformation:Macie:FindingsFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFindingsFilter gets an existing FindingsFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFindingsFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FindingsFilterState, opts ...pulumi.ResourceOption) (*FindingsFilter, error) {
	var resource FindingsFilter
	err := ctx.ReadResource("cloudformation:Macie:FindingsFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FindingsFilter resources.
type findingsFilterState struct {
	// The attributes associated with the resource
	Attributes *FindingsFilterAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *FindingsFilterProperties `pulumi:"properties"`
}

type FindingsFilterState struct {
	// The attributes associated with the resource
	Attributes FindingsFilterAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties FindingsFilterPropertiesPtrInput
}

func (FindingsFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*findingsFilterState)(nil)).Elem()
}

type findingsFilterArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties FindingsFilterProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a FindingsFilter resource.
type FindingsFilterArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties FindingsFilterPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (FindingsFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*findingsFilterArgs)(nil)).Elem()
}

type FindingsFilterInput interface {
	pulumi.Input

	ToFindingsFilterOutput() FindingsFilterOutput
	ToFindingsFilterOutputWithContext(ctx context.Context) FindingsFilterOutput
}

func (*FindingsFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*FindingsFilter)(nil))
}

func (i *FindingsFilter) ToFindingsFilterOutput() FindingsFilterOutput {
	return i.ToFindingsFilterOutputWithContext(context.Background())
}

func (i *FindingsFilter) ToFindingsFilterOutputWithContext(ctx context.Context) FindingsFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FindingsFilterOutput)
}

type FindingsFilterOutput struct {
	*pulumi.OutputState
}

func (FindingsFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FindingsFilter)(nil))
}

func (o FindingsFilterOutput) ToFindingsFilterOutput() FindingsFilterOutput {
	return o
}

func (o FindingsFilterOutput) ToFindingsFilterOutputWithContext(ctx context.Context) FindingsFilterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FindingsFilterOutput{})
}
