// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspacesweb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::WorkSpacesWeb::IdentityProvider Resource Type
type IdentityProvider struct {
	pulumi.CustomResourceState

	IdentityProviderArn     pulumi.StringOutput        `pulumi:"identityProviderArn"`
	IdentityProviderDetails pulumi.StringMapOutput     `pulumi:"identityProviderDetails"`
	IdentityProviderName    pulumi.StringOutput        `pulumi:"identityProviderName"`
	IdentityProviderType    IdentityProviderTypeOutput `pulumi:"identityProviderType"`
	PortalArn               pulumi.StringPtrOutput     `pulumi:"portalArn"`
}

// NewIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewIdentityProvider(ctx *pulumi.Context,
	name string, args *IdentityProviderArgs, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityProviderDetails == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderDetails'")
	}
	if args.IdentityProviderType == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProviderType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"portalArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IdentityProvider
	err := ctx.RegisterResource("aws-native:workspacesweb:IdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProvider gets an existing IdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderState, opts ...pulumi.ResourceOption) (*IdentityProvider, error) {
	var resource IdentityProvider
	err := ctx.ReadResource("aws-native:workspacesweb:IdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProvider resources.
type identityProviderState struct {
}

type IdentityProviderState struct {
}

func (IdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderState)(nil)).Elem()
}

type identityProviderArgs struct {
	IdentityProviderDetails map[string]string    `pulumi:"identityProviderDetails"`
	IdentityProviderName    *string              `pulumi:"identityProviderName"`
	IdentityProviderType    IdentityProviderType `pulumi:"identityProviderType"`
	PortalArn               *string              `pulumi:"portalArn"`
}

// The set of arguments for constructing a IdentityProvider resource.
type IdentityProviderArgs struct {
	IdentityProviderDetails pulumi.StringMapInput
	IdentityProviderName    pulumi.StringPtrInput
	IdentityProviderType    IdentityProviderTypeInput
	PortalArn               pulumi.StringPtrInput
}

func (IdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderArgs)(nil)).Elem()
}

type IdentityProviderInput interface {
	pulumi.Input

	ToIdentityProviderOutput() IdentityProviderOutput
	ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput
}

func (*IdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (i *IdentityProvider) ToIdentityProviderOutput() IdentityProviderOutput {
	return i.ToIdentityProviderOutputWithContext(context.Background())
}

func (i *IdentityProvider) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderOutput)
}

type IdentityProviderOutput struct{ *pulumi.OutputState }

func (IdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProvider)(nil)).Elem()
}

func (o IdentityProviderOutput) ToIdentityProviderOutput() IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) ToIdentityProviderOutputWithContext(ctx context.Context) IdentityProviderOutput {
	return o
}

func (o IdentityProviderOutput) IdentityProviderArn() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.IdentityProviderArn }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) IdentityProviderDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringMapOutput { return v.IdentityProviderDetails }).(pulumi.StringMapOutput)
}

func (o IdentityProviderOutput) IdentityProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringOutput { return v.IdentityProviderName }).(pulumi.StringOutput)
}

func (o IdentityProviderOutput) IdentityProviderType() IdentityProviderTypeOutput {
	return o.ApplyT(func(v *IdentityProvider) IdentityProviderTypeOutput { return v.IdentityProviderType }).(IdentityProviderTypeOutput)
}

func (o IdentityProviderOutput) PortalArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProvider) pulumi.StringPtrOutput { return v.PortalArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderInput)(nil)).Elem(), &IdentityProvider{})
	pulumi.RegisterOutputType(IdentityProviderOutput{})
}
