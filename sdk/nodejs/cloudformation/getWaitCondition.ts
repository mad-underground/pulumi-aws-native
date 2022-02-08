// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFormation::WaitCondition
 */
export function getWaitCondition(args: GetWaitConditionArgs, opts?: pulumi.InvokeOptions): Promise<GetWaitConditionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloudformation:getWaitCondition", {
        "id": args.id,
    }, opts);
}

export interface GetWaitConditionArgs {
    id: string;
}

export interface GetWaitConditionResult {
    readonly count?: number;
    readonly data?: any;
    readonly handle?: string;
    readonly id?: string;
    readonly timeout?: string;
}

export function getWaitConditionOutput(args: GetWaitConditionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetWaitConditionResult> {
    return pulumi.output(args).apply(a => getWaitCondition(a, opts))
}

export interface GetWaitConditionOutputArgs {
    id: pulumi.Input<string>;
}
