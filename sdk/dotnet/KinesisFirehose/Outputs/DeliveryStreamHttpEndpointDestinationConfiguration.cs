// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class DeliveryStreamHttpEndpointDestinationConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-bufferinghints
        /// </summary>
        public readonly Outputs.DeliveryStreamBufferingHints? BufferingHints;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-cloudwatchloggingoptions
        /// </summary>
        public readonly Outputs.DeliveryStreamCloudWatchLoggingOptions? CloudWatchLoggingOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-endpointconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamHttpEndpointConfiguration EndpointConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-processingconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamProcessingConfiguration? ProcessingConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-requestconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamHttpEndpointRequestConfiguration? RequestConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-retryoptions
        /// </summary>
        public readonly Outputs.DeliveryStreamRetryOptions? RetryOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-rolearn
        /// </summary>
        public readonly string? RoleARN;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-s3backupmode
        /// </summary>
        public readonly string? S3BackupMode;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration-s3configuration
        /// </summary>
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration S3Configuration;

        [OutputConstructor]
        private DeliveryStreamHttpEndpointDestinationConfiguration(
            Outputs.DeliveryStreamBufferingHints? bufferingHints,

            Outputs.DeliveryStreamCloudWatchLoggingOptions? cloudWatchLoggingOptions,

            Outputs.DeliveryStreamHttpEndpointConfiguration endpointConfiguration,

            Outputs.DeliveryStreamProcessingConfiguration? processingConfiguration,

            Outputs.DeliveryStreamHttpEndpointRequestConfiguration? requestConfiguration,

            Outputs.DeliveryStreamRetryOptions? retryOptions,

            string? roleARN,

            string? s3BackupMode,

            Outputs.DeliveryStreamS3DestinationConfiguration s3Configuration)
        {
            BufferingHints = bufferingHints;
            CloudWatchLoggingOptions = cloudWatchLoggingOptions;
            EndpointConfiguration = endpointConfiguration;
            ProcessingConfiguration = processingConfiguration;
            RequestConfiguration = requestConfiguration;
            RetryOptions = retryOptions;
            RoleARN = roleARN;
            S3BackupMode = s3BackupMode;
            S3Configuration = s3Configuration;
        }
    }
}
