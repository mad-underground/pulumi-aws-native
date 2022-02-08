// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EKS::Nodegroup
 */
export function getNodegroup(args: GetNodegroupArgs, opts?: pulumi.InvokeOptions): Promise<GetNodegroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:eks:getNodegroup", {
        "id": args.id,
    }, opts);
}

export interface GetNodegroupArgs {
    id: string;
}

export interface GetNodegroupResult {
    readonly arn?: string;
    readonly forceUpdateEnabled?: boolean;
    readonly id?: string;
    readonly labels?: any;
    readonly launchTemplate?: outputs.eks.NodegroupLaunchTemplateSpecification;
    readonly releaseVersion?: string;
    readonly scalingConfig?: outputs.eks.NodegroupScalingConfig;
    readonly tags?: any;
    readonly taints?: outputs.eks.NodegroupTaint[];
    readonly updateConfig?: outputs.eks.NodegroupUpdateConfig;
    readonly version?: string;
}

export function getNodegroupOutput(args: GetNodegroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNodegroupResult> {
    return pulumi.output(args).apply(a => getNodegroup(a, opts))
}

export interface GetNodegroupOutputArgs {
    id: pulumi.Input<string>;
}
