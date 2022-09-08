// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for Metric Stream
 */
export class MetricStream extends pulumi.CustomResource {
    /**
     * Get an existing MetricStream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MetricStream {
        return new MetricStream(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudwatch:MetricStream';

    /**
     * Returns true if the given object is an instance of MetricStream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MetricStream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MetricStream.__pulumiType;
    }

    /**
     * Amazon Resource Name of the metric stream.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The date of creation of the metric stream.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
     */
    public readonly excludeFilters!: pulumi.Output<outputs.cloudwatch.MetricStreamFilter[] | undefined>;
    /**
     * The ARN of the Kinesis Firehose where to stream the data.
     */
    public readonly firehoseArn!: pulumi.Output<string>;
    /**
     * Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
     */
    public readonly includeFilters!: pulumi.Output<outputs.cloudwatch.MetricStreamFilter[] | undefined>;
    /**
     * The date of the last update of the metric stream.
     */
    public /*out*/ readonly lastUpdateDate!: pulumi.Output<string>;
    /**
     * Name of the metric stream.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The output format of the data streamed to the Kinesis Firehose.
     */
    public readonly outputFormat!: pulumi.Output<string>;
    /**
     * The ARN of the role that provides access to the Kinesis Firehose.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * Displays the state of the Metric Stream.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.
     */
    public readonly statisticsConfigurations!: pulumi.Output<outputs.cloudwatch.MetricStreamStatisticsConfiguration[] | undefined>;
    /**
     * A set of tags to assign to the delivery stream.
     */
    public readonly tags!: pulumi.Output<outputs.cloudwatch.MetricStreamTag[] | undefined>;

    /**
     * Create a MetricStream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetricStreamArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.firehoseArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firehoseArn'");
            }
            if ((!args || args.outputFormat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'outputFormat'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            resourceInputs["excludeFilters"] = args ? args.excludeFilters : undefined;
            resourceInputs["firehoseArn"] = args ? args.firehoseArn : undefined;
            resourceInputs["includeFilters"] = args ? args.includeFilters : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["outputFormat"] = args ? args.outputFormat : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["statisticsConfigurations"] = args ? args.statisticsConfigurations : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["lastUpdateDate"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["excludeFilters"] = undefined /*out*/;
            resourceInputs["firehoseArn"] = undefined /*out*/;
            resourceInputs["includeFilters"] = undefined /*out*/;
            resourceInputs["lastUpdateDate"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["outputFormat"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statisticsConfigurations"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MetricStream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MetricStream resource.
 */
export interface MetricStreamArgs {
    /**
     * Define which metrics will be not streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
     */
    excludeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamFilterArgs>[]>;
    /**
     * The ARN of the Kinesis Firehose where to stream the data.
     */
    firehoseArn: pulumi.Input<string>;
    /**
     * Define which metrics will be streamed. Metrics matched by multiple instances of MetricStreamFilter are joined with an OR operation by default. If both IncludeFilters and ExcludeFilters are omitted, all metrics in the account will be streamed. IncludeFilters and ExcludeFilters are mutually exclusive. Default to null.
     */
    includeFilters?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamFilterArgs>[]>;
    /**
     * Name of the metric stream.
     */
    name?: pulumi.Input<string>;
    /**
     * The output format of the data streamed to the Kinesis Firehose.
     */
    outputFormat: pulumi.Input<string>;
    /**
     * The ARN of the role that provides access to the Kinesis Firehose.
     */
    roleArn: pulumi.Input<string>;
    /**
     * By default, a metric stream always sends the MAX, MIN, SUM, and SAMPLECOUNT statistics for each metric that is streamed. You can use this parameter to have the metric stream also send additional statistics in the stream. This array can have up to 100 members.
     */
    statisticsConfigurations?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamStatisticsConfigurationArgs>[]>;
    /**
     * A set of tags to assign to the delivery stream.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.cloudwatch.MetricStreamTagArgs>[]>;
}
