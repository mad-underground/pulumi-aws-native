// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html
type OriginRequestPolicy struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes OriginRequestPolicyAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties OriginRequestPolicyPropertiesOutput `pulumi:"properties"`
}

// NewOriginRequestPolicy registers a new resource with the given unique name, arguments, and options.
func NewOriginRequestPolicy(ctx *pulumi.Context,
	name string, args *OriginRequestPolicyArgs, opts ...pulumi.ResourceOption) (*OriginRequestPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource OriginRequestPolicy
	err := ctx.RegisterResource("cloudformation:CloudFront:OriginRequestPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOriginRequestPolicy gets an existing OriginRequestPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOriginRequestPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OriginRequestPolicyState, opts ...pulumi.ResourceOption) (*OriginRequestPolicy, error) {
	var resource OriginRequestPolicy
	err := ctx.ReadResource("cloudformation:CloudFront:OriginRequestPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OriginRequestPolicy resources.
type originRequestPolicyState struct {
	// The attributes associated with the resource
	Attributes *OriginRequestPolicyAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *OriginRequestPolicyProperties `pulumi:"properties"`
}

type OriginRequestPolicyState struct {
	// The attributes associated with the resource
	Attributes OriginRequestPolicyAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties OriginRequestPolicyPropertiesPtrInput
}

func (OriginRequestPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*originRequestPolicyState)(nil)).Elem()
}

type originRequestPolicyArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties OriginRequestPolicyProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a OriginRequestPolicy resource.
type OriginRequestPolicyArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties OriginRequestPolicyPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (OriginRequestPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*originRequestPolicyArgs)(nil)).Elem()
}

type OriginRequestPolicyInput interface {
	pulumi.Input

	ToOriginRequestPolicyOutput() OriginRequestPolicyOutput
	ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput
}

func (*OriginRequestPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*OriginRequestPolicy)(nil))
}

func (i *OriginRequestPolicy) ToOriginRequestPolicyOutput() OriginRequestPolicyOutput {
	return i.ToOriginRequestPolicyOutputWithContext(context.Background())
}

func (i *OriginRequestPolicy) ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OriginRequestPolicyOutput)
}

type OriginRequestPolicyOutput struct {
	*pulumi.OutputState
}

func (OriginRequestPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OriginRequestPolicy)(nil))
}

func (o OriginRequestPolicyOutput) ToOriginRequestPolicyOutput() OriginRequestPolicyOutput {
	return o
}

func (o OriginRequestPolicyOutput) ToOriginRequestPolicyOutputWithContext(ctx context.Context) OriginRequestPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OriginRequestPolicyOutput{})
}
