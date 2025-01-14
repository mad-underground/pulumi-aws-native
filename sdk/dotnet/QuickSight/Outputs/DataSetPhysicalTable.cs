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
    /// &lt;p&gt;A view of a data source that contains information about the shape of the data in the
    ///             underlying source. This is a variant type structure. For this structure to be valid,
    ///             only one of the attributes can be non-null.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class DataSetPhysicalTable
    {
        public readonly Outputs.DataSetCustomSql? CustomSql;
        public readonly Outputs.DataSetRelationalTable? RelationalTable;
        public readonly Outputs.DataSetS3Source? S3Source;

        [OutputConstructor]
        private DataSetPhysicalTable(
            Outputs.DataSetCustomSql? customSql,

            Outputs.DataSetRelationalTable? relationalTable,

            Outputs.DataSetS3Source? s3Source)
        {
            CustomSql = customSql;
            RelationalTable = relationalTable;
            S3Source = s3Source;
        }
    }
}
