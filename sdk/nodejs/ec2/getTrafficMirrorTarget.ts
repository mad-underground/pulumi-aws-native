// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TrafficMirrorTarget
 */
export function getTrafficMirrorTarget(args: GetTrafficMirrorTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetTrafficMirrorTargetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getTrafficMirrorTarget", {
        "id": args.id,
    }, opts);
}

export interface GetTrafficMirrorTargetArgs {
    id: string;
}

export interface GetTrafficMirrorTargetResult {
    readonly id?: string;
    readonly tags?: outputs.ec2.TrafficMirrorTargetTag[];
}

export function getTrafficMirrorTargetOutput(args: GetTrafficMirrorTargetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTrafficMirrorTargetResult> {
    return pulumi.output(args).apply(a => getTrafficMirrorTarget(a, opts))
}

export interface GetTrafficMirrorTargetOutputArgs {
    id: pulumi.Input<string>;
}
