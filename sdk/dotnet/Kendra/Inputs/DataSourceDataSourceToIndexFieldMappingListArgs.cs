// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourcetoindexfieldmappinglist.html
    /// </summary>
    public sealed class DataSourceDataSourceToIndexFieldMappingListArgs : Pulumi.ResourceArgs
    {
        [Input("DataSourceToIndexFieldMappingList")]
        private InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs>? _DataSourceToIndexFieldMappingList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-datasourcetoindexfieldmappinglist.html#cfn-kendra-datasource-datasourcetoindexfieldmappinglist-datasourcetoindexfieldmappinglist
        /// </summary>
        public InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs> DataSourceToIndexFieldMappingList
        {
            get => _DataSourceToIndexFieldMappingList ?? (_DataSourceToIndexFieldMappingList = new InputList<Inputs.DataSourceDataSourceToIndexFieldMappingArgs>());
            set => _DataSourceToIndexFieldMappingList = value;
        }

        public DataSourceDataSourceToIndexFieldMappingListArgs()
        {
        }
    }
}