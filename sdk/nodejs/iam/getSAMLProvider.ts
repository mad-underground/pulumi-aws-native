// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::IAM::SAMLProvider
 */
export function getSAMLProvider(args: GetSAMLProviderArgs, opts?: pulumi.InvokeOptions): Promise<GetSAMLProviderResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iam:getSAMLProvider", {
        "arn": args.arn,
    }, opts);
}

export interface GetSAMLProviderArgs {
    /**
     * Amazon Resource Name (ARN) of the SAML provider
     */
    arn: string;
}

export interface GetSAMLProviderResult {
    /**
     * Amazon Resource Name (ARN) of the SAML provider
     */
    readonly arn?: string;
    readonly samlMetadataDocument?: string;
    readonly tags?: outputs.iam.SAMLProviderTag[];
}
/**
 * Resource Type definition for AWS::IAM::SAMLProvider
 */
export function getSAMLProviderOutput(args: GetSAMLProviderOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSAMLProviderResult> {
    return pulumi.output(args).apply((a: any) => getSAMLProvider(a, opts))
}

export interface GetSAMLProviderOutputArgs {
    /**
     * Amazon Resource Name (ARN) of the SAML provider
     */
    arn: pulumi.Input<string>;
}
