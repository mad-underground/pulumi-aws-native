// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
 */
export function getDistribution(args: GetDistributionArgs, opts?: pulumi.InvokeOptions): Promise<GetDistributionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cloudfront:getDistribution", {
        "id": args.id,
    }, opts);
}

export interface GetDistributionArgs {
    id: string;
}

export interface GetDistributionResult {
    /**
     * The distribution's configuration.
     */
    readonly distributionConfig?: outputs.cloudfront.DistributionConfig;
    readonly domainName?: string;
    readonly id?: string;
    /**
     * A complex type that contains zero or more ``Tag`` elements.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.
 */
export function getDistributionOutput(args: GetDistributionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDistributionResult> {
    return pulumi.output(args).apply((a: any) => getDistribution(a, opts))
}

export interface GetDistributionOutputArgs {
    id: pulumi.Input<string>;
}
