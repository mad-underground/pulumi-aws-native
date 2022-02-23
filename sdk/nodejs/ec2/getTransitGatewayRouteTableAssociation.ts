// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGatewayRouteTableAssociation
 */
export function getTransitGatewayRouteTableAssociation(args: GetTransitGatewayRouteTableAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitGatewayRouteTableAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getTransitGatewayRouteTableAssociation", {
        "id": args.id,
    }, opts);
}

export interface GetTransitGatewayRouteTableAssociationArgs {
    id: string;
}

export interface GetTransitGatewayRouteTableAssociationResult {
    readonly id?: string;
}

export function getTransitGatewayRouteTableAssociationOutput(args: GetTransitGatewayRouteTableAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitGatewayRouteTableAssociationResult> {
    return pulumi.output(args).apply(a => getTransitGatewayRouteTableAssociation(a, opts))
}

export interface GetTransitGatewayRouteTableAssociationOutputArgs {
    id: pulumi.Input<string>;
}