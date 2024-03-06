// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Kinesis
{
    /// <summary>
    /// Resource Type definition for AWS::Kinesis::StreamConsumer
    /// </summary>
    [Obsolete(@"StreamConsumer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:kinesis:StreamConsumer")]
    public partial class StreamConsumer : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("consumerArn")]
        public Output<string> ConsumerArn { get; private set; } = null!;

        [Output("consumerCreationTimestamp")]
        public Output<string> ConsumerCreationTimestamp { get; private set; } = null!;

        [Output("consumerName")]
        public Output<string> ConsumerName { get; private set; } = null!;

        [Output("consumerStatus")]
        public Output<string> ConsumerStatus { get; private set; } = null!;

        [Output("streamArn")]
        public Output<string> StreamArn { get; private set; } = null!;


        /// <summary>
        /// Create a StreamConsumer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StreamConsumer(string name, StreamConsumerArgs args, CustomResourceOptions? options = null)
            : base("aws-native:kinesis:StreamConsumer", name, args ?? new StreamConsumerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StreamConsumer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:kinesis:StreamConsumer", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "consumerName",
                    "streamArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing StreamConsumer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StreamConsumer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StreamConsumer(name, id, options);
        }
    }

    public sealed class StreamConsumerArgs : global::Pulumi.ResourceArgs
    {
        [Input("consumerName", required: true)]
        public Input<string> ConsumerName { get; set; } = null!;

        [Input("streamArn", required: true)]
        public Input<string> StreamArn { get; set; } = null!;

        public StreamConsumerArgs()
        {
        }
        public static new StreamConsumerArgs Empty => new StreamConsumerArgs();
    }
}
