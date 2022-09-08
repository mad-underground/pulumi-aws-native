// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceDiscovery::PublicDnsNamespace
 */
export function getPublicDnsNamespace(args: GetPublicDnsNamespaceArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicDnsNamespaceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:servicediscovery:getPublicDnsNamespace", {
        "id": args.id,
    }, opts);
}

export interface GetPublicDnsNamespaceArgs {
    id: string;
}

export interface GetPublicDnsNamespaceResult {
    readonly arn?: string;
    readonly description?: string;
    readonly hostedZoneId?: string;
    readonly id?: string;
    readonly properties?: outputs.servicediscovery.PublicDnsNamespaceProperties;
    readonly tags?: outputs.servicediscovery.PublicDnsNamespaceTag[];
}

export function getPublicDnsNamespaceOutput(args: GetPublicDnsNamespaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublicDnsNamespaceResult> {
    return pulumi.output(args).apply(a => getPublicDnsNamespace(a, opts))
}

export interface GetPublicDnsNamespaceOutputArgs {
    id: pulumi.Input<string>;
}
