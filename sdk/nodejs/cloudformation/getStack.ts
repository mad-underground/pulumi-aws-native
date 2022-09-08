// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFormation::Stack
 */
export function getStack(args: GetStackArgs, opts?: pulumi.InvokeOptions): Promise<GetStackResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloudformation:getStack", {
        "id": args.id,
    }, opts);
}

export interface GetStackArgs {
    id: string;
}

export interface GetStackResult {
    readonly id?: string;
    readonly notificationARNs?: string[];
    readonly parameters?: any;
    readonly tags?: outputs.cloudformation.StackTag[];
    readonly templateURL?: string;
    readonly timeoutInMinutes?: number;
}

export function getStackOutput(args: GetStackOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStackResult> {
    return pulumi.output(args).apply(a => getStack(a, opts))
}

export interface GetStackOutputArgs {
    id: pulumi.Input<string>;
}
