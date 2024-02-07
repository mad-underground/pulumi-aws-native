// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package verifiedpermissions

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::VerifiedPermissions::IdentitySource Resource Type
type IdentitySource struct {
	pulumi.CustomResourceState

	Configuration       IdentitySourceConfigurationOutput `pulumi:"configuration"`
	Details             IdentitySourceDetailsOutput       `pulumi:"details"`
	IdentitySourceId    pulumi.StringOutput               `pulumi:"identitySourceId"`
	PolicyStoreId       pulumi.StringPtrOutput            `pulumi:"policyStoreId"`
	PrincipalEntityType pulumi.StringPtrOutput            `pulumi:"principalEntityType"`
}

// NewIdentitySource registers a new resource with the given unique name, arguments, and options.
func NewIdentitySource(ctx *pulumi.Context,
	name string, args *IdentitySourceArgs, opts ...pulumi.ResourceOption) (*IdentitySource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Configuration == nil {
		return nil, errors.New("invalid value for required argument 'Configuration'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"policyStoreId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IdentitySource
	err := ctx.RegisterResource("aws-native:verifiedpermissions:IdentitySource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentitySource gets an existing IdentitySource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentitySource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentitySourceState, opts ...pulumi.ResourceOption) (*IdentitySource, error) {
	var resource IdentitySource
	err := ctx.ReadResource("aws-native:verifiedpermissions:IdentitySource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentitySource resources.
type identitySourceState struct {
}

type IdentitySourceState struct {
}

func (IdentitySourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*identitySourceState)(nil)).Elem()
}

type identitySourceArgs struct {
	Configuration       IdentitySourceConfiguration `pulumi:"configuration"`
	PolicyStoreId       *string                     `pulumi:"policyStoreId"`
	PrincipalEntityType *string                     `pulumi:"principalEntityType"`
}

// The set of arguments for constructing a IdentitySource resource.
type IdentitySourceArgs struct {
	Configuration       IdentitySourceConfigurationInput
	PolicyStoreId       pulumi.StringPtrInput
	PrincipalEntityType pulumi.StringPtrInput
}

func (IdentitySourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identitySourceArgs)(nil)).Elem()
}

type IdentitySourceInput interface {
	pulumi.Input

	ToIdentitySourceOutput() IdentitySourceOutput
	ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput
}

func (*IdentitySource) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentitySource)(nil)).Elem()
}

func (i *IdentitySource) ToIdentitySourceOutput() IdentitySourceOutput {
	return i.ToIdentitySourceOutputWithContext(context.Background())
}

func (i *IdentitySource) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentitySourceOutput)
}

type IdentitySourceOutput struct{ *pulumi.OutputState }

func (IdentitySourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentitySource)(nil)).Elem()
}

func (o IdentitySourceOutput) ToIdentitySourceOutput() IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) ToIdentitySourceOutputWithContext(ctx context.Context) IdentitySourceOutput {
	return o
}

func (o IdentitySourceOutput) Configuration() IdentitySourceConfigurationOutput {
	return o.ApplyT(func(v *IdentitySource) IdentitySourceConfigurationOutput { return v.Configuration }).(IdentitySourceConfigurationOutput)
}

func (o IdentitySourceOutput) Details() IdentitySourceDetailsOutput {
	return o.ApplyT(func(v *IdentitySource) IdentitySourceDetailsOutput { return v.Details }).(IdentitySourceDetailsOutput)
}

func (o IdentitySourceOutput) IdentitySourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentitySource) pulumi.StringOutput { return v.IdentitySourceId }).(pulumi.StringOutput)
}

func (o IdentitySourceOutput) PolicyStoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentitySource) pulumi.StringPtrOutput { return v.PolicyStoreId }).(pulumi.StringPtrOutput)
}

func (o IdentitySourceOutput) PrincipalEntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentitySource) pulumi.StringPtrOutput { return v.PrincipalEntityType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentitySourceInput)(nil)).Elem(), &IdentitySource{})
	pulumi.RegisterOutputType(IdentitySourceOutput{})
}
