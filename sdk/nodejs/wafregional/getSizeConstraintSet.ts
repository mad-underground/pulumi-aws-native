// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::WAFRegional::SizeConstraintSet
 */
export function getSizeConstraintSet(args: GetSizeConstraintSetArgs, opts?: pulumi.InvokeOptions): Promise<GetSizeConstraintSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:wafregional:getSizeConstraintSet", {
        "id": args.id,
    }, opts);
}

export interface GetSizeConstraintSetArgs {
    id: string;
}

export interface GetSizeConstraintSetResult {
    readonly id?: string;
    readonly sizeConstraints?: outputs.wafregional.SizeConstraintSetSizeConstraint[];
}

export function getSizeConstraintSetOutput(args: GetSizeConstraintSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSizeConstraintSetResult> {
    return pulumi.output(args).apply(a => getSizeConstraintSet(a, opts))
}

export interface GetSizeConstraintSetOutputArgs {
    id: pulumi.Input<string>;
}
