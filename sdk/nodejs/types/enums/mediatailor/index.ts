// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ChannelLogType = {
    AsRun: "AS_RUN",
} as const;

export type ChannelLogType = (typeof ChannelLogType)[keyof typeof ChannelLogType];

export const ChannelPlaybackMode = {
    Loop: "LOOP",
    Linear: "LINEAR",
} as const;

export type ChannelPlaybackMode = (typeof ChannelPlaybackMode)[keyof typeof ChannelPlaybackMode];

export const ChannelTier = {
    Basic: "BASIC",
    Standard: "STANDARD",
} as const;

export type ChannelTier = (typeof ChannelTier)[keyof typeof ChannelTier];

export const LiveSourceType = {
    Dash: "DASH",
    Hls: "HLS",
} as const;

export type LiveSourceType = (typeof LiveSourceType)[keyof typeof LiveSourceType];

export const PlaybackConfigurationAvailSuppressionMode = {
    Off: "OFF",
    BehindLiveEdge: "BEHIND_LIVE_EDGE",
} as const;

/**
 * Sets the ad suppression mode. By default, ad suppression is set to OFF and all ad breaks are filled with ads or slate. When Mode is set to BEHIND_LIVE_EDGE, ad suppression is active and MediaTailor won't fill ad breaks on or behind the ad suppression Value time in the manifest lookback window.
 */
export type PlaybackConfigurationAvailSuppressionMode = (typeof PlaybackConfigurationAvailSuppressionMode)[keyof typeof PlaybackConfigurationAvailSuppressionMode];

export const PlaybackConfigurationDashConfigurationOriginManifestType = {
    SinglePeriod: "SINGLE_PERIOD",
    MultiPeriod: "MULTI_PERIOD",
} as const;

/**
 * The setting that controls whether MediaTailor handles manifests from the origin server as multi-period manifests or single-period manifests. If your origin server produces single-period manifests, set this to SINGLE_PERIOD. The default setting is MULTI_PERIOD. For multi-period manifests, omit this setting or set it to MULTI_PERIOD.
 */
export type PlaybackConfigurationDashConfigurationOriginManifestType = (typeof PlaybackConfigurationDashConfigurationOriginManifestType)[keyof typeof PlaybackConfigurationDashConfigurationOriginManifestType];

export const SourceLocationAccessType = {
    S3Sigv4: "S3_SIGV4",
    SecretsManagerAccessToken: "SECRETS_MANAGER_ACCESS_TOKEN",
} as const;

export type SourceLocationAccessType = (typeof SourceLocationAccessType)[keyof typeof SourceLocationAccessType];

export const VodSourceType = {
    Dash: "DASH",
    Hls: "HLS",
} as const;

export type VodSourceType = (typeof VodSourceType)[keyof typeof VodSourceType];
