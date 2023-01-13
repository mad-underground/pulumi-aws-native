// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::AppFlow::Flow.
 */
export function getFlow(args: GetFlowArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:appflow:getFlow", {
        "flowName": args.flowName,
    }, opts);
}

export interface GetFlowArgs {
    /**
     * Name of the flow.
     */
    flowName: string;
}

export interface GetFlowResult {
    /**
     * Description of the flow.
     */
    readonly description?: string;
    /**
     * List of Destination connectors of the flow.
     */
    readonly destinationFlowConfigList?: outputs.appflow.FlowDestinationFlowConfig[];
    /**
     * ARN identifier of the flow.
     */
    readonly flowArn?: string;
    /**
     * Configurations of metadata catalog of the flow.
     */
    readonly metadataCatalogConfig?: outputs.appflow.FlowMetadataCatalogConfig;
    /**
     * Configurations of Source connector of the flow.
     */
    readonly sourceFlowConfig?: outputs.appflow.FlowSourceFlowConfig;
    /**
     * List of Tags.
     */
    readonly tags?: outputs.appflow.FlowTag[];
    /**
     * List of tasks for the flow.
     */
    readonly tasks?: outputs.appflow.FlowTask[];
    /**
     * Trigger settings of the flow.
     */
    readonly triggerConfig?: outputs.appflow.FlowTriggerConfig;
}
/**
 * Resource schema for AWS::AppFlow::Flow.
 */
export function getFlowOutput(args: GetFlowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlowResult> {
    return pulumi.output(args).apply((a: any) => getFlow(a, opts))
}

export interface GetFlowOutputArgs {
    /**
     * Name of the flow.
     */
    flowName: pulumi.Input<string>;
}
