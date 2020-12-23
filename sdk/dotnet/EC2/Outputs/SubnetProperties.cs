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
    public sealed class SubnetProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-assignipv6addressoncreation
        /// </summary>
        public readonly bool? AssignIpv6AddressOnCreation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-availabilityzone
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-cidrblock
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv6cidrblock
        /// </summary>
        public readonly string? Ipv6CidrBlock;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-mappubliciponlaunch
        /// </summary>
        public readonly bool? MapPublicIpOnLaunch;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-outpostarn
        /// </summary>
        public readonly string? OutpostArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-awsec2subnet-prop-vpcid
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private SubnetProperties(
            bool? AssignIpv6AddressOnCreation,

            string? AvailabilityZone,

            string CidrBlock,

            string? Ipv6CidrBlock,

            bool? MapPublicIpOnLaunch,

            string? OutpostArn,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags,

            string VpcId)
        {
            this.AssignIpv6AddressOnCreation = AssignIpv6AddressOnCreation;
            this.AvailabilityZone = AvailabilityZone;
            this.CidrBlock = CidrBlock;
            this.Ipv6CidrBlock = Ipv6CidrBlock;
            this.MapPublicIpOnLaunch = MapPublicIpOnLaunch;
            this.OutpostArn = OutpostArn;
            this.Tags = Tags;
            this.VpcId = VpcId;
        }
    }
}
