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
type Topic struct {
	pulumi.CustomResourceState

	// The archive policy determines the number of days Amazon SNS retains messages. You can set a retention period from 1 to 365 days.
	ArchivePolicy pulumi.AnyOutput `pulumi:"archivePolicy"`
	// Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
	//
	// When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
	ContentBasedDeduplication pulumi.BoolPtrOutput `pulumi:"contentBasedDeduplication"`
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720
	DataProtectionPolicy pulumi.AnyOutput `pulumi:"dataProtectionPolicy"`
	// The display name to use for an Amazon SNS topic with SMS subscriptions.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Set to true to create a FIFO topic.
	FifoTopic pulumi.BoolPtrOutput `pulumi:"fifoTopic"`
	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
	KmsMasterKeyId pulumi.StringPtrOutput `pulumi:"kmsMasterKeyId"`
	// Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
	SignatureVersion pulumi.StringPtrOutput `pulumi:"signatureVersion"`
	// The SNS subscriptions (endpoints) for this topic.
	Subscription TopicSubscriptionArrayOutput `pulumi:"subscription"`
	Tags         TopicTagArrayOutput          `pulumi:"tags"`
	TopicArn     pulumi.StringOutput          `pulumi:"topicArn"`
	// The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with .fifo.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see Name Type.
	TopicName pulumi.StringPtrOutput `pulumi:"topicName"`
	// Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
	TracingConfig pulumi.StringPtrOutput `pulumi:"tracingConfig"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		args = &TopicArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"fifoTopic",
		"topicName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("aws-native:sns:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("aws-native:sns:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
}

type TopicState struct {
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
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
	// Set to true to create a FIFO topic.
	FifoTopic *bool `pulumi:"fifoTopic"`
	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
	KmsMasterKeyId *string `pulumi:"kmsMasterKeyId"`
	// Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
	SignatureVersion *string `pulumi:"signatureVersion"`
	// The SNS subscriptions (endpoints) for this topic.
	Subscription []TopicSubscription `pulumi:"subscription"`
	Tags         []TopicTag          `pulumi:"tags"`
	// The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with .fifo.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see Name Type.
	TopicName *string `pulumi:"topicName"`
	// Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
	TracingConfig *string `pulumi:"tracingConfig"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The archive policy determines the number of days Amazon SNS retains messages. You can set a retention period from 1 to 365 days.
	ArchivePolicy pulumi.Input
	// Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
	//
	// When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30720
	DataProtectionPolicy pulumi.Input
	// The display name to use for an Amazon SNS topic with SMS subscriptions.
	DisplayName pulumi.StringPtrInput
	// Set to true to create a FIFO topic.
	FifoTopic pulumi.BoolPtrInput
	// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
	KmsMasterKeyId pulumi.StringPtrInput
	// Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
	SignatureVersion pulumi.StringPtrInput
	// The SNS subscriptions (endpoints) for this topic.
	Subscription TopicSubscriptionArrayInput
	Tags         TopicTagArrayInput
	// The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with .fifo.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see Name Type.
	TopicName pulumi.StringPtrInput
	// Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
	TracingConfig pulumi.StringPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

func (i *Topic) ToOutput(ctx context.Context) pulumix.Output[*Topic] {
	return pulumix.Output[*Topic]{
		OutputState: i.ToTopicOutputWithContext(ctx).OutputState,
	}
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

func (o TopicOutput) ToOutput(ctx context.Context) pulumix.Output[*Topic] {
	return pulumix.Output[*Topic]{
		OutputState: o.OutputState,
	}
}

// The archive policy determines the number of days Amazon SNS retains messages. You can set a retention period from 1 to 365 days.
func (o TopicOutput) ArchivePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Topic) pulumi.AnyOutput { return v.ArchivePolicy }).(pulumi.AnyOutput)
}

// Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
//
// When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
//
// (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
func (o TopicOutput) ContentBasedDeduplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.ContentBasedDeduplication }).(pulumi.BoolPtrOutput)
}

// The body of the policy document you want to use for this topic.
//
// You can only add one policy per topic.
//
// The policy must be in JSON string format.
//
// Length Constraints: Maximum length of 30720
func (o TopicOutput) DataProtectionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Topic) pulumi.AnyOutput { return v.DataProtectionPolicy }).(pulumi.AnyOutput)
}

// The display name to use for an Amazon SNS topic with SMS subscriptions.
func (o TopicOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Set to true to create a FIFO topic.
func (o TopicOutput) FifoTopic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolPtrOutput { return v.FifoTopic }).(pulumi.BoolPtrOutput)
}

// The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
//
// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
func (o TopicOutput) KmsMasterKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.KmsMasterKeyId }).(pulumi.StringPtrOutput)
}

// Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
func (o TopicOutput) SignatureVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.SignatureVersion }).(pulumi.StringPtrOutput)
}

// The SNS subscriptions (endpoints) for this topic.
func (o TopicOutput) Subscription() TopicSubscriptionArrayOutput {
	return o.ApplyT(func(v *Topic) TopicSubscriptionArrayOutput { return v.Subscription }).(TopicSubscriptionArrayOutput)
}

func (o TopicOutput) Tags() TopicTagArrayOutput {
	return o.ApplyT(func(v *Topic) TopicTagArrayOutput { return v.Tags }).(TopicTagArrayOutput)
}

func (o TopicOutput) TopicArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TopicArn }).(pulumi.StringOutput)
}

// The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with .fifo.
//
// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see Name Type.
func (o TopicOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.TopicName }).(pulumi.StringPtrOutput)
}

// Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
func (o TopicOutput) TracingConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringPtrOutput { return v.TracingConfig }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterOutputType(TopicOutput{})
}
