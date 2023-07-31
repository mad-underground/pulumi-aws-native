// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::DataQualityJobDefinition
 */
export function getDataQualityJobDefinition(args: GetDataQualityJobDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetDataQualityJobDefinitionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:sagemaker:getDataQualityJobDefinition", {
        "jobDefinitionArn": args.jobDefinitionArn,
    }, opts);
}

export interface GetDataQualityJobDefinitionArgs {
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    jobDefinitionArn: string;
}

export interface GetDataQualityJobDefinitionResult {
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
 * Resource Type definition for AWS::SageMaker::DataQualityJobDefinition
 */
export function getDataQualityJobDefinitionOutput(args: GetDataQualityJobDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataQualityJobDefinitionResult> {
    return pulumi.output(args).apply((a: any) => getDataQualityJobDefinition(a, opts))
}

export interface GetDataQualityJobDefinitionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of job definition.
     */
    jobDefinitionArn: pulumi.Input<string>;
}
