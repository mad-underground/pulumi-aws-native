// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Specifies a VPC flow log, which enables you to capture IP traffic for a specific network interface, subnet, or VPC.
 */
export class FlowLog extends pulumi.CustomResource {
    /**
     * Get an existing FlowLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FlowLog {
        return new FlowLog(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:FlowLog';

    /**
     * Returns true if the given object is an instance of FlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowLog.__pulumiType;
    }

    /**
     * The Flow Log ID
     */
    public /*out*/ readonly awsId!: pulumi.Output<string>;
    /**
     * The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
     */
    public readonly deliverCrossAccountRole!: pulumi.Output<string | undefined>;
    /**
     * The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
     */
    public readonly deliverLogsPermissionArn!: pulumi.Output<string | undefined>;
    public readonly destinationOptions!: pulumi.Output<outputs.ec2.DestinationOptionsProperties | undefined>;
    /**
     * Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
     */
    public readonly logDestination!: pulumi.Output<string | undefined>;
    /**
     * Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
     */
    public readonly logDestinationType!: pulumi.Output<enums.ec2.FlowLogLogDestinationType | undefined>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    public readonly logFormat!: pulumi.Output<string | undefined>;
    /**
     * The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
     */
    public readonly logGroupName!: pulumi.Output<string | undefined>;
    /**
     * The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
     */
    public readonly maxAggregationInterval!: pulumi.Output<number | undefined>;
    /**
     * The ID of the subnet, network interface, or VPC for which you want to create a flow log.
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
     */
    public readonly resourceType!: pulumi.Output<enums.ec2.FlowLogResourceType>;
    /**
     * The tags to apply to the flow logs.
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
     */
    public readonly trafficType!: pulumi.Output<enums.ec2.FlowLogTrafficType | undefined>;

    /**
     * Create a FlowLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowLogArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceId'");
            }
            if ((!args || args.resourceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceType'");
            }
            resourceInputs["deliverCrossAccountRole"] = args ? args.deliverCrossAccountRole : undefined;
            resourceInputs["deliverLogsPermissionArn"] = args ? args.deliverLogsPermissionArn : undefined;
            resourceInputs["destinationOptions"] = args ? args.destinationOptions : undefined;
            resourceInputs["logDestination"] = args ? args.logDestination : undefined;
            resourceInputs["logDestinationType"] = args ? args.logDestinationType : undefined;
            resourceInputs["logFormat"] = args ? args.logFormat : undefined;
            resourceInputs["logGroupName"] = args ? args.logGroupName : undefined;
            resourceInputs["maxAggregationInterval"] = args ? args.maxAggregationInterval : undefined;
            resourceInputs["resourceId"] = args ? args.resourceId : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trafficType"] = args ? args.trafficType : undefined;
            resourceInputs["awsId"] = undefined /*out*/;
        } else {
            resourceInputs["awsId"] = undefined /*out*/;
            resourceInputs["deliverCrossAccountRole"] = undefined /*out*/;
            resourceInputs["deliverLogsPermissionArn"] = undefined /*out*/;
            resourceInputs["destinationOptions"] = undefined /*out*/;
            resourceInputs["logDestination"] = undefined /*out*/;
            resourceInputs["logDestinationType"] = undefined /*out*/;
            resourceInputs["logFormat"] = undefined /*out*/;
            resourceInputs["logGroupName"] = undefined /*out*/;
            resourceInputs["maxAggregationInterval"] = undefined /*out*/;
            resourceInputs["resourceId"] = undefined /*out*/;
            resourceInputs["resourceType"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["trafficType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["deliverCrossAccountRole", "deliverLogsPermissionArn", "destinationOptions", "logDestination", "logDestinationType", "logFormat", "logGroupName", "maxAggregationInterval", "resourceId", "resourceType", "trafficType"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(FlowLog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a FlowLog resource.
 */
export interface FlowLogArgs {
    /**
     * The ARN of the IAM role that allows Amazon EC2 to publish flow logs across accounts.
     */
    deliverCrossAccountRole?: pulumi.Input<string>;
    /**
     * The ARN for the IAM role that permits Amazon EC2 to publish flow logs to a CloudWatch Logs log group in your account. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
     */
    deliverLogsPermissionArn?: pulumi.Input<string>;
    destinationOptions?: pulumi.Input<inputs.ec2.DestinationOptionsPropertiesArgs>;
    /**
     * Specifies the destination to which the flow log data is to be published. Flow log data can be published to a CloudWatch Logs log group, an Amazon S3 bucket, or a Kinesis Firehose stream. The value specified for this parameter depends on the value specified for LogDestinationType.
     */
    logDestination?: pulumi.Input<string>;
    /**
     * Specifies the type of destination to which the flow log data is to be published. Flow log data can be published to CloudWatch Logs or Amazon S3.
     */
    logDestinationType?: pulumi.Input<enums.ec2.FlowLogLogDestinationType>;
    /**
     * The fields to include in the flow log record, in the order in which they should appear.
     */
    logFormat?: pulumi.Input<string>;
    /**
     * The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs. If you specify LogDestinationType as s3 or kinesis-data-firehose, do not specify DeliverLogsPermissionArn or LogGroupName.
     */
    logGroupName?: pulumi.Input<string>;
    /**
     * The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600 seconds (10 minutes).
     */
    maxAggregationInterval?: pulumi.Input<number>;
    /**
     * The ID of the subnet, network interface, or VPC for which you want to create a flow log.
     */
    resourceId: pulumi.Input<string>;
    /**
     * The type of resource for which to create the flow log. For example, if you specified a VPC ID for the ResourceId property, specify VPC for this property.
     */
    resourceType: pulumi.Input<enums.ec2.FlowLogResourceType>;
    /**
     * The tags to apply to the flow logs.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * The type of traffic to log. You can log traffic that the resource accepts or rejects, or all traffic.
     */
    trafficType?: pulumi.Input<enums.ec2.FlowLogTrafficType>;
}
