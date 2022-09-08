// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::KMS::ReplicaKey resource specifies a multi-region replica AWS KMS key in AWS Key Management Service (AWS KMS).
 */
export function getReplicaKey(args: GetReplicaKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetReplicaKeyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:kms:getReplicaKey", {
        "keyId": args.keyId,
    }, opts);
}

export interface GetReplicaKeyArgs {
    keyId: string;
}

export interface GetReplicaKeyResult {
    readonly arn?: string;
    /**
     * A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
     */
    readonly description?: string;
    /**
     * Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
     */
    readonly enabled?: boolean;
    readonly keyId?: string;
    /**
     * The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
     */
    readonly keyPolicy?: any;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.kms.ReplicaKeyTag[];
}

export function getReplicaKeyOutput(args: GetReplicaKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetReplicaKeyResult> {
    return pulumi.output(args).apply(a => getReplicaKey(a, opts))
}

export interface GetReplicaKeyOutputArgs {
    keyId: pulumi.Input<string>;
}
