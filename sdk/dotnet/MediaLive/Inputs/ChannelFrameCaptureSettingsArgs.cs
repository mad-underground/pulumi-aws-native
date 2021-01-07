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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturesettings.html
    /// </summary>
    public sealed class ChannelFrameCaptureSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturesettings.html#cfn-medialive-channel-framecapturesettings-captureinterval
        /// </summary>
        [Input("CaptureInterval")]
        public Input<int>? CaptureInterval { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-framecapturesettings.html#cfn-medialive-channel-framecapturesettings-captureintervalunits
        /// </summary>
        [Input("CaptureIntervalUnits")]
        public Input<string>? CaptureIntervalUnits { get; set; }

        public ChannelFrameCaptureSettingsArgs()
        {
        }
    }
}
