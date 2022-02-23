// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EMR::InstanceGroupConfig
 */
export function getInstanceGroupConfig(args: GetInstanceGroupConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceGroupConfigResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:emr:getInstanceGroupConfig", {
        "id": args.id,
    }, opts);
}

export interface GetInstanceGroupConfigArgs {
    id: string;
}

export interface GetInstanceGroupConfigResult {
    readonly autoScalingPolicy?: outputs.emr.InstanceGroupConfigAutoScalingPolicy;
    readonly id?: string;
    readonly instanceCount?: number;
}

export function getInstanceGroupConfigOutput(args: GetInstanceGroupConfigOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceGroupConfigResult> {
    return pulumi.output(args).apply(a => getInstanceGroupConfig(a, opts))
}

export interface GetInstanceGroupConfigOutputArgs {
    id: pulumi.Input<string>;
}