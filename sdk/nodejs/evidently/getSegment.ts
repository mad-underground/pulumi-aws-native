// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Evidently::Segment
 */
export function getSegment(args: GetSegmentArgs, opts?: pulumi.InvokeOptions): Promise<GetSegmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:evidently:getSegment", {
        "arn": args.arn,
    }, opts);
}

export interface GetSegmentArgs {
    arn: string;
}

export interface GetSegmentResult {
    readonly arn?: string;
    readonly description?: string;
    readonly name?: string;
    readonly pattern?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.evidently.SegmentTag[];
}

export function getSegmentOutput(args: GetSegmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSegmentResult> {
    return pulumi.output(args).apply(a => getSegment(a, opts))
}

export interface GetSegmentOutputArgs {
    arn: pulumi.Input<string>;
}