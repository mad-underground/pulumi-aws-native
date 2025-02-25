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
export function getAlarm(args: GetAlarmArgs, opts?: pulumi.InvokeOptions): Promise<GetAlarmResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudwatch:getAlarm", {
        "alarmName": args.alarmName,
    }, opts);
}

export interface GetAlarmArgs {
    /**
     * The name of the alarm.
     */
    alarmName: string;
}

export interface GetAlarmResult {
    /**
     * Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
     */
    readonly actionsEnabled?: boolean;
    /**
     * The list of actions to execute when this alarm transitions into an ALARM state from any other state.
     */
    readonly alarmActions?: string[];
    /**
     * The description of the alarm.
     */
    readonly alarmDescription?: string;
    /**
     * Amazon Resource Name is a unique name for each resource.
     */
    readonly arn?: string;
    /**
     * The arithmetic operation to use when comparing the specified statistic and threshold.
     */
    readonly comparisonOperator?: string;
    /**
     * The number of datapoints that must be breaching to trigger the alarm.
     */
    readonly datapointsToAlarm?: number;
    /**
     * The dimensions for the metric associated with the alarm. For an alarm based on a math expression, you can't specify Dimensions. Instead, you use Metrics.
     */
    readonly dimensions?: outputs.cloudwatch.AlarmDimension[];
    /**
     * Used only for alarms based on percentiles.
     */
    readonly evaluateLowSampleCountPercentile?: string;
    /**
     * The number of periods over which data is compared to the specified threshold.
     */
    readonly evaluationPeriods?: number;
    /**
     * The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
     */
    readonly extendedStatistic?: string;
    /**
     * The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.
     */
    readonly insufficientDataActions?: string[];
    /**
     * The name of the metric associated with the alarm.
     */
    readonly metricName?: string;
    /**
     * An array that enables you to create an alarm based on the result of a metric math expression.
     */
    readonly metrics?: outputs.cloudwatch.AlarmMetricDataQuery[];
    /**
     * The namespace of the metric associated with the alarm.
     */
    readonly namespace?: string;
    /**
     * The actions to execute when this alarm transitions to the OK state from any other state.
     */
    readonly okActions?: string[];
    /**
     * The period in seconds, over which the statistic is applied.
     */
    readonly period?: number;
    /**
     * The statistic for the metric associated with the alarm, other than percentile.
     */
    readonly statistic?: string;
    /**
     * In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
     */
    readonly threshold?: number;
    /**
     * In an alarm based on an anomaly detection model, this is the ID of the ANOMALY_DETECTION_BAND function used as the threshold for the alarm.
     */
    readonly thresholdMetricId?: string;
    /**
     * Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.
     */
    readonly treatMissingData?: string;
    /**
     * The unit of the metric associated with the alarm.
     */
    readonly unit?: string;
}
/**
 * Resource Type definition for AWS::CloudWatch::Alarm
 */
export function getAlarmOutput(args: GetAlarmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAlarmResult> {
    return pulumi.output(args).apply((a: any) => getAlarm(a, opts))
}

export interface GetAlarmOutputArgs {
    /**
     * The name of the alarm.
     */
    alarmName: pulumi.Input<string>;
}
