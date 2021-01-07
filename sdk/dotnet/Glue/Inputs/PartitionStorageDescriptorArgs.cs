// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Glue.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html
    /// </summary>
    public sealed class PartitionStorageDescriptorArgs : Pulumi.ResourceArgs
    {
        [Input("BucketColumns")]
        private InputList<string>? _BucketColumns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-bucketcolumns
        /// </summary>
        public InputList<string> BucketColumns
        {
            get => _BucketColumns ?? (_BucketColumns = new InputList<string>());
            set => _BucketColumns = value;
        }

        [Input("Columns")]
        private InputList<Inputs.PartitionColumnArgs>? _Columns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-columns
        /// </summary>
        public InputList<Inputs.PartitionColumnArgs> Columns
        {
            get => _Columns ?? (_Columns = new InputList<Inputs.PartitionColumnArgs>());
            set => _Columns = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-compressed
        /// </summary>
        [Input("Compressed")]
        public Input<bool>? Compressed { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-inputformat
        /// </summary>
        [Input("InputFormat")]
        public Input<string>? InputFormat { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-location
        /// </summary>
        [Input("Location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-numberofbuckets
        /// </summary>
        [Input("NumberOfBuckets")]
        public Input<int>? NumberOfBuckets { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-outputformat
        /// </summary>
        [Input("OutputFormat")]
        public Input<string>? OutputFormat { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-parameters
        /// </summary>
        [Input("Parameters")]
        public InputUnion<System.Text.Json.JsonElement, string>? Parameters { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-schemareference
        /// </summary>
        [Input("SchemaReference")]
        public Input<Inputs.PartitionSchemaReferenceArgs>? SchemaReference { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-serdeinfo
        /// </summary>
        [Input("SerdeInfo")]
        public Input<Inputs.PartitionSerdeInfoArgs>? SerdeInfo { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-skewedinfo
        /// </summary>
        [Input("SkewedInfo")]
        public Input<Inputs.PartitionSkewedInfoArgs>? SkewedInfo { get; set; }

        [Input("SortColumns")]
        private InputList<Inputs.PartitionOrderArgs>? _SortColumns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-sortcolumns
        /// </summary>
        public InputList<Inputs.PartitionOrderArgs> SortColumns
        {
            get => _SortColumns ?? (_SortColumns = new InputList<Inputs.PartitionOrderArgs>());
            set => _SortColumns = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-partition-storagedescriptor.html#cfn-glue-partition-storagedescriptor-storedassubdirectories
        /// </summary>
        [Input("StoredAsSubDirectories")]
        public Input<bool>? StoredAsSubDirectories { get; set; }

        public PartitionStorageDescriptorArgs()
        {
        }
    }
}
