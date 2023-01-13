// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::EventBusPolicy
 */
export function getEventBusPolicy(args: GetEventBusPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetEventBusPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:events:getEventBusPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetEventBusPolicyArgs {
    id: string;
}

export interface GetEventBusPolicyResult {
    readonly action?: string;
    readonly condition?: outputs.events.EventBusPolicyCondition;
    readonly id?: string;
    readonly principal?: string;
    readonly statement?: any;
}
/**
 * Resource Type definition for AWS::Events::EventBusPolicy
 */
export function getEventBusPolicyOutput(args: GetEventBusPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEventBusPolicyResult> {
    return pulumi.output(args).apply((a: any) => getEventBusPolicy(a, opts))
}

export interface GetEventBusPolicyOutputArgs {
    id: pulumi.Input<string>;
}
