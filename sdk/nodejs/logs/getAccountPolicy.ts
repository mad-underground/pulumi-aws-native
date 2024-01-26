// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.
 */
export function getAccountPolicy(args: GetAccountPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:logs:getAccountPolicy", {
        "accountId": args.accountId,
        "policyName": args.policyName,
        "policyType": args.policyType,
    }, opts);
}

export interface GetAccountPolicyArgs {
    /**
     * User account id
     */
    accountId: string;
    /**
     * The name of the account policy
     */
    policyName: string;
    /**
     * Type of the policy.
     */
    policyType: enums.logs.AccountPolicyPolicyType;
}

export interface GetAccountPolicyResult {
    /**
     * User account id
     */
    readonly accountId?: string;
    /**
     * The body of the policy document you want to use for this topic.
     *
     * You can only add one policy per PolicyType.
     *
     * The policy must be in JSON string format.
     *
     * Length Constraints: Maximum length of 30720
     */
    readonly policyDocument?: string;
    /**
     * Scope for policy application
     */
    readonly scope?: enums.logs.AccountPolicyScope;
    /**
     * Log group  selection criteria to apply policy only to a subset of log groups. SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes
     */
    readonly selectionCriteria?: string;
}
/**
 * The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.
 */
export function getAccountPolicyOutput(args: GetAccountPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccountPolicyResult> {
    return pulumi.output(args).apply((a: any) => getAccountPolicy(a, opts))
}

export interface GetAccountPolicyOutputArgs {
    /**
     * User account id
     */
    accountId: pulumi.Input<string>;
    /**
     * The name of the account policy
     */
    policyName: pulumi.Input<string>;
    /**
     * Type of the policy.
     */
    policyType: pulumi.Input<enums.logs.AccountPolicyPolicyType>;
}
