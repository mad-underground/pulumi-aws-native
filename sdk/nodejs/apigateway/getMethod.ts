// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::ApiGateway::Method
 */
export function getMethod(args: GetMethodArgs, opts?: pulumi.InvokeOptions): Promise<GetMethodResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:apigateway:getMethod", {
        "httpMethod": args.httpMethod,
        "resourceId": args.resourceId,
        "restApiId": args.restApiId,
    }, opts);
}

export interface GetMethodArgs {
    /**
     * The backend system that the method calls when it receives a request.
     */
    httpMethod: string;
    /**
     * The ID of an API Gateway resource.
     */
    resourceId: string;
    /**
     * The ID of the RestApi resource in which API Gateway creates the method.
     */
    restApiId: string;
}

export interface GetMethodResult {
    /**
     * Indicates whether the method requires clients to submit a valid API key.
     */
    readonly apiKeyRequired?: boolean;
    /**
     * A list of authorization scopes configured on the method.
     */
    readonly authorizationScopes?: string[];
    /**
     * The method's authorization type.
     */
    readonly authorizationType?: enums.apigateway.MethodAuthorizationType;
    /**
     * The identifier of the authorizer to use on this method.
     */
    readonly authorizerId?: string;
    /**
     * The backend system that the method calls when it receives a request.
     */
    readonly integration?: outputs.apigateway.MethodIntegration;
    /**
     * The responses that can be sent to the client who calls the method.
     */
    readonly methodResponses?: outputs.apigateway.MethodResponse[];
    /**
     * A friendly operation name for the method.
     */
    readonly operationName?: string;
    /**
     * The resources that are used for the request's content type. Specify request models as key-value pairs (string-to-string mapping), with a content type as the key and a Model resource name as the value.
     */
    readonly requestModels?: any;
    /**
     * The request parameters that API Gateway accepts. Specify request parameters as key-value pairs (string-to-Boolean mapping), with a source as the key and a Boolean as the value.
     */
    readonly requestParameters?: any;
    /**
     * The ID of the associated request validator.
     */
    readonly requestValidatorId?: string;
}

export function getMethodOutput(args: GetMethodOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMethodResult> {
    return pulumi.output(args).apply(a => getMethod(a, opts))
}

export interface GetMethodOutputArgs {
    /**
     * The backend system that the method calls when it receives a request.
     */
    httpMethod: pulumi.Input<string>;
    /**
     * The ID of an API Gateway resource.
     */
    resourceId: pulumi.Input<string>;
    /**
     * The ID of the RestApi resource in which API Gateway creates the method.
     */
    restApiId: pulumi.Input<string>;
}
