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
    public sealed class DataSourceConfluenceBlogFieldMappingsList
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluenceblogfieldmappingslist.html#cfn-kendra-datasource-confluenceblogfieldmappingslist-confluenceblogfieldmappingslist
        /// </summary>
        public readonly ImmutableArray<Outputs.DataSourceConfluenceBlogToIndexFieldMapping> ConfluenceBlogFieldMappingsList;

        [OutputConstructor]
        private DataSourceConfluenceBlogFieldMappingsList(ImmutableArray<Outputs.DataSourceConfluenceBlogToIndexFieldMapping> ConfluenceBlogFieldMappingsList)
        {
            this.ConfluenceBlogFieldMappingsList = ConfluenceBlogFieldMappingsList;
        }
    }
}