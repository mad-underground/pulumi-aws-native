// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class InstanceVolumeArgs : global::Pulumi.ResourceArgs
    {
        [Input("device", required: true)]
        public Input<string> Device { get; set; } = null!;

        [Input("volumeId", required: true)]
        public Input<string> VolumeId { get; set; } = null!;

        public InstanceVolumeArgs()
        {
        }
        public static new InstanceVolumeArgs Empty => new InstanceVolumeArgs();
    }
}