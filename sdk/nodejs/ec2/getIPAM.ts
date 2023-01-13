// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Schema of AWS::EC2::IPAM Type
 */
export function getIPAM(args: GetIPAMArgs, opts?: pulumi.InvokeOptions): Promise<GetIPAMResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ec2:getIPAM", {
        "ipamId": args.ipamId,
    }, opts);
}

export interface GetIPAMArgs {
    /**
     * Id of the IPAM.
     */
    ipamId: string;
}

export interface GetIPAMResult {
    /**
     * The Amazon Resource Name (ARN) of the IPAM.
     */
    readonly arn?: string;
    readonly description?: string;
    /**
     * Id of the IPAM.
     */
    readonly ipamId?: string;
    /**
     * The regions IPAM is enabled for. Allows pools to be created in these regions, as well as enabling monitoring
     */
    readonly operatingRegions?: outputs.ec2.IPAMIpamOperatingRegion[];
    /**
     * The Id of the default scope for publicly routable IP space, created with this IPAM.
     */
    readonly privateDefaultScopeId?: string;
    /**
     * The Id of the default scope for publicly routable IP space, created with this IPAM.
     */
    readonly publicDefaultScopeId?: string;
    /**
     * The number of scopes that currently exist in this IPAM.
     */
    readonly scopeCount?: number;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.ec2.IPAMTag[];
}
/**
 * Resource Schema of AWS::EC2::IPAM Type
 */
export function getIPAMOutput(args: GetIPAMOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIPAMResult> {
    return pulumi.output(args).apply((a: any) => getIPAM(a, opts))
}

export interface GetIPAMOutputArgs {
    /**
     * Id of the IPAM.
     */
    ipamId: pulumi.Input<string>;
}
