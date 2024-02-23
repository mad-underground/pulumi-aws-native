// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * AWS::RoboMaker::Fleet resource creates an AWS RoboMaker fleet. Fleets contain robots and can receive deployments.
 */
export function getFleet(args: GetFleetArgs, opts?: pulumi.InvokeOptions): Promise<GetFleetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:robomaker:getFleet", {
        "arn": args.arn,
    }, opts);
}

export interface GetFleetArgs {
    arn: string;
}

export interface GetFleetResult {
    readonly arn?: string;
    readonly tags?: {[key: string]: string};
}
/**
 * AWS::RoboMaker::Fleet resource creates an AWS RoboMaker fleet. Fleets contain robots and can receive deployments.
 */
export function getFleetOutput(args: GetFleetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFleetResult> {
    return pulumi.output(args).apply((a: any) => getFleet(a, opts))
}

export interface GetFleetOutputArgs {
    arn: pulumi.Input<string>;
}
