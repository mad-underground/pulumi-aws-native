// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-trafficmirrortarget.html
type TrafficMirrorTarget struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes TrafficMirrorTargetAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties TrafficMirrorTargetPropertiesOutput `pulumi:"properties"`
}

// NewTrafficMirrorTarget registers a new resource with the given unique name, arguments, and options.
func NewTrafficMirrorTarget(ctx *pulumi.Context,
	name string, args *TrafficMirrorTargetArgs, opts ...pulumi.ResourceOption) (*TrafficMirrorTarget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource TrafficMirrorTarget
	err := ctx.RegisterResource("cloudformation:EC2:TrafficMirrorTarget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrafficMirrorTarget gets an existing TrafficMirrorTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrafficMirrorTarget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrafficMirrorTargetState, opts ...pulumi.ResourceOption) (*TrafficMirrorTarget, error) {
	var resource TrafficMirrorTarget
	err := ctx.ReadResource("cloudformation:EC2:TrafficMirrorTarget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrafficMirrorTarget resources.
type trafficMirrorTargetState struct {
	// The attributes associated with the resource
	Attributes *TrafficMirrorTargetAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *TrafficMirrorTargetProperties `pulumi:"properties"`
}

type TrafficMirrorTargetState struct {
	// The attributes associated with the resource
	Attributes TrafficMirrorTargetAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties TrafficMirrorTargetPropertiesPtrInput
}

func (TrafficMirrorTargetState) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorTargetState)(nil)).Elem()
}

type trafficMirrorTargetArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties TrafficMirrorTargetProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a TrafficMirrorTarget resource.
type TrafficMirrorTargetArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties TrafficMirrorTargetPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (TrafficMirrorTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trafficMirrorTargetArgs)(nil)).Elem()
}

type TrafficMirrorTargetInput interface {
	pulumi.Input

	ToTrafficMirrorTargetOutput() TrafficMirrorTargetOutput
	ToTrafficMirrorTargetOutputWithContext(ctx context.Context) TrafficMirrorTargetOutput
}

func (*TrafficMirrorTarget) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficMirrorTarget)(nil))
}

func (i *TrafficMirrorTarget) ToTrafficMirrorTargetOutput() TrafficMirrorTargetOutput {
	return i.ToTrafficMirrorTargetOutputWithContext(context.Background())
}

func (i *TrafficMirrorTarget) ToTrafficMirrorTargetOutputWithContext(ctx context.Context) TrafficMirrorTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficMirrorTargetOutput)
}

type TrafficMirrorTargetOutput struct {
	*pulumi.OutputState
}

func (TrafficMirrorTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficMirrorTarget)(nil))
}

func (o TrafficMirrorTargetOutput) ToTrafficMirrorTargetOutput() TrafficMirrorTargetOutput {
	return o
}

func (o TrafficMirrorTargetOutput) ToTrafficMirrorTargetOutputWithContext(ctx context.Context) TrafficMirrorTargetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TrafficMirrorTargetOutput{})
}