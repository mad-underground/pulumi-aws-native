// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;A view of a data source that contains information about the shape of the data in the
    ///             underlying source. This is a variant type structure. For this structure to be valid,
    ///             only one of the attributes can be non-null.&lt;/p&gt;
    /// </summary>
    public sealed class DataSetPhysicalTableArgs : global::Pulumi.ResourceArgs
    {
        [Input("customSql")]
        public Input<Inputs.DataSetCustomSqlArgs>? CustomSql { get; set; }

        [Input("relationalTable")]
        public Input<Inputs.DataSetRelationalTableArgs>? RelationalTable { get; set; }

        [Input("s3Source")]
        public Input<Inputs.DataSetS3SourceArgs>? S3Source { get; set; }

        public DataSetPhysicalTableArgs()
        {
        }
        public static new DataSetPhysicalTableArgs Empty => new DataSetPhysicalTableArgs();
    }
}
