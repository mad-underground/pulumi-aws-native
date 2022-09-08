// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The AWS::Rekognition::StreamProcessor type is used to create an Amazon Rekognition StreamProcessor that you can use to analyze streaming videos.
 */
export function getStreamProcessor(args: GetStreamProcessorArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamProcessorResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:rekognition:getStreamProcessor", {
        "name": args.name,
    }, opts);
}

export interface GetStreamProcessorArgs {
    /**
     * Name of the stream processor. It's an identifier you assign to the stream processor. You can use it to manage the stream processor.
     */
    name: string;
}

export interface GetStreamProcessorResult {
    readonly arn?: string;
    /**
     * Current status of the stream processor.
     */
    readonly status?: string;
    /**
     * Detailed status message about the stream processor.
     */
    readonly statusMessage?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.rekognition.StreamProcessorTag[];
}

export function getStreamProcessorOutput(args: GetStreamProcessorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStreamProcessorResult> {
    return pulumi.output(args).apply(a => getStreamProcessor(a, opts))
}

export interface GetStreamProcessorOutputArgs {
    /**
     * Name of the stream processor. It's an identifier you assign to the stream processor. You can use it to manage the stream processor.
     */
    name: pulumi.Input<string>;
}
