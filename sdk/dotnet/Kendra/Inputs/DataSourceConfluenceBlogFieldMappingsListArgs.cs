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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogfieldmappingslist.html
    /// </summary>
    public sealed class DataSourceConfluenceBlogFieldMappingsListArgs : Pulumi.ResourceArgs
    {
        [Input("ConfluenceBlogFieldMappingsList")]
        private InputList<Inputs.DataSourceConfluenceBlogToIndexFieldMappingArgs>? _ConfluenceBlogFieldMappingsList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogfieldmappingslist.html#cfn-kendra-datasource-confluenceblogfieldmappingslist-confluenceblogfieldmappingslist
        /// </summary>
        public InputList<Inputs.DataSourceConfluenceBlogToIndexFieldMappingArgs> ConfluenceBlogFieldMappingsList
        {
            get => _ConfluenceBlogFieldMappingsList ?? (_ConfluenceBlogFieldMappingsList = new InputList<Inputs.DataSourceConfluenceBlogToIndexFieldMappingArgs>());
            set => _ConfluenceBlogFieldMappingsList = value;
        }

        public DataSourceConfluenceBlogFieldMappingsListArgs()
        {
        }
    }
}
