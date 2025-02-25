// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rolesanywhere

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::RolesAnywhere::TrustAnchor Resource Type.
type TrustAnchor struct {
	pulumi.CustomResourceState

	Enabled              pulumi.BoolPtrOutput                      `pulumi:"enabled"`
	Name                 pulumi.StringOutput                       `pulumi:"name"`
	NotificationSettings TrustAnchorNotificationSettingArrayOutput `pulumi:"notificationSettings"`
	Source               TrustAnchorSourceOutput                   `pulumi:"source"`
	Tags                 aws.TagArrayOutput                        `pulumi:"tags"`
	TrustAnchorArn       pulumi.StringOutput                       `pulumi:"trustAnchorArn"`
	TrustAnchorId        pulumi.StringOutput                       `pulumi:"trustAnchorId"`
}

// NewTrustAnchor registers a new resource with the given unique name, arguments, and options.
func NewTrustAnchor(ctx *pulumi.Context,
	name string, args *TrustAnchorArgs, opts ...pulumi.ResourceOption) (*TrustAnchor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Source == nil {
		return nil, errors.New("invalid value for required argument 'Source'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TrustAnchor
	err := ctx.RegisterResource("aws-native:rolesanywhere:TrustAnchor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrustAnchor gets an existing TrustAnchor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrustAnchor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrustAnchorState, opts ...pulumi.ResourceOption) (*TrustAnchor, error) {
	var resource TrustAnchor
	err := ctx.ReadResource("aws-native:rolesanywhere:TrustAnchor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrustAnchor resources.
type trustAnchorState struct {
}

type TrustAnchorState struct {
}

func (TrustAnchorState) ElementType() reflect.Type {
	return reflect.TypeOf((*trustAnchorState)(nil)).Elem()
}

type trustAnchorArgs struct {
	Enabled              *bool                            `pulumi:"enabled"`
	Name                 *string                          `pulumi:"name"`
	NotificationSettings []TrustAnchorNotificationSetting `pulumi:"notificationSettings"`
	Source               TrustAnchorSource                `pulumi:"source"`
	Tags                 []aws.Tag                        `pulumi:"tags"`
}

// The set of arguments for constructing a TrustAnchor resource.
type TrustAnchorArgs struct {
	Enabled              pulumi.BoolPtrInput
	Name                 pulumi.StringPtrInput
	NotificationSettings TrustAnchorNotificationSettingArrayInput
	Source               TrustAnchorSourceInput
	Tags                 aws.TagArrayInput
}

func (TrustAnchorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trustAnchorArgs)(nil)).Elem()
}

type TrustAnchorInput interface {
	pulumi.Input

	ToTrustAnchorOutput() TrustAnchorOutput
	ToTrustAnchorOutputWithContext(ctx context.Context) TrustAnchorOutput
}

func (*TrustAnchor) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchor)(nil)).Elem()
}

func (i *TrustAnchor) ToTrustAnchorOutput() TrustAnchorOutput {
	return i.ToTrustAnchorOutputWithContext(context.Background())
}

func (i *TrustAnchor) ToTrustAnchorOutputWithContext(ctx context.Context) TrustAnchorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorOutput)
}

type TrustAnchorOutput struct{ *pulumi.OutputState }

func (TrustAnchorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchor)(nil)).Elem()
}

func (o TrustAnchorOutput) ToTrustAnchorOutput() TrustAnchorOutput {
	return o
}

func (o TrustAnchorOutput) ToTrustAnchorOutputWithContext(ctx context.Context) TrustAnchorOutput {
	return o
}

func (o TrustAnchorOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o TrustAnchorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TrustAnchorOutput) NotificationSettings() TrustAnchorNotificationSettingArrayOutput {
	return o.ApplyT(func(v *TrustAnchor) TrustAnchorNotificationSettingArrayOutput { return v.NotificationSettings }).(TrustAnchorNotificationSettingArrayOutput)
}

func (o TrustAnchorOutput) Source() TrustAnchorSourceOutput {
	return o.ApplyT(func(v *TrustAnchor) TrustAnchorSourceOutput { return v.Source }).(TrustAnchorSourceOutput)
}

func (o TrustAnchorOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *TrustAnchor) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o TrustAnchorOutput) TrustAnchorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.StringOutput { return v.TrustAnchorArn }).(pulumi.StringOutput)
}

func (o TrustAnchorOutput) TrustAnchorId() pulumi.StringOutput {
	return o.ApplyT(func(v *TrustAnchor) pulumi.StringOutput { return v.TrustAnchorId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorInput)(nil)).Elem(), &TrustAnchor{})
	pulumi.RegisterOutputType(TrustAnchorOutput{})
}
