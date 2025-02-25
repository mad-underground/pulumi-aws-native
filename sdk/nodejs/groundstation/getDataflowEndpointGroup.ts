// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * AWS Ground Station DataflowEndpointGroup schema for CloudFormation
 */
export function getDataflowEndpointGroup(args: GetDataflowEndpointGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDataflowEndpointGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:groundstation:getDataflowEndpointGroup", {
        "id": args.id,
    }, opts);
}

export interface GetDataflowEndpointGroupArgs {
    id: string;
}

export interface GetDataflowEndpointGroupResult {
    readonly arn?: string;
    /**
     * Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.
     */
    readonly contactPostPassDurationSeconds?: number;
    /**
     * Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.
     */
    readonly contactPrePassDurationSeconds?: number;
    readonly endpointDetails?: outputs.groundstation.DataflowEndpointGroupEndpointDetails[];
    readonly id?: string;
    readonly tags?: outputs.Tag[];
}
/**
 * AWS Ground Station DataflowEndpointGroup schema for CloudFormation
 */
export function getDataflowEndpointGroupOutput(args: GetDataflowEndpointGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDataflowEndpointGroupResult> {
    return pulumi.output(args).apply((a: any) => getDataflowEndpointGroup(a, opts))
}

export interface GetDataflowEndpointGroupOutputArgs {
    id: pulumi.Input<string>;
}
