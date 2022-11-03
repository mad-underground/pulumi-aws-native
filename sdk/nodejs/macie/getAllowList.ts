// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Macie AllowList resource schema
 */
export function getAllowList(args: GetAllowListArgs, opts?: pulumi.InvokeOptions): Promise<GetAllowListResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:macie:getAllowList", {
        "id": args.id,
    }, opts);
}

export interface GetAllowListArgs {
    /**
     * AllowList ID.
     */
    id: string;
}

export interface GetAllowListResult {
    /**
     * AllowList ARN.
     */
    readonly arn?: string;
    /**
     * AllowList criteria.
     */
    readonly criteria?: outputs.macie.AllowListCriteria;
    /**
     * Description of AllowList.
     */
    readonly description?: string;
    /**
     * AllowList ID.
     */
    readonly id?: string;
    /**
     * Name of AllowList.
     */
    readonly name?: string;
    /**
     * AllowList status.
     */
    readonly status?: enums.macie.AllowListStatus;
    /**
     * A collection of tags associated with a resource
     */
    readonly tags?: outputs.macie.AllowListTag[];
}

export function getAllowListOutput(args: GetAllowListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAllowListResult> {
    return pulumi.output(args).apply(a => getAllowList(a, opts))
}

export interface GetAllowListOutputArgs {
    /**
     * AllowList ID.
     */
    id: pulumi.Input<string>;
}