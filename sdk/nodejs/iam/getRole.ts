// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::Role
 */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iam:getRole", {
        "roleName": args.roleName,
    }, opts);
}

export interface GetRoleArgs {
    /**
     * A name for the IAM role, up to 64 characters in length.
     */
    roleName: string;
}

export interface GetRoleResult {
    /**
     * The Amazon Resource Name (ARN) for the role.
     */
    readonly arn?: string;
    /**
     * The trust policy that is associated with this role.
     */
    readonly assumeRolePolicyDocument?: any;
    /**
     * A description of the role that you provide.
     */
    readonly description?: string;
    /**
     * A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role. 
     */
    readonly managedPolicyArns?: string[];
    /**
     * The maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours. 
     */
    readonly maxSessionDuration?: number;
    /**
     * The ARN of the policy used to set the permissions boundary for the role.
     */
    readonly permissionsBoundary?: string;
    /**
     * Adds or updates an inline policy document that is embedded in the specified IAM role. 
     */
    readonly policies?: outputs.iam.RolePolicy[];
    /**
     * The stable and unique string identifying the role.
     */
    readonly roleId?: string;
    /**
     * A list of tags that are attached to the role.
     */
    readonly tags?: outputs.iam.RoleTag[];
}

export function getRoleOutput(args: GetRoleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRoleResult> {
    return pulumi.output(args).apply(a => getRole(a, opts))
}

export interface GetRoleOutputArgs {
    /**
     * A name for the IAM role, up to 64 characters in length.
     */
    roleName: pulumi.Input<string>;
}
