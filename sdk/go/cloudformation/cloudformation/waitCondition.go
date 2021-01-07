// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudformation/sdk/go/cloudformation"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitcondition.html
type WaitCondition struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes WaitConditionAttributesOutput `pulumi:"attributes"`
	// The creation policy associated with the resource
	CreationPolicy cloudformation.CreationPolicyPtrOutput `pulumi:"creationPolicy"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties WaitConditionPropertiesOutput `pulumi:"properties"`
}

// NewWaitCondition registers a new resource with the given unique name, arguments, and options.
func NewWaitCondition(ctx *pulumi.Context,
	name string, args *WaitConditionArgs, opts ...pulumi.ResourceOption) (*WaitCondition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource WaitCondition
	err := ctx.RegisterResource("cloudformation:CloudFormation:WaitCondition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWaitCondition gets an existing WaitCondition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWaitCondition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WaitConditionState, opts ...pulumi.ResourceOption) (*WaitCondition, error) {
	var resource WaitCondition
	err := ctx.ReadResource("cloudformation:CloudFormation:WaitCondition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WaitCondition resources.
type waitConditionState struct {
	// The attributes associated with the resource
	Attributes *WaitConditionAttributes `pulumi:"attributes"`
	// The creation policy associated with the resource
	CreationPolicy *cloudformation.CreationPolicy `pulumi:"creationPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *WaitConditionProperties `pulumi:"properties"`
}

type WaitConditionState struct {
	// The attributes associated with the resource
	Attributes WaitConditionAttributesPtrInput
	// The creation policy associated with the resource
	CreationPolicy cloudformation.CreationPolicyPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties WaitConditionPropertiesPtrInput
}

func (WaitConditionState) ElementType() reflect.Type {
	return reflect.TypeOf((*waitConditionState)(nil)).Elem()
}

type waitConditionArgs struct {
	// The creation policy associated with the resource
	CreationPolicy *cloudformation.CreationPolicy `pulumi:"creationPolicy"`
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties WaitConditionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a WaitCondition resource.
type WaitConditionArgs struct {
	// The creation policy associated with the resource
	CreationPolicy cloudformation.CreationPolicyPtrInput
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties WaitConditionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (WaitConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*waitConditionArgs)(nil)).Elem()
}

type WaitConditionInput interface {
	pulumi.Input

	ToWaitConditionOutput() WaitConditionOutput
	ToWaitConditionOutputWithContext(ctx context.Context) WaitConditionOutput
}

func (*WaitCondition) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitCondition)(nil))
}

func (i *WaitCondition) ToWaitConditionOutput() WaitConditionOutput {
	return i.ToWaitConditionOutputWithContext(context.Background())
}

func (i *WaitCondition) ToWaitConditionOutputWithContext(ctx context.Context) WaitConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitConditionOutput)
}

type WaitConditionOutput struct {
	*pulumi.OutputState
}

func (WaitConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitCondition)(nil))
}

func (o WaitConditionOutput) ToWaitConditionOutput() WaitConditionOutput {
	return o
}

func (o WaitConditionOutput) ToWaitConditionOutputWithContext(ctx context.Context) WaitConditionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WaitConditionOutput{})
}