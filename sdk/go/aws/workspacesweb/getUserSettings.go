// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workspacesweb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::WorkSpacesWeb::UserSettings Resource Type
func LookupUserSettings(ctx *pulumi.Context, args *LookupUserSettingsArgs, opts ...pulumi.InvokeOption) (*LookupUserSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserSettingsResult
	err := ctx.Invoke("aws-native:workspacesweb:getUserSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserSettingsArgs struct {
	UserSettingsArn string `pulumi:"userSettingsArn"`
}

type LookupUserSettingsResult struct {
	AssociatedPortalArns               []string                                        `pulumi:"associatedPortalArns"`
	CookieSynchronizationConfiguration *UserSettingsCookieSynchronizationConfiguration `pulumi:"cookieSynchronizationConfiguration"`
	CopyAllowed                        *UserSettingsEnabledType                        `pulumi:"copyAllowed"`
	DisconnectTimeoutInMinutes         *float64                                        `pulumi:"disconnectTimeoutInMinutes"`
	DownloadAllowed                    *UserSettingsEnabledType                        `pulumi:"downloadAllowed"`
	IdleDisconnectTimeoutInMinutes     *float64                                        `pulumi:"idleDisconnectTimeoutInMinutes"`
	PasteAllowed                       *UserSettingsEnabledType                        `pulumi:"pasteAllowed"`
	PrintAllowed                       *UserSettingsEnabledType                        `pulumi:"printAllowed"`
	Tags                               []UserSettingsTag                               `pulumi:"tags"`
	UploadAllowed                      *UserSettingsEnabledType                        `pulumi:"uploadAllowed"`
	UserSettingsArn                    *string                                         `pulumi:"userSettingsArn"`
}

func LookupUserSettingsOutput(ctx *pulumi.Context, args LookupUserSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupUserSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserSettingsResult, error) {
			args := v.(LookupUserSettingsArgs)
			r, err := LookupUserSettings(ctx, &args, opts...)
			var s LookupUserSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserSettingsResultOutput)
}

type LookupUserSettingsOutputArgs struct {
	UserSettingsArn pulumi.StringInput `pulumi:"userSettingsArn"`
}

func (LookupUserSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsArgs)(nil)).Elem()
}

type LookupUserSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupUserSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserSettingsResult)(nil)).Elem()
}

func (o LookupUserSettingsResultOutput) ToLookupUserSettingsResultOutput() LookupUserSettingsResultOutput {
	return o
}

func (o LookupUserSettingsResultOutput) ToLookupUserSettingsResultOutputWithContext(ctx context.Context) LookupUserSettingsResultOutput {
	return o
}

func (o LookupUserSettingsResultOutput) AssociatedPortalArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) []string { return v.AssociatedPortalArns }).(pulumi.StringArrayOutput)
}

func (o LookupUserSettingsResultOutput) CookieSynchronizationConfiguration() UserSettingsCookieSynchronizationConfigurationPtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *UserSettingsCookieSynchronizationConfiguration {
		return v.CookieSynchronizationConfiguration
	}).(UserSettingsCookieSynchronizationConfigurationPtrOutput)
}

func (o LookupUserSettingsResultOutput) CopyAllowed() UserSettingsEnabledTypePtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *UserSettingsEnabledType { return v.CopyAllowed }).(UserSettingsEnabledTypePtrOutput)
}

func (o LookupUserSettingsResultOutput) DisconnectTimeoutInMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *float64 { return v.DisconnectTimeoutInMinutes }).(pulumi.Float64PtrOutput)
}

func (o LookupUserSettingsResultOutput) DownloadAllowed() UserSettingsEnabledTypePtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *UserSettingsEnabledType { return v.DownloadAllowed }).(UserSettingsEnabledTypePtrOutput)
}

func (o LookupUserSettingsResultOutput) IdleDisconnectTimeoutInMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *float64 { return v.IdleDisconnectTimeoutInMinutes }).(pulumi.Float64PtrOutput)
}

func (o LookupUserSettingsResultOutput) PasteAllowed() UserSettingsEnabledTypePtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *UserSettingsEnabledType { return v.PasteAllowed }).(UserSettingsEnabledTypePtrOutput)
}

func (o LookupUserSettingsResultOutput) PrintAllowed() UserSettingsEnabledTypePtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *UserSettingsEnabledType { return v.PrintAllowed }).(UserSettingsEnabledTypePtrOutput)
}

func (o LookupUserSettingsResultOutput) Tags() UserSettingsTagArrayOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) []UserSettingsTag { return v.Tags }).(UserSettingsTagArrayOutput)
}

func (o LookupUserSettingsResultOutput) UploadAllowed() UserSettingsEnabledTypePtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *UserSettingsEnabledType { return v.UploadAllowed }).(UserSettingsEnabledTypePtrOutput)
}

func (o LookupUserSettingsResultOutput) UserSettingsArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserSettingsResult) *string { return v.UserSettingsArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserSettingsResultOutput{})
}
