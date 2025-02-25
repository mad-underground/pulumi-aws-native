// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::GameLift::Location resource creates an Amazon GameLift (GameLift) custom location.
 */
export function getLocation(args: GetLocationArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:gamelift:getLocation", {
        "locationName": args.locationName,
    }, opts);
}

export interface GetLocationArgs {
    locationName: string;
}

export interface GetLocationResult {
    readonly locationArn?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * The AWS::GameLift::Location resource creates an Amazon GameLift (GameLift) custom location.
 */
export function getLocationOutput(args: GetLocationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationResult> {
    return pulumi.output(args).apply((a: any) => getLocation(a, opts))
}

export interface GetLocationOutputArgs {
    locationName: pulumi.Input<string>;
}
