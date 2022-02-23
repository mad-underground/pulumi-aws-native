// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::PinpointEmail::ConfigurationSet
 */
export function getConfigurationSet(args: GetConfigurationSetArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:pinpointemail:getConfigurationSet", {
        "id": args.id,
    }, opts);
}

export interface GetConfigurationSetArgs {
    id: string;
}

export interface GetConfigurationSetResult {
    readonly deliveryOptions?: outputs.pinpointemail.ConfigurationSetDeliveryOptions;
    readonly id?: string;
    readonly reputationOptions?: outputs.pinpointemail.ConfigurationSetReputationOptions;
    readonly sendingOptions?: outputs.pinpointemail.ConfigurationSetSendingOptions;
    readonly tags?: outputs.pinpointemail.ConfigurationSetTags[];
    readonly trackingOptions?: outputs.pinpointemail.ConfigurationSetTrackingOptions;
}

export function getConfigurationSetOutput(args: GetConfigurationSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConfigurationSetResult> {
    return pulumi.output(args).apply(a => getConfigurationSet(a, opts))
}

export interface GetConfigurationSetOutputArgs {
    id: pulumi.Input<string>;
}