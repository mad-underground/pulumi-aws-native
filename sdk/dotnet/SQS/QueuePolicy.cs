// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SQS
{
    /// <summary>
    /// Resource Type definition for AWS::SQS::QueuePolicy
    /// </summary>
    [Obsolete(@"QueuePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:sqs:QueuePolicy")]
    public partial class QueuePolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A policy document that contains the permissions for the specified Amazon SQS queues. For more information about Amazon SQS policies, see Creating Custom Policies Using the Access Policy Language in the Amazon Simple Queue Service Developer Guide.
        /// </summary>
        [Output("policyDocument")]
        public Output<object> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// The URLs of the queues to which you want to add the policy. You can use the Ref function to specify an AWS::SQS::Queue resource.
        /// </summary>
        [Output("queues")]
        public Output<ImmutableArray<string>> Queues { get; private set; } = null!;


        /// <summary>
        /// Create a QueuePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QueuePolicy(string name, QueuePolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:sqs:QueuePolicy", name, args ?? new QueuePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QueuePolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:sqs:QueuePolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing QueuePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QueuePolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new QueuePolicy(name, id, options);
        }
    }

    public sealed class QueuePolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A policy document that contains the permissions for the specified Amazon SQS queues. For more information about Amazon SQS policies, see Creating Custom Policies Using the Access Policy Language in the Amazon Simple Queue Service Developer Guide.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<object> PolicyDocument { get; set; } = null!;

        [Input("queues", required: true)]
        private InputList<string>? _queues;

        /// <summary>
        /// The URLs of the queues to which you want to add the policy. You can use the Ref function to specify an AWS::SQS::Queue resource.
        /// </summary>
        public InputList<string> Queues
        {
            get => _queues ?? (_queues = new InputList<string>());
            set => _queues = value;
        }

        public QueuePolicyArgs()
        {
        }
        public static new QueuePolicyArgs Empty => new QueuePolicyArgs();
    }
}
