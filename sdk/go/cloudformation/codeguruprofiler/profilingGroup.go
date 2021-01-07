// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codeguruprofiler

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeguruprofiler-profilinggroup.html
type ProfilingGroup struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ProfilingGroupAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ProfilingGroupPropertiesOutput `pulumi:"properties"`
}

// NewProfilingGroup registers a new resource with the given unique name, arguments, and options.
func NewProfilingGroup(ctx *pulumi.Context,
	name string, args *ProfilingGroupArgs, opts ...pulumi.ResourceOption) (*ProfilingGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ProfilingGroup
	err := ctx.RegisterResource("cloudformation:CodeGuruProfiler:ProfilingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfilingGroup gets an existing ProfilingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfilingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfilingGroupState, opts ...pulumi.ResourceOption) (*ProfilingGroup, error) {
	var resource ProfilingGroup
	err := ctx.ReadResource("cloudformation:CodeGuruProfiler:ProfilingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProfilingGroup resources.
type profilingGroupState struct {
	// The attributes associated with the resource
	Attributes *ProfilingGroupAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ProfilingGroupProperties `pulumi:"properties"`
}

type ProfilingGroupState struct {
	// The attributes associated with the resource
	Attributes ProfilingGroupAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ProfilingGroupPropertiesPtrInput
}

func (ProfilingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*profilingGroupState)(nil)).Elem()
}

type profilingGroupArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ProfilingGroupProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ProfilingGroup resource.
type ProfilingGroupArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ProfilingGroupPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ProfilingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profilingGroupArgs)(nil)).Elem()
}

type ProfilingGroupInput interface {
	pulumi.Input

	ToProfilingGroupOutput() ProfilingGroupOutput
	ToProfilingGroupOutputWithContext(ctx context.Context) ProfilingGroupOutput
}

func (*ProfilingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfilingGroup)(nil))
}

func (i *ProfilingGroup) ToProfilingGroupOutput() ProfilingGroupOutput {
	return i.ToProfilingGroupOutputWithContext(context.Background())
}

func (i *ProfilingGroup) ToProfilingGroupOutputWithContext(ctx context.Context) ProfilingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProfilingGroupOutput)
}

type ProfilingGroupOutput struct {
	*pulumi.OutputState
}

func (ProfilingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProfilingGroup)(nil))
}

func (o ProfilingGroupOutput) ToProfilingGroupOutput() ProfilingGroupOutput {
	return o
}

func (o ProfilingGroupOutput) ToProfilingGroupOutputWithContext(ctx context.Context) ProfilingGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ProfilingGroupOutput{})
}