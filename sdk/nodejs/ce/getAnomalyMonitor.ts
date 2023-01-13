// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.
 */
export function getAnomalyMonitor(args: GetAnomalyMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetAnomalyMonitorResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ce:getAnomalyMonitor", {
        "monitorArn": args.monitorArn,
    }, opts);
}

export interface GetAnomalyMonitorArgs {
    monitorArn: string;
}

export interface GetAnomalyMonitorResult {
    /**
     * The date when the monitor was created. 
     */
    readonly creationDate?: string;
    /**
     * The value for evaluated dimensions.
     */
    readonly dimensionalValueCount?: number;
    /**
     * The date when the monitor last evaluated for anomalies.
     */
    readonly lastEvaluatedDate?: string;
    /**
     * The date when the monitor was last updated.
     */
    readonly lastUpdatedDate?: string;
    readonly monitorArn?: string;
    /**
     * The name of the monitor.
     */
    readonly monitorName?: string;
}
/**
 * AWS Cost Anomaly Detection leverages advanced Machine Learning technologies to identify anomalous spend and root causes, so you can quickly take action. You can use Cost Anomaly Detection by creating monitor.
 */
export function getAnomalyMonitorOutput(args: GetAnomalyMonitorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAnomalyMonitorResult> {
    return pulumi.output(args).apply((a: any) => getAnomalyMonitor(a, opts))
}

export interface GetAnomalyMonitorOutputArgs {
    monitorArn: pulumi.Input<string>;
}
