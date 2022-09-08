// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationFSxWindows.
 */
export function getLocationFSxWindows(args: GetLocationFSxWindowsArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationFSxWindowsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datasync:getLocationFSxWindows", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationFSxWindowsArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.
     */
    locationArn: string;
}

export interface GetLocationFSxWindowsResult {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.
     */
    readonly locationArn?: string;
    /**
     * The URL of the FSx for Windows location that was described.
     */
    readonly locationUri?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationFSxWindowsTag[];
}

export function getLocationFSxWindowsOutput(args: GetLocationFSxWindowsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationFSxWindowsResult> {
    return pulumi.output(args).apply(a => getLocationFSxWindows(a, opts))
}

export interface GetLocationFSxWindowsOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx for Windows file system location that is created.
     */
    locationArn: pulumi.Input<string>;
}
