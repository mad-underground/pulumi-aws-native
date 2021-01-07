// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.MediaLive.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html
    /// </summary>
    public sealed class ChannelHlsSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-audioonlyhlssettings
        /// </summary>
        [Input("AudioOnlyHlsSettings")]
        public Input<Inputs.ChannelAudioOnlyHlsSettingsArgs>? AudioOnlyHlsSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-fmp4hlssettings
        /// </summary>
        [Input("Fmp4HlsSettings")]
        public Input<Inputs.ChannelFmp4HlsSettingsArgs>? Fmp4HlsSettings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-hlssettings.html#cfn-medialive-channel-hlssettings-standardhlssettings
        /// </summary>
        [Input("StandardHlsSettings")]
        public Input<Inputs.ChannelStandardHlsSettingsArgs>? StandardHlsSettings { get; set; }

        public ChannelHlsSettingsArgs()
        {
        }
    }
}
