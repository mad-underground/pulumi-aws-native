// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplifyuibuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::AmplifyUIBuilder::Theme Resource Type
type Theme struct {
	pulumi.CustomResourceState

	AppId           pulumi.StringPtrOutput `pulumi:"appId"`
	CreatedAt       pulumi.StringOutput    `pulumi:"createdAt"`
	EnvironmentName pulumi.StringPtrOutput `pulumi:"environmentName"`
	ModifiedAt      pulumi.StringOutput    `pulumi:"modifiedAt"`
	Name            pulumi.StringPtrOutput `pulumi:"name"`
	Overrides       ThemeValuesArrayOutput `pulumi:"overrides"`
	Tags            pulumi.StringMapOutput `pulumi:"tags"`
	Values          ThemeValuesArrayOutput `pulumi:"values"`
}

// NewTheme registers a new resource with the given unique name, arguments, and options.
func NewTheme(ctx *pulumi.Context,
	name string, args *ThemeArgs, opts ...pulumi.ResourceOption) (*Theme, error) {
	if args == nil {
		args = &ThemeArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"appId",
		"environmentName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Theme
	err := ctx.RegisterResource("aws-native:amplifyuibuilder:Theme", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTheme gets an existing Theme resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTheme(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThemeState, opts ...pulumi.ResourceOption) (*Theme, error) {
	var resource Theme
	err := ctx.ReadResource("aws-native:amplifyuibuilder:Theme", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Theme resources.
type themeState struct {
}

type ThemeState struct {
}

func (ThemeState) ElementType() reflect.Type {
	return reflect.TypeOf((*themeState)(nil)).Elem()
}

type themeArgs struct {
	AppId           *string           `pulumi:"appId"`
	EnvironmentName *string           `pulumi:"environmentName"`
	Name            *string           `pulumi:"name"`
	Overrides       []ThemeValues     `pulumi:"overrides"`
	Tags            map[string]string `pulumi:"tags"`
	Values          []ThemeValues     `pulumi:"values"`
}

// The set of arguments for constructing a Theme resource.
type ThemeArgs struct {
	AppId           pulumi.StringPtrInput
	EnvironmentName pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	Overrides       ThemeValuesArrayInput
	Tags            pulumi.StringMapInput
	Values          ThemeValuesArrayInput
}

func (ThemeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*themeArgs)(nil)).Elem()
}

type ThemeInput interface {
	pulumi.Input

	ToThemeOutput() ThemeOutput
	ToThemeOutputWithContext(ctx context.Context) ThemeOutput
}

func (*Theme) ElementType() reflect.Type {
	return reflect.TypeOf((**Theme)(nil)).Elem()
}

func (i *Theme) ToThemeOutput() ThemeOutput {
	return i.ToThemeOutputWithContext(context.Background())
}

func (i *Theme) ToThemeOutputWithContext(ctx context.Context) ThemeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThemeOutput)
}

type ThemeOutput struct{ *pulumi.OutputState }

func (ThemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Theme)(nil)).Elem()
}

func (o ThemeOutput) ToThemeOutput() ThemeOutput {
	return o
}

func (o ThemeOutput) ToThemeOutputWithContext(ctx context.Context) ThemeOutput {
	return o
}

func (o ThemeOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringPtrOutput { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o ThemeOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ThemeOutput) EnvironmentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringPtrOutput { return v.EnvironmentName }).(pulumi.StringPtrOutput)
}

func (o ThemeOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringOutput { return v.ModifiedAt }).(pulumi.StringOutput)
}

func (o ThemeOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ThemeOutput) Overrides() ThemeValuesArrayOutput {
	return o.ApplyT(func(v *Theme) ThemeValuesArrayOutput { return v.Overrides }).(ThemeValuesArrayOutput)
}

func (o ThemeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Theme) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ThemeOutput) Values() ThemeValuesArrayOutput {
	return o.ApplyT(func(v *Theme) ThemeValuesArrayOutput { return v.Values }).(ThemeValuesArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ThemeInput)(nil)).Elem(), &Theme{})
	pulumi.RegisterOutputType(ThemeOutput{})
}
