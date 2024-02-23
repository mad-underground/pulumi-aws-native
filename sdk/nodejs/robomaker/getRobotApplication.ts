// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This schema is for testing purpose only.
 */
export function getRobotApplication(args: GetRobotApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetRobotApplicationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:robomaker:getRobotApplication", {
        "arn": args.arn,
    }, opts);
}

export interface GetRobotApplicationArgs {
    arn: string;
}

export interface GetRobotApplicationResult {
    readonly arn?: string;
    /**
     * The revision ID of robot application.
     */
    readonly currentRevisionId?: string;
    /**
     * The URI of the Docker image for the robot application.
     */
    readonly environment?: string;
    readonly robotSoftwareSuite?: outputs.robomaker.RobotApplicationRobotSoftwareSuite;
    readonly tags?: {[key: string]: string};
}
/**
 * This schema is for testing purpose only.
 */
export function getRobotApplicationOutput(args: GetRobotApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRobotApplicationResult> {
    return pulumi.output(args).apply((a: any) => getRobotApplication(a, opts))
}

export interface GetRobotApplicationOutputArgs {
    arn: pulumi.Input<string>;
}
