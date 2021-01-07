// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html
type ByteMatchSet struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ByteMatchSetAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ByteMatchSetPropertiesOutput `pulumi:"properties"`
}

// NewByteMatchSet registers a new resource with the given unique name, arguments, and options.
func NewByteMatchSet(ctx *pulumi.Context,
	name string, args *ByteMatchSetArgs, opts ...pulumi.ResourceOption) (*ByteMatchSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ByteMatchSet
	err := ctx.RegisterResource("cloudformation:WAF:ByteMatchSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetByteMatchSet gets an existing ByteMatchSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetByteMatchSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ByteMatchSetState, opts ...pulumi.ResourceOption) (*ByteMatchSet, error) {
	var resource ByteMatchSet
	err := ctx.ReadResource("cloudformation:WAF:ByteMatchSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ByteMatchSet resources.
type byteMatchSetState struct {
	// The attributes associated with the resource
	Attributes *ByteMatchSetAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ByteMatchSetProperties `pulumi:"properties"`
}

type ByteMatchSetState struct {
	// The attributes associated with the resource
	Attributes ByteMatchSetAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ByteMatchSetPropertiesPtrInput
}

func (ByteMatchSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetState)(nil)).Elem()
}

type byteMatchSetArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ByteMatchSetProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ByteMatchSet resource.
type ByteMatchSetArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ByteMatchSetPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ByteMatchSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*byteMatchSetArgs)(nil)).Elem()
}

type ByteMatchSetInput interface {
	pulumi.Input

	ToByteMatchSetOutput() ByteMatchSetOutput
	ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput
}

func (*ByteMatchSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ByteMatchSet)(nil))
}

func (i *ByteMatchSet) ToByteMatchSetOutput() ByteMatchSetOutput {
	return i.ToByteMatchSetOutputWithContext(context.Background())
}

func (i *ByteMatchSet) ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ByteMatchSetOutput)
}

type ByteMatchSetOutput struct {
	*pulumi.OutputState
}

func (ByteMatchSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ByteMatchSet)(nil))
}

func (o ByteMatchSetOutput) ToByteMatchSetOutput() ByteMatchSetOutput {
	return o
}

func (o ByteMatchSetOutput) ToByteMatchSetOutputWithContext(ctx context.Context) ByteMatchSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ByteMatchSetOutput{})
}