// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource schema represents the LogAnomalyDetectionIntegration resource in the Amazon DevOps Guru.
 */
export function getLogAnomalyDetectionIntegration(args: GetLogAnomalyDetectionIntegrationArgs, opts?: pulumi.InvokeOptions): Promise<GetLogAnomalyDetectionIntegrationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:devopsguru:getLogAnomalyDetectionIntegration", {
        "accountId": args.accountId,
    }, opts);
}

export interface GetLogAnomalyDetectionIntegrationArgs {
    accountId: string;
}

export interface GetLogAnomalyDetectionIntegrationResult {
    readonly accountId?: string;
}
/**
 * This resource schema represents the LogAnomalyDetectionIntegration resource in the Amazon DevOps Guru.
 */
export function getLogAnomalyDetectionIntegrationOutput(args: GetLogAnomalyDetectionIntegrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogAnomalyDetectionIntegrationResult> {
    return pulumi.output(args).apply((a: any) => getLogAnomalyDetectionIntegration(a, opts))
}

export interface GetLogAnomalyDetectionIntegrationOutputArgs {
    accountId: pulumi.Input<string>;
}
