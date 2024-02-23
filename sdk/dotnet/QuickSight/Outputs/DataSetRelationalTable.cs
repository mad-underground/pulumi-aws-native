// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;A physical table type for relational data sources.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetRelationalTable
    {
        /// <summary>
        /// &lt;p&gt;The catalog associated with a table.&lt;/p&gt;
        /// </summary>
        public readonly string? Catalog;
        /// <summary>
        /// &lt;p&gt;The Amazon Resource Name (ARN) for the data source.&lt;/p&gt;
        /// </summary>
        public readonly string DataSourceArn;
        /// <summary>
        /// &lt;p&gt;The column schema of the table.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetInputColumn> InputColumns;
        /// <summary>
        /// &lt;p&gt;The name of the relational table.&lt;/p&gt;
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// &lt;p&gt;The schema name. This name applies to certain relational database engines.&lt;/p&gt;
        /// </summary>
        public readonly string? Schema;

        [OutputConstructor]
        private DataSetRelationalTable(
            string? catalog,

            string dataSourceArn,

            ImmutableArray<Outputs.DataSetInputColumn> inputColumns,

            string name,

            string? schema)
        {
            Catalog = catalog;
            DataSourceArn = dataSourceArn;
            InputColumns = inputColumns;
            Name = name;
            Schema = schema;
        }
    }
}
