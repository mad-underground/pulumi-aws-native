// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::NetworkInterfaceAttachment
 */
export function getNetworkInterfaceAttachment(args: GetNetworkInterfaceAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInterfaceAttachmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getNetworkInterfaceAttachment", {
        "id": args.id,
    }, opts);
}

export interface GetNetworkInterfaceAttachmentArgs {
    id: string;
}

export interface GetNetworkInterfaceAttachmentResult {
    readonly deleteOnTermination?: boolean;
    readonly deviceIndex?: string;
    readonly id?: string;
    readonly instanceId?: string;
    readonly networkInterfaceId?: string;
}

export function getNetworkInterfaceAttachmentOutput(args: GetNetworkInterfaceAttachmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkInterfaceAttachmentResult> {
    return pulumi.output(args).apply(a => getNetworkInterfaceAttachment(a, opts))
}

export interface GetNetworkInterfaceAttachmentOutputArgs {
    id: pulumi.Input<string>;
}