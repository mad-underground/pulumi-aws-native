// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package amplify

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("aws-native:amplify:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupDomainResult struct {
	Arn                           *string                    `pulumi:"arn"`
	AutoSubDomainCreationPatterns []string                   `pulumi:"autoSubDomainCreationPatterns"`
	AutoSubDomainIamRole          *string                    `pulumi:"autoSubDomainIamRole"`
	Certificate                   *DomainCertificate         `pulumi:"certificate"`
	CertificateRecord             *string                    `pulumi:"certificateRecord"`
	CertificateSettings           *DomainCertificateSettings `pulumi:"certificateSettings"`
	DomainStatus                  *string                    `pulumi:"domainStatus"`
	EnableAutoSubDomain           *bool                      `pulumi:"enableAutoSubDomain"`
	StatusReason                  *string                    `pulumi:"statusReason"`
	SubDomainSettings             []DomainSubDomainSetting   `pulumi:"subDomainSettings"`
	UpdateStatus                  *string                    `pulumi:"updateStatus"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) AutoSubDomainCreationPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []string { return v.AutoSubDomainCreationPatterns }).(pulumi.StringArrayOutput)
}

func (o LookupDomainResultOutput) AutoSubDomainIamRole() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.AutoSubDomainIamRole }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) Certificate() DomainCertificatePtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *DomainCertificate { return v.Certificate }).(DomainCertificatePtrOutput)
}

func (o LookupDomainResultOutput) CertificateRecord() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.CertificateRecord }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) CertificateSettings() DomainCertificateSettingsPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *DomainCertificateSettings { return v.CertificateSettings }).(DomainCertificateSettingsPtrOutput)
}

func (o LookupDomainResultOutput) DomainStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DomainStatus }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) EnableAutoSubDomain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *bool { return v.EnableAutoSubDomain }).(pulumi.BoolPtrOutput)
}

func (o LookupDomainResultOutput) StatusReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.StatusReason }).(pulumi.StringPtrOutput)
}

func (o LookupDomainResultOutput) SubDomainSettings() DomainSubDomainSettingArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []DomainSubDomainSetting { return v.SubDomainSettings }).(DomainSubDomainSettingArrayOutput)
}

func (o LookupDomainResultOutput) UpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.UpdateStatus }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
