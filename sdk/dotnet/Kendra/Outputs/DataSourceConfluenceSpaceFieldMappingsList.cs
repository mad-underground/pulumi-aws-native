// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacefieldmappingslist.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceConfluenceSpaceFieldMappingsList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespacefieldmappingslist.html#cfn-kendra-datasource-confluencespacefieldmappingslist-confluencespacefieldmappingslist
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceConfluenceSpaceToIndexFieldMapping> ConfluenceSpaceFieldMappingsList;

        [OutputConstructor]
        private DataSourceConfluenceSpaceFieldMappingsList(ImmutableArray<Outputs.DataSourceConfluenceSpaceToIndexFieldMapping> confluenceSpaceFieldMappingsList)
        {
            ConfluenceSpaceFieldMappingsList = confluenceSpaceFieldMappingsList;
        }
    }
}
