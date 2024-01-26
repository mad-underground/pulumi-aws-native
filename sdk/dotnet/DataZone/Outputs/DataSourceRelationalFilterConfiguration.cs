// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataZone.Outputs
{

    /// <summary>
    /// The relational filter configuration for the data source.
    /// </summary>
    [OutputType]
    public sealed class DataSourceRelationalFilterConfiguration
    {
        /// <summary>
        /// The database name specified in the relational filter configuration for the data source.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The filter expressions specified in the relational filter configuration for the data source.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceFilterExpression> FilterExpressions;
        /// <summary>
        /// The schema name specified in the relational filter configuration for the data source.
        /// </summary>
        public readonly string? SchemaName;

        [OutputConstructor]
        private DataSourceRelationalFilterConfiguration(
            string databaseName,

            ImmutableArray<Outputs.DataSourceFilterExpression> filterExpressions,

            string? schemaName)
        {
            DatabaseName = databaseName;
            FilterExpressions = filterExpressions;
            SchemaName = schemaName;
        }
    }
}
