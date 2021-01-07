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
    public sealed class ChannelH264ColorSpaceSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264colorspacesettings.html#cfn-medialive-channel-h264colorspacesettings-colorspacepassthroughsettings
        /// </summary>
        public readonly Outputs.ChannelColorSpacePassthroughSettings? ColorSpacePassthroughSettings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264colorspacesettings.html#cfn-medialive-channel-h264colorspacesettings-rec601settings
        /// </summary>
        public readonly Outputs.ChannelRec601Settings? Rec601Settings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-h264colorspacesettings.html#cfn-medialive-channel-h264colorspacesettings-rec709settings
        /// </summary>
        public readonly Outputs.ChannelRec709Settings? Rec709Settings;

        [OutputConstructor]
        private ChannelH264ColorSpaceSettings(
            Outputs.ChannelColorSpacePassthroughSettings? ColorSpacePassthroughSettings,

            Outputs.ChannelRec601Settings? Rec601Settings,

            Outputs.ChannelRec709Settings? Rec709Settings)
        {
            this.ColorSpacePassthroughSettings = ColorSpacePassthroughSettings;
            this.Rec601Settings = Rec601Settings;
            this.Rec709Settings = Rec709Settings;
        }
    }
}