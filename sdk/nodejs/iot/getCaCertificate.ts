// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Registers a CA Certificate in IoT.
 */
export function getCaCertificate(args: GetCaCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCaCertificateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iot:getCaCertificate", {
        "id": args.id,
    }, opts);
}

export interface GetCaCertificateArgs {
    id: string;
}

export interface GetCaCertificateResult {
    readonly arn?: string;
    readonly autoRegistrationStatus?: enums.iot.CaCertificateAutoRegistrationStatus;
    readonly id?: string;
    readonly registrationConfig?: outputs.iot.CaCertificateRegistrationConfig;
    readonly status?: enums.iot.CaCertificateStatus;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.Tag[];
}
/**
 * Registers a CA Certificate in IoT.
 */
export function getCaCertificateOutput(args: GetCaCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCaCertificateResult> {
    return pulumi.output(args).apply((a: any) => getCaCertificate(a, opts))
}

export interface GetCaCertificateOutputArgs {
    id: pulumi.Input<string>;
}
