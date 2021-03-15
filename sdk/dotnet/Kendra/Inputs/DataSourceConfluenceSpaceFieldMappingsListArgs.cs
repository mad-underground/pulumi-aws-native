// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacefieldmappingslist.html
    /// </summary>
    public sealed class DataSourceConfluenceSpaceFieldMappingsListArgs : Pulumi.ResourceArgs
    {
        [Input("confluenceSpaceFieldMappingsList")]
        private InputList<Inputs.DataSourceConfluenceSpaceToIndexFieldMappingArgs>? _confluenceSpaceFieldMappingsList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacefieldmappingslist.html#cfn-kendra-datasource-confluencespacefieldmappingslist-confluencespacefieldmappingslist
        /// </summary>
        public InputList<Inputs.DataSourceConfluenceSpaceToIndexFieldMappingArgs> ConfluenceSpaceFieldMappingsList
        {
            get => _confluenceSpaceFieldMappingsList ?? (_confluenceSpaceFieldMappingsList = new InputList<Inputs.DataSourceConfluenceSpaceToIndexFieldMappingArgs>());
            set => _confluenceSpaceFieldMappingsList = value;
        }

        public DataSourceConfluenceSpaceFieldMappingsListArgs()
        {
        }
    }
}
