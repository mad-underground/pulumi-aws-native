// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type TopicLoggingConfig struct {
	// The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch
	FailureFeedbackRoleArn *string `pulumi:"failureFeedbackRoleArn"`
	// Indicates one of the supported protocols for the SNS topic
	Protocol TopicLoggingConfigProtocol `pulumi:"protocol"`
	// The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch
	SuccessFeedbackRoleArn *string `pulumi:"successFeedbackRoleArn"`
	// The percentage of successful message deliveries to be logged in Amazon CloudWatch. Valid percentage values range from 0 to 100
	SuccessFeedbackSampleRate *string `pulumi:"successFeedbackSampleRate"`
}

// TopicLoggingConfigInput is an input type that accepts TopicLoggingConfigArgs and TopicLoggingConfigOutput values.
// You can construct a concrete instance of `TopicLoggingConfigInput` via:
//
//	TopicLoggingConfigArgs{...}
type TopicLoggingConfigInput interface {
	pulumi.Input

	ToTopicLoggingConfigOutput() TopicLoggingConfigOutput
	ToTopicLoggingConfigOutputWithContext(context.Context) TopicLoggingConfigOutput
}

type TopicLoggingConfigArgs struct {
	// The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch
	FailureFeedbackRoleArn pulumi.StringPtrInput `pulumi:"failureFeedbackRoleArn"`
	// Indicates one of the supported protocols for the SNS topic
	Protocol TopicLoggingConfigProtocolInput `pulumi:"protocol"`
	// The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch
	SuccessFeedbackRoleArn pulumi.StringPtrInput `pulumi:"successFeedbackRoleArn"`
	// The percentage of successful message deliveries to be logged in Amazon CloudWatch. Valid percentage values range from 0 to 100
	SuccessFeedbackSampleRate pulumi.StringPtrInput `pulumi:"successFeedbackSampleRate"`
}

func (TopicLoggingConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicLoggingConfig)(nil)).Elem()
}

func (i TopicLoggingConfigArgs) ToTopicLoggingConfigOutput() TopicLoggingConfigOutput {
	return i.ToTopicLoggingConfigOutputWithContext(context.Background())
}

func (i TopicLoggingConfigArgs) ToTopicLoggingConfigOutputWithContext(ctx context.Context) TopicLoggingConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicLoggingConfigOutput)
}

// TopicLoggingConfigArrayInput is an input type that accepts TopicLoggingConfigArray and TopicLoggingConfigArrayOutput values.
// You can construct a concrete instance of `TopicLoggingConfigArrayInput` via:
//
//	TopicLoggingConfigArray{ TopicLoggingConfigArgs{...} }
type TopicLoggingConfigArrayInput interface {
	pulumi.Input

	ToTopicLoggingConfigArrayOutput() TopicLoggingConfigArrayOutput
	ToTopicLoggingConfigArrayOutputWithContext(context.Context) TopicLoggingConfigArrayOutput
}

type TopicLoggingConfigArray []TopicLoggingConfigInput

func (TopicLoggingConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicLoggingConfig)(nil)).Elem()
}

func (i TopicLoggingConfigArray) ToTopicLoggingConfigArrayOutput() TopicLoggingConfigArrayOutput {
	return i.ToTopicLoggingConfigArrayOutputWithContext(context.Background())
}

func (i TopicLoggingConfigArray) ToTopicLoggingConfigArrayOutputWithContext(ctx context.Context) TopicLoggingConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicLoggingConfigArrayOutput)
}

type TopicLoggingConfigOutput struct{ *pulumi.OutputState }

func (TopicLoggingConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicLoggingConfig)(nil)).Elem()
}

func (o TopicLoggingConfigOutput) ToTopicLoggingConfigOutput() TopicLoggingConfigOutput {
	return o
}

func (o TopicLoggingConfigOutput) ToTopicLoggingConfigOutputWithContext(ctx context.Context) TopicLoggingConfigOutput {
	return o
}

