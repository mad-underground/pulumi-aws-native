// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// The AWS::EC2::NetworkInterface resource creates network interface
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:NetworkInterface")]
    public partial class NetworkInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A description for the network interface.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of security group IDs associated with this network interface.
        /// </summary>
        [Output("groupSet")]
        public Output<ImmutableArray<string>> GroupSet { get; private set; } = null!;

        /// <summary>
        /// Indicates the type of network interface.
        /// </summary>
        [Output("interfaceType")]
        public Output<string?> InterfaceType { get; private set; } = null!;

        /// <summary>
        /// The number of IPv4 prefixes to assign to a network interface. When you specify a number of IPv4 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /28 prefixes. You can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        /// </summary>
        [Output("ipv4PrefixCount")]
        public Output<int?> Ipv4PrefixCount { get; private set; } = null!;

        /// <summary>
        /// Assigns a list of IPv4 prefixes to the network interface. If you want EC2 to automatically assign IPv4 prefixes, use the Ipv4PrefixCount property and do not specify this property. Presently, only /28 prefixes are supported. You can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        /// </summary>
        [Output("ipv4Prefixes")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceIpv4PrefixSpecification>> Ipv4Prefixes { get; private set; } = null!;

        /// <summary>
        /// The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
        /// </summary>
        [Output("ipv6AddressCount")]
        public Output<int?> Ipv6AddressCount { get; private set; } = null!;

        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
        /// </summary>
        [Output("ipv6Addresses")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceInstanceIpv6Address>> Ipv6Addresses { get; private set; } = null!;

        /// <summary>
        /// The number of IPv6 prefixes to assign to a network interface. When you specify a number of IPv6 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /80 prefixes. You can't specify a count of IPv6 prefixes if you've specified one of the following: specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        /// </summary>
        [Output("ipv6PrefixCount")]
        public Output<int?> Ipv6PrefixCount { get; private set; } = null!;

        /// <summary>
        /// Assigns a list of IPv6 prefixes to the network interface. If you want EC2 to automatically assign IPv6 prefixes, use the Ipv6PrefixCount property and do not specify this property. Presently, only /80 prefixes are supported. You can't specify IPv6 prefixes if you've specified one of the following: a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        /// </summary>
        [Output("ipv6Prefixes")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceIpv6PrefixSpecification>> Ipv6Prefixes { get; private set; } = null!;

        /// <summary>
        /// Returns the primary private IP address of the network interface.
        /// </summary>
        [Output("primaryPrivateIpAddress")]
        public Output<string> PrimaryPrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. 
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string?> PrivateIpAddress { get; private set; } = null!;

        /// <summary>
        /// Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
        /// </summary>
        [Output("privateIpAddresses")]
        public Output<ImmutableArray<Outputs.NetworkInterfacePrivateIpAddressSpecification>> PrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
        /// </summary>
        [Output("secondaryPrivateIpAddressCount")]
        public Output<int?> SecondaryPrivateIpAddressCount { get; private set; } = null!;

        /// <summary>
        /// Returns the secondary private IP addresses of the network interface.
        /// </summary>
        [Output("secondaryPrivateIpAddresses")]
        public Output<ImmutableArray<string>> SecondaryPrivateIpAddresses { get; private set; } = null!;

        /// <summary>
        /// Indicates whether traffic to or from the instance is validated.
        /// </summary>
        [Output("sourceDestCheck")]
        public Output<bool?> SourceDestCheck { get; private set; } = null!;

        /// <summary>
        /// The ID of the subnet to associate with the network interface.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this network interface.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterface(string name, NetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInterface", name, args ?? new NetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterface(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInterface", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterface Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkInterface(name, id, options);
        }
    }

    public sealed class NetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description for the network interface.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("groupSet")]
        private InputList<string>? _groupSet;

        /// <summary>
        /// A list of security group IDs associated with this network interface.
        /// </summary>
        public InputList<string> GroupSet
        {
            get => _groupSet ?? (_groupSet = new InputList<string>());
            set => _groupSet = value;
        }

        /// <summary>
        /// Indicates the type of network interface.
        /// </summary>
        [Input("interfaceType")]
        public Input<string>? InterfaceType { get; set; }

        /// <summary>
        /// The number of IPv4 prefixes to assign to a network interface. When you specify a number of IPv4 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /28 prefixes. You can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        /// </summary>
        [Input("ipv4PrefixCount")]
        public Input<int>? Ipv4PrefixCount { get; set; }

        [Input("ipv4Prefixes")]
        private InputList<Inputs.NetworkInterfaceIpv4PrefixSpecificationArgs>? _ipv4Prefixes;

        /// <summary>
        /// Assigns a list of IPv4 prefixes to the network interface. If you want EC2 to automatically assign IPv4 prefixes, use the Ipv4PrefixCount property and do not specify this property. Presently, only /28 prefixes are supported. You can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceIpv4PrefixSpecificationArgs> Ipv4Prefixes
        {
            get => _ipv4Prefixes ?? (_ipv4Prefixes = new InputList<Inputs.NetworkInterfaceIpv4PrefixSpecificationArgs>());
            set => _ipv4Prefixes = value;
        }

        /// <summary>
        /// The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
        /// </summary>
        [Input("ipv6AddressCount")]
        public Input<int>? Ipv6AddressCount { get; set; }

        [Input("ipv6Addresses")]
        private InputList<Inputs.NetworkInterfaceInstanceIpv6AddressArgs>? _ipv6Addresses;

        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceInstanceIpv6AddressArgs> Ipv6Addresses
        {
            get => _ipv6Addresses ?? (_ipv6Addresses = new InputList<Inputs.NetworkInterfaceInstanceIpv6AddressArgs>());
            set => _ipv6Addresses = value;
        }

        /// <summary>
        /// The number of IPv6 prefixes to assign to a network interface. When you specify a number of IPv6 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /80 prefixes. You can't specify a count of IPv6 prefixes if you've specified one of the following: specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        /// </summary>
        [Input("ipv6PrefixCount")]
        public Input<int>? Ipv6PrefixCount { get; set; }

        [Input("ipv6Prefixes")]
        private InputList<Inputs.NetworkInterfaceIpv6PrefixSpecificationArgs>? _ipv6Prefixes;

        /// <summary>
        /// Assigns a list of IPv6 prefixes to the network interface. If you want EC2 to automatically assign IPv6 prefixes, use the Ipv6PrefixCount property and do not specify this property. Presently, only /80 prefixes are supported. You can't specify IPv6 prefixes if you've specified one of the following: a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceIpv6PrefixSpecificationArgs> Ipv6Prefixes
        {
            get => _ipv6Prefixes ?? (_ipv6Prefixes = new InputList<Inputs.NetworkInterfaceIpv6PrefixSpecificationArgs>());
            set => _ipv6Prefixes = value;
        }

        /// <summary>
        /// Assigns a single private IP address to the network interface, which is used as the primary private IP address. If you want to specify multiple private IP address, use the PrivateIpAddresses property. 
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("privateIpAddresses")]
        private InputList<Inputs.NetworkInterfacePrivateIpAddressSpecificationArgs>? _privateIpAddresses;

        /// <summary>
        /// Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
        /// </summary>
        public InputList<Inputs.NetworkInterfacePrivateIpAddressSpecificationArgs> PrivateIpAddresses
        {
            get => _privateIpAddresses ?? (_privateIpAddresses = new InputList<Inputs.NetworkInterfacePrivateIpAddressSpecificationArgs>());
            set => _privateIpAddresses = value;
        }

        /// <summary>
        /// The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
        /// </summary>
        [Input("secondaryPrivateIpAddressCount")]
        public Input<int>? SecondaryPrivateIpAddressCount { get; set; }

        /// <summary>
        /// Indicates whether traffic to or from the instance is validated.
        /// </summary>
        [Input("sourceDestCheck")]
        public Input<bool>? SourceDestCheck { get; set; }

        /// <summary>
        /// The ID of the subnet to associate with the network interface.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.NetworkInterfaceTagArgs>? _tags;

        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this network interface.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NetworkInterfaceTagArgs>());
            set => _tags = value;
        }

        public NetworkInterfaceArgs()
        {
        }
        public static new NetworkInterfaceArgs Empty => new NetworkInterfaceArgs();
    }
}
