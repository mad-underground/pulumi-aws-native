// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ServiceCatalog::TagOption
 */
export function getTagOption(args: GetTagOptionArgs, opts?: pulumi.InvokeOptions): Promise<GetTagOptionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:servicecatalog:getTagOption", {
        "id": args.id,
    }, opts);
}

export interface GetTagOptionArgs {
    id: string;
}

export interface GetTagOptionResult {
    readonly active?: boolean;
    readonly id?: string;
}

export function getTagOptionOutput(args: GetTagOptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTagOptionResult> {
    return pulumi.output(args).apply(a => getTagOption(a, opts))
}

export interface GetTagOptionOutputArgs {
    id: pulumi.Input<string>;
}