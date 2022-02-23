// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::ManagedPolicy
 */
export function getManagedPolicy(args: GetManagedPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iam:getManagedPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetManagedPolicyArgs {
    id: string;
}

export interface GetManagedPolicyResult {
    readonly groups?: string[];
    readonly id?: string;
    readonly policyDocument?: any;
    readonly roles?: string[];
    readonly users?: string[];
}

export function getManagedPolicyOutput(args: GetManagedPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetManagedPolicyResult> {
    return pulumi.output(args).apply(a => getManagedPolicy(a, opts))
}

export interface GetManagedPolicyOutputArgs {
    id: pulumi.Input<string>;
}