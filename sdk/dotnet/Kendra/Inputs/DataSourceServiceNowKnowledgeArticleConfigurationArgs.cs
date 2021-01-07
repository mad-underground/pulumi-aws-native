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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html
    /// </summary>
    public sealed class DataSourceServiceNowKnowledgeArticleConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-crawlattachments
        /// </summary>
        [Input("CrawlAttachments")]
        public Input<bool>? CrawlAttachments { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-documentdatafieldname
        /// </summary>
        [Input("DocumentDataFieldName", required: true)]
        public Input<string> DocumentDataFieldName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-documenttitlefieldname
        /// </summary>
        [Input("DocumentTitleFieldName")]
        public Input<string>? DocumentTitleFieldName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-excludeattachmentfilepatterns
        /// </summary>
        [Input("ExcludeAttachmentFilePatterns")]
        public Input<Inputs.DataSourceDataSourceInclusionsExclusionsStringsArgs>? ExcludeAttachmentFilePatterns { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-fieldmappings
        /// </summary>
        [Input("FieldMappings")]
        public Input<Inputs.DataSourceDataSourceToIndexFieldMappingListArgs>? FieldMappings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-includeattachmentfilepatterns
        /// </summary>
        [Input("IncludeAttachmentFilePatterns")]
        public Input<Inputs.DataSourceDataSourceInclusionsExclusionsStringsArgs>? IncludeAttachmentFilePatterns { get; set; }

        public DataSourceServiceNowKnowledgeArticleConfigurationArgs()
        {
        }
    }
}