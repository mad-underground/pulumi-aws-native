// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    /// <summary>
    /// The minimum and maximum amount of total accelerator memory, in MiB.
    /// </summary>
    [OutputType]
    public sealed class LaunchTemplateAcceleratorTotalMemoryMiB
    {
        /// <summary>
        /// The maximum amount of accelerator memory, in MiB. To specify no maximum limit, omit this parameter.
        /// </summary>
        public readonly int? Max;
        /// <summary>
        /// The minimum amount of accelerator memory, in MiB. To specify no minimum limit, omit this parameter.
        /// </summary>
        public readonly int? Min;

        [OutputConstructor]
        private LaunchTemplateAcceleratorTotalMemoryMiB(
            int? max,

            int? min)
        {
            Max = max;
            Min = min;
        }
    }
}
