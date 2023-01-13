// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use the AWS::IoT::Certificate resource to declare an AWS IoT X.509 certificate.
 */
export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getCertificate", {
        "id": args.id,
    }, opts);
}

export interface GetCertificateArgs {
    id: string;
}

export interface GetCertificateResult {
    readonly arn?: string;
    readonly id?: string;
    readonly status?: enums.iot.CertificateStatus;
}
/**
 * Use the AWS::IoT::Certificate resource to declare an AWS IoT X.509 certificate.
 */
export function getCertificateOutput(args: GetCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateResult> {
    return pulumi.output(args).apply((a: any) => getCertificate(a, opts))
}

export interface GetCertificateOutputArgs {
    id: pulumi.Input<string>;
}
