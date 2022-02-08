// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::RoboMaker::RobotApplication
 */
export function getRobotApplication(args: GetRobotApplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetRobotApplicationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:robomaker:getRobotApplication", {
        "id": args.id,
    }, opts);
}

export interface GetRobotApplicationArgs {
    id: string;
}

export interface GetRobotApplicationResult {
    readonly arn?: string;
    readonly currentRevisionId?: string;
    readonly id?: string;
    readonly sources?: outputs.robomaker.RobotApplicationSourceConfig[];
    readonly tags?: any;
}

export function getRobotApplicationOutput(args: GetRobotApplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRobotApplicationResult> {
    return pulumi.output(args).apply(a => getRobotApplication(a, opts))
}

export interface GetRobotApplicationOutputArgs {
    id: pulumi.Input<string>;
}
