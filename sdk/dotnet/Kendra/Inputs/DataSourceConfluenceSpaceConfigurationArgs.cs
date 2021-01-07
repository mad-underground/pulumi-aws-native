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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html
    /// </summary>
    public sealed class DataSourceConfluenceSpaceConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-crawlarchivedspaces
        /// </summary>
        [Input("CrawlArchivedSpaces")]
        public Input<bool>? CrawlArchivedSpaces { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-crawlpersonalspaces
        /// </summary>
        [Input("CrawlPersonalSpaces")]
        public Input<bool>? CrawlPersonalSpaces { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-excludespaces
        /// </summary>
        [Input("ExcludeSpaces")]
        public Input<Inputs.DataSourceConfluenceSpaceListArgs>? ExcludeSpaces { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-includespaces
        /// </summary>
        [Input("IncludeSpaces")]
        public Input<Inputs.DataSourceConfluenceSpaceListArgs>? IncludeSpaces { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencespaceconfiguration.html#cfn-kendra-datasource-confluencespaceconfiguration-spacefieldmappings
        /// </summary>
        [Input("SpaceFieldMappings")]
        public Input<Inputs.DataSourceConfluenceSpaceFieldMappingsListArgs>? SpaceFieldMappings { get; set; }

        public DataSourceConfluenceSpaceConfigurationArgs()
        {
        }
    }
}
