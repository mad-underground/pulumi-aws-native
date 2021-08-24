// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html
    /// </summary>
    public sealed class DataSetCustomSqlArgs : Pulumi.ResourceArgs
    {
        [Input("columns", required: true)]
        private InputList<Inputs.DataSetInputColumnArgs>? _columns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-columns
        /// </summary>
        public InputList<Inputs.DataSetInputColumnArgs> Columns
        {
            get => _columns ?? (_columns = new InputList<Inputs.DataSetInputColumnArgs>());
            set => _columns = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-datasourcearn
        /// </summary>
        [Input("dataSourceArn", required: true)]
        public Input<string> DataSourceArn { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-customsql.html#cfn-quicksight-dataset-customsql-sqlquery
        /// </summary>
        [Input("sqlQuery", required: true)]
        public Input<string> SqlQuery { get; set; } = null!;

        public DataSetCustomSqlArgs()
        {
        }
    }
}