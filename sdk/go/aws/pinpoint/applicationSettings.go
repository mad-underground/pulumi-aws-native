// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::ApplicationSettings
//
// Deprecated: ApplicationSettings is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ApplicationSettings struct {
	pulumi.CustomResourceState

	ApplicationId            pulumi.StringOutput                      `pulumi:"applicationId"`
	CampaignHook             ApplicationSettingsCampaignHookPtrOutput `pulumi:"campaignHook"`
	CloudWatchMetricsEnabled pulumi.BoolPtrOutput                     `pulumi:"cloudWatchMetricsEnabled"`
	Limits                   ApplicationSettingsLimitsPtrOutput       `pulumi:"limits"`
	QuietTime                ApplicationSettingsQuietTimePtrOutput    `pulumi:"quietTime"`
}

// NewApplicationSettings registers a new resource with the given unique name, arguments, and options.
func NewApplicationSettings(ctx *pulumi.Context,
	name string, args *ApplicationSettingsArgs, opts ...pulumi.ResourceOption) (*ApplicationSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApplicationSettings
	err := ctx.RegisterResource("aws-native:pinpoint:ApplicationSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSettings gets an existing ApplicationSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSettingsState, opts ...pulumi.ResourceOption) (*ApplicationSettings, error) {
	var resource ApplicationSettings
	err := ctx.ReadResource("aws-native:pinpoint:ApplicationSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSettings resources.
type applicationSettingsState struct {
}

type ApplicationSettingsState struct {
}

func (ApplicationSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSettingsState)(nil)).Elem()
}

type applicationSettingsArgs struct {
	ApplicationId            string                           `pulumi:"applicationId"`
	CampaignHook             *ApplicationSettingsCampaignHook `pulumi:"campaignHook"`
	CloudWatchMetricsEnabled *bool                            `pulumi:"cloudWatchMetricsEnabled"`
	Limits                   *ApplicationSettingsLimits       `pulumi:"limits"`
	QuietTime                *ApplicationSettingsQuietTime    `pulumi:"quietTime"`
}

// The set of arguments for constructing a ApplicationSettings resource.
type ApplicationSettingsArgs struct {
	ApplicationId            pulumi.StringInput
	CampaignHook             ApplicationSettingsCampaignHookPtrInput
	CloudWatchMetricsEnabled pulumi.BoolPtrInput
	Limits                   ApplicationSettingsLimitsPtrInput
	QuietTime                ApplicationSettingsQuietTimePtrInput
}

func (ApplicationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSettingsArgs)(nil)).Elem()
}

type ApplicationSettingsInput interface {
	pulumi.Input

	ToApplicationSettingsOutput() ApplicationSettingsOutput
	ToApplicationSettingsOutputWithContext(ctx context.Context) ApplicationSettingsOutput
}

func (*ApplicationSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSettings)(nil)).Elem()
}

func (i *ApplicationSettings) ToApplicationSettingsOutput() ApplicationSettingsOutput {
	return i.ToApplicationSettingsOutputWithContext(context.Background())
}

func (i *ApplicationSettings) ToApplicationSettingsOutputWithContext(ctx context.Context) ApplicationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationSettingsOutput)
}

type ApplicationSettingsOutput struct{ *pulumi.OutputState }

func (ApplicationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationSettings)(nil)).Elem()
}

func (o ApplicationSettingsOutput) ToApplicationSettingsOutput() ApplicationSettingsOutput {
	return o
}

func (o ApplicationSettingsOutput) ToApplicationSettingsOutputWithContext(ctx context.Context) ApplicationSettingsOutput {
	return o
}

func (o ApplicationSettingsOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationSettings) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o ApplicationSettingsOutput) CampaignHook() ApplicationSettingsCampaignHookPtrOutput {
	return o.ApplyT(func(v *ApplicationSettings) ApplicationSettingsCampaignHookPtrOutput { return v.CampaignHook }).(ApplicationSettingsCampaignHookPtrOutput)
}

func (o ApplicationSettingsOutput) CloudWatchMetricsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApplicationSettings) pulumi.BoolPtrOutput { return v.CloudWatchMetricsEnabled }).(pulumi.BoolPtrOutput)
}

func (o ApplicationSettingsOutput) Limits() ApplicationSettingsLimitsPtrOutput {
	return o.ApplyT(func(v *ApplicationSettings) ApplicationSettingsLimitsPtrOutput { return v.Limits }).(ApplicationSettingsLimitsPtrOutput)
}

func (o ApplicationSettingsOutput) QuietTime() ApplicationSettingsQuietTimePtrOutput {
	return o.ApplyT(func(v *ApplicationSettings) ApplicationSettingsQuietTimePtrOutput { return v.QuietTime }).(ApplicationSettingsQuietTimePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationSettingsInput)(nil)).Elem(), &ApplicationSettings{})
	pulumi.RegisterOutputType(ApplicationSettingsOutput{})
}
