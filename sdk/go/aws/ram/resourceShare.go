// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RAM::ResourceShare
//
// Deprecated: ResourceShare is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ResourceShare struct {
	pulumi.CustomResourceState

	AllowExternalPrincipals pulumi.BoolPtrOutput        `pulumi:"allowExternalPrincipals"`
	Arn                     pulumi.StringOutput         `pulumi:"arn"`
	Name                    pulumi.StringOutput         `pulumi:"name"`
	PermissionArns          pulumi.StringArrayOutput    `pulumi:"permissionArns"`
	Principals              pulumi.StringArrayOutput    `pulumi:"principals"`
	ResourceArns            pulumi.StringArrayOutput    `pulumi:"resourceArns"`
	Tags                    ResourceShareTagArrayOutput `pulumi:"tags"`
}

// NewResourceShare registers a new resource with the given unique name, arguments, and options.
func NewResourceShare(ctx *pulumi.Context,
	name string, args *ResourceShareArgs, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource ResourceShare
	err := ctx.RegisterResource("aws-native:ram:ResourceShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceShare gets an existing ResourceShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceShareState, opts ...pulumi.ResourceOption) (*ResourceShare, error) {
	var resource ResourceShare
	err := ctx.ReadResource("aws-native:ram:ResourceShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceShare resources.
type resourceShareState struct {
}

type ResourceShareState struct {
}

func (ResourceShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareState)(nil)).Elem()
}

type resourceShareArgs struct {
	AllowExternalPrincipals *bool              `pulumi:"allowExternalPrincipals"`
	Name                    string             `pulumi:"name"`
	PermissionArns          []string           `pulumi:"permissionArns"`
	Principals              []string           `pulumi:"principals"`
	ResourceArns            []string           `pulumi:"resourceArns"`
	Tags                    []ResourceShareTag `pulumi:"tags"`
}

// The set of arguments for constructing a ResourceShare resource.
type ResourceShareArgs struct {
	AllowExternalPrincipals pulumi.BoolPtrInput
	Name                    pulumi.StringInput
	PermissionArns          pulumi.StringArrayInput
	Principals              pulumi.StringArrayInput
	ResourceArns            pulumi.StringArrayInput
	Tags                    ResourceShareTagArrayInput
}

func (ResourceShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareArgs)(nil)).Elem()
}

type ResourceShareInput interface {
	pulumi.Input

	ToResourceShareOutput() ResourceShareOutput
	ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput
}

func (*ResourceShare) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceShare)(nil))
}

func (i *ResourceShare) ToResourceShareOutput() ResourceShareOutput {
	return i.ToResourceShareOutputWithContext(context.Background())
}

func (i *ResourceShare) ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceShareOutput)
}

type ResourceShareOutput struct{ *pulumi.OutputState }

func (ResourceShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceShare)(nil))
}

func (o ResourceShareOutput) ToResourceShareOutput() ResourceShareOutput {
	return o
}

func (o ResourceShareOutput) ToResourceShareOutputWithContext(ctx context.Context) ResourceShareOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceShareOutput{})
}