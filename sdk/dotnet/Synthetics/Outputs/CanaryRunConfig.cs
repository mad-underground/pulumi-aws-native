// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html
    /// </summary>
    [OutputType]
    public sealed class CanaryRunConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-activetracing
        /// </summary>
        public readonly bool? ActiveTracing;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-environmentvariables
        /// </summary>
        public readonly ImmutableDictionary<string, string>? EnvironmentVariables;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-memoryinmb
        /// </summary>
        public readonly int? MemoryInMB;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-timeoutinseconds
        /// </summary>
        public readonly int TimeoutInSeconds;

        [OutputConstructor]
        private CanaryRunConfig(
            bool? activeTracing,

            ImmutableDictionary<string, string>? environmentVariables,

            int? memoryInMB,

            int timeoutInSeconds)
        {
            ActiveTracing = activeTracing;
            EnvironmentVariables = environmentVariables;
            MemoryInMB = memoryInMB;
            TimeoutInSeconds = timeoutInSeconds;
        }
    }
}
