// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html
type FunctionDefinition struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes FunctionDefinitionAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties FunctionDefinitionPropertiesOutput `pulumi:"properties"`
}

// NewFunctionDefinition registers a new resource with the given unique name, arguments, and options.
func NewFunctionDefinition(ctx *pulumi.Context,
	name string, args *FunctionDefinitionArgs, opts ...pulumi.ResourceOption) (*FunctionDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource FunctionDefinition
	err := ctx.RegisterResource("cloudformation:Greengrass:FunctionDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionDefinition gets an existing FunctionDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionDefinitionState, opts ...pulumi.ResourceOption) (*FunctionDefinition, error) {
	var resource FunctionDefinition
	err := ctx.ReadResource("cloudformation:Greengrass:FunctionDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionDefinition resources.
type functionDefinitionState struct {
	// The attributes associated with the resource
	Attributes *FunctionDefinitionAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *FunctionDefinitionProperties `pulumi:"properties"`
}

type FunctionDefinitionState struct {
	// The attributes associated with the resource
	Attributes FunctionDefinitionAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties FunctionDefinitionPropertiesPtrInput
}

func (FunctionDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDefinitionState)(nil)).Elem()
}

type functionDefinitionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties FunctionDefinitionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a FunctionDefinition resource.
type FunctionDefinitionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties FunctionDefinitionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (FunctionDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionDefinitionArgs)(nil)).Elem()
}

type FunctionDefinitionInput interface {
	pulumi.Input

	ToFunctionDefinitionOutput() FunctionDefinitionOutput
	ToFunctionDefinitionOutputWithContext(ctx context.Context) FunctionDefinitionOutput
}

func (*FunctionDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionDefinition)(nil))
}

func (i *FunctionDefinition) ToFunctionDefinitionOutput() FunctionDefinitionOutput {
	return i.ToFunctionDefinitionOutputWithContext(context.Background())
}

func (i *FunctionDefinition) ToFunctionDefinitionOutputWithContext(ctx context.Context) FunctionDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FunctionDefinitionOutput)
}

type FunctionDefinitionOutput struct {
	*pulumi.OutputState
}

func (FunctionDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionDefinition)(nil))
}

func (o FunctionDefinitionOutput) ToFunctionDefinitionOutput() FunctionDefinitionOutput {
	return o
}

func (o FunctionDefinitionOutput) ToFunctionDefinitionOutputWithContext(ctx context.Context) FunctionDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FunctionDefinitionOutput{})
}
