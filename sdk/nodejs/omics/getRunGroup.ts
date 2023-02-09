// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Omics::RunGroup Resource Type
 */
export function getRunGroup(args: GetRunGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetRunGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:omics:getRunGroup", {
        "id": args.id,
    }, opts);
}

export interface GetRunGroupArgs {
    id: string;
}

export interface GetRunGroupResult {
    readonly arn?: string;
    readonly creationTime?: string;
    readonly id?: string;
    readonly maxCpus?: number;
    readonly maxDuration?: number;
    readonly maxRuns?: number;
    readonly name?: string;
    readonly tags?: outputs.omics.RunGroupTagMap;
}
/**
 * Definition of AWS::Omics::RunGroup Resource Type
 */
export function getRunGroupOutput(args: GetRunGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRunGroupResult> {
    return pulumi.output(args).apply((a: any) => getRunGroup(a, opts))
}

export interface GetRunGroupOutputArgs {
    id: pulumi.Input<string>;
}
