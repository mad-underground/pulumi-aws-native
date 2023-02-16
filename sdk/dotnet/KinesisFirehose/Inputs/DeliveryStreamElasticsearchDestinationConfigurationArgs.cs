// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamElasticsearchDestinationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bufferingHints")]
        public Input<Inputs.DeliveryStreamElasticsearchBufferingHintsArgs>? BufferingHints { get; set; }

        [Input("cloudWatchLoggingOptions")]
        public Input<Inputs.DeliveryStreamCloudWatchLoggingOptionsArgs>? CloudWatchLoggingOptions { get; set; }

        [Input("clusterEndpoint")]
        public Input<string>? ClusterEndpoint { get; set; }

        [Input("domainARN")]
        public Input<string>? DomainARN { get; set; }

        [Input("indexName", required: true)]
        public Input<string> IndexName { get; set; } = null!;

        [Input("indexRotationPeriod")]
        public Input<Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationIndexRotationPeriod>? IndexRotationPeriod { get; set; }

        [Input("processingConfiguration")]
        public Input<Inputs.DeliveryStreamProcessingConfigurationArgs>? ProcessingConfiguration { get; set; }

        [Input("retryOptions")]
        public Input<Inputs.DeliveryStreamElasticsearchRetryOptionsArgs>? RetryOptions { get; set; }

        [Input("roleARN", required: true)]
        public Input<string> RoleARN { get; set; } = null!;

        [Input("s3BackupMode")]
        public Input<Pulumi.AwsNative.KinesisFirehose.DeliveryStreamElasticsearchDestinationConfigurationS3BackupMode>? S3BackupMode { get; set; }

        [Input("s3Configuration", required: true)]
        public Input<Inputs.DeliveryStreamS3DestinationConfigurationArgs> S3Configuration { get; set; } = null!;

        [Input("typeName")]
        public Input<string>? TypeName { get; set; }

        [Input("vpcConfiguration")]
        public Input<Inputs.DeliveryStreamVpcConfigurationArgs>? VpcConfiguration { get; set; }

        public DeliveryStreamElasticsearchDestinationConfigurationArgs()
        {
        }
        public static new DeliveryStreamElasticsearchDestinationConfigurationArgs Empty => new DeliveryStreamElasticsearchDestinationConfigurationArgs();
    }
}
