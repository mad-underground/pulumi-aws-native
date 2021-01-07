// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-moduledefaultversion.html
type ModuleDefaultVersion struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ModuleDefaultVersionAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ModuleDefaultVersionPropertiesOutput `pulumi:"properties"`
}

// NewModuleDefaultVersion registers a new resource with the given unique name, arguments, and options.
func NewModuleDefaultVersion(ctx *pulumi.Context,
	name string, args *ModuleDefaultVersionArgs, opts ...pulumi.ResourceOption) (*ModuleDefaultVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ModuleDefaultVersion
	err := ctx.RegisterResource("cloudformation:CloudFormation:ModuleDefaultVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModuleDefaultVersion gets an existing ModuleDefaultVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModuleDefaultVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModuleDefaultVersionState, opts ...pulumi.ResourceOption) (*ModuleDefaultVersion, error) {
	var resource ModuleDefaultVersion
	err := ctx.ReadResource("cloudformation:CloudFormation:ModuleDefaultVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModuleDefaultVersion resources.
type moduleDefaultVersionState struct {
	// The attributes associated with the resource
	Attributes *ModuleDefaultVersionAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ModuleDefaultVersionProperties `pulumi:"properties"`
}

type ModuleDefaultVersionState struct {
	// The attributes associated with the resource
	Attributes ModuleDefaultVersionAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ModuleDefaultVersionPropertiesPtrInput
}

func (ModuleDefaultVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleDefaultVersionState)(nil)).Elem()
}

type moduleDefaultVersionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ModuleDefaultVersionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ModuleDefaultVersion resource.
type ModuleDefaultVersionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ModuleDefaultVersionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ModuleDefaultVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moduleDefaultVersionArgs)(nil)).Elem()
}

type ModuleDefaultVersionInput interface {
	pulumi.Input

	ToModuleDefaultVersionOutput() ModuleDefaultVersionOutput
	ToModuleDefaultVersionOutputWithContext(ctx context.Context) ModuleDefaultVersionOutput
}

func (*ModuleDefaultVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleDefaultVersion)(nil))
}

func (i *ModuleDefaultVersion) ToModuleDefaultVersionOutput() ModuleDefaultVersionOutput {
	return i.ToModuleDefaultVersionOutputWithContext(context.Background())
}

func (i *ModuleDefaultVersion) ToModuleDefaultVersionOutputWithContext(ctx context.Context) ModuleDefaultVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleDefaultVersionOutput)
}

type ModuleDefaultVersionOutput struct {
	*pulumi.OutputState
}

func (ModuleDefaultVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleDefaultVersion)(nil))
}

func (o ModuleDefaultVersionOutput) ToModuleDefaultVersionOutput() ModuleDefaultVersionOutput {
	return o
}

func (o ModuleDefaultVersionOutput) ToModuleDefaultVersionOutputWithContext(ctx context.Context) ModuleDefaultVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ModuleDefaultVersionOutput{})
}
