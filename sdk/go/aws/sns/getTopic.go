// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::SNS::Topic
func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTopicResult
	err := ctx.Invoke("aws-native:sns:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	TopicArn string `pulumi:"topicArn"`
}

type LookupTopicResult struct {
	// The archive policy determines the number of days Amazon SNS retains messages. You can set a retention period from 1 to 365 days.
	ArchivePolicy interface{} `pulumi:"archivePolicy"`
	// Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
	//
	// When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720
	DataProtectionPolicy interface{} `pulumi:"dataProtectionPolicy"`
	// The display name to use for an Amazon SNS topic with SMS subscriptions.
	DisplayName *string `pulumi:"displayName"`
	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
	KmsMasterKeyId *string `pulumi:"kmsMasterKeyId"`
	// Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
	SignatureVersion *string `pulumi:"signatureVersion"`
	// The SNS subscriptions (endpoints) for this topic.
	Subscription []TopicSubscription `pulumi:"subscription"`
	Tags         []TopicTag          `pulumi:"tags"`
	TopicArn     *string             `pulumi:"topicArn"`
	// Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
	TracingConfig *string `pulumi:"tracingConfig"`
}

func LookupTopicOutput(ctx *pulumi.Context, args LookupTopicOutputArgs, opts ...pulumi.InvokeOption) LookupTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicResult, error) {
			args := v.(LookupTopicArgs)
			r, err := LookupTopic(ctx, &args, opts...)
			var s LookupTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicResultOutput)
}

type LookupTopicOutputArgs struct {
	TopicArn pulumi.StringInput `pulumi:"topicArn"`
}

func (LookupTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicArgs)(nil)).Elem()
}

type LookupTopicResultOutput struct{ *pulumi.OutputState }

func (LookupTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicResult)(nil)).Elem()
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutput() LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutputWithContext(ctx context.Context) LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupTopicResult] {
	return pulumix.Output[LookupTopicResult]{
		OutputState: o.OutputState,
	}
}

// The archive policy determines the number of days Amazon SNS retains messages. You can set a retention period from 1 to 365 days.
func (o LookupTopicResultOutput) ArchivePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTopicResult) interface{} { return v.ArchivePolicy }).(pulumi.AnyOutput)
}

// Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
//
// When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
//
// (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
func (o LookupTopicResultOutput) ContentBasedDeduplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.ContentBasedDeduplication }).(pulumi.BoolPtrOutput)
}

// The body of the policy document you want to use for this topic.
//
// You can only add one policy per topic.
//
// The policy must be in JSON string format.
//
// Length Constraints: Maximum length of 30720
func (o LookupTopicResultOutput) DataProtectionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupTopicResult) interface{} { return v.DataProtectionPolicy }).(pulumi.AnyOutput)
}

// The display name to use for an Amazon SNS topic with SMS subscriptions.
func (o LookupTopicResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
//
// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
func (o LookupTopicResultOutput) KmsMasterKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.KmsMasterKeyId }).(pulumi.StringPtrOutput)
}

// Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
func (o LookupTopicResultOutput) SignatureVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.SignatureVersion }).(pulumi.StringPtrOutput)
}

// The SNS subscriptions (endpoints) for this topic.
func (o LookupTopicResultOutput) Subscription() TopicSubscriptionArrayOutput {
	return o.ApplyT(func(v LookupTopicResult) []TopicSubscription { return v.Subscription }).(TopicSubscriptionArrayOutput)
}

func (o LookupTopicResultOutput) Tags() TopicTagArrayOutput {
	return o.ApplyT(func(v LookupTopicResult) []TopicTag { return v.Tags }).(TopicTagArrayOutput)
}

func (o LookupTopicResultOutput) TopicArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.TopicArn }).(pulumi.StringPtrOutput)
}

// Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
func (o LookupTopicResultOutput) TracingConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.TracingConfig }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicResultOutput{})
}
