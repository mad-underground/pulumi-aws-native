// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGateway::DocumentationPart`` resource creates a documentation part for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.
 */
export function getDocumentationPart(args: GetDocumentationPartArgs, opts?: pulumi.InvokeOptions): Promise<GetDocumentationPartResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apigateway:getDocumentationPart", {
        "documentationPartId": args.documentationPartId,
        "restApiId": args.restApiId,
    }, opts);
}

export interface GetDocumentationPartArgs {
    /**
     * The identifier of the documentation Part.
     */
    documentationPartId: string;
    /**
     * The string identifier of the associated RestApi.
     */
    restApiId: string;
}

export interface GetDocumentationPartResult {
    /**
     * The identifier of the documentation Part.
     */
    readonly documentationPartId?: string;
    /**
     * The new documentation content map of the targeted API entity. Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value pairs can be exported and, hence, published.
     */
    readonly properties?: string;
}
/**
 * The ``AWS::ApiGateway::DocumentationPart`` resource creates a documentation part for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.
 */
export function getDocumentationPartOutput(args: GetDocumentationPartOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDocumentationPartResult> {
    return pulumi.output(args).apply((a: any) => getDocumentationPart(a, opts))
}

export interface GetDocumentationPartOutputArgs {
    /**
     * The identifier of the documentation Part.
     */
    documentationPartId: pulumi.Input<string>;
    /**
     * The string identifier of the associated RestApi.
     */
    restApiId: pulumi.Input<string>;
}
