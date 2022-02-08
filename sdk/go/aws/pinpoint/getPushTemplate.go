// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::PushTemplate
func LookupPushTemplate(ctx *pulumi.Context, args *LookupPushTemplateArgs, opts ...pulumi.InvokeOption) (*LookupPushTemplateResult, error) {
	var rv LookupPushTemplateResult
	err := ctx.Invoke("aws-native:pinpoint:getPushTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPushTemplateArgs struct {
	Id string `pulumi:"id"`
}

type LookupPushTemplateResult struct {
	ADM                  *PushTemplateAndroidPushNotificationTemplate `pulumi:"aDM"`
	APNS                 *PushTemplateAPNSPushNotificationTemplate    `pulumi:"aPNS"`
	Arn                  *string                                      `pulumi:"arn"`
	Baidu                *PushTemplateAndroidPushNotificationTemplate `pulumi:"baidu"`
	Default              *PushTemplateDefaultPushNotificationTemplate `pulumi:"default"`
	DefaultSubstitutions *string                                      `pulumi:"defaultSubstitutions"`
	GCM                  *PushTemplateAndroidPushNotificationTemplate `pulumi:"gCM"`
	Id                   *string                                      `pulumi:"id"`
	Tags                 interface{}                                  `pulumi:"tags"`
	TemplateDescription  *string                                      `pulumi:"templateDescription"`
}

func LookupPushTemplateOutput(ctx *pulumi.Context, args LookupPushTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupPushTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPushTemplateResult, error) {
			args := v.(LookupPushTemplateArgs)
			r, err := LookupPushTemplate(ctx, &args, opts...)
			return *r, err
		}).(LookupPushTemplateResultOutput)
}

type LookupPushTemplateOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupPushTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPushTemplateArgs)(nil)).Elem()
}

type LookupPushTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupPushTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPushTemplateResult)(nil)).Elem()
}

func (o LookupPushTemplateResultOutput) ToLookupPushTemplateResultOutput() LookupPushTemplateResultOutput {
	return o
}

func (o LookupPushTemplateResultOutput) ToLookupPushTemplateResultOutputWithContext(ctx context.Context) LookupPushTemplateResultOutput {
	return o
}

func (o LookupPushTemplateResultOutput) ADM() PushTemplateAndroidPushNotificationTemplatePtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *PushTemplateAndroidPushNotificationTemplate { return v.ADM }).(PushTemplateAndroidPushNotificationTemplatePtrOutput)
}

func (o LookupPushTemplateResultOutput) APNS() PushTemplateAPNSPushNotificationTemplatePtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *PushTemplateAPNSPushNotificationTemplate { return v.APNS }).(PushTemplateAPNSPushNotificationTemplatePtrOutput)
}

func (o LookupPushTemplateResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupPushTemplateResultOutput) Baidu() PushTemplateAndroidPushNotificationTemplatePtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *PushTemplateAndroidPushNotificationTemplate { return v.Baidu }).(PushTemplateAndroidPushNotificationTemplatePtrOutput)
}

func (o LookupPushTemplateResultOutput) Default() PushTemplateDefaultPushNotificationTemplatePtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *PushTemplateDefaultPushNotificationTemplate { return v.Default }).(PushTemplateDefaultPushNotificationTemplatePtrOutput)
}

func (o LookupPushTemplateResultOutput) DefaultSubstitutions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *string { return v.DefaultSubstitutions }).(pulumi.StringPtrOutput)
}

func (o LookupPushTemplateResultOutput) GCM() PushTemplateAndroidPushNotificationTemplatePtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *PushTemplateAndroidPushNotificationTemplate { return v.GCM }).(PushTemplateAndroidPushNotificationTemplatePtrOutput)
}

func (o LookupPushTemplateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupPushTemplateResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupPushTemplateResultOutput) TemplateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPushTemplateResult) *string { return v.TemplateDescription }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPushTemplateResultOutput{})
}
