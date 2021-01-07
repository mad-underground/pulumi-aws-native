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
    public sealed class ClientVpnEndpointProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-authenticationoptions
        /// </summary>
        public readonly ImmutableArray<Outputs.ClientVpnEndpointClientAuthenticationRequest> AuthenticationOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientcidrblock
        /// </summary>
        public readonly string ClientCidrBlock;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-clientconnectoptions
        /// </summary>
        public readonly Outputs.ClientVpnEndpointClientConnectOptions? ClientConnectOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-connectionlogoptions
        /// </summary>
        public readonly Outputs.ClientVpnEndpointConnectionLogOptions ConnectionLogOptions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-dnsservers
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-securitygroupids
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-selfserviceportal
        /// </summary>
        public readonly string? SelfServicePortal;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-servercertificatearn
        /// </summary>
        public readonly string ServerCertificateArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-splittunnel
        /// </summary>
        public readonly bool? SplitTunnel;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-tagspecifications
        /// </summary>
        public readonly ImmutableArray<Outputs.ClientVpnEndpointTagSpecification> TagSpecifications;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-transportprotocol
        /// </summary>
        public readonly string? TransportProtocol;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-vpcid
        /// </summary>
        public readonly string? VpcId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpnendpoint.html#cfn-ec2-clientvpnendpoint-vpnport
        /// </summary>
        public readonly int? VpnPort;

        [OutputConstructor]
        private ClientVpnEndpointProperties(
            ImmutableArray<Outputs.ClientVpnEndpointClientAuthenticationRequest> AuthenticationOptions,

            string ClientCidrBlock,

            Outputs.ClientVpnEndpointClientConnectOptions? ClientConnectOptions,

            Outputs.ClientVpnEndpointConnectionLogOptions ConnectionLogOptions,

            string? Description,

            ImmutableArray<string> DnsServers,

            ImmutableArray<string> SecurityGroupIds,

            string? SelfServicePortal,

            string ServerCertificateArn,

            bool? SplitTunnel,

            ImmutableArray<Outputs.ClientVpnEndpointTagSpecification> TagSpecifications,

            string? TransportProtocol,

            string? VpcId,

            int? VpnPort)
        {
            this.AuthenticationOptions = AuthenticationOptions;
            this.ClientCidrBlock = ClientCidrBlock;
            this.ClientConnectOptions = ClientConnectOptions;
            this.ConnectionLogOptions = ConnectionLogOptions;
            this.Description = Description;
            this.DnsServers = DnsServers;
            this.SecurityGroupIds = SecurityGroupIds;
            this.SelfServicePortal = SelfServicePortal;
            this.ServerCertificateArn = ServerCertificateArn;
            this.SplitTunnel = SplitTunnel;
            this.TagSpecifications = TagSpecifications;
            this.TransportProtocol = TransportProtocol;
            this.VpcId = VpcId;
            this.VpnPort = VpnPort;
        }
    }
}