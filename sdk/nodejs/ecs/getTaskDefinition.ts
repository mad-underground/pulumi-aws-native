// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Schema describing various properties for ECS TaskDefinition
 */
export function getTaskDefinition(args: GetTaskDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ecs:getTaskDefinition", {
        "taskDefinitionArn": args.taskDefinitionArn,
    }, opts);
}

export interface GetTaskDefinitionArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon ECS task definition
     */
    taskDefinitionArn: string;
}

export interface GetTaskDefinitionResult {
    readonly tags?: outputs.ecs.TaskDefinitionTag[];
    /**
     * The Amazon Resource Name (ARN) of the Amazon ECS task definition
     */
    readonly taskDefinitionArn?: string;
}

export function getTaskDefinitionOutput(args: GetTaskDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTaskDefinitionResult> {
    return pulumi.output(args).apply(a => getTaskDefinition(a, opts))
}

export interface GetTaskDefinitionOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the Amazon ECS task definition
     */
    taskDefinitionArn: pulumi.Input<string>;
}
