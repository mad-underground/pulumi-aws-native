// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package chatbot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Chatbot::SlackChannelConfiguration.
func LookupSlackChannelConfiguration(ctx *pulumi.Context, args *LookupSlackChannelConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupSlackChannelConfigurationResult, error) {
	var rv LookupSlackChannelConfigurationResult
	err := ctx.Invoke("aws-native:chatbot:getSlackChannelConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSlackChannelConfigurationArgs struct {
	// Amazon Resource Name (ARN) of the configuration
	Arn string `pulumi:"arn"`
}

type LookupSlackChannelConfigurationResult struct {
	// Amazon Resource Name (ARN) of the configuration
	Arn *string `pulumi:"arn"`
	// The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	GuardrailPolicies []string `pulumi:"guardrailPolicies"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot
	IamRoleArn *string `pulumi:"iamRoleArn"`
	// Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
	LoggingLevel *string `pulumi:"loggingLevel"`
	// The id of the Slack channel
	SlackChannelId *string `pulumi:"slackChannelId"`
	// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
	SnsTopicArns []string `pulumi:"snsTopicArns"`
	// Enables use of a user role requirement in your chat configuration
	UserRoleRequired *bool `pulumi:"userRoleRequired"`
}

func LookupSlackChannelConfigurationOutput(ctx *pulumi.Context, args LookupSlackChannelConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupSlackChannelConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSlackChannelConfigurationResult, error) {
			args := v.(LookupSlackChannelConfigurationArgs)
			r, err := LookupSlackChannelConfiguration(ctx, &args, opts...)
			return *r, err
		}).(LookupSlackChannelConfigurationResultOutput)
}

type LookupSlackChannelConfigurationOutputArgs struct {
	// Amazon Resource Name (ARN) of the configuration
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupSlackChannelConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSlackChannelConfigurationArgs)(nil)).Elem()
}

type LookupSlackChannelConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupSlackChannelConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSlackChannelConfigurationResult)(nil)).Elem()
}

func (o LookupSlackChannelConfigurationResultOutput) ToLookupSlackChannelConfigurationResultOutput() LookupSlackChannelConfigurationResultOutput {
	return o
}

func (o LookupSlackChannelConfigurationResultOutput) ToLookupSlackChannelConfigurationResultOutputWithContext(ctx context.Context) LookupSlackChannelConfigurationResultOutput {
	return o
}

// Amazon Resource Name (ARN) of the configuration
func (o LookupSlackChannelConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The list of IAM policy ARNs that are applied as channel guardrails. The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
func (o LookupSlackChannelConfigurationResultOutput) GuardrailPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) []string { return v.GuardrailPolicies }).(pulumi.StringArrayOutput)
}

// The ARN of the IAM role that defines the permissions for AWS Chatbot
func (o LookupSlackChannelConfigurationResultOutput) IamRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *string { return v.IamRoleArn }).(pulumi.StringPtrOutput)
}

// Specifies the logging level for this configuration:ERROR,INFO or NONE. This property affects the log entries pushed to Amazon CloudWatch logs
func (o LookupSlackChannelConfigurationResultOutput) LoggingLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *string { return v.LoggingLevel }).(pulumi.StringPtrOutput)
}

// The id of the Slack channel
func (o LookupSlackChannelConfigurationResultOutput) SlackChannelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *string { return v.SlackChannelId }).(pulumi.StringPtrOutput)
}

// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
func (o LookupSlackChannelConfigurationResultOutput) SnsTopicArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) []string { return v.SnsTopicArns }).(pulumi.StringArrayOutput)
}

// Enables use of a user role requirement in your chat configuration
func (o LookupSlackChannelConfigurationResultOutput) UserRoleRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSlackChannelConfigurationResult) *bool { return v.UserRoleRequired }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSlackChannelConfigurationResultOutput{})
}
