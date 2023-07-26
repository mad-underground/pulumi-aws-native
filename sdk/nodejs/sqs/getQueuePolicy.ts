// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SQS::QueuePolicy
 */
export function getQueuePolicy(args: GetQueuePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetQueuePolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sqs:getQueuePolicy", {
        "id": args.id,
    }, opts);
}

export interface GetQueuePolicyArgs {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    id: string;
}

export interface GetQueuePolicyResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id?: string;
    /**
     * A policy document that contains the permissions for the specified Amazon SQS queues. For more information about Amazon SQS policies, see Creating Custom Policies Using the Access Policy Language in the Amazon Simple Queue Service Developer Guide.
     */
    readonly policyDocument?: any;
    /**
     * The URLs of the queues to which you want to add the policy. You can use the Ref function to specify an AWS::SQS::Queue resource.
     */
    readonly queues?: string[];
}
/**
 * Resource Type definition for AWS::SQS::QueuePolicy
 */
export function getQueuePolicyOutput(args: GetQueuePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueuePolicyResult> {
    return pulumi.output(args).apply((a: any) => getQueuePolicy(a, opts))
}

export interface GetQueuePolicyOutputArgs {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    id: pulumi.Input<string>;
}
