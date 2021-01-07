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
    public sealed class ChannelAacSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-bitrate
        /// </summary>
        public readonly double? Bitrate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-codingmode
        /// </summary>
        public readonly string? CodingMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-inputtype
        /// </summary>
        public readonly string? InputType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-profile
        /// </summary>
        public readonly string? Profile;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-ratecontrolmode
        /// </summary>
        public readonly string? RateControlMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-rawformat
        /// </summary>
        public readonly string? RawFormat;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-samplerate
        /// </summary>
        public readonly double? SampleRate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-spec
        /// </summary>
        public readonly string? Spec;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-aacsettings.html#cfn-medialive-channel-aacsettings-vbrquality
        /// </summary>
        public readonly string? VbrQuality;

        [OutputConstructor]
        private ChannelAacSettings(
            double? Bitrate,

            string? CodingMode,

            string? InputType,

            string? Profile,

            string? RateControlMode,

            string? RawFormat,

            double? SampleRate,

            string? Spec,

            string? VbrQuality)
        {
            this.Bitrate = Bitrate;
            this.CodingMode = CodingMode;
            this.InputType = InputType;
            this.Profile = Profile;
            this.RateControlMode = RateControlMode;
            this.RawFormat = RawFormat;
            this.SampleRate = SampleRate;
            this.Spec = Spec;
            this.VbrQuality = VbrQuality;
        }
    }
}
