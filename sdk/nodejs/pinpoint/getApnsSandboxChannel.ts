// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::APNSSandboxChannel
 */
export function getApnsSandboxChannel(args: GetApnsSandboxChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetApnsSandboxChannelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:pinpoint:getApnsSandboxChannel", {
        "id": args.id,
    }, opts);
}

export interface GetApnsSandboxChannelArgs {
    id: string;
}

export interface GetApnsSandboxChannelResult {
    readonly bundleId?: string;
    readonly certificate?: string;
    readonly defaultAuthenticationMethod?: string;
    readonly enabled?: boolean;
    readonly id?: string;
    readonly privateKey?: string;
    readonly teamId?: string;
    readonly tokenKey?: string;
    readonly tokenKeyId?: string;
}
/**
 * Resource Type definition for AWS::Pinpoint::APNSSandboxChannel
 */
export function getApnsSandboxChannelOutput(args: GetApnsSandboxChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApnsSandboxChannelResult> {
    return pulumi.output(args).apply((a: any) => getApnsSandboxChannel(a, opts))
}

export interface GetApnsSandboxChannelOutputArgs {
    id: pulumi.Input<string>;
}
