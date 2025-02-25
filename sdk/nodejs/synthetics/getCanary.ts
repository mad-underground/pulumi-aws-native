// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Synthetics::Canary
 */
export function getCanary(args: GetCanaryArgs, opts?: pulumi.InvokeOptions): Promise<GetCanaryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:synthetics:getCanary", {
        "name": args.name,
    }, opts);
}

export interface GetCanaryArgs {
    /**
     * Name of the canary.
     */
    name: string;
}

export interface GetCanaryResult {
    /**
     * Provide artifact configuration
     */
    readonly artifactConfig?: outputs.synthetics.CanaryArtifactConfig;
    /**
     * Provide the s3 bucket output location for test results
     */
    readonly artifactS3Location?: string;
    /**
     * Provide the canary script source
     */
    readonly code?: outputs.synthetics.CanaryCode;
    /**
     * Lambda Execution role used to run your canaries
     */
    readonly executionRoleArn?: string;
    /**
     * Retention period of failed canary runs represented in number of days
     */
    readonly failureRetentionPeriod?: number;
    /**
     * Id of the canary
     */
    readonly id?: string;
    /**
     * Provide canary run configuration
     */
    readonly runConfig?: outputs.synthetics.CanaryRunConfig;
    /**
     * Runtime version of Synthetics Library
     */
    readonly runtimeVersion?: string;
    /**
     * Frequency to run your canaries
     */
    readonly schedule?: outputs.synthetics.CanarySchedule;
    /**
     * State of the canary
     */
    readonly state?: string;
    /**
     * Retention period of successful canary runs represented in number of days
     */
    readonly successRetentionPeriod?: number;
    readonly tags?: outputs.Tag[];
    /**
     * Provide VPC Configuration if enabled.
     */
    readonly vpcConfig?: outputs.synthetics.CanaryVpcConfig;
}
/**
 * Resource Type definition for AWS::Synthetics::Canary
 */
export function getCanaryOutput(args: GetCanaryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCanaryResult> {
    return pulumi.output(args).apply((a: any) => getCanary(a, opts))
}

export interface GetCanaryOutputArgs {
    /**
     * Name of the canary.
     */
    name: pulumi.Input<string>;
}
