// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Batch::JobDefinition
 */
export function getJobDefinition(args: GetJobDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetJobDefinitionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:batch:getJobDefinition", {
        "id": args.id,
    }, opts);
}

export interface GetJobDefinitionArgs {
    id: string;
}

export interface GetJobDefinitionResult {
    readonly containerProperties?: outputs.batch.JobDefinitionContainerProperties;
    readonly ecsProperties?: outputs.batch.JobDefinitionEcsProperties;
    readonly eksProperties?: outputs.batch.JobDefinitionEksProperties;
    readonly id?: string;
    readonly nodeProperties?: outputs.batch.JobDefinitionNodeProperties;
    /**
     * Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
     */
    readonly parameters?: any;
    readonly platformCapabilities?: string[];
    readonly propagateTags?: boolean;
    readonly retryStrategy?: outputs.batch.JobDefinitionRetryStrategy;
    readonly schedulingPriority?: number;
    readonly timeout?: outputs.batch.JobDefinitionTimeout;
    readonly type?: string;
}
/**
 * Resource Type definition for AWS::Batch::JobDefinition
 */
export function getJobDefinitionOutput(args: GetJobDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobDefinitionResult> {
    return pulumi.output(args).apply((a: any) => getJobDefinition(a, opts))
}

export interface GetJobDefinitionOutputArgs {
    id: pulumi.Input<string>;
}
