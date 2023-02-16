// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Create and manage partner account
 */
export function getPartnerAccount(args: GetPartnerAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetPartnerAccountResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:iotwireless:getPartnerAccount", {
        "partnerAccountId": args.partnerAccountId,
    }, opts);
}

export interface GetPartnerAccountArgs {
    /**
     * The partner account ID to disassociate from the AWS account
     */
    partnerAccountId: string;
}

export interface GetPartnerAccountResult {
    /**
     * Whether the partner account is linked to the AWS account.
     */
    readonly accountLinked?: boolean;
    /**
     * PartnerAccount arn. Returned after successful create.
     */
    readonly arn?: string;
    /**
     * The fingerprint of the Sidewalk application server private key.
     */
    readonly fingerprint?: string;
    /**
     * The partner type
     */
    readonly partnerType?: enums.iotwireless.PartnerAccountPartnerType;
    /**
     * The Sidewalk account credentials.
     */
    readonly sidewalkResponse?: outputs.iotwireless.PartnerAccountSidewalkAccountInfoWithFingerprint;
    /**
     * A list of key-value pairs that contain metadata for the destination.
     */
    readonly tags?: outputs.iotwireless.PartnerAccountTag[];
}
/**
 * Create and manage partner account
 */
export function getPartnerAccountOutput(args: GetPartnerAccountOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPartnerAccountResult> {
    return pulumi.output(args).apply((a: any) => getPartnerAccount(a, opts))
}

export interface GetPartnerAccountOutputArgs {
    /**
     * The partner account ID to disassociate from the AWS account
     */
    partnerAccountId: pulumi.Input<string>;
}
