// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-segment.html
type Segment struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes SegmentAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties SegmentPropertiesOutput `pulumi:"properties"`
}

// NewSegment registers a new resource with the given unique name, arguments, and options.
func NewSegment(ctx *pulumi.Context,
	name string, args *SegmentArgs, opts ...pulumi.ResourceOption) (*Segment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Segment
	err := ctx.RegisterResource("cloudformation:Pinpoint:Segment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSegment gets an existing Segment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SegmentState, opts ...pulumi.ResourceOption) (*Segment, error) {
	var resource Segment
	err := ctx.ReadResource("cloudformation:Pinpoint:Segment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Segment resources.
type segmentState struct {
	// The attributes associated with the resource
	Attributes *SegmentAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *SegmentProperties `pulumi:"properties"`
}

type SegmentState struct {
	// The attributes associated with the resource
	Attributes SegmentAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties SegmentPropertiesPtrInput
}

func (SegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentState)(nil)).Elem()
}

type segmentArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties SegmentProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Segment resource.
type SegmentArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties SegmentPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (SegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentArgs)(nil)).Elem()
}

type SegmentInput interface {
	pulumi.Input

	ToSegmentOutput() SegmentOutput
	ToSegmentOutputWithContext(ctx context.Context) SegmentOutput
}

func (*Segment) ElementType() reflect.Type {
	return reflect.TypeOf((*Segment)(nil))
}

func (i *Segment) ToSegmentOutput() SegmentOutput {
	return i.ToSegmentOutputWithContext(context.Background())
}

func (i *Segment) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentOutput)
}

type SegmentOutput struct {
	*pulumi.OutputState
}

func (SegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Segment)(nil))
}

func (o SegmentOutput) ToSegmentOutput() SegmentOutput {
	return o
}

func (o SegmentOutput) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SegmentOutput{})
}