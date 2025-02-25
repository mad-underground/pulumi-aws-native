// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamHttpEndpointDestinationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("bufferingHints")]
        public Input<Inputs.DeliveryStreamBufferingHintsArgs>? BufferingHints { get; set; }

        [Input("cloudWatchLoggingOptions")]
        public Input<Inputs.DeliveryStreamCloudWatchLoggingOptionsArgs>? CloudWatchLoggingOptions { get; set; }

        [Input("endpointConfiguration", required: true)]
        public Input<Inputs.DeliveryStreamHttpEndpointConfigurationArgs> EndpointConfiguration { get; set; } = null!;

        [Input("processingConfiguration")]
        public Input<Inputs.DeliveryStreamProcessingConfigurationArgs>? ProcessingConfiguration { get; set; }

        [Input("requestConfiguration")]
        public Input<Inputs.DeliveryStreamHttpEndpointRequestConfigurationArgs>? RequestConfiguration { get; set; }

        [Input("retryOptions")]
        public Input<Inputs.DeliveryStreamRetryOptionsArgs>? RetryOptions { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("s3BackupMode")]
        public Input<string>? S3BackupMode { get; set; }

        [Input("s3Configuration", required: true)]
        public Input<Inputs.DeliveryStreamS3DestinationConfigurationArgs> S3Configuration { get; set; } = null!;

        public DeliveryStreamHttpEndpointDestinationConfigurationArgs()
        {
        }
        public static new DeliveryStreamHttpEndpointDestinationConfigurationArgs Empty => new DeliveryStreamHttpEndpointDestinationConfigurationArgs();
    }
}
