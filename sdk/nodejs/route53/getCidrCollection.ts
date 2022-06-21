// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Route53::CidrCollection.
 */
export function getCidrCollection(args: GetCidrCollectionArgs, opts?: pulumi.InvokeOptions): Promise<GetCidrCollectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:route53:getCidrCollection", {
        "id": args.id,
    }, opts);
}

export interface GetCidrCollectionArgs {
    /**
     * UUID of the CIDR collection.
     */
    id: string;
}

export interface GetCidrCollectionResult {
    /**
     * The Amazon resource name (ARN) to uniquely identify the AWS resource.
     */
    readonly arn?: string;
    /**
     * UUID of the CIDR collection.
     */
    readonly id?: string;
    /**
     * A complex type that contains information about the list of CIDR locations.
     */
    readonly locations?: outputs.route53.CidrCollectionLocation[];
}

export function getCidrCollectionOutput(args: GetCidrCollectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCidrCollectionResult> {
    return pulumi.output(args).apply(a => getCidrCollection(a, opts))
}

export interface GetCidrCollectionOutputArgs {
    /**
     * UUID of the CIDR collection.
     */
    id: pulumi.Input<string>;
}
