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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DataSourceS3DataSourceConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-accesscontrollistconfiguration
        /// </summary>
        public readonly Outputs.DataSourceAccessControlListConfiguration? AccessControlListConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-bucketname
        /// </summary>
        public readonly string BucketName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-documentsmetadataconfiguration
        /// </summary>
        public readonly Outputs.DataSourceDocumentsMetadataConfiguration? DocumentsMetadataConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-exclusionpatterns
        /// </summary>
        public readonly Outputs.DataSourceDataSourceInclusionsExclusionsStrings? ExclusionPatterns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-inclusionpatterns
        /// </summary>
        public readonly Outputs.DataSourceDataSourceInclusionsExclusionsStrings? InclusionPatterns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-s3datasourceconfiguration.html#cfn-kendra-datasource-s3datasourceconfiguration-inclusionprefixes
        /// </summary>
        public readonly Outputs.DataSourceDataSourceInclusionsExclusionsStrings? InclusionPrefixes;

        [OutputConstructor]
        private DataSourceS3DataSourceConfiguration(
            Outputs.DataSourceAccessControlListConfiguration? accessControlListConfiguration,

            string bucketName,

            Outputs.DataSourceDocumentsMetadataConfiguration? documentsMetadataConfiguration,

            Outputs.DataSourceDataSourceInclusionsExclusionsStrings? exclusionPatterns,

            Outputs.DataSourceDataSourceInclusionsExclusionsStrings? inclusionPatterns,

            Outputs.DataSourceDataSourceInclusionsExclusionsStrings? inclusionPrefixes)
        {
            AccessControlListConfiguration = accessControlListConfiguration;
            BucketName = bucketName;
            DocumentsMetadataConfiguration = documentsMetadataConfiguration;
            ExclusionPatterns = exclusionPatterns;
            InclusionPatterns = inclusionPatterns;
            InclusionPrefixes = inclusionPrefixes;
        }
    }
}
