// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Schema for IAM Role Policy
 */
export function getRolePolicy(args: GetRolePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetRolePolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iam:getRolePolicy", {
        "policyName": args.policyName,
        "roleName": args.roleName,
    }, opts);
}

export interface GetRolePolicyArgs {
    /**
     * The friendly name (not ARN) identifying the policy.
     */
    policyName: string;
    /**
     * The name of the policy document.
     */
    roleName: string;
}

export interface GetRolePolicyResult {
    /**
     * The policy document.
     */
    readonly policyDocument?: any;
}
/**
 * Schema for IAM Role Policy
 */
export function getRolePolicyOutput(args: GetRolePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRolePolicyResult> {
    return pulumi.output(args).apply((a: any) => getRolePolicy(a, opts))
}

export interface GetRolePolicyOutputArgs {
    /**
     * The friendly name (not ARN) identifying the policy.
     */
    policyName: pulumi.Input<string>;
    /**
     * The name of the policy document.
     */
    roleName: pulumi.Input<string>;
}
