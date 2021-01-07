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
    public sealed class ChannelRtmpGroupSettings
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-authenticationscheme
        /// </summary>
        public readonly string? AuthenticationScheme;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-cachefullbehavior
        /// </summary>
        public readonly string? CacheFullBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-cachelength
        /// </summary>
        public readonly int? CacheLength;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-captiondata
        /// </summary>
        public readonly string? CaptionData;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-inputlossaction
        /// </summary>
        public readonly string? InputLossAction;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-rtmpgroupsettings.html#cfn-medialive-channel-rtmpgroupsettings-restartdelay
        /// </summary>
        public readonly int? RestartDelay;

        [OutputConstructor]
        private ChannelRtmpGroupSettings(
            string? AuthenticationScheme,

            string? CacheFullBehavior,

            int? CacheLength,

            string? CaptionData,

            string? InputLossAction,

            int? RestartDelay)
        {
            this.AuthenticationScheme = AuthenticationScheme;
            this.CacheFullBehavior = CacheFullBehavior;
            this.CacheLength = CacheLength;
            this.CaptionData = CaptionData;
            this.InputLossAction = InputLossAction;
            this.RestartDelay = RestartDelay;
        }
    }
}
