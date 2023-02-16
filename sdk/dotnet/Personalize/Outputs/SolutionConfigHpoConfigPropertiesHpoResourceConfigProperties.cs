// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Outputs
{

    /// <summary>
    /// Describes the resource configuration for hyperparameter optimization (HPO).
    /// </summary>
    [OutputType]
    public sealed class SolutionConfigHpoConfigPropertiesHpoResourceConfigProperties
    {
        /// <summary>
        /// The maximum number of training jobs when you create a solution version. The maximum value for maxNumberOfTrainingJobs is 40.
        /// </summary>
        public readonly string? MaxNumberOfTrainingJobs;
        /// <summary>
        /// The maximum number of parallel training jobs when you create a solution version. The maximum value for maxParallelTrainingJobs is 10.
        /// </summary>
        public readonly string? MaxParallelTrainingJobs;

        [OutputConstructor]
        private SolutionConfigHpoConfigPropertiesHpoResourceConfigProperties(
            string? maxNumberOfTrainingJobs,

            string? maxParallelTrainingJobs)
        {
            MaxNumberOfTrainingJobs = maxNumberOfTrainingJobs;
            MaxParallelTrainingJobs = maxParallelTrainingJobs;
        }
    }
}
