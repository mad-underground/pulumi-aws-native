// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticsconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class JobStatisticsConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticsconfiguration.html#cfn-databrew-job-statisticsconfiguration-includedstatistics
        /// </summary>
        public readonly ImmutableArray<string> IncludedStatistics;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticsconfiguration.html#cfn-databrew-job-statisticsconfiguration-overrides
        /// </summary>
        public readonly ImmutableArray<Outputs.JobStatisticOverride> Overrides;

        [OutputConstructor]
        private JobStatisticsConfiguration(
            ImmutableArray<string> includedStatistics,

            ImmutableArray<Outputs.JobStatisticOverride> overrides)
        {
            IncludedStatistics = includedStatistics;
            Overrides = overrides;
        }
    }
}