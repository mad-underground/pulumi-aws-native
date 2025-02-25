// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGateway::ClientCertificate`` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.
 */
export function getClientCertificate(args: GetClientCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetClientCertificateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:apigateway:getClientCertificate", {
        "clientCertificateId": args.clientCertificateId,
    }, opts);
}

export interface GetClientCertificateArgs {
    clientCertificateId: string;
}

export interface GetClientCertificateResult {
    readonly clientCertificateId?: string;
    /**
     * The description of the client certificate.
     */
    readonly description?: string;
    /**
     * The collection of tags. Each tag element is associated with a given resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * The ``AWS::ApiGateway::ClientCertificate`` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.
 */
export function getClientCertificateOutput(args: GetClientCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClientCertificateResult> {
    return pulumi.output(args).apply((a: any) => getClientCertificate(a, opts))
}

export interface GetClientCertificateOutputArgs {
    clientCertificateId: pulumi.Input<string>;
}
