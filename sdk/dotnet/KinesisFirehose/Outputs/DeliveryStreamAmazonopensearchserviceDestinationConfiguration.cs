// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamAmazonopensearchserviceDestinationConfiguration
    {
        public readonly Outputs.DeliveryStreamAmazonopensearchserviceBufferingHints? BufferingHints;
        public readonly Outputs.DeliveryStreamBulkRequestCustomizationConfiguration? BulkRequestCustomizationConfiguration;
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        public readonly string? ClusterEndpoint;
        public readonly Outputs.DeliveryStreamDocumentIdOptions? DocumentIdOptions;
        public readonly string? DomainARN;
        public readonly string IndexName;
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamAmazonopensearchserviceDestinationConfigurationIndexRotationPeriod? IndexRotationPeriod;
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        public readonly Outputs.DeliveryStreamAmazonopensearchserviceRetryOptions? RetryOptions;
        public readonly string RoleARN;
        public readonly Pulumi.AwsNative.KinesisFirehose.DeliveryStreamAmazonopensearchserviceDestinationConfigurationS3BackupMode? S3BackupMode;
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration S3Configuration;
        public readonly string? TypeName;
        public readonly Outputs.DeliveryStreamVpcConfiguration? VpcConfiguration;

        [OutputConstructor]
        private DeliveryStreamAmazonopensearchserviceDestinationConfiguration(
            Outputs.DeliveryStreamAmazonopensearchserviceBufferingHints? bufferingHints,

            Outputs.DeliveryStreamBulkRequestCustomizationConfiguration? bulkRequestCustomizationConfiguration,

            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            string? clusterEndpoint,

            Outputs.DeliveryStreamDocumentIdOptions? documentIdOptions,

            string? domainARN,

            string indexName,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamAmazonopensearchserviceDestinationConfigurationIndexRotationPeriod? indexRotationPeriod,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            Outputs.DeliveryStreamAmazonopensearchserviceRetryOptions? retryOptions,

            string roleARN,

            Pulumi.AwsNative.KinesisFirehose.DeliveryStreamAmazonopensearchserviceDestinationConfigurationS3BackupMode? s3BackupMode,

            Outputs.DeliveryStreamS3DestinationConfiguration s3Configuration,

            string? typeName,

            Outputs.DeliveryStreamVpcConfiguration? vpcConfiguration)
        {
            BufferingHints = bufferingHints;
            BulkRequestCustomizationConfiguration = bulkRequestCustomizationConfiguration;
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            ClusterEndpoint = clusterEndpoint;
            DocumentIdOptions = documentIdOptions;
            DomainARN = domainARN;
            IndexName = indexName;
            IndexRotationPeriod = indexRotationPeriod;
            ProcessingConfiguration = processingConfiguration;
            RetryOptions = retryOptions;
            RoleARN = roleARN;
            S3BackupMode = s3BackupMode;
            S3Configuration = s3Configuration;
            TypeName = typeName;
            VpcConfiguration = vpcConfiguration;
        }
    }
}
