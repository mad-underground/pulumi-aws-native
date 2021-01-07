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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html
    /// </summary>
    public sealed class ChannelEncoderSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("AudioDescriptions")]
        private InputList<Inputs.ChannelAudioDescriptionArgs>? _AudioDescriptions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-audiodescriptions
        /// </summary>
        public InputList<Inputs.ChannelAudioDescriptionArgs> AudioDescriptions
        {
            get => _AudioDescriptions ?? (_AudioDescriptions = new InputList<Inputs.ChannelAudioDescriptionArgs>());
            set => _AudioDescriptions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availblanking
        /// </summary>
        [Input("AvailBlanking")]
        public Input<Inputs.ChannelAvailBlankingArgs>? AvailBlanking { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-availconfiguration
        /// </summary>
        [Input("AvailConfiguration")]
        public Input<Inputs.ChannelAvailConfigurationArgs>? AvailConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-blackoutslate
        /// </summary>
        [Input("BlackoutSlate")]
        public Input<Inputs.ChannelBlackoutSlateArgs>? BlackoutSlate { get; set; }

        [Input("CaptionDescriptions")]
        private InputList<Inputs.ChannelCaptionDescriptionArgs>? _CaptionDescriptions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-captiondescriptions
        /// </summary>
        public InputList<Inputs.ChannelCaptionDescriptionArgs> CaptionDescriptions
        {
            get => _CaptionDescriptions ?? (_CaptionDescriptions = new InputList<Inputs.ChannelCaptionDescriptionArgs>());
            set => _CaptionDescriptions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-featureactivations
        /// </summary>
        [Input("FeatureActivations")]
        public Input<Inputs.ChannelFeatureActivationsArgs>? FeatureActivations { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-globalconfiguration
        /// </summary>
        [Input("GlobalConfiguration")]
        public Input<Inputs.ChannelGlobalConfigurationArgs>? GlobalConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-nielsenconfiguration
        /// </summary>
        [Input("NielsenConfiguration")]
        public Input<Inputs.ChannelNielsenConfigurationArgs>? NielsenConfiguration { get; set; }

        [Input("OutputGroups")]
        private InputList<Inputs.ChannelOutputGroupArgs>? _OutputGroups;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-outputgroups
        /// </summary>
        public InputList<Inputs.ChannelOutputGroupArgs> OutputGroups
        {
            get => _OutputGroups ?? (_OutputGroups = new InputList<Inputs.ChannelOutputGroupArgs>());
            set => _OutputGroups = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-timecodeconfig
        /// </summary>
        [Input("TimecodeConfig")]
        public Input<Inputs.ChannelTimecodeConfigArgs>? TimecodeConfig { get; set; }

        [Input("VideoDescriptions")]
        private InputList<Inputs.ChannelVideoDescriptionArgs>? _VideoDescriptions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-encodersettings.html#cfn-medialive-channel-encodersettings-videodescriptions
        /// </summary>
        public InputList<Inputs.ChannelVideoDescriptionArgs> VideoDescriptions
        {
            get => _VideoDescriptions ?? (_VideoDescriptions = new InputList<Inputs.ChannelVideoDescriptionArgs>());
            set => _VideoDescriptions = value;
        }

        public ChannelEncoderSettingsArgs()
        {
        }
    }
}
