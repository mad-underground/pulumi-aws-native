// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Inspector::ResourceGroup
//
// Deprecated: ResourceGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ResourceGroup struct {
	pulumi.CustomResourceState

	Arn               pulumi.StringOutput         `pulumi:"arn"`
	ResourceGroupTags ResourceGroupTagArrayOutput `pulumi:"resourceGroupTags"`
}

// NewResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewResourceGroup(ctx *pulumi.Context,
	name string, args *ResourceGroupArgs, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupTags == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupTags'")
	}
	var resource ResourceGroup
	err := ctx.RegisterResource("aws-native:inspector:ResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceGroup gets an existing ResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceGroupState, opts ...pulumi.ResourceOption) (*ResourceGroup, error) {
	var resource ResourceGroup
	err := ctx.ReadResource("aws-native:inspector:ResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceGroup resources.
type resourceGroupState struct {
}

type ResourceGroupState struct {
}

func (ResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupState)(nil)).Elem()
}

type resourceGroupArgs struct {
	ResourceGroupTags []ResourceGroupTag `pulumi:"resourceGroupTags"`
}

// The set of arguments for constructing a ResourceGroup resource.
type ResourceGroupArgs struct {
	ResourceGroupTags ResourceGroupTagArrayInput
}

func (ResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceGroupArgs)(nil)).Elem()
}

type ResourceGroupInput interface {
	pulumi.Input

	ToResourceGroupOutput() ResourceGroupOutput
	ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput
}

func (*ResourceGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroup)(nil))
}

func (i *ResourceGroup) ToResourceGroupOutput() ResourceGroupOutput {
	return i.ToResourceGroupOutputWithContext(context.Background())
}

func (i *ResourceGroup) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupOutput)
}

type ResourceGroupOutput struct{ *pulumi.OutputState }

func (ResourceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroup)(nil))
}

func (o ResourceGroupOutput) ToResourceGroupOutput() ResourceGroupOutput {
	return o
}

func (o ResourceGroupOutput) ToResourceGroupOutputWithContext(ctx context.Context) ResourceGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceGroupOutput{})
}