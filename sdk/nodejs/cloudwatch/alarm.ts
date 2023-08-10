// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudWatch::Alarm
 */
export class Alarm extends pulumi.CustomResource {
    /**
     * Get an existing Alarm resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Alarm {
        return new Alarm(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:cloudwatch:Alarm';

    /**
     * Returns true if the given object is an instance of Alarm.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Alarm {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Alarm.__pulumiType;
    }

    /**
     * Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
     */
    public readonly actionsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The list of actions to execute when this alarm transitions into an ALARM state from any other state.
     */
    public readonly alarmActions!: pulumi.Output<string[] | undefined>;
    /**
     * The description of the alarm.
     */
    public readonly alarmDescription!: pulumi.Output<string | undefined>;
    /**
     * The name of the alarm.
     */
    public readonly alarmName!: pulumi.Output<string | undefined>;
    /**
     * Amazon Resource Name is a unique name for each resource.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The arithmetic operation to use when comparing the specified statistic and threshold.
     */
    public readonly comparisonOperator!: pulumi.Output<string>;
    /**
     * The number of datapoints that must be breaching to trigger the alarm.
     */
    public readonly datapointsToAlarm!: pulumi.Output<number | undefined>;
    /**
     * The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify Dimensions. Instead, you use Metrics.
     */
    public readonly dimensions!: pulumi.Output<outputs.cloudwatch.AlarmDimension[] | undefined>;
    /**
     * Used only for alarms based on percentiles.
     */
    public readonly evaluateLowSampleCountPercentile!: pulumi.Output<string | undefined>;
    /**
     * The number of periods over which data is compared to the specified threshold.
     */
    public readonly evaluationPeriods!: pulumi.Output<number>;
    /**
     * The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
     */
    public readonly extendedStatistic!: pulumi.Output<string | undefined>;
    /**
     * The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
     */
    public readonly insufficientDataActions!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the metric associated with the alarm.
     */
    public readonly metricName!: pulumi.Output<string | undefined>;
    /**
     * An array that enables you to create an alarm based on the result of a metric math expression.
     */
    public readonly metrics!: pulumi.Output<outputs.cloudwatch.AlarmMetricDataQuery[] | undefined>;
    /**
     * The namespace of the metric associated with the alarm.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The actions to execute when this alarm transitions to the OK state from any other state.
     */
    public readonly okActions!: pulumi.Output<string[] | undefined>;
    /**
     * The period in seconds, over which the statistic is applied.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The statistic for the metric associated with the alarm, other than percentile.
     */
    public readonly statistic!: pulumi.Output<string | undefined>;
    /**
     * In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
     */
    public readonly threshold!: pulumi.Output<number | undefined>;
    /**
     * In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
     */
    public readonly thresholdMetricId!: pulumi.Output<string | undefined>;
    /**
     * Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.
     */
    public readonly treatMissingData!: pulumi.Output<string | undefined>;
    /**
     * The unit of the metric associated with the alarm.
     */
    public readonly unit!: pulumi.Output<string | undefined>;

