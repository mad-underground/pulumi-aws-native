// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackage-packaginggroup.html
type PackagingGroup struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes PackagingGroupAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties PackagingGroupPropertiesOutput `pulumi:"properties"`
}

// NewPackagingGroup registers a new resource with the given unique name, arguments, and options.
func NewPackagingGroup(ctx *pulumi.Context,
	name string, args *PackagingGroupArgs, opts ...pulumi.ResourceOption) (*PackagingGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource PackagingGroup
	err := ctx.RegisterResource("cloudformation:MediaPackage:PackagingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPackagingGroup gets an existing PackagingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPackagingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PackagingGroupState, opts ...pulumi.ResourceOption) (*PackagingGroup, error) {
	var resource PackagingGroup
	err := ctx.ReadResource("cloudformation:MediaPackage:PackagingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PackagingGroup resources.
type packagingGroupState struct {
	// The attributes associated with the resource
	Attributes *PackagingGroupAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *PackagingGroupProperties `pulumi:"properties"`
}

type PackagingGroupState struct {
	// The attributes associated with the resource
	Attributes PackagingGroupAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties PackagingGroupPropertiesPtrInput
}

func (PackagingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*packagingGroupState)(nil)).Elem()
}

type packagingGroupArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties PackagingGroupProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a PackagingGroup resource.
type PackagingGroupArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties PackagingGroupPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (PackagingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packagingGroupArgs)(nil)).Elem()
}

type PackagingGroupInput interface {
	pulumi.Input

	ToPackagingGroupOutput() PackagingGroupOutput
	ToPackagingGroupOutputWithContext(ctx context.Context) PackagingGroupOutput
}

func (*PackagingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*PackagingGroup)(nil))
}

func (i *PackagingGroup) ToPackagingGroupOutput() PackagingGroupOutput {
	return i.ToPackagingGroupOutputWithContext(context.Background())
}

func (i *PackagingGroup) ToPackagingGroupOutputWithContext(ctx context.Context) PackagingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackagingGroupOutput)
}

type PackagingGroupOutput struct {
	*pulumi.OutputState
}

func (PackagingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PackagingGroup)(nil))
}

func (o PackagingGroupOutput) ToPackagingGroupOutput() PackagingGroupOutput {
	return o
}

func (o PackagingGroupOutput) ToPackagingGroupOutputWithContext(ctx context.Context) PackagingGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PackagingGroupOutput{})
}
