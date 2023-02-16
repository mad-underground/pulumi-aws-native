// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGatewayRouteTablePropagation
 */
export function getTransitGatewayRouteTablePropagation(args: GetTransitGatewayRouteTablePropagationArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitGatewayRouteTablePropagationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getTransitGatewayRouteTablePropagation", {
        "id": args.id,
    }, opts);
}

export interface GetTransitGatewayRouteTablePropagationArgs {
    id: string;
}

export interface GetTransitGatewayRouteTablePropagationResult {
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::EC2::TransitGatewayRouteTablePropagation
 */
export function getTransitGatewayRouteTablePropagationOutput(args: GetTransitGatewayRouteTablePropagationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitGatewayRouteTablePropagationResult> {
    return pulumi.output(args).apply((a: any) => getTransitGatewayRouteTablePropagation(a, opts))
}

export interface GetTransitGatewayRouteTablePropagationOutputArgs {
    id: pulumi.Input<string>;
}
