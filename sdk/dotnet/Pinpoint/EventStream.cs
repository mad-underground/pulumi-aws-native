// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    /// <summary>
    /// Resource Type definition for AWS::Pinpoint::EventStream
    /// </summary>
    [Obsolete(@"EventStream is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:pinpoint:EventStream")]
    public partial class EventStream : global::Pulumi.CustomResource
    {
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("destinationStreamArn")]
        public Output<string> DestinationStreamArn { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;


        /// <summary>
        /// Create a EventStream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventStream(string name, EventStreamArgs args, CustomResourceOptions? options = null)
            : base("aws-native:pinpoint:EventStream", name, args ?? new EventStreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventStream(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:pinpoint:EventStream", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "applicationId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventStream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventStream Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventStream(name, id, options);
        }
    }

    public sealed class EventStreamArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        [Input("destinationStreamArn", required: true)]
        public Input<string> DestinationStreamArn { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        public EventStreamArgs()
        {
        }
        public static new EventStreamArgs Empty => new EventStreamArgs();
    }
}
