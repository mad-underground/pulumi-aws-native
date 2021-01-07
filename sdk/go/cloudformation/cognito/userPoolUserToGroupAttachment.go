// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpoolusertogroupattachment.html
type UserPoolUserToGroupAttachment struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes UserPoolUserToGroupAttachmentAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties UserPoolUserToGroupAttachmentPropertiesOutput `pulumi:"properties"`
}

// NewUserPoolUserToGroupAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserPoolUserToGroupAttachment(ctx *pulumi.Context,
	name string, args *UserPoolUserToGroupAttachmentArgs, opts ...pulumi.ResourceOption) (*UserPoolUserToGroupAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource UserPoolUserToGroupAttachment
	err := ctx.RegisterResource("cloudformation:Cognito:UserPoolUserToGroupAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPoolUserToGroupAttachment gets an existing UserPoolUserToGroupAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPoolUserToGroupAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPoolUserToGroupAttachmentState, opts ...pulumi.ResourceOption) (*UserPoolUserToGroupAttachment, error) {
	var resource UserPoolUserToGroupAttachment
	err := ctx.ReadResource("cloudformation:Cognito:UserPoolUserToGroupAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPoolUserToGroupAttachment resources.
type userPoolUserToGroupAttachmentState struct {
	// The attributes associated with the resource
	Attributes *UserPoolUserToGroupAttachmentAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *UserPoolUserToGroupAttachmentProperties `pulumi:"properties"`
}

type UserPoolUserToGroupAttachmentState struct {
	// The attributes associated with the resource
	Attributes UserPoolUserToGroupAttachmentAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties UserPoolUserToGroupAttachmentPropertiesPtrInput
}

func (UserPoolUserToGroupAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolUserToGroupAttachmentState)(nil)).Elem()
}

type userPoolUserToGroupAttachmentArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties UserPoolUserToGroupAttachmentProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a UserPoolUserToGroupAttachment resource.
type UserPoolUserToGroupAttachmentArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties UserPoolUserToGroupAttachmentPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (UserPoolUserToGroupAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolUserToGroupAttachmentArgs)(nil)).Elem()
}

type UserPoolUserToGroupAttachmentInput interface {
	pulumi.Input

	ToUserPoolUserToGroupAttachmentOutput() UserPoolUserToGroupAttachmentOutput
	ToUserPoolUserToGroupAttachmentOutputWithContext(ctx context.Context) UserPoolUserToGroupAttachmentOutput
}

func (*UserPoolUserToGroupAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolUserToGroupAttachment)(nil))
}

func (i *UserPoolUserToGroupAttachment) ToUserPoolUserToGroupAttachmentOutput() UserPoolUserToGroupAttachmentOutput {
	return i.ToUserPoolUserToGroupAttachmentOutputWithContext(context.Background())
}

func (i *UserPoolUserToGroupAttachment) ToUserPoolUserToGroupAttachmentOutputWithContext(ctx context.Context) UserPoolUserToGroupAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolUserToGroupAttachmentOutput)
}

type UserPoolUserToGroupAttachmentOutput struct {
	*pulumi.OutputState
}

func (UserPoolUserToGroupAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolUserToGroupAttachment)(nil))
}

func (o UserPoolUserToGroupAttachmentOutput) ToUserPoolUserToGroupAttachmentOutput() UserPoolUserToGroupAttachmentOutput {
	return o
}

func (o UserPoolUserToGroupAttachmentOutput) ToUserPoolUserToGroupAttachmentOutputWithContext(ctx context.Context) UserPoolUserToGroupAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserPoolUserToGroupAttachmentOutput{})
}
