// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS::DeviceFarm::VPCEConfiguration creates a new Device Farm VPCE Configuration
 */
export function getVPCEConfiguration(args: GetVPCEConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetVPCEConfigurationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:devicefarm:getVPCEConfiguration", {
        "arn": args.arn,
    }, opts);
}

export interface GetVPCEConfigurationArgs {
    arn: string;
}

export interface GetVPCEConfigurationResult {
    readonly arn?: string;
    readonly serviceDnsName?: string;
    readonly tags?: outputs.devicefarm.VPCEConfigurationTag[];
    readonly vpceConfigurationDescription?: string;
    readonly vpceConfigurationName?: string;
    readonly vpceServiceName?: string;
}
/**
 * AWS::DeviceFarm::VPCEConfiguration creates a new Device Farm VPCE Configuration
 */
export function getVPCEConfigurationOutput(args: GetVPCEConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVPCEConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getVPCEConfiguration(a, opts))
}

export interface GetVPCEConfigurationOutputArgs {
    arn: pulumi.Input<string>;
}