// The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch
func (o TopicLoggingConfigOutput) FailureFeedbackRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TopicLoggingConfig) *string { return v.FailureFeedbackRoleArn }).(pulumi.StringPtrOutput)
}

// Indicates one of the supported protocols for the SNS topic
func (o TopicLoggingConfigOutput) Protocol() TopicLoggingConfigProtocolOutput {
	return o.ApplyT(func(v TopicLoggingConfig) TopicLoggingConfigProtocol { return v.Protocol }).(TopicLoggingConfigProtocolOutput)
}

// The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch
func (o TopicLoggingConfigOutput) SuccessFeedbackRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TopicLoggingConfig) *string { return v.SuccessFeedbackRoleArn }).(pulumi.StringPtrOutput)
}

// The percentage of successful message deliveries to be logged in Amazon CloudWatch. Valid percentage values range from 0 to 100
func (o TopicLoggingConfigOutput) SuccessFeedbackSampleRate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TopicLoggingConfig) *string { return v.SuccessFeedbackSampleRate }).(pulumi.StringPtrOutput)
}

type TopicLoggingConfigArrayOutput struct{ *pulumi.OutputState }

func (TopicLoggingConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicLoggingConfig)(nil)).Elem()
}

func (o TopicLoggingConfigArrayOutput) ToTopicLoggingConfigArrayOutput() TopicLoggingConfigArrayOutput {
	return o
}

func (o TopicLoggingConfigArrayOutput) ToTopicLoggingConfigArrayOutputWithContext(ctx context.Context) TopicLoggingConfigArrayOutput {
	return o
}

func (o TopicLoggingConfigArrayOutput) Index(i pulumi.IntInput) TopicLoggingConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicLoggingConfig {
		return vs[0].([]TopicLoggingConfig)[vs[1].(int)]
	}).(TopicLoggingConfigOutput)
}

type TopicSubscription struct {
	Endpoint string `pulumi:"endpoint"`
	Protocol string `pulumi:"protocol"`
}

// TopicSubscriptionInput is an input type that accepts TopicSubscriptionArgs and TopicSubscriptionOutput values.
// You can construct a concrete instance of `TopicSubscriptionInput` via:
//
//	TopicSubscriptionArgs{...}
type TopicSubscriptionInput interface {
	pulumi.Input

	ToTopicSubscriptionOutput() TopicSubscriptionOutput
	ToTopicSubscriptionOutputWithContext(context.Context) TopicSubscriptionOutput
}

type TopicSubscriptionArgs struct {
	Endpoint pulumi.StringInput `pulumi:"endpoint"`
	Protocol pulumi.StringInput `pulumi:"protocol"`
}

func (TopicSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil)).Elem()
}

func (i TopicSubscriptionArgs) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return i.ToTopicSubscriptionOutputWithContext(context.Background())
}

func (i TopicSubscriptionArgs) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionOutput)
}

// TopicSubscriptionArrayInput is an input type that accepts TopicSubscriptionArray and TopicSubscriptionArrayOutput values.
// You can construct a concrete instance of `TopicSubscriptionArrayInput` via:
//
//	TopicSubscriptionArray{ TopicSubscriptionArgs{...} }
type TopicSubscriptionArrayInput interface {
	pulumi.Input

	ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput
	ToTopicSubscriptionArrayOutputWithContext(context.Context) TopicSubscriptionArrayOutput
}

type TopicSubscriptionArray []TopicSubscriptionInput

func (TopicSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicSubscription)(nil)).Elem()
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return i.ToTopicSubscriptionArrayOutputWithContext(context.Background())
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionArrayOutput)
}

type TopicSubscriptionOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil)).Elem()
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v TopicSubscription) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o TopicSubscriptionOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v TopicSubscription) string { return v.Protocol }).(pulumi.StringOutput)
}

type TopicSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicSubscription)(nil)).Elem()
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) Index(i pulumi.IntInput) TopicSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicSubscription {
		return vs[0].([]TopicSubscription)[vs[1].(int)]
	}).(TopicSubscriptionOutput)
}

type TopicTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 characters in length.
	Value string `pulumi:"value"`
}

// TopicTagInput is an input type that accepts TopicTagArgs and TopicTagOutput values.
// You can construct a concrete instance of `TopicTagInput` via:
//
//	TopicTagArgs{...}
type TopicTagInput interface {
	pulumi.Input

	ToTopicTagOutput() TopicTagOutput
	ToTopicTagOutputWithContext(context.Context) TopicTagOutput
}

type TopicTagArgs struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.
	Key pulumi.StringInput `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 characters in length.
	Value pulumi.StringInput `pulumi:"value"`
}

func (TopicTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicTag)(nil)).Elem()
}

func (i TopicTagArgs) ToTopicTagOutput() TopicTagOutput {
	return i.ToTopicTagOutputWithContext(context.Background())
}

func (i TopicTagArgs) ToTopicTagOutputWithContext(ctx context.Context) TopicTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicTagOutput)
}

// TopicTagArrayInput is an input type that accepts TopicTagArray and TopicTagArrayOutput values.
// You can construct a concrete instance of `TopicTagArrayInput` via:
//
//	TopicTagArray{ TopicTagArgs{...} }
type TopicTagArrayInput interface {
	pulumi.Input

	ToTopicTagArrayOutput() TopicTagArrayOutput
	ToTopicTagArrayOutputWithContext(context.Context) TopicTagArrayOutput
}

type TopicTagArray []TopicTagInput

func (TopicTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicTag)(nil)).Elem()
}

func (i TopicTagArray) ToTopicTagArrayOutput() TopicTagArrayOutput {
	return i.ToTopicTagArrayOutputWithContext(context.Background())
}

func (i TopicTagArray) ToTopicTagArrayOutputWithContext(ctx context.Context) TopicTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicTagArrayOutput)
}

type TopicTagOutput struct{ *pulumi.OutputState }

func (TopicTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicTag)(nil)).Elem()
}

func (o TopicTagOutput) ToTopicTagOutput() TopicTagOutput {
	return o
}

func (o TopicTagOutput) ToTopicTagOutputWithContext(ctx context.Context) TopicTagOutput {
	return o
}

// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, `_`, `.`, `/`, `=`, `+`, and `-`.
func (o TopicTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v TopicTag) string { return v.Key }).(pulumi.StringOutput)
}

// The value for the tag. You can specify a value that is 0 to 256 characters in length.
func (o TopicTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TopicTag) string { return v.Value }).(pulumi.StringOutput)
}

type TopicTagArrayOutput struct{ *pulumi.OutputState }

func (TopicTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicTag)(nil)).Elem()
}

func (o TopicTagArrayOutput) ToTopicTagArrayOutput() TopicTagArrayOutput {
	return o
}

func (o TopicTagArrayOutput) ToTopicTagArrayOutputWithContext(ctx context.Context) TopicTagArrayOutput {
	return o
}

func (o TopicTagArrayOutput) Index(i pulumi.IntInput) TopicTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicTag {
		return vs[0].([]TopicTag)[vs[1].(int)]
	}).(TopicTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicLoggingConfigInput)(nil)).Elem(), TopicLoggingConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicLoggingConfigArrayInput)(nil)).Elem(), TopicLoggingConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionInput)(nil)).Elem(), TopicSubscriptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicSubscriptionArrayInput)(nil)).Elem(), TopicSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicTagInput)(nil)).Elem(), TopicTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicTagArrayInput)(nil)).Elem(), TopicTagArray{})
	pulumi.RegisterOutputType(TopicLoggingConfigOutput{})
	pulumi.RegisterOutputType(TopicLoggingConfigArrayOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(TopicTagOutput{})
	pulumi.RegisterOutputType(TopicTagArrayOutput{})
}
