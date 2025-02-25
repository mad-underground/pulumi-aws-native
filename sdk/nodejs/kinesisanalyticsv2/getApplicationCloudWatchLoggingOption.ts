// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption
 */
export function getApplicationCloudWatchLoggingOption(args: GetApplicationCloudWatchLoggingOptionArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationCloudWatchLoggingOptionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:kinesisanalyticsv2:getApplicationCloudWatchLoggingOption", {
        "id": args.id,
    }, opts);
}

export interface GetApplicationCloudWatchLoggingOptionArgs {
    id: string;
}

export interface GetApplicationCloudWatchLoggingOptionResult {
    readonly cloudWatchLoggingOption?: outputs.kinesisanalyticsv2.ApplicationCloudWatchLoggingOptionCloudWatchLoggingOption;
    readonly id?: string;
}
/**
 * Resource Type definition for AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption
 */
export function getApplicationCloudWatchLoggingOptionOutput(args: GetApplicationCloudWatchLoggingOptionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationCloudWatchLoggingOptionResult> {
    return pulumi.output(args).apply((a: any) => getApplicationCloudWatchLoggingOption(a, opts))
}

export interface GetApplicationCloudWatchLoggingOptionOutputArgs {
    id: pulumi.Input<string>;
}
