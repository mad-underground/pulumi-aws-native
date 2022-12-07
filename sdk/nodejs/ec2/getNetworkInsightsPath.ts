// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::EC2::NetworkInsightsPath
 */
export function getNetworkInsightsPath(args: GetNetworkInsightsPathArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInsightsPathResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getNetworkInsightsPath", {
        "networkInsightsPathId": args.networkInsightsPathId,
    }, opts);
}

export interface GetNetworkInsightsPathArgs {
    networkInsightsPathId: string;
}

export interface GetNetworkInsightsPathResult {
    readonly createdDate?: string;
    readonly destinationArn?: string;
    readonly networkInsightsPathArn?: string;
    readonly networkInsightsPathId?: string;
    readonly sourceArn?: string;
    readonly tags?: outputs.ec2.NetworkInsightsPathTag[];
}

export function getNetworkInsightsPathOutput(args: GetNetworkInsightsPathOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkInsightsPathResult> {
    return pulumi.output(args).apply(a => getNetworkInsightsPath(a, opts))
}

export interface GetNetworkInsightsPathOutputArgs {
    networkInsightsPathId: pulumi.Input<string>;
}
