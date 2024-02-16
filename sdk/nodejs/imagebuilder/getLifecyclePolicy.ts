// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::ImageBuilder::LifecyclePolicy
 */
export function getLifecyclePolicy(args: GetLifecyclePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetLifecyclePolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:imagebuilder:getLifecyclePolicy", {
        "arn": args.arn,
    }, opts);
}

export interface GetLifecyclePolicyArgs {
    /**
     * The Amazon Resource Name (ARN) of the lifecycle policy.
     */
    arn: string;
}

export interface GetLifecyclePolicyResult {
    /**
     * The Amazon Resource Name (ARN) of the lifecycle policy.
     */
    readonly arn?: string;
    /**
     * The description of the lifecycle policy.
     */
    readonly description?: string;
    /**
     * The execution role of the lifecycle policy.
     */
    readonly executionRole?: string;
    /**
     * The policy details of the lifecycle policy.
     */
    readonly policyDetails?: outputs.imagebuilder.LifecyclePolicyPolicyDetail[];
    /**
     * The resource selection of the lifecycle policy.
     */
    readonly resourceSelection?: outputs.imagebuilder.LifecyclePolicyResourceSelection;
    /**
     * The resource type of the lifecycle policy.
     */
    readonly resourceType?: enums.imagebuilder.LifecyclePolicyResourceType;
    /**
     * The status of the lifecycle policy.
     */
    readonly status?: enums.imagebuilder.LifecyclePolicyStatus;
    /**
     * The tags associated with the lifecycle policy.
     */
    readonly tags?: {[key: string]: string};
}
/**
 * Resource schema for AWS::ImageBuilder::LifecyclePolicy
 */
export function getLifecyclePolicyOutput(args: GetLifecyclePolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLifecyclePolicyResult> {
    return pulumi.output(args).apply((a: any) => getLifecyclePolicy(a, opts))
}

export interface GetLifecyclePolicyOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the lifecycle policy.
     */
    arn: pulumi.Input<string>;
}
