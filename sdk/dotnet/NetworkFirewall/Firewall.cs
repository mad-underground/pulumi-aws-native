// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall
{
    /// <summary>
    /// Resource type definition for AWS::NetworkFirewall::Firewall
    /// </summary>
    [AwsNativeResourceType("aws-native:networkfirewall:Firewall")]
    public partial class Firewall : global::Pulumi.CustomResource
    {
        [Output("deleteProtection")]
        public Output<bool?> DeleteProtection { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("endpointIds")]
        public Output<ImmutableArray<string>> EndpointIds { get; private set; } = null!;

        [Output("firewallArn")]
        public Output<string> FirewallArn { get; private set; } = null!;

        [Output("firewallId")]
        public Output<string> FirewallId { get; private set; } = null!;

        [Output("firewallName")]
        public Output<string> FirewallName { get; private set; } = null!;

        [Output("firewallPolicyArn")]
        public Output<string> FirewallPolicyArn { get; private set; } = null!;

        [Output("firewallPolicyChangeProtection")]
        public Output<bool?> FirewallPolicyChangeProtection { get; private set; } = null!;

        [Output("subnetChangeProtection")]
        public Output<bool?> SubnetChangeProtection { get; private set; } = null!;

        [Output("subnetMappings")]
        public Output<ImmutableArray<Outputs.FirewallSubnetMapping>> SubnetMappings { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a Firewall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Firewall(string name, FirewallArgs args, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:Firewall", name, args ?? new FirewallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Firewall(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:networkfirewall:Firewall", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "firewallName",
                    "vpcId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Firewall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Firewall Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Firewall(name, id, options);
        }
    }

    public sealed class FirewallArgs : global::Pulumi.ResourceArgs
    {
        [Input("deleteProtection")]
        public Input<bool>? DeleteProtection { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("firewallName")]
        public Input<string>? FirewallName { get; set; }

        [Input("firewallPolicyArn", required: true)]
        public Input<string> FirewallPolicyArn { get; set; } = null!;

        [Input("firewallPolicyChangeProtection")]
        public Input<bool>? FirewallPolicyChangeProtection { get; set; }

        [Input("subnetChangeProtection")]
        public Input<bool>? SubnetChangeProtection { get; set; }

        [Input("subnetMappings", required: true)]
        private InputList<Inputs.FirewallSubnetMappingArgs>? _subnetMappings;
        public InputList<Inputs.FirewallSubnetMappingArgs> SubnetMappings
        {
            get => _subnetMappings ?? (_subnetMappings = new InputList<Inputs.FirewallSubnetMappingArgs>());
            set => _subnetMappings = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public FirewallArgs()
        {
        }
        public static new FirewallArgs Empty => new FirewallArgs();
    }
}
