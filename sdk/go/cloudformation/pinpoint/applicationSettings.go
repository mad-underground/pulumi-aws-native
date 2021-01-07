// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpoint-applicationsettings.html
type ApplicationSettings struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ApplicationSettingsAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ApplicationSettingsPropertiesOutput `pulumi:"properties"`
}

// NewApplicationSettings registers a new resource with the given unique name, arguments, and options.
func NewApplicationSettings(ctx *pulumi.Context,
	name string, args *ApplicationSettingsArgs, opts ...pulumi.ResourceOption) (*ApplicationSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ApplicationSettings
	err := ctx.RegisterResource("cloudformation:Pinpoint:ApplicationSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSettings gets an existing ApplicationSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSettingsState, opts ...pulumi.ResourceOption) (*ApplicationSettings, error) {
	var resource ApplicationSettings
	err := ctx.ReadResource("cloudformation:Pinpoint:ApplicationSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSettings resources.
type applicationSettingsState struct {
	// The attributes associated with the resource
	Attributes *ApplicationSettingsAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ApplicationSettingsProperties `pulumi:"properties"`
}

type ApplicationSettingsState struct {
	// The attributes associated with the resource
	Attributes ApplicationSettingsAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ApplicationSettingsPropertiesPtrInput
}

func (ApplicationSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSettingsState)(nil)).Elem()
}

type applicationSettingsArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ApplicationSettingsProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ApplicationSettings resource.
type ApplicationSettingsArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ApplicationSettingsPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ApplicationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSettingsArgs)(nil)).Elem()
}

type ApplicationSettingsInput interface {
	pulumi.Input

	ToApplicationSettingsOutput() ApplicationSettingsOutput
	ToApplicationSettingsOutputWithContext(ctx context.Context) ApplicationSettingsOutput
}

func (*ApplicationSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSettings)(nil))
}

func (i *ApplicationSettings) ToApplicationSettingsOutput() ApplicationSettingsOutput {
	return i.ToApplicationSettingsOutputWithContext(context.Background())
}

func (i *ApplicationSettings) ToApplicationSettingsOutputWithContext(ctx context.Context) ApplicationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSettingsOutput)
}

type ApplicationSettingsOutput struct {
	*pulumi.OutputState
}

func (ApplicationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationSettings)(nil))
}

func (o ApplicationSettingsOutput) ToApplicationSettingsOutput() ApplicationSettingsOutput {
	return o
}

func (o ApplicationSettingsOutput) ToApplicationSettingsOutputWithContext(ctx context.Context) ApplicationSettingsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApplicationSettingsOutput{})
}