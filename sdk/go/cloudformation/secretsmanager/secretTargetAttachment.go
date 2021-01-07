// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html
type SecretTargetAttachment struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes SecretTargetAttachmentAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties SecretTargetAttachmentPropertiesOutput `pulumi:"properties"`
}

// NewSecretTargetAttachment registers a new resource with the given unique name, arguments, and options.
func NewSecretTargetAttachment(ctx *pulumi.Context,
	name string, args *SecretTargetAttachmentArgs, opts ...pulumi.ResourceOption) (*SecretTargetAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource SecretTargetAttachment
	err := ctx.RegisterResource("cloudformation:SecretsManager:SecretTargetAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretTargetAttachment gets an existing SecretTargetAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretTargetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretTargetAttachmentState, opts ...pulumi.ResourceOption) (*SecretTargetAttachment, error) {
	var resource SecretTargetAttachment
	err := ctx.ReadResource("cloudformation:SecretsManager:SecretTargetAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretTargetAttachment resources.
type secretTargetAttachmentState struct {
	// The attributes associated with the resource
	Attributes *SecretTargetAttachmentAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *SecretTargetAttachmentProperties `pulumi:"properties"`
}

type SecretTargetAttachmentState struct {
	// The attributes associated with the resource
	Attributes SecretTargetAttachmentAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties SecretTargetAttachmentPropertiesPtrInput
}

func (SecretTargetAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretTargetAttachmentState)(nil)).Elem()
}

type secretTargetAttachmentArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties SecretTargetAttachmentProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a SecretTargetAttachment resource.
type SecretTargetAttachmentArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties SecretTargetAttachmentPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (SecretTargetAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretTargetAttachmentArgs)(nil)).Elem()
}

type SecretTargetAttachmentInput interface {
	pulumi.Input

	ToSecretTargetAttachmentOutput() SecretTargetAttachmentOutput
	ToSecretTargetAttachmentOutputWithContext(ctx context.Context) SecretTargetAttachmentOutput
}

func (*SecretTargetAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretTargetAttachment)(nil))
}

func (i *SecretTargetAttachment) ToSecretTargetAttachmentOutput() SecretTargetAttachmentOutput {
	return i.ToSecretTargetAttachmentOutputWithContext(context.Background())
}

func (i *SecretTargetAttachment) ToSecretTargetAttachmentOutputWithContext(ctx context.Context) SecretTargetAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretTargetAttachmentOutput)
}

type SecretTargetAttachmentOutput struct {
	*pulumi.OutputState
}

func (SecretTargetAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretTargetAttachment)(nil))
}

func (o SecretTargetAttachmentOutput) ToSecretTargetAttachmentOutput() SecretTargetAttachmentOutput {
	return o
}

func (o SecretTargetAttachmentOutput) ToSecretTargetAttachmentOutputWithContext(ctx context.Context) SecretTargetAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecretTargetAttachmentOutput{})
}
