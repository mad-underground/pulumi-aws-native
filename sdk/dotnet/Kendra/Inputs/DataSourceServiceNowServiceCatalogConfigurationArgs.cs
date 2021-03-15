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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.html
    /// </summary>
    public sealed class DataSourceServiceNowServiceCatalogConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.html#cfn-kendra-datasource-servicenowservicecatalogconfiguration-crawlattachments
        /// </summary>
        [Input("crawlAttachments")]
        public Input<bool>? CrawlAttachments { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.html#cfn-kendra-datasource-servicenowservicecatalogconfiguration-documentdatafieldname
        /// </summary>
        [Input("documentDataFieldName", required: true)]
        public Input<string> DocumentDataFieldName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.html#cfn-kendra-datasource-servicenowservicecatalogconfiguration-documenttitlefieldname
        /// </summary>
        [Input("documentTitleFieldName")]
        public Input<string>? DocumentTitleFieldName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.html#cfn-kendra-datasource-servicenowservicecatalogconfiguration-excludeattachmentfilepatterns
        /// </summary>
        [Input("excludeAttachmentFilePatterns")]
        public Input<Inputs.DataSourceDataSourceInclusionsExclusionsStringsArgs>? ExcludeAttachmentFilePatterns { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.html#cfn-kendra-datasource-servicenowservicecatalogconfiguration-fieldmappings
        /// </summary>
        [Input("fieldMappings")]
        public Input<Inputs.DataSourceDataSourceToIndexFieldMappingListArgs>? FieldMappings { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowservicecatalogconfiguration.html#cfn-kendra-datasource-servicenowservicecatalogconfiguration-includeattachmentfilepatterns
        /// </summary>
        [Input("includeAttachmentFilePatterns")]
        public Input<Inputs.DataSourceDataSourceInclusionsExclusionsStringsArgs>? IncludeAttachmentFilePatterns { get; set; }

        public DataSourceServiceNowServiceCatalogConfigurationArgs()
        {
        }
    }
}
