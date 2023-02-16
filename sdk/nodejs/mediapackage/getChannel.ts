// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaPackage::Channel
 */
export function getChannel(args: GetChannelArgs, opts?: pulumi.InvokeOptions): Promise<GetChannelResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:mediapackage:getChannel", {
        "id": args.id,
    }, opts);
}

export interface GetChannelArgs {
    /**
     * The ID of the Channel.
     */
    id: string;
}

export interface GetChannelResult {
    /**
     * The Amazon Resource Name (ARN) assigned to the Channel.
     */
    readonly arn?: string;
    /**
     * A short text description of the Channel.
     */
    readonly description?: string;
    /**
     * The configuration parameters for egress access logging.
     */
    readonly egressAccessLogs?: outputs.mediapackage.ChannelLogConfiguration;
    /**
     * An HTTP Live Streaming (HLS) ingest resource configuration.
     */
    readonly hlsIngest?: outputs.mediapackage.ChannelHlsIngest;
    /**
     * The configuration parameters for egress access logging.
     */
    readonly ingressAccessLogs?: outputs.mediapackage.ChannelLogConfiguration;
}
/**
 * Resource schema for AWS::MediaPackage::Channel
 */
export function getChannelOutput(args: GetChannelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChannelResult> {
    return pulumi.output(args).apply((a: any) => getChannel(a, opts))
}

export interface GetChannelOutputArgs {
    /**
     * The ID of the Channel.
     */
    id: pulumi.Input<string>;
}
