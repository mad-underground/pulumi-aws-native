// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// specifies the CPU options for an instance.
    /// </summary>
    [OutputType]
    public sealed class LaunchTemplateCpuOptions
    {
        /// <summary>
        /// Indicates whether to enable the instance for AMD SEV-SNP. AMD SEV-SNP is supported with M6a, R6a, and C6a instance types only.
        /// </summary>
        public readonly Pulumi.AwsNative.EC2.LaunchTemplateCpuOptionsAmdSevSnp? AmdSevSnp;
        /// <summary>
        /// The number of CPU cores for the instance.
        /// </summary>
        public readonly int? CoreCount;
        /// <summary>
        /// The number of threads per CPU core. To disable multithreading for the instance, specify a value of 1. Otherwise, specify the default value of 2.
        /// </summary>
        public readonly int? ThreadsPerCore;

        [OutputConstructor]
        private LaunchTemplateCpuOptions(
            Pulumi.AwsNative.EC2.LaunchTemplateCpuOptionsAmdSevSnp? amdSevSnp,

            int? coreCount,

            int? threadsPerCore)
        {
            AmdSevSnp = amdSevSnp;
            CoreCount = coreCount;
            ThreadsPerCore = threadsPerCore;
        }
    }
}
