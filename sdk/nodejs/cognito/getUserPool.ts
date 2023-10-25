// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::UserPool
 */
export function getUserPool(args: GetUserPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPoolResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cognito:getUserPool", {
        "userPoolId": args.userPoolId,
    }, opts);
}

export interface GetUserPoolArgs {
    userPoolId: string;
}

export interface GetUserPoolResult {
    readonly accountRecoverySetting?: outputs.cognito.UserPoolAccountRecoverySetting;
    readonly adminCreateUserConfig?: outputs.cognito.UserPoolAdminCreateUserConfig;
    readonly aliasAttributes?: string[];
    readonly arn?: string;
    readonly autoVerifiedAttributes?: string[];
    readonly deletionProtection?: string;
    readonly deviceConfiguration?: outputs.cognito.UserPoolDeviceConfiguration;
    readonly emailConfiguration?: outputs.cognito.UserPoolEmailConfiguration;
    readonly emailVerificationMessage?: string;
    readonly emailVerificationSubject?: string;
    readonly lambdaConfig?: outputs.cognito.UserPoolLambdaConfig;
    readonly mfaConfiguration?: string;
    readonly policies?: outputs.cognito.UserPoolPolicies;
    readonly providerName?: string;
    readonly providerUrl?: string;
    readonly schema?: outputs.cognito.UserPoolSchemaAttribute[];
    readonly smsAuthenticationMessage?: string;
    readonly smsConfiguration?: outputs.cognito.UserPoolSmsConfiguration;
    readonly smsVerificationMessage?: string;
    readonly userAttributeUpdateSettings?: outputs.cognito.UserPoolUserAttributeUpdateSettings;
    readonly userPoolAddOns?: outputs.cognito.UserPoolAddOns;
    readonly userPoolId?: string;
    readonly userPoolName?: string;
    readonly userPoolTags?: any;
    readonly usernameAttributes?: string[];
    readonly usernameConfiguration?: outputs.cognito.UserPoolUsernameConfiguration;
    readonly verificationMessageTemplate?: outputs.cognito.UserPoolVerificationMessageTemplate;
}
/**
 * Resource Type definition for AWS::Cognito::UserPool
 */
export function getUserPoolOutput(args: GetUserPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserPoolResult> {
    return pulumi.output(args).apply((a: any) => getUserPool(a, opts))
}

export interface GetUserPoolOutputArgs {
    userPoolId: pulumi.Input<string>;
}
