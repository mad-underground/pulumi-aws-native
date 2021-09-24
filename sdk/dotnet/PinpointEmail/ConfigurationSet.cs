// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PinpointEmail
{
    /// <summary>
    /// Resource Type definition for AWS::PinpointEmail::ConfigurationSet
    /// </summary>
    [Obsolete(@"ConfigurationSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:pinpointemail:ConfigurationSet")]
    public partial class ConfigurationSet : Pulumi.CustomResource
    {
        [Output("deliveryOptions")]
        public Output<Outputs.ConfigurationSetDeliveryOptions?> DeliveryOptions { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("reputationOptions")]
        public Output<Outputs.ConfigurationSetReputationOptions?> ReputationOptions { get; private set; } = null!;

        [Output("sendingOptions")]
        public Output<Outputs.ConfigurationSetSendingOptions?> SendingOptions { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.ConfigurationSetTags>> Tags { get; private set; } = null!;

        [Output("trackingOptions")]
        public Output<Outputs.ConfigurationSetTrackingOptions?> TrackingOptions { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigurationSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigurationSet(string name, ConfigurationSetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:pinpointemail:ConfigurationSet", name, args ?? new ConfigurationSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConfigurationSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:pinpointemail:ConfigurationSet", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConfigurationSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigurationSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigurationSet(name, id, options);
        }
    }

    public sealed class ConfigurationSetArgs : Pulumi.ResourceArgs
    {
        [Input("deliveryOptions")]
        public Input<Inputs.ConfigurationSetDeliveryOptionsArgs>? DeliveryOptions { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("reputationOptions")]
        public Input<Inputs.ConfigurationSetReputationOptionsArgs>? ReputationOptions { get; set; }

        [Input("sendingOptions")]
        public Input<Inputs.ConfigurationSetSendingOptionsArgs>? SendingOptions { get; set; }

        [Input("tags")]
        private InputList<Inputs.ConfigurationSetTagsArgs>? _tags;
        public InputList<Inputs.ConfigurationSetTagsArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ConfigurationSetTagsArgs>());
            set => _tags = value;
        }

        [Input("trackingOptions")]
        public Input<Inputs.ConfigurationSetTrackingOptionsArgs>? TrackingOptions { get; set; }

        public ConfigurationSetArgs()
        {
        }
    }
}