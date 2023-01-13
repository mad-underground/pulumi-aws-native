// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Specifies a metric filter that describes how CloudWatch Logs extracts information from logs and transforms it into Amazon CloudWatch metrics.
 */
export function getMetricFilter(args: GetMetricFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetMetricFilterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:logs:getMetricFilter", {
        "filterName": args.filterName,
        "logGroupName": args.logGroupName,
    }, opts);
}

export interface GetMetricFilterArgs {
    /**
     * A name for the metric filter.
     */
    filterName: string;
    /**
     * Existing log group that you want to associate with this filter.
     */
    logGroupName: string;
}

export interface GetMetricFilterResult {
    /**
     * Pattern that Logs follows to interpret each entry in a log.
     */
    readonly filterPattern?: string;
    /**
     * A collection of information that defines how metric data gets emitted.
     */
    readonly metricTransformations?: outputs.logs.MetricFilterMetricTransformation[];
}
/**
 * Specifies a metric filter that describes how CloudWatch Logs extracts information from logs and transforms it into Amazon CloudWatch metrics.
 */
export function getMetricFilterOutput(args: GetMetricFilterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMetricFilterResult> {
    return pulumi.output(args).apply((a: any) => getMetricFilter(a, opts))
}

export interface GetMetricFilterOutputArgs {
    /**
     * A name for the metric filter.
     */
    filterName: pulumi.Input<string>;
    /**
     * Existing log group that you want to associate with this filter.
     */
    logGroupName: pulumi.Input<string>;
}
