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
    public sealed class LaunchTemplateNetworkInterface
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-associatecarrieripaddress
        /// </summary>
        public readonly bool? AssociateCarrierIpAddress;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-associatepublicipaddress
        /// </summary>
        public readonly bool? AssociatePublicIpAddress;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-deleteontermination
        /// </summary>
        public readonly bool? DeleteOnTermination;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-deviceindex
        /// </summary>
        public readonly int? DeviceIndex;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-groups
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-interfacetype
        /// </summary>
        public readonly string? InterfaceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv6addresscount
        /// </summary>
        public readonly int? Ipv6AddressCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv6addresses
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchTemplateIpv6Add> Ipv6Addresses;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-networkcardindex
        /// </summary>
        public readonly int? NetworkCardIndex;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-networkinterfaceid
        /// </summary>
        public readonly string? NetworkInterfaceId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-privateipaddress
        /// </summary>
        public readonly string? PrivateIpAddress;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-privateipaddresses
        /// </summary>
        public readonly ImmutableArray<Outputs.LaunchTemplatePrivateIpAdd> PrivateIpAddresses;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-secondaryprivateipaddresscount
        /// </summary>
        public readonly int? SecondaryPrivateIpAddressCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-subnetid
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private LaunchTemplateNetworkInterface(
            bool? AssociateCarrierIpAddress,

            bool? AssociatePublicIpAddress,

            bool? DeleteOnTermination,

            string? Description,

            int? DeviceIndex,

            ImmutableArray<string> Groups,

            string? InterfaceType,

            int? Ipv6AddressCount,

            ImmutableArray<Outputs.LaunchTemplateIpv6Add> Ipv6Addresses,

            int? NetworkCardIndex,

            string? NetworkInterfaceId,

            string? PrivateIpAddress,

            ImmutableArray<Outputs.LaunchTemplatePrivateIpAdd> PrivateIpAddresses,

            int? SecondaryPrivateIpAddressCount,

            string? SubnetId)
        {
            this.AssociateCarrierIpAddress = AssociateCarrierIpAddress;
            this.AssociatePublicIpAddress = AssociatePublicIpAddress;
            this.DeleteOnTermination = DeleteOnTermination;
            this.Description = Description;
            this.DeviceIndex = DeviceIndex;
            this.Groups = Groups;
            this.InterfaceType = InterfaceType;
            this.Ipv6AddressCount = Ipv6AddressCount;
            this.Ipv6Addresses = Ipv6Addresses;
            this.NetworkCardIndex = NetworkCardIndex;
            this.NetworkInterfaceId = NetworkInterfaceId;
            this.PrivateIpAddress = PrivateIpAddress;
            this.PrivateIpAddresses = PrivateIpAddresses;
            this.SecondaryPrivateIpAddressCount = SecondaryPrivateIpAddressCount;
            this.SubnetId = SubnetId;
        }
    }
}
