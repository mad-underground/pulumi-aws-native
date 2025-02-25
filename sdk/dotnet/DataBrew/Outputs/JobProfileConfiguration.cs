// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Outputs
{

    [OutputType]
    public sealed class JobProfileConfiguration
    {
        public readonly ImmutableArray<Outputs.JobColumnStatisticsConfiguration> ColumnStatisticsConfigurations;
        public readonly Outputs.JobStatisticsConfiguration? DatasetStatisticsConfiguration;
        public readonly Outputs.JobEntityDetectorConfiguration? EntityDetectorConfiguration;
        public readonly ImmutableArray<Outputs.JobColumnSelector> ProfileColumns;

        [OutputConstructor]
        private JobProfileConfiguration(
            ImmutableArray<Outputs.JobColumnStatisticsConfiguration> columnStatisticsConfigurations,

            Outputs.JobStatisticsConfiguration? datasetStatisticsConfiguration,

            Outputs.JobEntityDetectorConfiguration? entityDetectorConfiguration,

            ImmutableArray<Outputs.JobColumnSelector> profileColumns)
        {
            ColumnStatisticsConfigurations = columnStatisticsConfigurations;
            DatasetStatisticsConfiguration = datasetStatisticsConfiguration;
            EntityDetectorConfiguration = entityDetectorConfiguration;
            ProfileColumns = profileColumns;
        }
    }
}
