// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGatewayV2::IntegrationResponse
 */
export function getIntegrationResponse(args: GetIntegrationResponseArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationResponseResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigatewayv2:getIntegrationResponse", {
        "id": args.id,
    }, opts);
}

export interface GetIntegrationResponseArgs {
    id: string;
}

export interface GetIntegrationResponseResult {
    readonly apiId?: string;
    readonly contentHandlingStrategy?: string;
    readonly id?: string;
    readonly integrationId?: string;
    readonly integrationResponseKey?: string;
    readonly responseParameters?: any;
    readonly responseTemplates?: any;
    readonly templateSelectionExpression?: string;
}

export function getIntegrationResponseOutput(args: GetIntegrationResponseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIntegrationResponseResult> {
    return pulumi.output(args).apply(a => getIntegrationResponse(a, opts))
}

export interface GetIntegrationResponseOutputArgs {
    id: pulumi.Input<string>;
}