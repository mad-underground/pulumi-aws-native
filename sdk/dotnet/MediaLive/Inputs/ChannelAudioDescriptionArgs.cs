// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelAudioDescriptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("audioNormalizationSettings")]
        public Input<Inputs.ChannelAudioNormalizationSettingsArgs>? AudioNormalizationSettings { get; set; }

        [Input("audioSelectorName")]
        public Input<string>? AudioSelectorName { get; set; }

        [Input("audioType")]
        public Input<string>? AudioType { get; set; }

        [Input("audioTypeControl")]
        public Input<string>? AudioTypeControl { get; set; }

        [Input("audioWatermarkingSettings")]
        public Input<Inputs.ChannelAudioWatermarkSettingsArgs>? AudioWatermarkingSettings { get; set; }

        [Input("codecSettings")]
        public Input<Inputs.ChannelAudioCodecSettingsArgs>? CodecSettings { get; set; }

        [Input("languageCode")]
        public Input<string>? LanguageCode { get; set; }

        [Input("languageCodeControl")]
        public Input<string>? LanguageCodeControl { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("remixSettings")]
        public Input<Inputs.ChannelRemixSettingsArgs>? RemixSettings { get; set; }

        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        public ChannelAudioDescriptionArgs()
        {
        }
        public static new ChannelAudioDescriptionArgs Empty => new ChannelAudioDescriptionArgs();
    }
}
