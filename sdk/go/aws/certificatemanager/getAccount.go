// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package certificatemanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::CertificateManager::Account.
func LookupAccount(ctx *pulumi.Context, args *LookupAccountArgs, opts ...pulumi.InvokeOption) (*LookupAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccountResult
	err := ctx.Invoke("aws-native:certificatemanager:getAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAccountArgs struct {
	AccountId string `pulumi:"accountId"`
}

type LookupAccountResult struct {
	AccountId                 *string                           `pulumi:"accountId"`
	ExpiryEventsConfiguration *AccountExpiryEventsConfiguration `pulumi:"expiryEventsConfiguration"`
}

func LookupAccountOutput(ctx *pulumi.Context, args LookupAccountOutputArgs, opts ...pulumi.InvokeOption) LookupAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccountResult, error) {
			args := v.(LookupAccountArgs)
			r, err := LookupAccount(ctx, &args, opts...)
			var s LookupAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccountResultOutput)
}

type LookupAccountOutputArgs struct {
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (LookupAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountArgs)(nil)).Elem()
}

type LookupAccountResultOutput struct{ *pulumi.OutputState }

func (LookupAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccountResult)(nil)).Elem()
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutput() LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) ToLookupAccountResultOutputWithContext(ctx context.Context) LookupAccountResultOutput {
	return o
}

func (o LookupAccountResultOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o LookupAccountResultOutput) ExpiryEventsConfiguration() AccountExpiryEventsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupAccountResult) *AccountExpiryEventsConfiguration { return v.ExpiryEventsConfiguration }).(AccountExpiryEventsConfigurationPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccountResultOutput{})
}
