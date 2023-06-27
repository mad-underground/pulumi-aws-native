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
    /// The minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferential chips) on an instance.
    /// </summary>
    public sealed class LaunchTemplateAcceleratorCountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of accelerators.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// The minimum number of accelerators.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        public LaunchTemplateAcceleratorCountArgs()
        {
        }
        public static new LaunchTemplateAcceleratorCountArgs Empty => new LaunchTemplateAcceleratorCountArgs();
    }
}
