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
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sns:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * The archive policy determines the number of days Amazon SNS retains messages. You can set a retention period from 1 to 365 days.
     */
    public readonly archivePolicy!: pulumi.Output<any | undefined>;
    /**
     * Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
     *
     * When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
     *
     * (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
     */
    public readonly contentBasedDeduplication!: pulumi.Output<boolean | undefined>;
    /**
     * The body of the policy document you want to use for this topic.
     *
     * You can only add one policy per topic.
     *
     * The policy must be in JSON string format.
     *
     * Length Constraints: Maximum length of 30720
     */
    public readonly dataProtectionPolicy!: pulumi.Output<any | undefined>;
    /**
     * Delivery status logging configuration for supported protocols for an Amazon SNS topic.
     */
    public readonly deliveryStatusLogging!: pulumi.Output<outputs.sns.TopicLoggingConfig[] | undefined>;
    /**
     * The display name to use for an Amazon SNS topic with SMS subscriptions.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Set to true to create a FIFO topic.
     */
    public readonly fifoTopic!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
     *
     * This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
     */
    public readonly kmsMasterKeyId!: pulumi.Output<string | undefined>;
    /**
     * Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
     */
    public readonly signatureVersion!: pulumi.Output<string | undefined>;
    /**
     * The SNS subscriptions (endpoints) for this topic.
     */
    public readonly subscription!: pulumi.Output<outputs.sns.TopicSubscription[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.sns.TopicTag[] | undefined>;
    public /*out*/ readonly topicArn!: pulumi.Output<string>;
    /**
     * The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with .fifo.
     *
     * If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see Name Type.
     */
    public readonly topicName!: pulumi.Output<string | undefined>;
    /**
     * Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
     */
    public readonly tracingConfig!: pulumi.Output<string | undefined>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TopicArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["archivePolicy"] = args ? args.archivePolicy : undefined;
            resourceInputs["contentBasedDeduplication"] = args ? args.contentBasedDeduplication : undefined;
            resourceInputs["dataProtectionPolicy"] = args ? args.dataProtectionPolicy : undefined;
            resourceInputs["deliveryStatusLogging"] = args ? args.deliveryStatusLogging : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["fifoTopic"] = args ? args.fifoTopic : undefined;
            resourceInputs["kmsMasterKeyId"] = args ? args.kmsMasterKeyId : undefined;
            resourceInputs["signatureVersion"] = args ? args.signatureVersion : undefined;
            resourceInputs["subscription"] = args ? args.subscription : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["topicName"] = args ? args.topicName : undefined;
            resourceInputs["tracingConfig"] = args ? args.tracingConfig : undefined;
            resourceInputs["topicArn"] = undefined /*out*/;
        } else {
            resourceInputs["archivePolicy"] = undefined /*out*/;
            resourceInputs["contentBasedDeduplication"] = undefined /*out*/;
            resourceInputs["dataProtectionPolicy"] = undefined /*out*/;
            resourceInputs["deliveryStatusLogging"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["fifoTopic"] = undefined /*out*/;
            resourceInputs["kmsMasterKeyId"] = undefined /*out*/;
            resourceInputs["signatureVersion"] = undefined /*out*/;
            resourceInputs["subscription"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["topicArn"] = undefined /*out*/;
            resourceInputs["topicName"] = undefined /*out*/;
            resourceInputs["tracingConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["fifoTopic", "topicName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Topic.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * The archive policy determines the number of days Amazon SNS retains messages. You can set a retention period from 1 to 365 days.
     */
    archivePolicy?: any;
    /**
     * Enables content-based deduplication for FIFO topics. By default, ContentBasedDeduplication is set to false. If you create a FIFO topic and this attribute is false, you must specify a value for the MessageDeduplicationId parameter for the Publish action.
     *
     * When you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to generate the MessageDeduplicationId using the body of the message (but not the attributes of the message).
     *
     * (Optional) To override the generated value, you can specify a value for the the MessageDeduplicationId parameter for the Publish action.
     */
    contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * The body of the policy document you want to use for this topic.
     *
     * You can only add one policy per topic.
     *
     * The policy must be in JSON string format.
     *
     * Length Constraints: Maximum length of 30720
     */
    dataProtectionPolicy?: any;
    /**
     * Delivery status logging configuration for supported protocols for an Amazon SNS topic.
     */
    deliveryStatusLogging?: pulumi.Input<pulumi.Input<inputs.sns.TopicLoggingConfigArgs>[]>;
    /**
     * The display name to use for an Amazon SNS topic with SMS subscriptions.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Set to true to create a FIFO topic.
     */
    fifoTopic?: pulumi.Input<boolean>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see Key Terms. For more examples, see KeyId in the AWS Key Management Service API Reference.
     *
     * This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html).
     */
    kmsMasterKeyId?: pulumi.Input<string>;
    /**
     * Version of the Amazon SNS signature used. If the SignatureVersion is 1, Signature is a Base64-encoded SHA1withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values. If the SignatureVersion is 2, Signature is a Base64-encoded SHA256withRSA signature of the Message, MessageId, Type, Timestamp, and TopicArn values.
     */
    signatureVersion?: pulumi.Input<string>;
    /**
     * The SNS subscriptions (endpoints) for this topic.
     */
    subscription?: pulumi.Input<pulumi.Input<inputs.sns.TopicSubscriptionArgs>[]>;
    tags?: pulumi.Input<pulumi.Input<inputs.sns.TopicTagArgs>[]>;
    /**
     * The name of the topic you want to create. Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with .fifo.
     *
     * If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see Name Type.
     */
    topicName?: pulumi.Input<string>;
    /**
     * Tracing mode of an Amazon SNS topic. By default TracingConfig is set to PassThrough, and the topic passes through the tracing header it receives from an SNS publisher to its subscriptions. If set to Active, SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true. Only supported on standard topics.
     */
    tracingConfig?: pulumi.Input<string>;
}
