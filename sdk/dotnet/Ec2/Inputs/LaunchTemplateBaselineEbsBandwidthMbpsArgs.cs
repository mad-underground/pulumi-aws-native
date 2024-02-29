// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    /// <summary>
    /// The minimum and maximum baseline bandwidth to Amazon EBS, in Mbps. For more information, see [Amazon EBS–optimized instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-optimized.html) in the *Amazon EC2 User Guide*.
    /// </summary>
    public sealed class LaunchTemplateBaselineEbsBandwidthMbpsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum baseline bandwidth, in Mbps. To specify no maximum limit, omit this parameter.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// The minimum baseline bandwidth, in Mbps. To specify no minimum limit, omit this parameter.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        public LaunchTemplateBaselineEbsBandwidthMbpsArgs()
        {
        }
        public static new LaunchTemplateBaselineEbsBandwidthMbpsArgs Empty => new LaunchTemplateBaselineEbsBandwidthMbpsArgs();
    }
}
