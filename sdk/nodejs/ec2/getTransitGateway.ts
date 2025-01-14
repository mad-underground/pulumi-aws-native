// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGateway
 */
export function getTransitGateway(args: GetTransitGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitGatewayResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getTransitGateway", {
        "id": args.id,
    }, opts);
}

export interface GetTransitGatewayArgs {
    id: string;
}

export interface GetTransitGatewayResult {
    readonly associationDefaultRouteTableId?: string;
    readonly autoAcceptSharedAttachments?: string;
    readonly defaultRouteTableAssociation?: string;
    readonly defaultRouteTablePropagation?: string;
    readonly description?: string;
    readonly dnsSupport?: string;
    readonly id?: string;
    readonly propagationDefaultRouteTableId?: string;
    readonly tags?: outputs.Tag[];
    readonly transitGatewayArn?: string;
    readonly transitGatewayCidrBlocks?: string[];
    readonly vpnEcmpSupport?: string;
}
/**
 * Resource Type definition for AWS::EC2::TransitGateway
 */
export function getTransitGatewayOutput(args: GetTransitGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitGatewayResult> {
    return pulumi.output(args).apply((a: any) => getTransitGateway(a, opts))
}

export interface GetTransitGatewayOutputArgs {
    id: pulumi.Input<string>;
}
