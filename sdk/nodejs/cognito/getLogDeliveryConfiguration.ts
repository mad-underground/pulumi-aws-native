// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cognito::LogDeliveryConfiguration
 */
export function getLogDeliveryConfiguration(args: GetLogDeliveryConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetLogDeliveryConfigurationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:cognito:getLogDeliveryConfiguration", {
        "id": args.id,
    }, opts);
}

export interface GetLogDeliveryConfigurationArgs {
    id: string;
}

export interface GetLogDeliveryConfigurationResult {
    readonly id?: string;
    readonly logConfigurations?: outputs.cognito.LogDeliveryConfigurationLogConfiguration[];
}
/**
 * Resource Type definition for AWS::Cognito::LogDeliveryConfiguration
 */
export function getLogDeliveryConfigurationOutput(args: GetLogDeliveryConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLogDeliveryConfigurationResult> {
    return pulumi.output(args).apply((a: any) => getLogDeliveryConfiguration(a, opts))
}

export interface GetLogDeliveryConfigurationOutputArgs {
    id: pulumi.Input<string>;
}
