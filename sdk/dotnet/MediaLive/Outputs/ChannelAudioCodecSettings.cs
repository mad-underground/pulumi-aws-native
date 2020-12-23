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
    public sealed class ChannelAudioCodecSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-aacsettings
        /// </summary>
        public readonly Outputs.ChannelAacSettings? AacSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-ac3settings
        /// </summary>
        public readonly Outputs.ChannelAc3Settings? Ac3Settings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-eac3settings
        /// </summary>
        public readonly Outputs.ChannelEac3Settings? Eac3Settings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-mp2settings
        /// </summary>
        public readonly Outputs.ChannelMp2Settings? Mp2Settings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-audiocodecsettings.html#cfn-medialive-channel-audiocodecsettings-passthroughsettings
        /// </summary>
        public readonly Outputs.ChannelPassThroughSettings? PassThroughSettings;

        [OutputConstructor]
        private ChannelAudioCodecSettings(
            Outputs.ChannelAacSettings? AacSettings,

            Outputs.ChannelAc3Settings? Ac3Settings,

            Outputs.ChannelEac3Settings? Eac3Settings,

            Outputs.ChannelMp2Settings? Mp2Settings,

            Outputs.ChannelPassThroughSettings? PassThroughSettings)
        {
            this.AacSettings = AacSettings;
            this.Ac3Settings = Ac3Settings;
            this.Eac3Settings = Eac3Settings;
            this.Mp2Settings = Mp2Settings;
            this.PassThroughSettings = PassThroughSettings;
        }
    }
}
