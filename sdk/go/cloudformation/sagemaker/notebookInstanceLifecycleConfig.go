// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html
type NotebookInstanceLifecycleConfig struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes NotebookInstanceLifecycleConfigAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties NotebookInstanceLifecycleConfigPropertiesOutput `pulumi:"properties"`
}

// NewNotebookInstanceLifecycleConfig registers a new resource with the given unique name, arguments, and options.
func NewNotebookInstanceLifecycleConfig(ctx *pulumi.Context,
	name string, args *NotebookInstanceLifecycleConfigArgs, opts ...pulumi.ResourceOption) (*NotebookInstanceLifecycleConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource NotebookInstanceLifecycleConfig
	err := ctx.RegisterResource("cloudformation:SageMaker:NotebookInstanceLifecycleConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookInstanceLifecycleConfig gets an existing NotebookInstanceLifecycleConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookInstanceLifecycleConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookInstanceLifecycleConfigState, opts ...pulumi.ResourceOption) (*NotebookInstanceLifecycleConfig, error) {
	var resource NotebookInstanceLifecycleConfig
	err := ctx.ReadResource("cloudformation:SageMaker:NotebookInstanceLifecycleConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookInstanceLifecycleConfig resources.
type notebookInstanceLifecycleConfigState struct {
	// The attributes associated with the resource
	Attributes *NotebookInstanceLifecycleConfigAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *NotebookInstanceLifecycleConfigProperties `pulumi:"properties"`
}

type NotebookInstanceLifecycleConfigState struct {
	// The attributes associated with the resource
	Attributes NotebookInstanceLifecycleConfigAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties NotebookInstanceLifecycleConfigPropertiesPtrInput
}

func (NotebookInstanceLifecycleConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceLifecycleConfigState)(nil)).Elem()
}

type notebookInstanceLifecycleConfigArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties NotebookInstanceLifecycleConfigProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a NotebookInstanceLifecycleConfig resource.
type NotebookInstanceLifecycleConfigArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties NotebookInstanceLifecycleConfigPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (NotebookInstanceLifecycleConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookInstanceLifecycleConfigArgs)(nil)).Elem()
}

type NotebookInstanceLifecycleConfigInput interface {
	pulumi.Input

	ToNotebookInstanceLifecycleConfigOutput() NotebookInstanceLifecycleConfigOutput
	ToNotebookInstanceLifecycleConfigOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigOutput
}

func (*NotebookInstanceLifecycleConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstanceLifecycleConfig)(nil))
}

func (i *NotebookInstanceLifecycleConfig) ToNotebookInstanceLifecycleConfigOutput() NotebookInstanceLifecycleConfigOutput {
	return i.ToNotebookInstanceLifecycleConfigOutputWithContext(context.Background())
}

func (i *NotebookInstanceLifecycleConfig) ToNotebookInstanceLifecycleConfigOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookInstanceLifecycleConfigOutput)
}

type NotebookInstanceLifecycleConfigOutput struct {
	*pulumi.OutputState
}

func (NotebookInstanceLifecycleConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookInstanceLifecycleConfig)(nil))
}

func (o NotebookInstanceLifecycleConfigOutput) ToNotebookInstanceLifecycleConfigOutput() NotebookInstanceLifecycleConfigOutput {
	return o
}

func (o NotebookInstanceLifecycleConfigOutput) ToNotebookInstanceLifecycleConfigOutputWithContext(ctx context.Context) NotebookInstanceLifecycleConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NotebookInstanceLifecycleConfigOutput{})
}