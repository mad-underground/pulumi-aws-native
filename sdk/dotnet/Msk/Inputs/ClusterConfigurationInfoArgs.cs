// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterConfigurationInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("revision", required: true)]
        public Input<int> Revision { get; set; } = null!;

        public ClusterConfigurationInfoArgs()
        {
        }
        public static new ClusterConfigurationInfoArgs Empty => new ClusterConfigurationInfoArgs();
    }
}
