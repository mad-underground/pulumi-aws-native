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
    /// The minimum and maximum amount of total accelerator memory, in MiB.
    /// </summary>
    public sealed class LaunchTemplateAcceleratorTotalMemoryMiBArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum amount of accelerator memory, in MiB. To specify no maximum limit, omit this parameter.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// The minimum amount of accelerator memory, in MiB. To specify no minimum limit, omit this parameter.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        public LaunchTemplateAcceleratorTotalMemoryMiBArgs()
        {
        }
        public static new LaunchTemplateAcceleratorTotalMemoryMiBArgs Empty => new LaunchTemplateAcceleratorTotalMemoryMiBArgs();
    }
}
