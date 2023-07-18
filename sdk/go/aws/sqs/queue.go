// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SQS::Queue
type Queue struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the queue.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.
	ContentBasedDeduplication pulumi.BoolPtrOutput `pulumi:"contentBasedDeduplication"`
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.
	DeduplicationScope pulumi.StringPtrOutput `pulumi:"deduplicationScope"`
	// The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.
	DelaySeconds pulumi.IntPtrOutput `pulumi:"delaySeconds"`
	// If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.
	FifoQueue pulumi.BoolPtrOutput `pulumi:"fifoQueue"`
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.
	FifoThroughputLimit pulumi.StringPtrOutput `pulumi:"fifoThroughputLimit"`
	// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds pulumi.IntPtrOutput `pulumi:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.
	KmsMasterKeyId pulumi.StringPtrOutput `pulumi:"kmsMasterKeyId"`
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).
	MaximumMessageSize pulumi.IntPtrOutput `pulumi:"maximumMessageSize"`
	// The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).
	MessageRetentionPeriod pulumi.IntPtrOutput `pulumi:"messageRetentionPeriod"`
	// A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.
	QueueName pulumi.StringPtrOutput `pulumi:"queueName"`
	// URL of the source queue.
	QueueUrl pulumi.StringOutput `pulumi:"queueUrl"`
	// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.
	ReceiveMessageWaitTimeSeconds pulumi.IntPtrOutput `pulumi:"receiveMessageWaitTimeSeconds"`
	// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
	RedriveAllowPolicy pulumi.AnyOutput `pulumi:"redriveAllowPolicy"`
	// A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.
	RedrivePolicy pulumi.AnyOutput `pulumi:"redrivePolicy"`
	// Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).
	SqsManagedSseEnabled pulumi.BoolPtrOutput `pulumi:"sqsManagedSseEnabled"`
	// The tags that you attach to this queue.
	Tags QueueTagArrayOutput `pulumi:"tags"`
	// The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
	VisibilityTimeout pulumi.IntPtrOutput `pulumi:"visibilityTimeout"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil {
		args = &QueueArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Queue
	err := ctx.RegisterResource("aws-native:sqs:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("aws-native:sqs:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
}

type QueueState struct {
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.
	ContentBasedDeduplication *bool `pulumi:"contentBasedDeduplication"`
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.
	DeduplicationScope *string `pulumi:"deduplicationScope"`
	// The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.
	DelaySeconds *int `pulumi:"delaySeconds"`
	// If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.
	FifoQueue *bool `pulumi:"fifoQueue"`
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.
	FifoThroughputLimit *string `pulumi:"fifoThroughputLimit"`
	// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds *int `pulumi:"kmsDataKeyReusePeriodSeconds"`
	// The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.
	KmsMasterKeyId *string `pulumi:"kmsMasterKeyId"`
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).
	MaximumMessageSize *int `pulumi:"maximumMessageSize"`
	// The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).
	MessageRetentionPeriod *int `pulumi:"messageRetentionPeriod"`
	// A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.
	QueueName *string `pulumi:"queueName"`
	// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.
	ReceiveMessageWaitTimeSeconds *int `pulumi:"receiveMessageWaitTimeSeconds"`
	// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
	RedriveAllowPolicy interface{} `pulumi:"redriveAllowPolicy"`
	// A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.
	RedrivePolicy interface{} `pulumi:"redrivePolicy"`
	// Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).
	SqsManagedSseEnabled *bool `pulumi:"sqsManagedSseEnabled"`
	// The tags that you attach to this queue.
	Tags []QueueTag `pulumi:"tags"`
	// The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
	VisibilityTimeout *int `pulumi:"visibilityTimeout"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.
	ContentBasedDeduplication pulumi.BoolPtrInput
	// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.
	DeduplicationScope pulumi.StringPtrInput
	// The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.
	DelaySeconds pulumi.IntPtrInput
	// If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.
	FifoQueue pulumi.BoolPtrInput
	// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.
	FifoThroughputLimit pulumi.StringPtrInput
	// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
	KmsDataKeyReusePeriodSeconds pulumi.IntPtrInput
	// The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.
	KmsMasterKeyId pulumi.StringPtrInput
	// The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).
	MaximumMessageSize pulumi.IntPtrInput
	// The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).
	MessageRetentionPeriod pulumi.IntPtrInput
	// A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.
	QueueName pulumi.StringPtrInput
	// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.
	ReceiveMessageWaitTimeSeconds pulumi.IntPtrInput
	// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
	RedriveAllowPolicy pulumi.Input
	// A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.
	RedrivePolicy pulumi.Input
	// Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).
	SqsManagedSseEnabled pulumi.BoolPtrInput
	// The tags that you attach to this queue.
	Tags QueueTagArrayInput
	// The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
	VisibilityTimeout pulumi.IntPtrInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}

