// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationNFS
 */
export function getLocationNFS(args: GetLocationNFSArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationNFSResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:datasync:getLocationNFS", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationNFSArgs {
    /**
     * The Amazon Resource Name (ARN) of the NFS location.
     */
    locationArn: string;
}

export interface GetLocationNFSResult {
    /**
     * The Amazon Resource Name (ARN) of the NFS location.
     */
    readonly locationArn?: string;
    /**
     * The URL of the NFS location that was described.
     */
    readonly locationUri?: string;
    readonly mountOptions?: outputs.datasync.LocationNFSMountOptions;
    readonly onPremConfig?: outputs.datasync.LocationNFSOnPremConfig;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationNFSTag[];
}
/**
 * Resource schema for AWS::DataSync::LocationNFS
 */
export function getLocationNFSOutput(args: GetLocationNFSOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationNFSResult> {
    return pulumi.output(args).apply((a: any) => getLocationNFS(a, opts))
}

export interface GetLocationNFSOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the NFS location.
     */
    locationArn: pulumi.Input<string>;
}
