// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaTailor
{
    public static class GetPlaybackConfiguration
    {
        /// <summary>
        /// Resource schema for AWS::MediaTailor::PlaybackConfiguration
        /// </summary>
        public static Task<GetPlaybackConfigurationResult> InvokeAsync(GetPlaybackConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPlaybackConfigurationResult>("aws-native:mediatailor:getPlaybackConfiguration", args ?? new GetPlaybackConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::MediaTailor::PlaybackConfiguration
        /// </summary>
        public static Output<GetPlaybackConfigurationResult> Invoke(GetPlaybackConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPlaybackConfigurationResult>("aws-native:mediatailor:getPlaybackConfiguration", args ?? new GetPlaybackConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPlaybackConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the playback configuration.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetPlaybackConfigurationArgs()
        {
        }
        public static new GetPlaybackConfigurationArgs Empty => new GetPlaybackConfigurationArgs();
    }

    public sealed class GetPlaybackConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the playback configuration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetPlaybackConfigurationInvokeArgs()
        {
        }
        public static new GetPlaybackConfigurationInvokeArgs Empty => new GetPlaybackConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetPlaybackConfigurationResult
    {
        /// <summary>
        /// The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
        /// </summary>
        public readonly string? AdDecisionServerUrl;
        /// <summary>
        /// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        /// </summary>
        public readonly Outputs.PlaybackConfigurationAvailSuppression? AvailSuppression;
        /// <summary>
        /// The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
        /// </summary>
        public readonly Outputs.PlaybackConfigurationBumper? Bumper;
        /// <summary>
        /// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
        /// </summary>
        public readonly Outputs.PlaybackConfigurationCdnConfiguration? CdnConfiguration;
        /// <summary>
        /// The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables. 
        /// </summary>
        public readonly Outputs.PlaybackConfigurationConfigurationAliases? ConfigurationAliases;
        /// <summary>
        /// The configuration for DASH content.
        /// </summary>
        public readonly Outputs.PlaybackConfigurationDashConfiguration? DashConfiguration;
        /// <summary>
        /// The configuration for HLS content.
        /// </summary>
        public readonly Outputs.PlaybackConfigurationHlsConfiguration? HlsConfiguration;
        /// <summary>
        /// The configuration for pre-roll ad insertion.
        /// </summary>
        public readonly Outputs.PlaybackConfigurationLivePreRollConfiguration? LivePreRollConfiguration;
        /// <summary>
        /// The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
        /// </summary>
        public readonly Outputs.PlaybackConfigurationManifestProcessingRules? ManifestProcessingRules;
        /// <summary>
        /// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
        /// </summary>
        public readonly int? PersonalizationThresholdSeconds;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the playback configuration.
        /// </summary>
        public readonly string? PlaybackConfigurationArn;
        /// <summary>
        /// The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
        /// </summary>
        public readonly string? PlaybackEndpointPrefix;
        /// <summary>
        /// The URL that the player uses to initialize a session that uses client-side reporting.
        /// </summary>
        public readonly string? SessionInitializationEndpointPrefix;
        /// <summary>
        /// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
        /// </summary>
        public readonly string? SlateAdUrl;
        /// <summary>
        /// The tags to assign to the playback configuration.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;
        /// <summary>
        /// The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
        /// </summary>
        public readonly string? TranscodeProfileName;
        /// <summary>
        /// The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
        /// </summary>
        public readonly string? VideoContentSourceUrl;

        [OutputConstructor]
        private GetPlaybackConfigurationResult(
            string? adDecisionServerUrl,

            Outputs.PlaybackConfigurationAvailSuppression? availSuppression,

            Outputs.PlaybackConfigurationBumper? bumper,

            Outputs.PlaybackConfigurationCdnConfiguration? cdnConfiguration,

            Outputs.PlaybackConfigurationConfigurationAliases? configurationAliases,

            Outputs.PlaybackConfigurationDashConfiguration? dashConfiguration,

            Outputs.PlaybackConfigurationHlsConfiguration? hlsConfiguration,

            Outputs.PlaybackConfigurationLivePreRollConfiguration? livePreRollConfiguration,

            Outputs.PlaybackConfigurationManifestProcessingRules? manifestProcessingRules,

            int? personalizationThresholdSeconds,

            string? playbackConfigurationArn,

            string? playbackEndpointPrefix,

            string? sessionInitializationEndpointPrefix,

            string? slateAdUrl,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags,

            string? transcodeProfileName,

            string? videoContentSourceUrl)
        {
            AdDecisionServerUrl = adDecisionServerUrl;
            AvailSuppression = availSuppression;
            Bumper = bumper;
            CdnConfiguration = cdnConfiguration;
            ConfigurationAliases = configurationAliases;
            DashConfiguration = dashConfiguration;
            HlsConfiguration = hlsConfiguration;
            LivePreRollConfiguration = livePreRollConfiguration;
            ManifestProcessingRules = manifestProcessingRules;
            PersonalizationThresholdSeconds = personalizationThresholdSeconds;
            PlaybackConfigurationArn = playbackConfigurationArn;
            PlaybackEndpointPrefix = playbackEndpointPrefix;
            SessionInitializationEndpointPrefix = sessionInitializationEndpointPrefix;
            SlateAdUrl = slateAdUrl;
            Tags = tags;
            TranscodeProfileName = transcodeProfileName;
            VideoContentSourceUrl = videoContentSourceUrl;
        }
    }
}
