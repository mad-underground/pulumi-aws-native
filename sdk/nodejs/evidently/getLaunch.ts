// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Evidently::Launch.
 */
export function getLaunch(args: GetLaunchArgs, opts?: pulumi.InvokeOptions): Promise<GetLaunchResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:evidently:getLaunch", {
        "arn": args.arn,
    }, opts);
}

export interface GetLaunchArgs {
    arn: string;
}

export interface GetLaunchResult {
    readonly arn?: string;
    readonly description?: string;
    /**
     * Start or Stop Launch Launch. Default is not started.
     */
    readonly executionStatus?: outputs.evidently.LaunchExecutionStatusObject;
    readonly groups?: outputs.evidently.LaunchGroupObject[];
    readonly metricMonitors?: outputs.evidently.LaunchMetricDefinitionObject[];
    readonly randomizationSalt?: string;
    readonly scheduledSplitsConfig?: outputs.evidently.LaunchStepConfig[];
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.evidently.LaunchTag[];
}
/**
 * Resource Type definition for AWS::Evidently::Launch.
 */
export function getLaunchOutput(args: GetLaunchOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLaunchResult> {
    return pulumi.output(args).apply((a: any) => getLaunch(a, opts))
}

export interface GetLaunchOutputArgs {
    arn: pulumi.Input<string>;
}
