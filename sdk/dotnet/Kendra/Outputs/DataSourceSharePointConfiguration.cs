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
            bool? crawlAttachments,

            bool? disableLocalGroups,

            string? documentTitleFieldName,

            Outputs.DataSourceDataSourceInclusionsExclusionsStrings? exclusionPatterns,

            Outputs.DataSourceDataSourceToIndexFieldMappingList? fieldMappings,

            Outputs.DataSourceDataSourceInclusionsExclusionsStrings? inclusionPatterns,

            string secretArn,

            string sharePointVersion,

            ImmutableArray<string> urls,

            bool? useChangeLog,

            Outputs.DataSourceDataSourceVpcConfiguration? vpcConfiguration)
        {
            CrawlAttachments = crawlAttachments;
            DisableLocalGroups = disableLocalGroups;
            DocumentTitleFieldName = documentTitleFieldName;
            ExclusionPatterns = exclusionPatterns;
            FieldMappings = fieldMappings;
            InclusionPatterns = inclusionPatterns;
            SecretArn = secretArn;
            SharePointVersion = sharePointVersion;
            Urls = urls;
            UseChangeLog = useChangeLog;
            VpcConfiguration = vpcConfiguration;
        }
    }
}
