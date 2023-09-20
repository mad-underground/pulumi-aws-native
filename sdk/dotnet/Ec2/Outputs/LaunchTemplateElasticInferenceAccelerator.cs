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
    /// Specifies an elastic inference accelerator.
    /// </summary>
    [OutputType]
    public sealed class LaunchTemplateElasticInferenceAccelerator
    {
        /// <summary>
        /// The number of elastic inference accelerators to attach to the instance.
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// The type of elastic inference accelerator.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private LaunchTemplateElasticInferenceAccelerator(
            int? count,

            string? type)
        {
            Count = count;
            Type = type;
        }
    }
}