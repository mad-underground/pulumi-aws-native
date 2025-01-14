// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppMesh::VirtualNode
 */
export function getVirtualNode(args: GetVirtualNodeArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualNodeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:appmesh:getVirtualNode", {
        "id": args.id,
    }, opts);
}

export interface GetVirtualNodeArgs {
    id: string;
}

export interface GetVirtualNodeResult {
    readonly arn?: string;
    readonly id?: string;
    readonly resourceOwner?: string;
    readonly spec?: outputs.appmesh.VirtualNodeSpec;
    readonly tags?: outputs.Tag[];
    readonly uid?: string;
}
/**
 * Resource Type definition for AWS::AppMesh::VirtualNode
 */
export function getVirtualNodeOutput(args: GetVirtualNodeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVirtualNodeResult> {
    return pulumi.output(args).apply((a: any) => getVirtualNode(a, opts))
}

export interface GetVirtualNodeOutputArgs {
    id: pulumi.Input<string>;
}
