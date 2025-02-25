// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a streaming session machine image that can be used to launch a streaming session
 */
export function getStreamingImage(args: GetStreamingImageArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamingImageResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("aws-native:nimblestudio:getStreamingImage", {
        "streamingImageId": args.streamingImageId,
        "studioId": args.studioId,
    }, opts);
}

export interface GetStreamingImageArgs {
    streamingImageId: string;
    /**
     * <p>The studioId. </p>
     */
    studioId: string;
}

export interface GetStreamingImageResult {
    /**
     * <p>A human-readable description of the streaming image.</p>
     */
    readonly description?: string;
    readonly encryptionConfiguration?: outputs.nimblestudio.StreamingImageEncryptionConfiguration;
    /**
     * <p>The list of EULAs that must be accepted before a Streaming Session can be started using this streaming image.</p>
     */
    readonly eulaIds?: string[];
    /**
     * <p>A friendly name for a streaming image resource.</p>
     */
    readonly name?: string;
    /**
     * <p>The owner of the streaming image, either the studioId that contains the streaming image, or 'amazon' for images that are provided by Amazon Nimble Studio.</p>
     */
    readonly owner?: string;
    /**
     * <p>The platform of the streaming image, either WINDOWS or LINUX.</p>
     */
    readonly platform?: string;
    readonly streamingImageId?: string;
}
/**
 * Represents a streaming session machine image that can be used to launch a streaming session
 */
export function getStreamingImageOutput(args: GetStreamingImageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStreamingImageResult> {
    return pulumi.output(args).apply((a: any) => getStreamingImage(a, opts))
}

export interface GetStreamingImageOutputArgs {
    streamingImageId: pulumi.Input<string>;
    /**
     * <p>The studioId. </p>
     */
    studioId: pulumi.Input<string>;
}
