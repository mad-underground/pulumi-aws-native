// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.KinesisFirehose.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html
    /// </summary>
    public sealed class DeliveryStreamPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamencryptionconfigurationinput
        /// </summary>
        [Input("DeliveryStreamEncryptionConfigurationInput")]
        public Input<Inputs.DeliveryStreamDeliveryStreamEncryptionConfigurationInputArgs>? DeliveryStreamEncryptionConfigurationInput { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamname
        /// </summary>
        [Input("DeliveryStreamName")]
        public Input<string>? DeliveryStreamName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-deliverystreamtype
        /// </summary>
        [Input("DeliveryStreamType")]
        public Input<string>? DeliveryStreamType { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-elasticsearchdestinationconfiguration
        /// </summary>
        [Input("ElasticsearchDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamElasticsearchDestinationConfigurationArgs>? ElasticsearchDestinationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-extendeds3destinationconfiguration
        /// </summary>
        [Input("ExtendedS3DestinationConfiguration")]
        public Input<Inputs.DeliveryStreamExtendedS3DestinationConfigurationArgs>? ExtendedS3DestinationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-httpendpointdestinationconfiguration
        /// </summary>
        [Input("HttpEndpointDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamHttpEndpointDestinationConfigurationArgs>? HttpEndpointDestinationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration
        /// </summary>
        [Input("KinesisStreamSourceConfiguration")]
        public Input<Inputs.DeliveryStreamKinesisStreamSourceConfigurationArgs>? KinesisStreamSourceConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-redshiftdestinationconfiguration
        /// </summary>
        [Input("RedshiftDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamRedshiftDestinationConfigurationArgs>? RedshiftDestinationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-s3destinationconfiguration
        /// </summary>
        [Input("S3DestinationConfiguration")]
        public Input<Inputs.DeliveryStreamS3DestinationConfigurationArgs>? S3DestinationConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-splunkdestinationconfiguration
        /// </summary>
        [Input("SplunkDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamSplunkDestinationConfigurationArgs>? SplunkDestinationConfiguration { get; set; }

        [Input("Tags")]
        private InputList<Pulumi.Cloudformation.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html#cfn-kinesisfirehose-deliverystream-tags
        /// </summary>
        public InputList<Pulumi.Cloudformation.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.Cloudformation.Inputs.TagArgs>());
            set => _Tags = value;
        }

        public DeliveryStreamPropertiesArgs()
        {
        }
    }
}
