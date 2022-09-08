// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::Pipeline
 */
export function getPipeline(args: GetPipelineArgs, opts?: pulumi.InvokeOptions): Promise<GetPipelineResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sagemaker:getPipeline", {
        "pipelineName": args.pipelineName,
    }, opts);
}

export interface GetPipelineArgs {
    /**
     * The name of the Pipeline.
     */
    pipelineName: string;
}

export interface GetPipelineResult {
    readonly parallelismConfiguration?: outputs.sagemaker.ParallelismConfigurationProperties;
    readonly pipelineDefinition?: outputs.sagemaker.PipelineDefinition0Properties | outputs.sagemaker.PipelineDefinition1Properties;
    /**
     * The description of the Pipeline.
     */
    readonly pipelineDescription?: string;
    /**
     * The display name of the Pipeline.
     */
    readonly pipelineDisplayName?: string;
    /**
     * Role Arn
     */
    readonly roleArn?: string;
    readonly tags?: outputs.sagemaker.PipelineTag[];
}

export function getPipelineOutput(args: GetPipelineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPipelineResult> {
    return pulumi.output(args).apply(a => getPipeline(a, opts))
}

export interface GetPipelineOutputArgs {
    /**
     * The name of the Pipeline.
     */
    pipelineName: pulumi.Input<string>;
}
