// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Inputs
{

    /// <summary>
    /// The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps.
    /// </summary>
    public sealed class LaunchTemplateBaselineEbsBandwidthMbpsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum baseline bandwidth, in Mbps.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// The minimum baseline bandwidth, in Mbps.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        public LaunchTemplateBaselineEbsBandwidthMbpsArgs()
        {
        }
        public static new LaunchTemplateBaselineEbsBandwidthMbpsArgs Empty => new LaunchTemplateBaselineEbsBandwidthMbpsArgs();
    }
}
