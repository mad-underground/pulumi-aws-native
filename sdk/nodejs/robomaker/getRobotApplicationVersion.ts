// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * AWS::RoboMaker::RobotApplicationVersion resource creates an AWS RoboMaker RobotApplicationVersion. This helps you control which code your robot uses.
 */
export function getRobotApplicationVersion(args: GetRobotApplicationVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetRobotApplicationVersionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:robomaker:getRobotApplicationVersion", {
        "arn": args.arn,
    }, opts);
}

export interface GetRobotApplicationVersionArgs {
    arn: string;
}

export interface GetRobotApplicationVersionResult {
    readonly applicationVersion?: string;
    readonly arn?: string;
}
/**
 * AWS::RoboMaker::RobotApplicationVersion resource creates an AWS RoboMaker RobotApplicationVersion. This helps you control which code your robot uses.
 */
export function getRobotApplicationVersionOutput(args: GetRobotApplicationVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRobotApplicationVersionResult> {
    return pulumi.output(args).apply((a: any) => getRobotApplicationVersion(a, opts))
}

export interface GetRobotApplicationVersionOutputArgs {
    arn: pulumi.Input<string>;
}