    /**
     * Create a Alarm resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlarmArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.comparisonOperator === undefined) && !opts.urn) {
                throw new Error("Missing required property 'comparisonOperator'");
            }
            if ((!args || args.evaluationPeriods === undefined) && !opts.urn) {
                throw new Error("Missing required property 'evaluationPeriods'");
            }
            resourceInputs["actionsEnabled"] = args ? args.actionsEnabled : undefined;
            resourceInputs["alarmActions"] = args ? args.alarmActions : undefined;
            resourceInputs["alarmDescription"] = args ? args.alarmDescription : undefined;
            resourceInputs["alarmName"] = args ? args.alarmName : undefined;
            resourceInputs["comparisonOperator"] = args ? args.comparisonOperator : undefined;
            resourceInputs["datapointsToAlarm"] = args ? args.datapointsToAlarm : undefined;
            resourceInputs["dimensions"] = args ? args.dimensions : undefined;
            resourceInputs["evaluateLowSampleCountPercentile"] = args ? args.evaluateLowSampleCountPercentile : undefined;
            resourceInputs["evaluationPeriods"] = args ? args.evaluationPeriods : undefined;
            resourceInputs["extendedStatistic"] = args ? args.extendedStatistic : undefined;
            resourceInputs["insufficientDataActions"] = args ? args.insufficientDataActions : undefined;
            resourceInputs["metricName"] = args ? args.metricName : undefined;
            resourceInputs["metrics"] = args ? args.metrics : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["okActions"] = args ? args.okActions : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["statistic"] = args ? args.statistic : undefined;
            resourceInputs["threshold"] = args ? args.threshold : undefined;
            resourceInputs["thresholdMetricId"] = args ? args.thresholdMetricId : undefined;
            resourceInputs["treatMissingData"] = args ? args.treatMissingData : undefined;
            resourceInputs["unit"] = args ? args.unit : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["actionsEnabled"] = undefined /*out*/;
            resourceInputs["alarmActions"] = undefined /*out*/;
            resourceInputs["alarmDescription"] = undefined /*out*/;
            resourceInputs["alarmName"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["comparisonOperator"] = undefined /*out*/;
            resourceInputs["datapointsToAlarm"] = undefined /*out*/;
            resourceInputs["dimensions"] = undefined /*out*/;
            resourceInputs["evaluateLowSampleCountPercentile"] = undefined /*out*/;
            resourceInputs["evaluationPeriods"] = undefined /*out*/;
            resourceInputs["extendedStatistic"] = undefined /*out*/;
            resourceInputs["insufficientDataActions"] = undefined /*out*/;
            resourceInputs["metricName"] = undefined /*out*/;
            resourceInputs["metrics"] = undefined /*out*/;
            resourceInputs["namespace"] = undefined /*out*/;
            resourceInputs["okActions"] = undefined /*out*/;
            resourceInputs["period"] = undefined /*out*/;
            resourceInputs["statistic"] = undefined /*out*/;
            resourceInputs["threshold"] = undefined /*out*/;
            resourceInputs["thresholdMetricId"] = undefined /*out*/;
            resourceInputs["treatMissingData"] = undefined /*out*/;
            resourceInputs["unit"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Alarm.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Alarm resource.
 */
export interface AlarmArgs {
    /**
     * Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
     */
    actionsEnabled?: pulumi.Input<boolean>;
    /**
     * The list of actions to execute when this alarm transitions into an ALARM state from any other state.
     */
    alarmActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The description of the alarm.
     */
    alarmDescription?: pulumi.Input<string>;
    /**
     * The name of the alarm.
     */
    alarmName?: pulumi.Input<string>;
    /**
     * The arithmetic operation to use when comparing the specified statistic and threshold.
     */
    comparisonOperator: pulumi.Input<string>;
    /**
     * The number of datapoints that must be breaching to trigger the alarm.
     */
    datapointsToAlarm?: pulumi.Input<number>;
    /**
     * The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify Dimensions. Instead, you use Metrics.
     */
    dimensions?: pulumi.Input<pulumi.Input<inputs.cloudwatch.AlarmDimensionArgs>[]>;
    /**
     * Used only for alarms based on percentiles.
     */
    evaluateLowSampleCountPercentile?: pulumi.Input<string>;
    /**
     * The number of periods over which data is compared to the specified threshold.
     */
    evaluationPeriods: pulumi.Input<number>;
    /**
     * The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
     */
    extendedStatistic?: pulumi.Input<string>;
    /**
     * The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
     */
    insufficientDataActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the metric associated with the alarm.
     */
    metricName?: pulumi.Input<string>;
    /**
     * An array that enables you to create an alarm based on the result of a metric math expression.
     */
    metrics?: pulumi.Input<pulumi.Input<inputs.cloudwatch.AlarmMetricDataQueryArgs>[]>;
    /**
     * The namespace of the metric associated with the alarm.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The actions to execute when this alarm transitions to the OK state from any other state.
     */
    okActions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The period in seconds, over which the statistic is applied.
     */
    period?: pulumi.Input<number>;
    /**
     * The statistic for the metric associated with the alarm, other than percentile.
     */
    statistic?: pulumi.Input<string>;
    /**
     * In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
     */
    threshold?: pulumi.Input<number>;
    /**
     * In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
     */
    thresholdMetricId?: pulumi.Input<string>;
    /**
     * Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.
     */
    treatMissingData?: pulumi.Input<string>;
    /**
     * The unit of the metric associated with the alarm.
     */
    unit?: pulumi.Input<string>;
}
