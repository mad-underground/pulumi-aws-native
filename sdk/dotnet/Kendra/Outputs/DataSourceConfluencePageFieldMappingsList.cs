// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceConfluencePageFieldMappingsList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepagefieldmappingslist.html#cfn-kendra-datasource-confluencepagefieldmappingslist-confluencepagefieldmappingslist
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceConfluencePageToIndexFieldMapping> ConfluencePageFieldMappingsList;

        [OutputConstructor]
        private DataSourceConfluencePageFieldMappingsList(ImmutableArray<Outputs.DataSourceConfluencePageToIndexFieldMapping> ConfluencePageFieldMappingsList)
        {
            this.ConfluencePageFieldMappingsList = ConfluencePageFieldMappingsList;
        }
    }
}