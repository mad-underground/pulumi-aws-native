// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAF::ByteMatchSet
 */
export function getByteMatchSet(args: GetByteMatchSetArgs, opts?: pulumi.InvokeOptions): Promise<GetByteMatchSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:waf:getByteMatchSet", {
        "id": args.id,
    }, opts);
}

export interface GetByteMatchSetArgs {
    id: string;
}

export interface GetByteMatchSetResult {
    readonly byteMatchTuples?: outputs.waf.ByteMatchSetByteMatchTuple[];
    readonly id?: string;
}

export function getByteMatchSetOutput(args: GetByteMatchSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetByteMatchSetResult> {
    return pulumi.output(args).apply(a => getByteMatchSet(a, opts))
}

export interface GetByteMatchSetOutputArgs {
    id: pulumi.Input<string>;
}
