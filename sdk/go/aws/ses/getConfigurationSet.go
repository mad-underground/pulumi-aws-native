// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::SES::ConfigurationSet.
func LookupConfigurationSet(ctx *pulumi.Context, args *LookupConfigurationSetArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationSetResult, error) {
	var rv LookupConfigurationSetResult
	err := ctx.Invoke("aws-native:ses:getConfigurationSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationSetArgs struct {
	// The name of the configuration set.
	Name string `pulumi:"name"`
}

type LookupConfigurationSetResult struct {
	DeliveryOptions    *ConfigurationSetDeliveryOptions    `pulumi:"deliveryOptions"`
	ReputationOptions  *ConfigurationSetReputationOptions  `pulumi:"reputationOptions"`
	SendingOptions     *ConfigurationSetSendingOptions     `pulumi:"sendingOptions"`
	SuppressionOptions *ConfigurationSetSuppressionOptions `pulumi:"suppressionOptions"`
	TrackingOptions    *ConfigurationSetTrackingOptions    `pulumi:"trackingOptions"`
}

func LookupConfigurationSetOutput(ctx *pulumi.Context, args LookupConfigurationSetOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationSetResult, error) {
			args := v.(LookupConfigurationSetArgs)
			r, err := LookupConfigurationSet(ctx, &args, opts...)
			var s LookupConfigurationSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationSetResultOutput)
}

type LookupConfigurationSetOutputArgs struct {
	// The name of the configuration set.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupConfigurationSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationSetArgs)(nil)).Elem()
}

type LookupConfigurationSetResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationSetResult)(nil)).Elem()
}

func (o LookupConfigurationSetResultOutput) ToLookupConfigurationSetResultOutput() LookupConfigurationSetResultOutput {
	return o
}

func (o LookupConfigurationSetResultOutput) ToLookupConfigurationSetResultOutputWithContext(ctx context.Context) LookupConfigurationSetResultOutput {
	return o
}

func (o LookupConfigurationSetResultOutput) DeliveryOptions() ConfigurationSetDeliveryOptionsPtrOutput {
	return o.ApplyT(func(v LookupConfigurationSetResult) *ConfigurationSetDeliveryOptions { return v.DeliveryOptions }).(ConfigurationSetDeliveryOptionsPtrOutput)
}

func (o LookupConfigurationSetResultOutput) ReputationOptions() ConfigurationSetReputationOptionsPtrOutput {
	return o.ApplyT(func(v LookupConfigurationSetResult) *ConfigurationSetReputationOptions { return v.ReputationOptions }).(ConfigurationSetReputationOptionsPtrOutput)
}

func (o LookupConfigurationSetResultOutput) SendingOptions() ConfigurationSetSendingOptionsPtrOutput {
	return o.ApplyT(func(v LookupConfigurationSetResult) *ConfigurationSetSendingOptions { return v.SendingOptions }).(ConfigurationSetSendingOptionsPtrOutput)
}

func (o LookupConfigurationSetResultOutput) SuppressionOptions() ConfigurationSetSuppressionOptionsPtrOutput {
	return o.ApplyT(func(v LookupConfigurationSetResult) *ConfigurationSetSuppressionOptions { return v.SuppressionOptions }).(ConfigurationSetSuppressionOptionsPtrOutput)
}

func (o LookupConfigurationSetResultOutput) TrackingOptions() ConfigurationSetTrackingOptionsPtrOutput {
	return o.ApplyT(func(v LookupConfigurationSetResult) *ConfigurationSetTrackingOptions { return v.TrackingOptions }).(ConfigurationSetTrackingOptionsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationSetResultOutput{})
}
