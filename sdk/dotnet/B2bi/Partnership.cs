// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.B2bi
{
    /// <summary>
    /// Definition of AWS::B2BI::Partnership Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:b2bi:Partnership")]
    public partial class Partnership : global::Pulumi.CustomResource
    {
        [Output("capabilities")]
        public Output<ImmutableArray<string>> Capabilities { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        [Output("modifiedAt")]
        public Output<string> ModifiedAt { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("partnershipArn")]
        public Output<string> PartnershipArn { get; private set; } = null!;

        [Output("partnershipId")]
        public Output<string> PartnershipId { get; private set; } = null!;

        [Output("phone")]
        public Output<string?> Phone { get; private set; } = null!;

        [Output("profileId")]
        public Output<string> ProfileId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("tradingPartnerId")]
        public Output<string> TradingPartnerId { get; private set; } = null!;


        /// <summary>
        /// Create a Partnership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Partnership(string name, PartnershipArgs args, CustomResourceOptions? options = null)
            : base("aws-native:b2bi:Partnership", name, args ?? new PartnershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Partnership(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:b2bi:Partnership", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "email",
                    "phone",
                    "profileId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Partnership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Partnership Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Partnership(name, id, options);
        }
    }

    public sealed class PartnershipArgs : global::Pulumi.ResourceArgs
    {
        [Input("capabilities")]
        private InputList<string>? _capabilities;
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("phone")]
        public Input<string>? Phone { get; set; }

        [Input("profileId", required: true)]
        public Input<string> ProfileId { get; set; } = null!;

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public PartnershipArgs()
        {
        }
        public static new PartnershipArgs Empty => new PartnershipArgs();
    }
}
