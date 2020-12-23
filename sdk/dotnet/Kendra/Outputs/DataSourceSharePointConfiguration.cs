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
    public sealed class DataSourceSharePointConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-crawlattachments
        /// </summary>
        public readonly bool? CrawlAttachments;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-disablelocalgroups
        /// </summary>
        public readonly bool? DisableLocalGroups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-documenttitlefieldname
        /// </summary>
        public readonly string? DocumentTitleFieldName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-exclusionpatterns
        /// </summary>
        public readonly Outputs.DataSourceDataSourceInclusionsExclusionsStrings? ExclusionPatterns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-fieldmappings
        /// </summary>
        public readonly Outputs.DataSourceDataSourceToIndexFieldMappingList? FieldMappings;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-inclusionpatterns
        /// </summary>
        public readonly Outputs.DataSourceDataSourceInclusionsExclusionsStrings? InclusionPatterns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-secretarn
        /// </summary>
        public readonly string SecretArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-sharepointversion
        /// </summary>
        public readonly string SharePointVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-urls
        /// </summary>
        public readonly ImmutableArray<string> Urls;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-usechangelog
        /// </summary>
        public readonly bool? UseChangeLog;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-sharepointconfiguration.html#cfn-kendra-datasource-sharepointconfiguration-vpcconfiguration
        /// </summary>
        public readonly Outputs.DataSourceDataSourceVpcConfiguration? VpcConfiguration;

        [OutputConstructor]
        private DataSourceSharePointConfiguration(
            bool? CrawlAttachments,

            bool? DisableLocalGroups,

            string? DocumentTitleFieldName,

            Outputs.DataSourceDataSourceInclusionsExclusionsStrings? ExclusionPatterns,

            Outputs.DataSourceDataSourceToIndexFieldMappingList? FieldMappings,

            Outputs.DataSourceDataSourceInclusionsExclusionsStrings? InclusionPatterns,

            string SecretArn,

            string SharePointVersion,

            ImmutableArray<string> Urls,

            bool? UseChangeLog,

            Outputs.DataSourceDataSourceVpcConfiguration? VpcConfiguration)
        {
            this.CrawlAttachments = CrawlAttachments;
            this.DisableLocalGroups = DisableLocalGroups;
            this.DocumentTitleFieldName = DocumentTitleFieldName;
            this.ExclusionPatterns = ExclusionPatterns;
            this.FieldMappings = FieldMappings;
            this.InclusionPatterns = InclusionPatterns;
            this.SecretArn = SecretArn;
            this.SharePointVersion = SharePointVersion;
            this.Urls = Urls;
            this.UseChangeLog = UseChangeLog;
            this.VpcConfiguration = VpcConfiguration;
        }
    }
}
