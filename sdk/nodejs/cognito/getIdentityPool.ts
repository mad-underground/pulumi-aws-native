// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::IdentityPool
 */
export function getIdentityPool(args: GetIdentityPoolArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentityPoolResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cognito:getIdentityPool", {
        "id": args.id,
    }, opts);
}

export interface GetIdentityPoolArgs {
    id: string;
}

export interface GetIdentityPoolResult {
    readonly allowClassicFlow?: boolean;
    readonly allowUnauthenticatedIdentities?: boolean;
    readonly cognitoIdentityProviders?: outputs.cognito.IdentityPoolCognitoIdentityProvider[];
    readonly developerProviderName?: string;
    readonly id?: string;
    readonly identityPoolName?: string;
    readonly name?: string;
    readonly openIdConnectProviderArns?: string[];
    readonly samlProviderArns?: string[];
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Cognito::IdentityPool` for more information about the expected schema for this property.
     */
    readonly supportedLoginProviders?: any;
}
/**
 * Resource Type definition for AWS::Cognito::IdentityPool
 */
export function getIdentityPoolOutput(args: GetIdentityPoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdentityPoolResult> {
    return pulumi.output(args).apply((a: any) => getIdentityPool(a, opts))
}

export interface GetIdentityPoolOutputArgs {
    id: pulumi.Input<string>;
}
