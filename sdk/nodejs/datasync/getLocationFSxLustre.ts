// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationFSxLustre.
 */
export function getLocationFSxLustre(args: GetLocationFSxLustreArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationFSxLustreResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datasync:getLocationFSxLustre", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationFSxLustreArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
     */
    locationArn: string;
}

export interface GetLocationFSxLustreResult {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
     */
    readonly locationArn?: string;
    /**
     * The URL of the FSx for Lustre location that was described.
     */
    readonly locationUri?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationFSxLustreTag[];
}

export function getLocationFSxLustreOutput(args: GetLocationFSxLustreOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationFSxLustreResult> {
    return pulumi.output(args).apply(a => getLocationFSxLustre(a, opts))
}

export interface GetLocationFSxLustreOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon FSx for Lustre file system location that is created.
     */
    locationArn: pulumi.Input<string>;
}
