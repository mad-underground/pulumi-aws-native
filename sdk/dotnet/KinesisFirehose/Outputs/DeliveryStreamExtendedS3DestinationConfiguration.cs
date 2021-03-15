// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamExtendedS3DestinationConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-bucketarn
        /// </summary>
        public readonly string BucketARN;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-bufferinghints
        /// </summary>
        public readonly Outputs.DeliveryStreamBufferingHints? BufferingHints;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-cloudwatchloggingoptions
        /// </summary>
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-compressionformat
        /// </summary>
        public readonly string? CompressionFormat;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-dataformatconversionconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamDataFormatConversionConfiguration? DataFormatConversionConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-encryptionconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamEncryptionConfiguration? EncryptionConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-erroroutputprefix
        /// </summary>
        public readonly string? ErrorOutputPrefix;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-prefix
        /// </summary>
        public readonly string? Prefix;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-processingconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-rolearn
        /// </summary>
        public readonly string RoleARN;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-s3backupconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration? S3BackupConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-extendeds3destinationconfiguration.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration-s3backupmode
        /// </summary>
        public readonly string? S3BackupMode;

        [OutputConstructor]
        private DeliveryStreamExtendedS3DestinationConfiguration(
            string bucketARN,

            Outputs.DeliveryStreamBufferingHints? bufferingHints,

            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            string? compressionFormat,

            Outputs.DeliveryStreamDataFormatConversionConfiguration? dataFormatConversionConfiguration,

            Outputs.DeliveryStreamEncryptionConfiguration? encryptionConfiguration,

            string? errorOutputPrefix,

            string? prefix,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            string roleARN,

            Outputs.DeliveryStreamS3DestinationConfiguration? s3BackupConfiguration,

            string? s3BackupMode)
        {
            BucketARN = bucketARN;
            BufferingHints = bufferingHints;
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            CompressionFormat = compressionFormat;
            DataFormatConversionConfiguration = dataFormatConversionConfiguration;
            EncryptionConfiguration = encryptionConfiguration;
            ErrorOutputPrefix = errorOutputPrefix;
            Prefix = prefix;
            ProcessingConfiguration = processingConfiguration;
            RoleARN = roleARN;
            S3BackupConfiguration = s3BackupConfiguration;
            S3BackupMode = s3BackupMode;
        }
    }
}
