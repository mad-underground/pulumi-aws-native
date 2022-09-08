// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CodePipeline::CustomActionType
 */
export function getCustomActionType(args: GetCustomActionTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomActionTypeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:codepipeline:getCustomActionType", {
        "id": args.id,
    }, opts);
}

export interface GetCustomActionTypeArgs {
    id: string;
}

export interface GetCustomActionTypeResult {
    readonly id?: string;
    readonly tags?: outputs.codepipeline.CustomActionTypeTag[];
}

export function getCustomActionTypeOutput(args: GetCustomActionTypeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomActionTypeResult> {
    return pulumi.output(args).apply(a => getCustomActionType(a, opts))
}

export interface GetCustomActionTypeOutputArgs {
    id: pulumi.Input<string>;
}
