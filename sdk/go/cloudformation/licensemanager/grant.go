// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package licensemanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-licensemanager-grant.html
type Grant struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes GrantAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GrantPropertiesOutput `pulumi:"properties"`
}

// NewGrant registers a new resource with the given unique name, arguments, and options.
func NewGrant(ctx *pulumi.Context,
	name string, args *GrantArgs, opts ...pulumi.ResourceOption) (*Grant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Grant
	err := ctx.RegisterResource("cloudformation:LicenseManager:Grant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrant gets an existing Grant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrantState, opts ...pulumi.ResourceOption) (*Grant, error) {
	var resource Grant
	err := ctx.ReadResource("cloudformation:LicenseManager:Grant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Grant resources.
type grantState struct {
	// The attributes associated with the resource
	Attributes *GrantAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *GrantProperties `pulumi:"properties"`
}

type GrantState struct {
	// The attributes associated with the resource
	Attributes GrantAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GrantPropertiesPtrInput
}

func (GrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*grantState)(nil)).Elem()
}

type grantArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties GrantProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Grant resource.
type GrantArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties GrantPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (GrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grantArgs)(nil)).Elem()
}

type GrantInput interface {
	pulumi.Input

	ToGrantOutput() GrantOutput
	ToGrantOutputWithContext(ctx context.Context) GrantOutput
}

func (*Grant) ElementType() reflect.Type {
	return reflect.TypeOf((*Grant)(nil))
}

func (i *Grant) ToGrantOutput() GrantOutput {
	return i.ToGrantOutputWithContext(context.Background())
}

func (i *Grant) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantOutput)
}

type GrantOutput struct {
	*pulumi.OutputState
}

func (GrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Grant)(nil))
}

func (o GrantOutput) ToGrantOutput() GrantOutput {
	return o
}

func (o GrantOutput) ToGrantOutputWithContext(ctx context.Context) GrantOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GrantOutput{})
}
