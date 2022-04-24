// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Connect::ContactFlow
type ContactFlow struct {
	pulumi.CustomResourceState

	// The identifier of the contact flow (ARN).
	ContactFlowArn pulumi.StringOutput `pulumi:"contactFlowArn"`
	// The content of the contact flow in JSON format.
	Content pulumi.StringOutput `pulumi:"content"`
	// The description of the contact flow.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The identifier of the Amazon Connect instance (ARN).
	InstanceArn pulumi.StringOutput `pulumi:"instanceArn"`
	// The name of the contact flow.
	Name pulumi.StringOutput `pulumi:"name"`
	// The state of the contact flow.
	State ContactFlowStateEnumPtrOutput `pulumi:"state"`
	// One or more tags.
	Tags ContactFlowTagArrayOutput `pulumi:"tags"`
	// The type of the contact flow.
	Type ContactFlowTypePtrOutput `pulumi:"type"`
}

// NewContactFlow registers a new resource with the given unique name, arguments, and options.
func NewContactFlow(ctx *pulumi.Context,
	name string, args *ContactFlowArgs, opts ...pulumi.ResourceOption) (*ContactFlow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.InstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'InstanceArn'")
	}
	var resource ContactFlow
	err := ctx.RegisterResource("aws-native:connect:ContactFlow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContactFlow gets an existing ContactFlow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContactFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactFlowState, opts ...pulumi.ResourceOption) (*ContactFlow, error) {
	var resource ContactFlow
	err := ctx.ReadResource("aws-native:connect:ContactFlow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContactFlow resources.
type contactFlowState struct {
}

type ContactFlowState struct {
}

func (ContactFlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactFlowState)(nil)).Elem()
}

type contactFlowArgs struct {
	// The content of the contact flow in JSON format.
	Content string `pulumi:"content"`
	// The description of the contact flow.
	Description *string `pulumi:"description"`
	// The identifier of the Amazon Connect instance (ARN).
	InstanceArn string `pulumi:"instanceArn"`
	// The name of the contact flow.
	Name *string `pulumi:"name"`
	// The state of the contact flow.
	State *ContactFlowStateEnum `pulumi:"state"`
	// One or more tags.
	Tags []ContactFlowTag `pulumi:"tags"`
	// The type of the contact flow.
	Type *ContactFlowType `pulumi:"type"`
}

// The set of arguments for constructing a ContactFlow resource.
type ContactFlowArgs struct {
	// The content of the contact flow in JSON format.
	Content pulumi.StringInput
	// The description of the contact flow.
	Description pulumi.StringPtrInput
	// The identifier of the Amazon Connect instance (ARN).
	InstanceArn pulumi.StringInput
	// The name of the contact flow.
	Name pulumi.StringPtrInput
	// The state of the contact flow.
	State ContactFlowStateEnumPtrInput
	// One or more tags.
	Tags ContactFlowTagArrayInput
	// The type of the contact flow.
	Type ContactFlowTypePtrInput
}

func (ContactFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactFlowArgs)(nil)).Elem()
}

type ContactFlowInput interface {
	pulumi.Input

	ToContactFlowOutput() ContactFlowOutput
	ToContactFlowOutputWithContext(ctx context.Context) ContactFlowOutput
}

func (*ContactFlow) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactFlow)(nil)).Elem()
}

func (i *ContactFlow) ToContactFlowOutput() ContactFlowOutput {
	return i.ToContactFlowOutputWithContext(context.Background())
}

func (i *ContactFlow) ToContactFlowOutputWithContext(ctx context.Context) ContactFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactFlowOutput)
}

type ContactFlowOutput struct{ *pulumi.OutputState }

func (ContactFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactFlow)(nil)).Elem()
}

func (o ContactFlowOutput) ToContactFlowOutput() ContactFlowOutput {
	return o
}

func (o ContactFlowOutput) ToContactFlowOutputWithContext(ctx context.Context) ContactFlowOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactFlowInput)(nil)).Elem(), &ContactFlow{})
	pulumi.RegisterOutputType(ContactFlowOutput{})
}
