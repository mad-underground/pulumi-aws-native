// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::Volume
 */
export function getVolume(args: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getVolume", {
        "id": args.id,
    }, opts);
}

export interface GetVolumeArgs {
    id: string;
}

export interface GetVolumeResult {
    readonly autoEnableIO?: boolean;
    readonly availabilityZone?: string;
    readonly encrypted?: boolean;
    readonly id?: string;
    readonly iops?: number;
    readonly kmsKeyId?: string;
    readonly multiAttachEnabled?: boolean;
    readonly outpostArn?: string;
    readonly size?: number;
    readonly snapshotId?: string;
    readonly tags?: outputs.ec2.VolumeTag[];
    readonly throughput?: number;
    readonly volumeType?: string;
}

export function getVolumeOutput(args: GetVolumeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeResult> {
    return pulumi.output(args).apply(a => getVolume(a, opts))
}

export interface GetVolumeOutputArgs {
    id: pulumi.Input<string>;
}
