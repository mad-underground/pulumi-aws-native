// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    /// <summary>
    /// Resource Type definition for AWS::IoT::BillingGroup
    /// </summary>
    [AwsNativeResourceType("aws-native:iot:BillingGroup")]
    public partial class BillingGroup : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("billingGroupName")]
        public Output<string?> BillingGroupName { get; private set; } = null!;

        [Output("billingGroupProperties")]
        public Output<Outputs.BillingGroupPropertiesProperties?> BillingGroupProperties { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a BillingGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BillingGroup(string name, BillingGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:iot:BillingGroup", name, args ?? new BillingGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BillingGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:iot:BillingGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "billingGroupName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BillingGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BillingGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BillingGroup(name, id, options);
        }
    }

    public sealed class BillingGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("billingGroupName")]
        public Input<string>? BillingGroupName { get; set; }

        [Input("billingGroupProperties")]
        public Input<Inputs.BillingGroupPropertiesPropertiesArgs>? BillingGroupProperties { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public BillingGroupArgs()
        {
        }
        public static new BillingGroupArgs Empty => new BillingGroupArgs();
    }
}
