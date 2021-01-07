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
    public sealed class ChannelInputChannelLevel
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputchannellevel.html#cfn-medialive-channel-inputchannellevel-gain
        /// </summary>
        public readonly int? Gain;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-inputchannellevel.html#cfn-medialive-channel-inputchannellevel-inputchannel
        /// </summary>
        public readonly int? InputChannel;

        [OutputConstructor]
        private ChannelInputChannelLevel(
            int? Gain,

            int? InputChannel)
        {
            this.Gain = Gain;
            this.InputChannel = InputChannel;
        }
    }
}