type QueueInput interface {
	pulumi.Input

	ToQueueOutput() QueueOutput
	ToQueueOutputWithContext(ctx context.Context) QueueOutput
}

func (*Queue) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (i *Queue) ToQueueOutput() QueueOutput {
	return i.ToQueueOutputWithContext(context.Background())
}

func (i *Queue) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueOutput)
}

type QueueOutput struct{ *pulumi.OutputState }

func (QueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Queue)(nil)).Elem()
}

func (o QueueOutput) ToQueueOutput() QueueOutput {
	return o
}

func (o QueueOutput) ToQueueOutputWithContext(ctx context.Context) QueueOutput {
	return o
}

// Amazon Resource Name (ARN) of the queue.
func (o QueueOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.
func (o QueueOutput) ContentBasedDeduplication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.ContentBasedDeduplication }).(pulumi.BoolPtrOutput)
}

// Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.
func (o QueueOutput) DeduplicationScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.DeduplicationScope }).(pulumi.StringPtrOutput)
}

// The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.
func (o QueueOutput) DelaySeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.DelaySeconds }).(pulumi.IntPtrOutput)
}

// If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.
func (o QueueOutput) FifoQueue() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.FifoQueue }).(pulumi.BoolPtrOutput)
}

// Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.
func (o QueueOutput) FifoThroughputLimit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.FifoThroughputLimit }).(pulumi.StringPtrOutput)
}

// The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).
func (o QueueOutput) KmsDataKeyReusePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.KmsDataKeyReusePeriodSeconds }).(pulumi.IntPtrOutput)
}

// The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.
func (o QueueOutput) KmsMasterKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.KmsMasterKeyId }).(pulumi.StringPtrOutput)
}

// The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).
func (o QueueOutput) MaximumMessageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.MaximumMessageSize }).(pulumi.IntPtrOutput)
}

// The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).
func (o QueueOutput) MessageRetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.MessageRetentionPeriod }).(pulumi.IntPtrOutput)
}

// A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.
func (o QueueOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringPtrOutput { return v.QueueName }).(pulumi.StringPtrOutput)
}

// URL of the source queue.
func (o QueueOutput) QueueUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Queue) pulumi.StringOutput { return v.QueueUrl }).(pulumi.StringOutput)
}

// Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.
func (o QueueOutput) ReceiveMessageWaitTimeSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.ReceiveMessageWaitTimeSeconds }).(pulumi.IntPtrOutput)
}

// The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.
func (o QueueOutput) RedriveAllowPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Queue) pulumi.AnyOutput { return v.RedriveAllowPolicy }).(pulumi.AnyOutput)
}

// A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.
func (o QueueOutput) RedrivePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Queue) pulumi.AnyOutput { return v.RedrivePolicy }).(pulumi.AnyOutput)
}

// Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).
func (o QueueOutput) SqsManagedSseEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.BoolPtrOutput { return v.SqsManagedSseEnabled }).(pulumi.BoolPtrOutput)
}

// The tags that you attach to this queue.
func (o QueueOutput) Tags() QueueTagArrayOutput {
	return o.ApplyT(func(v *Queue) QueueTagArrayOutput { return v.Tags }).(QueueTagArrayOutput)
}

// The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.
func (o QueueOutput) VisibilityTimeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Queue) pulumi.IntPtrOutput { return v.VisibilityTimeout }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueueInput)(nil)).Elem(), &Queue{})
	pulumi.RegisterOutputType(QueueOutput{})
}
