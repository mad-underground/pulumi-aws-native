// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::ApplicationSettings
func LookupApplicationSettings(ctx *pulumi.Context, args *LookupApplicationSettingsArgs, opts ...pulumi.InvokeOption) (*LookupApplicationSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationSettingsResult
	err := ctx.Invoke("aws-native:pinpoint:getApplicationSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApplicationSettingsArgs struct {
	Id string `pulumi:"id"`
}

type LookupApplicationSettingsResult struct {
	CampaignHook             *ApplicationSettingsCampaignHook `pulumi:"campaignHook"`
	CloudWatchMetricsEnabled *bool                            `pulumi:"cloudWatchMetricsEnabled"`
	Id                       *string                          `pulumi:"id"`
	Limits                   *ApplicationSettingsLimits       `pulumi:"limits"`
	QuietTime                *ApplicationSettingsQuietTime    `pulumi:"quietTime"`
}

func LookupApplicationSettingsOutput(ctx *pulumi.Context, args LookupApplicationSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationSettingsResult, error) {
			args := v.(LookupApplicationSettingsArgs)
			r, err := LookupApplicationSettings(ctx, &args, opts...)
			var s LookupApplicationSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationSettingsResultOutput)
}

type LookupApplicationSettingsOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApplicationSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSettingsArgs)(nil)).Elem()
}

type LookupApplicationSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationSettingsResult)(nil)).Elem()
}

func (o LookupApplicationSettingsResultOutput) ToLookupApplicationSettingsResultOutput() LookupApplicationSettingsResultOutput {
	return o
}

func (o LookupApplicationSettingsResultOutput) ToLookupApplicationSettingsResultOutputWithContext(ctx context.Context) LookupApplicationSettingsResultOutput {
	return o
}

func (o LookupApplicationSettingsResultOutput) CampaignHook() ApplicationSettingsCampaignHookPtrOutput {
	return o.ApplyT(func(v LookupApplicationSettingsResult) *ApplicationSettingsCampaignHook { return v.CampaignHook }).(ApplicationSettingsCampaignHookPtrOutput)
}

func (o LookupApplicationSettingsResultOutput) CloudWatchMetricsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationSettingsResult) *bool { return v.CloudWatchMetricsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupApplicationSettingsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationSettingsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationSettingsResultOutput) Limits() ApplicationSettingsLimitsPtrOutput {
	return o.ApplyT(func(v LookupApplicationSettingsResult) *ApplicationSettingsLimits { return v.Limits }).(ApplicationSettingsLimitsPtrOutput)
}

func (o LookupApplicationSettingsResultOutput) QuietTime() ApplicationSettingsQuietTimePtrOutput {
	return o.ApplyT(func(v LookupApplicationSettingsResult) *ApplicationSettingsQuietTime { return v.QuietTime }).(ApplicationSettingsQuietTimePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationSettingsResultOutput{})
}
