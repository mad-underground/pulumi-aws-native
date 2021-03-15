// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kendra.Outputs
{

    [OutputType]
    public sealed class DataSourceConfluenceSpaceConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-crawlarchivedspaces
        /// </summary>
        public readonly bool? CrawlArchivedSpaces;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-crawlpersonalspaces
        /// </summary>
        public readonly bool? CrawlPersonalSpaces;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-excludespaces
        /// </summary>
        public readonly Outputs.DataSourceConfluenceSpaceList? ExcludeSpaces;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-includespaces
        /// </summary>
        public readonly Outputs.DataSourceConfluenceSpaceList? IncludeSpaces;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-spacefieldmappings
        /// </summary>
        public readonly Outputs.DataSourceConfluenceSpaceFieldMappingsList? SpaceFieldMappings;

        [OutputConstructor]
        private DataSourceConfluenceSpaceConfiguration(
            bool? crawlArchivedSpaces,

            bool? crawlPersonalSpaces,

            Outputs.DataSourceConfluenceSpaceList? excludeSpaces,

            Outputs.DataSourceConfluenceSpaceList? includeSpaces,

            Outputs.DataSourceConfluenceSpaceFieldMappingsList? spaceFieldMappings)
        {
            CrawlArchivedSpaces = crawlArchivedSpaces;
            CrawlPersonalSpaces = crawlPersonalSpaces;
            ExcludeSpaces = excludeSpaces;
            IncludeSpaces = includeSpaces;
            SpaceFieldMappings = spaceFieldMappings;
        }
    }
}
