// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverrule.html
type ResolverRule struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ResolverRuleAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResolverRulePropertiesOutput `pulumi:"properties"`
}

// NewResolverRule registers a new resource with the given unique name, arguments, and options.
func NewResolverRule(ctx *pulumi.Context,
	name string, args *ResolverRuleArgs, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ResolverRule
	err := ctx.RegisterResource("cloudformation:Route53Resolver:ResolverRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverRule gets an existing ResolverRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverRuleState, opts ...pulumi.ResourceOption) (*ResolverRule, error) {
	var resource ResolverRule
	err := ctx.ReadResource("cloudformation:Route53Resolver:ResolverRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverRule resources.
type resolverRuleState struct {
	// The attributes associated with the resource
	Attributes *ResolverRuleAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ResolverRuleProperties `pulumi:"properties"`
}

type ResolverRuleState struct {
	// The attributes associated with the resource
	Attributes ResolverRuleAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResolverRulePropertiesPtrInput
}

func (ResolverRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleState)(nil)).Elem()
}

type resolverRuleArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResolverRuleProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ResolverRule resource.
type ResolverRuleArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResolverRulePropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ResolverRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleArgs)(nil)).Elem()
}

type ResolverRuleInput interface {
	pulumi.Input

	ToResolverRuleOutput() ResolverRuleOutput
	ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput
}

func (*ResolverRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRule)(nil))
}

func (i *ResolverRule) ToResolverRuleOutput() ResolverRuleOutput {
	return i.ToResolverRuleOutputWithContext(context.Background())
}

func (i *ResolverRule) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleOutput)
}

type ResolverRuleOutput struct {
	*pulumi.OutputState
}

func (ResolverRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRule)(nil))
}

func (o ResolverRuleOutput) ToResolverRuleOutput() ResolverRuleOutput {
	return o
}

func (o ResolverRuleOutput) ToResolverRuleOutputWithContext(ctx context.Context) ResolverRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResolverRuleOutput{})
}
