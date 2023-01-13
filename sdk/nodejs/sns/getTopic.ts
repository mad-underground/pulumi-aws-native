// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SNS::Topic
 */
export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sns:getTopic", {
        "topicArn": args.topicArn,
    }, opts);
}

export interface GetTopicArgs {
    topicArn: string;
}

export interface GetTopicResult {
    /**
     * Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
     *
     * When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
     *
     * (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
     */
    readonly contentBasedDeduplication?: boolean;
    /**
     * The body of the policy document you want to use for this topic.
     *
     * You can only add one policy per topic.
     *
     * The policy must be in JSON string format.
     *
     * Length Constraints: Maximum length of 30720
     */
    readonly dataProtectionPolicy?: any;
    /**
     * The display name to use for an Amazon SNS topic with SMS subscriptions.
     */
    readonly displayName?: string;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
     *
     * This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
     */
    readonly kmsMasterKeyId?: string;
    /**
     * Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
     */
    readonly signatureVersion?: string;
    /**
     * The SNS subscriptions (endpoints) for this topic.
     */
    readonly subscription?: outputs.sns.TopicSubscription[];
    readonly tags?: outputs.sns.TopicTag[];
    readonly topicArn?: string;
}
/**
 * Resource Type definition for AWS::SNS::Topic
 */
export function getTopicOutput(args: GetTopicOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTopicResult> {
    return pulumi.output(args).apply((a: any) => getTopic(a, opts))
}

export interface GetTopicOutputArgs {
    topicArn: pulumi.Input<string>;
}
