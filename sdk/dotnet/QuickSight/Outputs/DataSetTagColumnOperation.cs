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
    /// &lt;p&gt;A transform operation that tags a column with additional information.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetTagColumnOperation
    {
        /// <summary>
        /// &lt;p&gt;The column that this operation acts on.&lt;/p&gt;
        /// </summary>
        public readonly string ColumnName;
        /// <summary>
        /// &lt;p&gt;The dataset column tag, currently only used for geospatial type tagging. .&lt;/p&gt;
        ///         &lt;note&gt;
        ///             &lt;p&gt;This is not tags for the AWS tagging feature. .&lt;/p&gt;
        ///         &lt;/note&gt;
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSetColumnTag> Tags;

        [OutputConstructor]
        private DataSetTagColumnOperation(
            string columnName,

            ImmutableArray<Outputs.DataSetColumnTag> tags)
        {
            ColumnName = columnName;
            Tags = tags;
        }
    }
}
