// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelInputLossFailoverSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("inputLossThresholdMsec")]
        public Input<int>? InputLossThresholdMsec { get; set; }

        public ChannelInputLossFailoverSettingsArgs()
        {
        }
        public static new ChannelInputLossFailoverSettingsArgs Empty => new ChannelInputLossFailoverSettingsArgs();
    }
}
