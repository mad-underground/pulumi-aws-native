// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisFirehose.Outputs
{

    [OutputType]
    public sealed class DeliveryStreamProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput
        /// </summary>
        public readonly Outputs.DeliveryStreamDeliveryStreamEncryptionConfigurationInput? DeliveryStreamEncryptionConfigurationInput;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamname
        /// </summary>
        public readonly string? DeliveryStreamName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamtype
        /// </summary>
        public readonly string? DeliveryStreamType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamElasticsearchDestinationConfiguration? ElasticsearchDestinationConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamExtendedS3DestinationConfiguration? ExtendedS3DestinationConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamHttpEndpointDestinationConfiguration? HttpEndpointDestinationConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamKinesisStreamSourceConfiguration? KinesisStreamSourceConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamRedshiftDestinationConfiguration? RedshiftDestinationConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamS3DestinationConfiguration? S3DestinationConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-splunkdestinationconfiguration
        /// </summary>
        public readonly Outputs.DeliveryStreamSplunkDestinationConfiguration? SplunkDestinationConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private DeliveryStreamProperties(
            Outputs.DeliveryStreamDeliveryStreamEncryptionConfigurationInput? DeliveryStreamEncryptionConfigurationInput,

            string? DeliveryStreamName,

            string? DeliveryStreamType,

            Outputs.DeliveryStreamElasticsearchDestinationConfiguration? ElasticsearchDestinationConfiguration,

            Outputs.DeliveryStreamExtendedS3DestinationConfiguration? ExtendedS3DestinationConfiguration,

            Outputs.DeliveryStreamHttpEndpointDestinationConfiguration? HttpEndpointDestinationConfiguration,

            Outputs.DeliveryStreamKinesisStreamSourceConfiguration? KinesisStreamSourceConfiguration,

            Outputs.DeliveryStreamRedshiftDestinationConfiguration? RedshiftDestinationConfiguration,

            Outputs.DeliveryStreamS3DestinationConfiguration? S3DestinationConfiguration,

            Outputs.DeliveryStreamSplunkDestinationConfiguration? SplunkDestinationConfiguration,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.DeliveryStreamEncryptionConfigurationInput = DeliveryStreamEncryptionConfigurationInput;
            this.DeliveryStreamName = DeliveryStreamName;
            this.DeliveryStreamType = DeliveryStreamType;
            this.ElasticsearchDestinationConfiguration = ElasticsearchDestinationConfiguration;
            this.ExtendedS3DestinationConfiguration = ExtendedS3DestinationConfiguration;
            this.HttpEndpointDestinationConfiguration = HttpEndpointDestinationConfiguration;
            this.KinesisStreamSourceConfiguration = KinesisStreamSourceConfiguration;
            this.RedshiftDestinationConfiguration = RedshiftDestinationConfiguration;
            this.S3DestinationConfiguration = S3DestinationConfiguration;
            this.SplunkDestinationConfiguration = SplunkDestinationConfiguration;
            this.Tags = Tags;
        }
    }
}
