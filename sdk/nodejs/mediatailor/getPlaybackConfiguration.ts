// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaTailor::PlaybackConfiguration
 */
export function getPlaybackConfiguration(args: GetPlaybackConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetPlaybackConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:mediatailor:getPlaybackConfiguration", {
        "name": args.name,
    }, opts);
}

export interface GetPlaybackConfigurationArgs {
    /**
     * The identifier for the playback configuration.
     */
    name: string;
}

export interface GetPlaybackConfigurationResult {
    /**
     * The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
     */
    readonly adDecisionServerUrl?: string;
    /**
     * The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
     */
    readonly availSuppression?: outputs.mediatailor.PlaybackConfigurationAvailSuppression;
    /**
     * The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
     */
    readonly bumper?: outputs.mediatailor.PlaybackConfigurationBumper;
    /**
     * The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
     */
    readonly cdnConfiguration?: outputs.mediatailor.PlaybackConfigurationCdnConfiguration;
    /**
     * The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
     */
    readonly configurationAliases?: outputs.mediatailor.PlaybackConfigurationConfigurationAliases;
    /**
     * The configuration for DASH content.
     */
    readonly dashConfiguration?: outputs.mediatailor.PlaybackConfigurationDashConfiguration;
    /**
     * The configuration for HLS content.
     */
    readonly hlsConfiguration?: outputs.mediatailor.PlaybackConfigurationHlsConfiguration;
    /**
     * The configuration for pre-roll ad insertion.
     */
    readonly livePreRollConfiguration?: outputs.mediatailor.PlaybackConfigurationLivePreRollConfiguration;
    /**
     * The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
     */
    readonly manifestProcessingRules?: outputs.mediatailor.PlaybackConfigurationManifestProcessingRules;
    /**
     * Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
     */
    readonly personalizationThresholdSeconds?: number;
    /**
     * The Amazon Resource Name (ARN) for the playback configuration.
     */
    readonly playbackConfigurationArn?: string;
    /**
     * The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
     */
    readonly playbackEndpointPrefix?: string;
    /**
     * The URL that the player uses to initialize a session that uses client-side reporting.
     */
    readonly sessionInitializationEndpointPrefix?: string;
    /**
     * The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
     */
    readonly slateAdUrl?: string;
    /**
     * The tags to assign to the playback configuration.
     */
    readonly tags?: outputs.mediatailor.PlaybackConfigurationTag[];
    /**
     * The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
     */
    readonly transcodeProfileName?: string;
    /**
     * The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
     */
    readonly videoContentSourceUrl?: string;
}

export function getPlaybackConfigurationOutput(args: GetPlaybackConfigurationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPlaybackConfigurationResult> {
    return pulumi.output(args).apply(a => getPlaybackConfiguration(a, opts))
}

export interface GetPlaybackConfigurationOutputArgs {
    /**
     * The identifier for the playback configuration.
     */
    name: pulumi.Input<string>;
}
