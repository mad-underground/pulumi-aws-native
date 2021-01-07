// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securityhub

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-securityhub-hub.html
type Hub struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes HubAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties HubPropertiesOutput `pulumi:"properties"`
}

// NewHub registers a new resource with the given unique name, arguments, and options.
func NewHub(ctx *pulumi.Context,
	name string, args *HubArgs, opts ...pulumi.ResourceOption) (*Hub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Hub
	err := ctx.RegisterResource("cloudformation:SecurityHub:Hub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHub gets an existing Hub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubState, opts ...pulumi.ResourceOption) (*Hub, error) {
	var resource Hub
	err := ctx.ReadResource("cloudformation:SecurityHub:Hub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Hub resources.
type hubState struct {
	// The attributes associated with the resource
	Attributes *HubAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *HubProperties `pulumi:"properties"`
}

type HubState struct {
	// The attributes associated with the resource
	Attributes HubAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties HubPropertiesPtrInput
}

func (HubState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubState)(nil)).Elem()
}

type hubArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties HubProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Hub resource.
type HubArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties HubPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (HubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubArgs)(nil)).Elem()
}

type HubInput interface {
	pulumi.Input

	ToHubOutput() HubOutput
	ToHubOutputWithContext(ctx context.Context) HubOutput
}

func (*Hub) ElementType() reflect.Type {
	return reflect.TypeOf((*Hub)(nil))
}

func (i *Hub) ToHubOutput() HubOutput {
	return i.ToHubOutputWithContext(context.Background())
}

func (i *Hub) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubOutput)
}

type HubOutput struct {
	*pulumi.OutputState
}

func (HubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Hub)(nil))
}

func (o HubOutput) ToHubOutput() HubOutput {
	return o
}

func (o HubOutput) ToHubOutputWithContext(ctx context.Context) HubOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HubOutput{})
}
