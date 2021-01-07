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
    public sealed class ChannelInputLossBehavior
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossbehavior.html#cfn-medialive-channel-inputlossbehavior-blackframemsec
        /// </summary>
        public readonly int? BlackFrameMsec;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossbehavior.html#cfn-medialive-channel-inputlossbehavior-inputlossimagecolor
        /// </summary>
        public readonly string? InputLossImageColor;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossbehavior.html#cfn-medialive-channel-inputlossbehavior-inputlossimageslate
        /// </summary>
        public readonly Outputs.ChannelInputLocation? InputLossImageSlate;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossbehavior.html#cfn-medialive-channel-inputlossbehavior-inputlossimagetype
        /// </summary>
        public readonly string? InputLossImageType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputlossbehavior.html#cfn-medialive-channel-inputlossbehavior-repeatframemsec
        /// </summary>
        public readonly int? RepeatFrameMsec;

        [OutputConstructor]
        private ChannelInputLossBehavior(
            int? BlackFrameMsec,

            string? InputLossImageColor,

            Outputs.ChannelInputLocation? InputLossImageSlate,

            string? InputLossImageType,

            int? RepeatFrameMsec)
        {
            this.BlackFrameMsec = BlackFrameMsec;
            this.InputLossImageColor = InputLossImageColor;
            this.InputLossImageSlate = InputLossImageSlate;
            this.InputLossImageType = InputLossImageType;
            this.RepeatFrameMsec = RepeatFrameMsec;
        }
    }
}