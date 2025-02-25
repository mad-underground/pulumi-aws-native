// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::IoTSiteWise::Gateway
 */
export function getGateway(args: GetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotsitewise:getGateway", {
        "gatewayId": args.gatewayId,
    }, opts);
}

export interface GetGatewayArgs {
    /**
     * The ID of the gateway device.
     */
    gatewayId: string;
}

export interface GetGatewayResult {
    /**
     * A list of gateway capability summaries that each contain a namespace and status.
     */
    readonly gatewayCapabilitySummaries?: outputs.iotsitewise.GatewayCapabilitySummary[];
    /**
     * The ID of the gateway device.
     */
    readonly gatewayId?: string;
    /**
     * A unique, friendly name for the gateway.
     */
    readonly gatewayName?: string;
    /**
     * A list of key-value pairs that contain metadata for the gateway.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Resource schema for AWS::IoTSiteWise::Gateway
 */
export function getGatewayOutput(args: GetGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayResult> {
    return pulumi.output(args).apply((a: any) => getGateway(a, opts))
}

export interface GetGatewayOutputArgs {
    /**
     * The ID of the gateway device.
     */
    gatewayId: pulumi.Input<string>;
}
