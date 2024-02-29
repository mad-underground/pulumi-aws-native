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
    /// Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load Balancer.
    /// </summary>
    [AwsNativeResourceType("aws-native:elasticloadbalancingv2:LoadBalancer")]
    public partial class LoadBalancer : global::Pulumi.CustomResource
    {
        [Output("canonicalHostedZoneId")]
        public Output<string> CanonicalHostedZoneId { get; private set; } = null!;

        [Output("dnsName")]
        public Output<string> DnsName { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.
        /// </summary>
        [Output("enforceSecurityGroupInboundRulesOnPrivateLinkTraffic")]
        public Output<string?> EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic { get; private set; } = null!;

        /// <summary>
        /// The IP address type. The possible values are ``ipv4`` (for IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses). You can’t specify ``dualstack`` for a load balancer with a UDP or TCP_UDP listener.
        /// </summary>
        [Output("ipAddressType")]
        public Output<string?> IpAddressType { get; private set; } = null!;

        [Output("loadBalancerArn")]
        public Output<string> LoadBalancerArn { get; private set; } = null!;

        /// <summary>
        /// The load balancer attributes.
        /// </summary>
        [Output("loadBalancerAttributes")]
        public Output<ImmutableArray<Outputs.LoadBalancerAttribute>> LoadBalancerAttributes { get; private set; } = null!;

        [Output("loadBalancerFullName")]
        public Output<string> LoadBalancerFullName { get; private set; } = null!;

        [Output("loadBalancerName")]
        public Output<string> LoadBalancerName { get; private set; } = null!;

        /// <summary>
        /// The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
        ///  If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
        ///  The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
        ///  The default is an Internet-facing load balancer.
        ///  You cannot specify a scheme for a Gateway Load Balancer.
        /// </summary>
        [Output("scheme")]
        public Output<string?> Scheme { get; private set; } = null!;

        /// <summary>
        /// [Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.
        /// </summary>
        [Output("securityGroups")]
        public Output<ImmutableArray<string>> SecurityGroups { get; private set; } = null!;

        /// <summary>
        /// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
        ///  [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.
        ///  [Application Load Balancers on Outposts] You must specify one Outpost subnet.
        ///  [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
        ///  [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.
        ///  [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP
        /// </summary>
        [Output("subnetMappings")]
        public Output<ImmutableArray<Outputs.LoadBalancerSubnetMapping>> SubnetMappings { get; private set; } = null!;

        /// <summary>
        /// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
        ///  [Application Load Balancers] You must specify subnets from at least two Availability Zones.
        ///  [Application Load Balancers on Outposts] You must specify one Outpost subnet.
        ///  [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
        ///  [Network Load Balancers] You can specify subnets from one or more Availability Zones.
        ///  [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;

        /// <summary>
        /// The tags to assign to the load balancer.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of load balancer. The default is ``application``.
        /// </summary>
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
                ReplaceOnChanges =
                {
                    "name",
                    "scheme",
                    "type",
                },
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
        /// <summary>
        /// Indicates whether to evaluate inbound security group rules for traffic sent to a Network Load Balancer through privatelink.
        /// </summary>
        [Input("enforceSecurityGroupInboundRulesOnPrivateLinkTraffic")]
        public Input<string>? EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic { get; set; }

        /// <summary>
        /// The IP address type. The possible values are ``ipv4`` (for IPv4 addresses) and ``dualstack`` (for IPv4 and IPv6 addresses). You can’t specify ``dualstack`` for a load balancer with a UDP or TCP_UDP listener.
        /// </summary>
        [Input("ipAddressType")]
        public Input<string>? IpAddressType { get; set; }

        [Input("loadBalancerAttributes")]
        private InputList<Inputs.LoadBalancerAttributeArgs>? _loadBalancerAttributes;

        /// <summary>
        /// The load balancer attributes.
        /// </summary>
        public InputList<Inputs.LoadBalancerAttributeArgs> LoadBalancerAttributes
        {
            get => _loadBalancerAttributes ?? (_loadBalancerAttributes = new InputList<Inputs.LoadBalancerAttributeArgs>());
            set => _loadBalancerAttributes = value;
        }

        /// <summary>
        /// The name of the load balancer. This name must be unique per region per account, can have a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not begin or end with a hyphen, and must not begin with "internal-".
        ///  If you don't specify a name, AWS CloudFormation generates a unique physical ID for the load balancer. If you specify a name, you cannot perform updates that require replacement of this resource, but you can perform other updates. To replace the resource, specify a new name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes. Therefore, Internet-facing load balancers can route requests from clients over the internet.
        ///  The nodes of an internal load balancer have only private IP addresses. The DNS name of an internal load balancer is publicly resolvable to the private IP addresses of the nodes. Therefore, internal load balancers can route requests only from clients with access to the VPC for the load balancer.
        ///  The default is an Internet-facing load balancer.
        ///  You cannot specify a scheme for a Gateway Load Balancer.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;

        /// <summary>
        /// [Application Load Balancers and Network Load Balancers] The IDs of the security groups for the load balancer.
        /// </summary>
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("subnetMappings")]
        private InputList<Inputs.LoadBalancerSubnetMappingArgs>? _subnetMappings;

        /// <summary>
        /// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both.
        ///  [Application Load Balancers] You must specify subnets from at least two Availability Zones. You cannot specify Elastic IP addresses for your subnets.
        ///  [Application Load Balancers on Outposts] You must specify one Outpost subnet.
        ///  [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
        ///  [Network Load Balancers] You can specify subnets from one or more Availability Zones. You can specify one Elastic IP address per subnet if you need static IP addresses for your internet-facing load balancer. For internal load balancers, you can specify one private IP address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you can specify one IPv6 address per subnet.
        ///  [Gateway Load Balancers] You can specify subnets from one or more Availability Zones. You cannot specify Elastic IP
        /// </summary>
        public InputList<Inputs.LoadBalancerSubnetMappingArgs> SubnetMappings
        {
            get => _subnetMappings ?? (_subnetMappings = new InputList<Inputs.LoadBalancerSubnetMappingArgs>());
            set => _subnetMappings = value;
        }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// The IDs of the public subnets. You can specify only one subnet per Availability Zone. You must specify either subnets or subnet mappings, but not both. To specify an Elastic IP address, specify subnet mappings instead of subnets.
        ///  [Application Load Balancers] You must specify subnets from at least two Availability Zones.
        ///  [Application Load Balancers on Outposts] You must specify one Outpost subnet.
        ///  [Application Load Balancers on Local Zones] You can specify subnets from one or more Local Zones.
        ///  [Network Load Balancers] You can specify subnets from one or more Availability Zones.
        ///  [Gateway Load Balancers] You can specify subnets from one or more Availability Zones.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// The tags to assign to the load balancer.
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of load balancer. The default is ``application``.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public LoadBalancerArgs()
        {
        }
        public static new LoadBalancerArgs Empty => new LoadBalancerArgs();
    }
}
