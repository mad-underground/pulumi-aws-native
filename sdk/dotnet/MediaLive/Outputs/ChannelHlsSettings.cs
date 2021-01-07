// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelHlsSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-audioonlyhlssettings
        /// </summary>
        public readonly Outputs.ChannelAudioOnlyHlsSettings? AudioOnlyHlsSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-fmp4hlssettings
        /// </summary>
        public readonly Outputs.ChannelFmp4HlsSettings? Fmp4HlsSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-standardhlssettings
        /// </summary>
        public readonly Outputs.ChannelStandardHlsSettings? StandardHlsSettings;

        [OutputConstructor]
        private ChannelHlsSettings(
            Outputs.ChannelAudioOnlyHlsSettings? AudioOnlyHlsSettings,

            Outputs.ChannelFmp4HlsSettings? Fmp4HlsSettings,

            Outputs.ChannelStandardHlsSettings? StandardHlsSettings)
        {
            this.AudioOnlyHlsSettings = AudioOnlyHlsSettings;
            this.Fmp4HlsSettings = Fmp4HlsSettings;
            this.StandardHlsSettings = StandardHlsSettings;
        }
    }
}