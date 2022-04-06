// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationFSxOpenZFS.
 */
export function getLocationFSxOpenZFS(args: GetLocationFSxOpenZFSArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationFSxOpenZFSResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datasync:getLocationFSxOpenZFS", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationFSxOpenZFSArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
     */
    locationArn: string;
}

export interface GetLocationFSxOpenZFSResult {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
     */
    readonly locationArn?: string;
    /**
     * The URL of the FSx OpenZFS that was described.
     */
    readonly locationUri?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationFSxOpenZFSTag[];
}

export function getLocationFSxOpenZFSOutput(args: GetLocationFSxOpenZFSOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationFSxOpenZFSResult> {
    return pulumi.output(args).apply(a => getLocationFSxOpenZFS(a, opts))
}

export interface GetLocationFSxOpenZFSOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx OpenZFS file system location that is created.
     */
    locationArn: pulumi.Input<string>;
}