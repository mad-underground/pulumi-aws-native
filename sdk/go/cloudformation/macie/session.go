// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-macie-session.html
type Session struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes SessionAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties SessionPropertiesOutput `pulumi:"properties"`
}

// NewSession registers a new resource with the given unique name, arguments, and options.
func NewSession(ctx *pulumi.Context,
	name string, args *SessionArgs, opts ...pulumi.ResourceOption) (*Session, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Session
	err := ctx.RegisterResource("cloudformation:Macie:Session", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSession gets an existing Session resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSession(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SessionState, opts ...pulumi.ResourceOption) (*Session, error) {
	var resource Session
	err := ctx.ReadResource("cloudformation:Macie:Session", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Session resources.
type sessionState struct {
	// The attributes associated with the resource
	Attributes *SessionAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *SessionProperties `pulumi:"properties"`
}

type SessionState struct {
	// The attributes associated with the resource
	Attributes SessionAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties SessionPropertiesPtrInput
}

func (SessionState) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionState)(nil)).Elem()
}

type sessionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties SessionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Session resource.
type SessionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties SessionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (SessionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sessionArgs)(nil)).Elem()
}

type SessionInput interface {
	pulumi.Input

	ToSessionOutput() SessionOutput
	ToSessionOutputWithContext(ctx context.Context) SessionOutput
}

func (*Session) ElementType() reflect.Type {
	return reflect.TypeOf((*Session)(nil))
}

func (i *Session) ToSessionOutput() SessionOutput {
	return i.ToSessionOutputWithContext(context.Background())
}

func (i *Session) ToSessionOutputWithContext(ctx context.Context) SessionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SessionOutput)
}

type SessionOutput struct {
	*pulumi.OutputState
}

func (SessionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Session)(nil))
}

func (o SessionOutput) ToSessionOutput() SessionOutput {
	return o
}

func (o SessionOutput) ToSessionOutputWithContext(ctx context.Context) SessionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SessionOutput{})
}
