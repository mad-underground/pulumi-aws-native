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
    /// The minimum and maximum amount of memory, in MiB.
    /// </summary>
    public sealed class LaunchTemplateMemoryMiBArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum amount of memory, in MiB. To specify no maximum limit, omit this parameter.
        /// </summary>
        [Input("max")]
        public Input<int>? Max { get; set; }

        /// <summary>
        /// The minimum amount of memory, in MiB. To specify no minimum limit, specify ``0``.
        /// </summary>
        [Input("min")]
        public Input<int>? Min { get; set; }

        public LaunchTemplateMemoryMiBArgs()
        {
        }
        public static new LaunchTemplateMemoryMiBArgs Empty => new LaunchTemplateMemoryMiBArgs();
    }
}
