// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EC2.Outputs
{

    [OutputType]
    public sealed class NetworkInterfaceProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-groupset
        /// </summary>
        public readonly ImmutableArray<string> GroupSet;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-interfacetype
        /// </summary>
        public readonly string? InterfaceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-ipv6addresscount
        /// </summary>
        public readonly int? Ipv6AddressCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-ec2-networkinterface-ipv6addresses
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceInstanceIpv6Address> Ipv6Addresses;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-privateipaddress
        /// </summary>
        public readonly string? PrivateIpAddress;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-privateipaddresses
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfacePrivateIpAddressSpecification> PrivateIpAddresses;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-secondaryprivateipcount
        /// </summary>
        public readonly int? SecondaryPrivateIpAddressCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-sourcedestcheck
        /// </summary>
        public readonly bool? SourceDestCheck;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-subnetid
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#cfn-awsec2networkinterface-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private NetworkInterfaceProperties(
            string? Description,

            ImmutableArray<string> GroupSet,

            string? InterfaceType,

            int? Ipv6AddressCount,

            ImmutableArray<Outputs.NetworkInterfaceInstanceIpv6Address> Ipv6Addresses,

            string? PrivateIpAddress,

            ImmutableArray<Outputs.NetworkInterfacePrivateIpAddressSpecification> PrivateIpAddresses,

            int? SecondaryPrivateIpAddressCount,

            bool? SourceDestCheck,

            string SubnetId,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.Description = Description;
            this.GroupSet = GroupSet;
            this.InterfaceType = InterfaceType;
            this.Ipv6AddressCount = Ipv6AddressCount;
            this.Ipv6Addresses = Ipv6Addresses;
            this.PrivateIpAddress = PrivateIpAddress;
            this.PrivateIpAddresses = PrivateIpAddresses;
            this.SecondaryPrivateIpAddressCount = SecondaryPrivateIpAddressCount;
            this.SourceDestCheck = SourceDestCheck;
            this.SubnetId = SubnetId;
            this.Tags = Tags;
        }
    }
}
