// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra.Inputs
{

    public sealed class TableColumnArgs : global::Pulumi.ResourceArgs
    {
        [Input("columnName", required: true)]
        public Input<string> ColumnName { get; set; } = null!;

        [Input("columnType", required: true)]
        public Input<string> ColumnType { get; set; } = null!;

        public TableColumnArgs()
        {
        }
        public static new TableColumnArgs Empty => new TableColumnArgs();
    }
}
