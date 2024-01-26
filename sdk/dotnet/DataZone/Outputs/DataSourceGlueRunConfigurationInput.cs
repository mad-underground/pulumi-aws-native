// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    [OutputType]
    public sealed class DataSourceGlueRunConfigurationInput
    {
        /// <summary>
        /// The data access role included in the configuration details of the AWS Glue data source.
        /// </summary>
        public readonly string? DataAccessRole;
        /// <summary>
        /// The relational filter configurations included in the configuration details of the AWS Glue data source.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceRelationalFilterConfiguration> RelationalFilterConfigurations;

        [OutputConstructor]
        private DataSourceGlueRunConfigurationInput(
            string? dataAccessRole,

            ImmutableArray<Outputs.DataSourceRelationalFilterConfiguration> relationalFilterConfigurations)
        {
            DataAccessRole = dataAccessRole;
            RelationalFilterConfigurations = relationalFilterConfigurations;
        }
    }
}
