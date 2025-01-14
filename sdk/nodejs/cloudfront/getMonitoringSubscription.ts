// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFront::MonitoringSubscription
 */
export function getMonitoringSubscription(args: GetMonitoringSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetMonitoringSubscriptionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudfront:getMonitoringSubscription", {
        "distributionId": args.distributionId,
    }, opts);
}

export interface GetMonitoringSubscriptionArgs {
    distributionId: string;
}

export interface GetMonitoringSubscriptionResult {
    readonly monitoringSubscription?: outputs.cloudfront.MonitoringSubscription;
}
/**
 * Resource Type definition for AWS::CloudFront::MonitoringSubscription
 */
export function getMonitoringSubscriptionOutput(args: GetMonitoringSubscriptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMonitoringSubscriptionResult> {
    return pulumi.output(args).apply((a: any) => getMonitoringSubscription(a, opts))
}

export interface GetMonitoringSubscriptionOutputArgs {
    distributionId: pulumi.Input<string>;
}
