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
    public sealed class ChannelScte35TimeSignalApos
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte35timesignalapos.html#cfn-medialive-channel-scte35timesignalapos-adavailoffset
        /// </summary>
        public readonly int? AdAvailOffset;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte35timesignalapos.html#cfn-medialive-channel-scte35timesignalapos-noregionalblackoutflag
        /// </summary>
        public readonly string? NoRegionalBlackoutFlag;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-scte35timesignalapos.html#cfn-medialive-channel-scte35timesignalapos-webdeliveryallowedflag
        /// </summary>
        public readonly string? WebDeliveryAllowedFlag;

        [OutputConstructor]
        private ChannelScte35TimeSignalApos(
            int? AdAvailOffset,

            string? NoRegionalBlackoutFlag,

            string? WebDeliveryAllowedFlag)
        {
            this.AdAvailOffset = AdAvailOffset;
            this.NoRegionalBlackoutFlag = NoRegionalBlackoutFlag;
            this.WebDeliveryAllowedFlag = WebDeliveryAllowedFlag;
        }
    }
}