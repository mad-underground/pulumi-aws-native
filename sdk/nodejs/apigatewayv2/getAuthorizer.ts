// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGatewayV2::Authorizer
 */
export function getAuthorizer(args: GetAuthorizerArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthorizerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apigatewayv2:getAuthorizer", {
        "apiId": args.apiId,
        "authorizerId": args.authorizerId,
    }, opts);
}

export interface GetAuthorizerArgs {
    apiId: string;
    authorizerId: string;
}

export interface GetAuthorizerResult {
    readonly authorizerCredentialsArn?: string;
    readonly authorizerId?: string;
    readonly authorizerPayloadFormatVersion?: string;
    readonly authorizerResultTtlInSeconds?: number;
    readonly authorizerType?: string;
    readonly authorizerUri?: string;
    readonly enableSimpleResponses?: boolean;
    readonly identitySource?: string[];
    readonly identityValidationExpression?: string;
    readonly jwtConfiguration?: outputs.apigatewayv2.AuthorizerJWTConfiguration;
    readonly name?: string;
}
/**
 * Resource Type definition for AWS::ApiGatewayV2::Authorizer
 */
export function getAuthorizerOutput(args: GetAuthorizerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAuthorizerResult> {
    return pulumi.output(args).apply((a: any) => getAuthorizer(a, opts))
}

export interface GetAuthorizerOutputArgs {
    apiId: pulumi.Input<string>;
    authorizerId: pulumi.Input<string>;
}
