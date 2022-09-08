// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::Segment
 */
export function getSegment(args: GetSegmentArgs, opts?: pulumi.InvokeOptions): Promise<GetSegmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:pinpoint:getSegment", {
        "segmentId": args.segmentId,
    }, opts);
}

export interface GetSegmentArgs {
    segmentId: string;
}

export interface GetSegmentResult {
    readonly arn?: string;
    readonly dimensions?: outputs.pinpoint.SegmentDimensions;
    readonly name?: string;
    readonly segmentGroups?: outputs.pinpoint.SegmentGroups;
    readonly segmentId?: string;
    readonly tags?: any;
}

export function getSegmentOutput(args: GetSegmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSegmentResult> {
    return pulumi.output(args).apply(a => getSegment(a, opts))
}

export interface GetSegmentOutputArgs {
    segmentId: pulumi.Input<string>;
}
