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
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloudwatch:getAlarm", {
        "id": args.id,
    }, opts);
}

export interface GetAlarmArgs {
    id: string;
}

export interface GetAlarmResult {
    readonly actionsEnabled?: boolean;
    readonly alarmActions?: string[];
    readonly alarmDescription?: string;
    readonly arn?: string;
    readonly comparisonOperator?: string;
    readonly datapointsToAlarm?: number;
    readonly dimensions?: outputs.cloudwatch.AlarmDimension[];
    readonly evaluateLowSampleCountPercentile?: string;
    readonly evaluationPeriods?: number;
    readonly extendedStatistic?: string;
    readonly id?: string;
    readonly insufficientDataActions?: string[];
    readonly metricName?: string;
    readonly metrics?: outputs.cloudwatch.AlarmMetricDataQuery[];
    readonly namespace?: string;
    readonly oKActions?: string[];
    readonly period?: number;
    readonly statistic?: string;
    readonly threshold?: number;
    readonly thresholdMetricId?: string;
    readonly treatMissingData?: string;
    readonly unit?: string;
}

export function getAlarmOutput(args: GetAlarmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAlarmResult> {
    return pulumi.output(args).apply(a => getAlarm(a, opts))
}

export interface GetAlarmOutputArgs {
    id: pulumi.Input<string>;
}
