// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS::DeviceFarm::DevicePool creates a new Device Pool for a given DF Project
 */
export function getDevicePool(args: GetDevicePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetDevicePoolResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:devicefarm:getDevicePool", {
        "arn": args.arn,
    }, opts);
}

export interface GetDevicePoolArgs {
    arn: string;
}

export interface GetDevicePoolResult {
    readonly arn?: string;
    readonly description?: string;
    readonly maxDevices?: number;
    readonly name?: string;
    readonly rules?: outputs.devicefarm.DevicePoolRule[];
    readonly tags?: outputs.Tag[];
}
/**
 * AWS::DeviceFarm::DevicePool creates a new Device Pool for a given DF Project
 */
export function getDevicePoolOutput(args: GetDevicePoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDevicePoolResult> {
    return pulumi.output(args).apply((a: any) => getDevicePool(a, opts))
}

export interface GetDevicePoolOutputArgs {
    arn: pulumi.Input<string>;
}
