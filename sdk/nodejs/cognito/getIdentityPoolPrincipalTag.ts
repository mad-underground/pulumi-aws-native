// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::IdentityPoolPrincipalTag
 */
export function getIdentityPoolPrincipalTag(args: GetIdentityPoolPrincipalTagArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityPoolPrincipalTagResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cognito:getIdentityPoolPrincipalTag", {
        "identityPoolId": args.identityPoolId,
        "identityProviderName": args.identityProviderName,
    }, opts);
}

export interface GetIdentityPoolPrincipalTagArgs {
    identityPoolId: string;
    identityProviderName: string;
}

export interface GetIdentityPoolPrincipalTagResult {
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::IdentityPoolPrincipalTag` for more information about the expected schema for this property.
     */
    readonly principalTags?: any;
    readonly useDefaults?: boolean;
}
/**
 * Resource Type definition for AWS::Cognito::IdentityPoolPrincipalTag
 */
export function getIdentityPoolPrincipalTagOutput(args: GetIdentityPoolPrincipalTagOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdentityPoolPrincipalTagResult> {
    return pulumi.output(args).apply((a: any) => getIdentityPoolPrincipalTag(a, opts))
}

export interface GetIdentityPoolPrincipalTagOutputArgs {
    identityPoolId: pulumi.Input<string>;
    identityProviderName: pulumi.Input<string>;
}
