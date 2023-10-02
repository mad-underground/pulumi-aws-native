// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose
{
    /// <summary>
    /// Resource Type definition for AWS::KinesisFirehose::DeliveryStream
    /// </summary>
    [AwsNativeResourceType("aws-native:kinesisfirehose:DeliveryStream")]
    public partial class DeliveryStream : global::Pulumi.CustomResource
    {
        [Output("amazonOpenSearchServerlessDestinationConfiguration")]
        public Output<Outputs.DeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration?> AmazonOpenSearchServerlessDestinationConfiguration { get; private set; } = null!;

        [Output("amazonopensearchserviceDestinationConfiguration")]
        public Output<Outputs.DeliveryStreamAmazonopensearchserviceDestinationConfiguration?> AmazonopensearchserviceDestinationConfiguration { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("deliveryStreamEncryptionConfigurationInput")]
        public Output<Outputs.DeliveryStreamEncryptionConfigurationInput?> DeliveryStreamEncryptionConfigurationInput { get; private set; } = null!;

        [Output("deliveryStreamName")]
        public Output<string?> DeliveryStreamName { get; private set; } = null!;

        [Output("deliveryStreamType")]
        public Output<Pulumi.AwsNative.KinesisFirehose.DeliveryStreamType?> DeliveryStreamType { get; private set; } = null!;

        [Output("elasticsearchDestinationConfiguration")]
        public Output<Outputs.DeliveryStreamElasticsearchDestinationConfiguration?> ElasticsearchDestinationConfiguration { get; private set; } = null!;

        [Output("extendedS3DestinationConfiguration")]
        public Output<Outputs.DeliveryStreamExtendedS3DestinationConfiguration?> ExtendedS3DestinationConfiguration { get; private set; } = null!;

        [Output("httpEndpointDestinationConfiguration")]
        public Output<Outputs.DeliveryStreamHttpEndpointDestinationConfiguration?> HttpEndpointDestinationConfiguration { get; private set; } = null!;

        [Output("kinesisStreamSourceConfiguration")]
        public Output<Outputs.DeliveryStreamKinesisStreamSourceConfiguration?> KinesisStreamSourceConfiguration { get; private set; } = null!;

        [Output("mskSourceConfiguration")]
        public Output<Outputs.DeliveryStreamMskSourceConfiguration?> MskSourceConfiguration { get; private set; } = null!;

        [Output("redshiftDestinationConfiguration")]
        public Output<Outputs.DeliveryStreamRedshiftDestinationConfiguration?> RedshiftDestinationConfiguration { get; private set; } = null!;

        [Output("s3DestinationConfiguration")]
        public Output<Outputs.DeliveryStreamS3DestinationConfiguration?> S3DestinationConfiguration { get; private set; } = null!;

        [Output("splunkDestinationConfiguration")]
        public Output<Outputs.DeliveryStreamSplunkDestinationConfiguration?> SplunkDestinationConfiguration { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.DeliveryStreamTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DeliveryStream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeliveryStream(string name, DeliveryStreamArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:kinesisfirehose:DeliveryStream", name, args ?? new DeliveryStreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeliveryStream(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:kinesisfirehose:DeliveryStream", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "amazonOpenSearchServerlessDestinationConfiguration.vpcConfiguration",
                    "amazonopensearchserviceDestinationConfiguration.vpcConfiguration",
                    "deliveryStreamName",
                    "deliveryStreamType",
                    "elasticsearchDestinationConfiguration.vpcConfiguration",
                    "kinesisStreamSourceConfiguration",
                    "mskSourceConfiguration",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DeliveryStream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeliveryStream Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeliveryStream(name, id, options);
        }
    }

    public sealed class DeliveryStreamArgs : global::Pulumi.ResourceArgs
    {
        [Input("amazonOpenSearchServerlessDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationArgs>? AmazonOpenSearchServerlessDestinationConfiguration { get; set; }

        [Input("amazonopensearchserviceDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamAmazonopensearchserviceDestinationConfigurationArgs>? AmazonopensearchserviceDestinationConfiguration { get; set; }

        [Input("deliveryStreamEncryptionConfigurationInput")]
        public Input<Inputs.DeliveryStreamEncryptionConfigurationInputArgs>? DeliveryStreamEncryptionConfigurationInput { get; set; }

        [Input("deliveryStreamName")]
        public Input<string>? DeliveryStreamName { get; set; }

        [Input("deliveryStreamType")]
        public Input<Pulumi.AwsNative.KinesisFirehose.DeliveryStreamType>? DeliveryStreamType { get; set; }

        [Input("elasticsearchDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamElasticsearchDestinationConfigurationArgs>? ElasticsearchDestinationConfiguration { get; set; }

        [Input("extendedS3DestinationConfiguration")]
        public Input<Inputs.DeliveryStreamExtendedS3DestinationConfigurationArgs>? ExtendedS3DestinationConfiguration { get; set; }

        [Input("httpEndpointDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamHttpEndpointDestinationConfigurationArgs>? HttpEndpointDestinationConfiguration { get; set; }

        [Input("kinesisStreamSourceConfiguration")]
        public Input<Inputs.DeliveryStreamKinesisStreamSourceConfigurationArgs>? KinesisStreamSourceConfiguration { get; set; }

        [Input("mskSourceConfiguration")]
        public Input<Inputs.DeliveryStreamMskSourceConfigurationArgs>? MskSourceConfiguration { get; set; }

        [Input("redshiftDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamRedshiftDestinationConfigurationArgs>? RedshiftDestinationConfiguration { get; set; }

        [Input("s3DestinationConfiguration")]
        public Input<Inputs.DeliveryStreamS3DestinationConfigurationArgs>? S3DestinationConfiguration { get; set; }

        [Input("splunkDestinationConfiguration")]
        public Input<Inputs.DeliveryStreamSplunkDestinationConfigurationArgs>? SplunkDestinationConfiguration { get; set; }

        [Input("tags")]
        private InputList<Inputs.DeliveryStreamTagArgs>? _tags;
        public InputList<Inputs.DeliveryStreamTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.DeliveryStreamTagArgs>());
            set => _tags = value;
        }

        public DeliveryStreamArgs()
        {
        }
        public static new DeliveryStreamArgs Empty => new DeliveryStreamArgs();
    }
}
