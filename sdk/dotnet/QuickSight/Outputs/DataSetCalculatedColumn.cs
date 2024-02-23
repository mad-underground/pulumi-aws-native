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
    /// &lt;p&gt;A calculated column for a dataset.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetCalculatedColumn
    {
        /// <summary>
        /// &lt;p&gt;A unique ID to identify a calculated column. During a dataset update, if the column ID
        ///             of a calculated column matches that of an existing calculated column, Amazon QuickSight
        ///             preserves the existing calculated column.&lt;/p&gt;
        /// </summary>
        public readonly string ColumnId;
        /// <summary>
        /// &lt;p&gt;Column name.&lt;/p&gt;
        /// </summary>
        public readonly string ColumnName;
        /// <summary>
        /// &lt;p&gt;An expression that defines the calculated column.&lt;/p&gt;
        /// </summary>
        public readonly string Expression;

        [OutputConstructor]
        private DataSetCalculatedColumn(
            string columnId,

            string columnName,

            string expression)
        {
            ColumnId = columnId;
            ColumnName = columnName;
            Expression = expression;
        }
    }
}
