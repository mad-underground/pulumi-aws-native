// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Schema for IAM User Policy
 */
export function getUserPolicy(args: GetUserPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iam:getUserPolicy", {
        "policyName": args.policyName,
        "userName": args.userName,
    }, opts);
}

export interface GetUserPolicyArgs {
    /**
     * The name of the policy document.
     */
    policyName: string;
    /**
     * The name of the user to associate the policy with.
     */
    userName: string;
}

export interface GetUserPolicyResult {
    /**
     * The policy document.
     */
    readonly policyDocument?: any;
}
/**
 * Schema for IAM User Policy
 */
export function getUserPolicyOutput(args: GetUserPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserPolicyResult> {
    return pulumi.output(args).apply((a: any) => getUserPolicy(a, opts))
}

export interface GetUserPolicyOutputArgs {
    /**
     * The name of the policy document.
     */
    policyName: pulumi.Input<string>;
    /**
     * The name of the user to associate the policy with.
     */
    userName: pulumi.Input<string>;
}
