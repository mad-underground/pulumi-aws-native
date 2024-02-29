// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.
 */
export function getTaskSet(args: GetTaskSetArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskSetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:ecs:getTaskSet", {
        "cluster": args.cluster,
        "id": args.id,
        "service": args.service,
    }, opts);
}

export interface GetTaskSetArgs {
    /**
     * The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
     */
    cluster: string;
    /**
     * The ID of the task set.
     */
    id: string;
    /**
     * The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
     */
    service: string;
}

export interface GetTaskSetResult {
    /**
     * The ID of the task set.
     */
    readonly id?: string;
    /**
     * A floating-point percentage of the desired number of tasks to place and keep running in the task set.
     */
    readonly scale?: outputs.ecs.TaskSetScale;
    readonly tags?: outputs.Tag[];
}
/**
 * Create a task set in the specified cluster and service. This is used when a service uses the EXTERNAL deployment controller type. For more information, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.htmlin the Amazon Elastic Container Service Developer Guide.
 */
export function getTaskSetOutput(args: GetTaskSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTaskSetResult> {
    return pulumi.output(args).apply((a: any) => getTaskSet(a, opts))
}

export interface GetTaskSetOutputArgs {
    /**
     * The short name or full Amazon Resource Name (ARN) of the cluster that hosts the service to create the task set in.
     */
    cluster: pulumi.Input<string>;
    /**
     * The ID of the task set.
     */
    id: pulumi.Input<string>;
    /**
     * The short name or full Amazon Resource Name (ARN) of the service to create the task set in.
     */
    service: pulumi.Input<string>;
}
