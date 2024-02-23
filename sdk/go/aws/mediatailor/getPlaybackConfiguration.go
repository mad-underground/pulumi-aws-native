// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediatailor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::MediaTailor::PlaybackConfiguration
func LookupPlaybackConfiguration(ctx *pulumi.Context, args *LookupPlaybackConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupPlaybackConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPlaybackConfigurationResult
	err := ctx.Invoke("aws-native:mediatailor:getPlaybackConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPlaybackConfigurationArgs struct {
	// The identifier for the playback configuration.
	Name string `pulumi:"name"`
}

type LookupPlaybackConfigurationResult struct {
	// The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `pulumi:"adDecisionServerUrl"`
	// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	AvailSuppression *PlaybackConfigurationAvailSuppression `pulumi:"availSuppression"`
	// The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
	Bumper *PlaybackConfigurationBumper `pulumi:"bumper"`
	// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
	CdnConfiguration *PlaybackConfigurationCdnConfiguration `pulumi:"cdnConfiguration"`
	// The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables.
	ConfigurationAliases map[string]interface{} `pulumi:"configurationAliases"`
	// The configuration for DASH content.
	DashConfiguration *PlaybackConfigurationDashConfiguration `pulumi:"dashConfiguration"`
	// The configuration for HLS content.
	HlsConfiguration *PlaybackConfigurationHlsConfiguration `pulumi:"hlsConfiguration"`
	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *PlaybackConfigurationLivePreRollConfiguration `pulumi:"livePreRollConfiguration"`
	// The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *PlaybackConfigurationManifestProcessingRules `pulumi:"manifestProcessingRules"`
	// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
	PersonalizationThresholdSeconds *int `pulumi:"personalizationThresholdSeconds"`
	// The Amazon Resource Name (ARN) for the playback configuration.
	PlaybackConfigurationArn *string `pulumi:"playbackConfigurationArn"`
	// The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
	PlaybackEndpointPrefix *string `pulumi:"playbackEndpointPrefix"`
	// The URL that the player uses to initialize a session that uses client-side reporting.
	SessionInitializationEndpointPrefix *string `pulumi:"sessionInitializationEndpointPrefix"`
	// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `pulumi:"slateAdUrl"`
	// The tags to assign to the playback configuration.
	Tags []aws.Tag `pulumi:"tags"`
	// The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
	TranscodeProfileName *string `pulumi:"transcodeProfileName"`
	// The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
	VideoContentSourceUrl *string `pulumi:"videoContentSourceUrl"`
}

func LookupPlaybackConfigurationOutput(ctx *pulumi.Context, args LookupPlaybackConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupPlaybackConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPlaybackConfigurationResult, error) {
			args := v.(LookupPlaybackConfigurationArgs)
			r, err := LookupPlaybackConfiguration(ctx, &args, opts...)
			var s LookupPlaybackConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPlaybackConfigurationResultOutput)
}

type LookupPlaybackConfigurationOutputArgs struct {
	// The identifier for the playback configuration.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupPlaybackConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlaybackConfigurationArgs)(nil)).Elem()
}

type LookupPlaybackConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupPlaybackConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPlaybackConfigurationResult)(nil)).Elem()
}

func (o LookupPlaybackConfigurationResultOutput) ToLookupPlaybackConfigurationResultOutput() LookupPlaybackConfigurationResultOutput {
	return o
}

func (o LookupPlaybackConfigurationResultOutput) ToLookupPlaybackConfigurationResultOutputWithContext(ctx context.Context) LookupPlaybackConfigurationResultOutput {
	return o
}

// The URL for the ad decision server (ADS). This includes the specification of static parameters and placeholders for dynamic parameters. AWS Elemental MediaTailor substitutes player-specific and session-specific parameters as needed when calling the ADS. Alternately, for testing you can provide a static VAST URL. The maximum length is 25,000 characters.
func (o LookupPlaybackConfigurationResultOutput) AdDecisionServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *string { return v.AdDecisionServerUrl }).(pulumi.StringPtrOutput)
}

// The configuration for avail suppression, also known as ad suppression. For more information about ad suppression, see Ad Suppression (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
func (o LookupPlaybackConfigurationResultOutput) AvailSuppression() PlaybackConfigurationAvailSuppressionPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *PlaybackConfigurationAvailSuppression {
		return v.AvailSuppression
	}).(PlaybackConfigurationAvailSuppressionPtrOutput)
}

// The configuration for bumpers. Bumpers are short audio or video clips that play at the start or before the end of an ad break. To learn more about bumpers, see Bumpers (https://docs.aws.amazon.com/mediatailor/latest/ug/bumpers.html).
func (o LookupPlaybackConfigurationResultOutput) Bumper() PlaybackConfigurationBumperPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *PlaybackConfigurationBumper { return v.Bumper }).(PlaybackConfigurationBumperPtrOutput)
}

// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
func (o LookupPlaybackConfigurationResultOutput) CdnConfiguration() PlaybackConfigurationCdnConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *PlaybackConfigurationCdnConfiguration {
		return v.CdnConfiguration
	}).(PlaybackConfigurationCdnConfigurationPtrOutput)
}

// The player parameters and aliases used as dynamic variables during session initialization. For more information, see Domain Variables.
func (o LookupPlaybackConfigurationResultOutput) ConfigurationAliases() pulumi.MapOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) map[string]interface{} { return v.ConfigurationAliases }).(pulumi.MapOutput)
}

// The configuration for DASH content.
func (o LookupPlaybackConfigurationResultOutput) DashConfiguration() PlaybackConfigurationDashConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *PlaybackConfigurationDashConfiguration {
		return v.DashConfiguration
	}).(PlaybackConfigurationDashConfigurationPtrOutput)
}

// The configuration for HLS content.
func (o LookupPlaybackConfigurationResultOutput) HlsConfiguration() PlaybackConfigurationHlsConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *PlaybackConfigurationHlsConfiguration {
		return v.HlsConfiguration
	}).(PlaybackConfigurationHlsConfigurationPtrOutput)
}

// The configuration for pre-roll ad insertion.
func (o LookupPlaybackConfigurationResultOutput) LivePreRollConfiguration() PlaybackConfigurationLivePreRollConfigurationPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *PlaybackConfigurationLivePreRollConfiguration {
		return v.LivePreRollConfiguration
	}).(PlaybackConfigurationLivePreRollConfigurationPtrOutput)
}

// The configuration for manifest processing rules. Manifest processing rules enable customization of the personalized manifests created by MediaTailor.
func (o LookupPlaybackConfigurationResultOutput) ManifestProcessingRules() PlaybackConfigurationManifestProcessingRulesPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *PlaybackConfigurationManifestProcessingRules {
		return v.ManifestProcessingRules
	}).(PlaybackConfigurationManifestProcessingRulesPtrOutput)
}

// Defines the maximum duration of underfilled ad time (in seconds) allowed in an ad break. If the duration of underfilled ad time exceeds the personalization threshold, then the personalization of the ad break is abandoned and the underlying content is shown. This feature applies to ad replacement in live and VOD streams, rather than ad insertion, because it relies on an underlying content stream. For more information about ad break behavior, including ad replacement and insertion, see Ad Behavior in AWS Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/ad-behavior.html).
func (o LookupPlaybackConfigurationResultOutput) PersonalizationThresholdSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *int { return v.PersonalizationThresholdSeconds }).(pulumi.IntPtrOutput)
}

// The Amazon Resource Name (ARN) for the playback configuration.
func (o LookupPlaybackConfigurationResultOutput) PlaybackConfigurationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *string { return v.PlaybackConfigurationArn }).(pulumi.StringPtrOutput)
}

// The URL that the player accesses to get a manifest from MediaTailor. This session will use server-side reporting.
func (o LookupPlaybackConfigurationResultOutput) PlaybackEndpointPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *string { return v.PlaybackEndpointPrefix }).(pulumi.StringPtrOutput)
}

// The URL that the player uses to initialize a session that uses client-side reporting.
func (o LookupPlaybackConfigurationResultOutput) SessionInitializationEndpointPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *string { return v.SessionInitializationEndpointPrefix }).(pulumi.StringPtrOutput)
}

// The URL for a high-quality video asset to transcode and use to fill in time that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps in media content. Configuring the slate is optional for non-VPAID configurations. For VPAID, the slate is required because MediaTailor provides it in the slots that are designated for dynamic ad content. The slate must be a high-quality asset that contains both audio and video.
func (o LookupPlaybackConfigurationResultOutput) SlateAdUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *string { return v.SlateAdUrl }).(pulumi.StringPtrOutput)
}

// The tags to assign to the playback configuration.
func (o LookupPlaybackConfigurationResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

// The name that is used to associate this playback configuration with a custom transcode profile. This overrides the dynamic transcoding defaults of MediaTailor. Use this only if you have already set up custom profiles with the help of AWS Support.
func (o LookupPlaybackConfigurationResultOutput) TranscodeProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *string { return v.TranscodeProfileName }).(pulumi.StringPtrOutput)
}

// The URL prefix for the parent manifest for the stream, minus the asset ID. The maximum length is 512 characters.
func (o LookupPlaybackConfigurationResultOutput) VideoContentSourceUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPlaybackConfigurationResult) *string { return v.VideoContentSourceUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPlaybackConfigurationResultOutput{})
}
