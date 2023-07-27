// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.FSx.Inputs
{

    public sealed class VolumeSnaplockRetentionPeriodArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultRetention", required: true)]
        public Input<Inputs.VolumeRetentionPeriodArgs> DefaultRetention { get; set; } = null!;

        [Input("maximumRetention", required: true)]
        public Input<Inputs.VolumeRetentionPeriodArgs> MaximumRetention { get; set; } = null!;

        [Input("minimumRetention", required: true)]
        public Input<Inputs.VolumeRetentionPeriodArgs> MinimumRetention { get; set; } = null!;

        public VolumeSnaplockRetentionPeriodArgs()
        {
        }
        public static new VolumeSnaplockRetentionPeriodArgs Empty => new VolumeSnaplockRetentionPeriodArgs();
    }
}