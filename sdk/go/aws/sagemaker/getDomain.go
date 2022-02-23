// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::Domain
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("aws-native:sagemaker:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	// The domain name.
	DomainId string `pulumi:"domainId"`
}

type LookupDomainResult struct {
	// The default user settings.
	DefaultUserSettings *DomainUserSettings `pulumi:"defaultUserSettings"`
	// The Amazon Resource Name (ARN) of the created domain.
	DomainArn *string `pulumi:"domainArn"`
	// The domain name.
	DomainId *string `pulumi:"domainId"`
	// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
	HomeEfsFileSystemId *string `pulumi:"homeEfsFileSystemId"`
	// The SSO managed application instance ID.
	SingleSignOnManagedApplicationInstanceId *string `pulumi:"singleSignOnManagedApplicationInstanceId"`
	// The URL to the created domain.
	Url *string `pulumi:"url"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			return *r, err
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	// The domain name.
	DomainId pulumi.StringInput `pulumi:"domainId"`
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

// The default user settings.
func (o LookupDomainResultOutput) DefaultUserSettings() DomainUserSettingsPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *DomainUserSettings { return v.DefaultUserSettings }).(DomainUserSettingsPtrOutput)
}

// The Amazon Resource Name (ARN) of the created domain.
func (o LookupDomainResultOutput) DomainArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DomainArn }).(pulumi.StringPtrOutput)
}

// The domain name.
func (o LookupDomainResultOutput) DomainId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.DomainId }).(pulumi.StringPtrOutput)
}

// The ID of the Amazon Elastic File System (EFS) managed by this Domain.
func (o LookupDomainResultOutput) HomeEfsFileSystemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.HomeEfsFileSystemId }).(pulumi.StringPtrOutput)
}

// The SSO managed application instance ID.
func (o LookupDomainResultOutput) SingleSignOnManagedApplicationInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.SingleSignOnManagedApplicationInstanceId }).(pulumi.StringPtrOutput)
}

// The URL to the created domain.
func (o LookupDomainResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}