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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepagefieldmappingslist.html
    /// </summary>
    public sealed class DataSourceConfluencePageFieldMappingsListArgs : Pulumi.ResourceArgs
    {
        [Input("ConfluencePageFieldMappingsList")]
        private InputList<Inputs.DataSourceConfluencePageToIndexFieldMappingArgs>? _ConfluencePageFieldMappingsList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepagefieldmappingslist.html#cfn-kendra-datasource-confluencepagefieldmappingslist-confluencepagefieldmappingslist
        /// </summary>
        public InputList<Inputs.DataSourceConfluencePageToIndexFieldMappingArgs> ConfluencePageFieldMappingsList
        {
            get => _ConfluencePageFieldMappingsList ?? (_ConfluencePageFieldMappingsList = new InputList<Inputs.DataSourceConfluencePageToIndexFieldMappingArgs>());
            set => _ConfluencePageFieldMappingsList = value;
        }

        public DataSourceConfluencePageFieldMappingsListArgs()
        {
        }
    }
}
