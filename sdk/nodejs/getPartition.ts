// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getPartition(opts?: pulumi.InvokeOptions): Promise<GetPartitionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:index:getPartition", {
    }, opts);
}

export interface GetPartitionResult {
    /**
     * Base DNS domain name for the current partition (e.g., `amazonaws.com` in AWS Commercial, `amazonaws.com.cn` in AWS China).
     */
    readonly dnsSuffix: string;
    /**
     * Identifier of the current partition (e.g., `aws` in AWS Commercial, `aws-cn` in AWS China).
     */
    readonly partition: string;
}
export function getPartitionOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetPartitionResult> {
    return pulumi.output(getPartition(opts))
}
