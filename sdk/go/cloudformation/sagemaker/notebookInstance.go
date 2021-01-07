// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html
type NotebookInstance struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes NotebookInstanceAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties NotebookInstancePropertiesOutput `pulumi:"properties"`
}

// NewNotebookInstance registers a new resource with the given unique name, arguments, and options.
func NewNotebookInstance(ctx *pulumi.Context,
	name string, args *NotebookInstanceArgs, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource NotebookInstance
	err := ctx.RegisterResource("cloudformation:SageMaker:NotebookInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookInstance gets an existing NotebookInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookInstanceState, opts ...pulumi.ResourceOption) (*NotebookInstance, error) {
	var resource NotebookInstance
	err := ctx.ReadResource("cloudformation:SageMaker:NotebookInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookInstance resources.
type notebookInstanceState struct {
	// The attributes associated with the resource
	Attributes *NotebookInstanceAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *NotebookInstanceProperties `pulumi:"properties"`
}

type NotebookInstanceState struct {
	// The attributes associated with the resource
	Attributes NotebookInstanceAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties NotebookInstancePropertiesPtrInput
}

func (NotebookInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceState)(nil)).Elem()
}

type notebookInstanceArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties NotebookInstanceProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a NotebookInstance resource.
type NotebookInstanceArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties NotebookInstancePropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (NotebookInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceArgs)(nil)).Elem()
}

type NotebookInstanceInput interface {
	pulumi.Input

	ToNotebookInstanceOutput() NotebookInstanceOutput
	ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput
}

func (*NotebookInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstance)(nil))
}

func (i *NotebookInstance) ToNotebookInstanceOutput() NotebookInstanceOutput {
	return i.ToNotebookInstanceOutputWithContext(context.Background())
}

func (i *NotebookInstance) ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceOutput)
}

type NotebookInstanceOutput struct {
	*pulumi.OutputState
}

func (NotebookInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstance)(nil))
}

func (o NotebookInstanceOutput) ToNotebookInstanceOutput() NotebookInstanceOutput {
	return o
}

func (o NotebookInstanceOutput) ToNotebookInstanceOutputWithContext(ctx context.Context) NotebookInstanceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotebookInstanceOutput{})
}