// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::ModelQualityJobDefinition
 */
export function getModelQualityJobDefinition(args: GetModelQualityJobDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetModelQualityJobDefinitionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getModelQualityJobDefinition", {
        "jobDefinitionArn": args.jobDefinitionArn,
    }, opts);
}

export interface GetModelQualityJobDefinitionArgs {
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    jobDefinitionArn: string;
}

export interface GetModelQualityJobDefinitionResult {
    /**
     * The time at which the job definition was created.
     */
    readonly creationTime?: string;
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    readonly jobDefinitionArn?: string;
}
/**
 * Resource Type definition for AWS::SageMaker::ModelQualityJobDefinition
 */
export function getModelQualityJobDefinitionOutput(args: GetModelQualityJobDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetModelQualityJobDefinitionResult> {
    return pulumi.output(args).apply((a: any) => getModelQualityJobDefinition(a, opts))
}

export interface GetModelQualityJobDefinitionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    jobDefinitionArn: pulumi.Input<string>;
}
