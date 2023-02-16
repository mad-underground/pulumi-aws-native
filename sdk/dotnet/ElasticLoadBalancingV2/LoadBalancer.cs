// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2
{
    /// <summary>
    /// Resource Type definition for AWS::ElasticLoadBalancingV2::LoadBalancer
    /// </summary>
    [Obsolete(@"LoadBalancer is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:elasticloadbalancingv2:LoadBalancer")]
    public partial class LoadBalancer : global::Pulumi.CustomResource
    {
        [Output("canonicalHostedZoneID")]
        public Output<string> CanonicalHostedZoneID { get; private set; } = null!;

        [Output("dNSName")]
        public Output<string> DNSName { get; private set; } = null!;

        [Output("ipAddressType")]
        public Output<string?> IpAddressType { get; private set; } = null!;

        [Output("loadBalancerAttributes")]
        public Output<ImmutableArray<Outputs.LoadBalancerAttribute>> LoadBalancerAttributes { get; private set; } = null!;

        [Output("loadBalancerFullName")]
        public Output<string> LoadBalancerFullName { get; private set; } = null!;

        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("scheme")]
        public Output<string?> Scheme { get; private set; } = null!;

        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        [Output("subnetMappings")]
        public Output<ImmutableArray<Outputs.LoadBalancerSubnetMapping>> SubnetMappings { get; private set; } = null!;

        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.LoadBalancerTag>> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:elasticloadbalancingv2:LoadBalancer", name, args ?? new LoadBalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:elasticloadbalancingv2:LoadBalancer", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, options);
        }
    }

    public sealed class LoadBalancerArgs : global::Pulumi.ResourceArgs
    {
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        [Input("loadBalancerAttributes")]
        private InputList<Inputs.LoadBalancerAttributeArgs>? _loadBalancerAttributes;
        public InputList<Inputs.LoadBalancerAttributeArgs> LoadBalancerAttributes
        {
            get => _loadBalancerAttributes ?? (_loadBalancerAttributes = new InputList<Inputs.LoadBalancerAttributeArgs>());
            set => _loadBalancerAttributes = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnetMappings")]
        private InputList<Inputs.LoadBalancerSubnetMappingArgs>? _subnetMappings;
        public InputList<Inputs.LoadBalancerSubnetMappingArgs> SubnetMappings
        {
            get => _subnetMappings ?? (_subnetMappings = new InputList<Inputs.LoadBalancerSubnetMappingArgs>());
            set => _subnetMappings = value;
        }

        [Input("subnets")]
        private InputList<string>? _subnets;
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputList<Inputs.LoadBalancerTagArgs>? _tags;
        public InputList<Inputs.LoadBalancerTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LoadBalancerTagArgs>());
            set => _tags = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public LoadBalancerArgs()
        {
        }
        public static new LoadBalancerArgs Empty => new LoadBalancerArgs();
    }
}
