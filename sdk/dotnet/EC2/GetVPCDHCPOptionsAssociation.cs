// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetVPCDHCPOptionsAssociation
    {
        /// <summary>
        /// Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.
        /// </summary>
        public static Task<GetVPCDHCPOptionsAssociationResult> InvokeAsync(GetVPCDHCPOptionsAssociationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVPCDHCPOptionsAssociationResult>("aws-native:ec2:getVPCDHCPOptionsAssociation", args ?? new GetVPCDHCPOptionsAssociationArgs(), options.WithDefaults());

        /// <summary>
        /// Associates a set of DHCP options with a VPC, or associates no DHCP options with the VPC.
        /// </summary>
        public static Output<GetVPCDHCPOptionsAssociationResult> Invoke(GetVPCDHCPOptionsAssociationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVPCDHCPOptionsAssociationResult>("aws-native:ec2:getVPCDHCPOptionsAssociation", args ?? new GetVPCDHCPOptionsAssociationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVPCDHCPOptionsAssociationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
        /// </summary>
        [Input("dhcpOptionsId", required: true)]
        public string DhcpOptionsId { get; set; } = null!;

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public string VpcId { get; set; } = null!;

        public GetVPCDHCPOptionsAssociationArgs()
        {
        }
        public static new GetVPCDHCPOptionsAssociationArgs Empty => new GetVPCDHCPOptionsAssociationArgs();
    }

    public sealed class GetVPCDHCPOptionsAssociationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the DHCP options set, or default to associate no DHCP options with the VPC.
        /// </summary>
        [Input("dhcpOptionsId", required: true)]
        public Input<string> DhcpOptionsId { get; set; } = null!;

        /// <summary>
        /// The ID of the VPC.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public GetVPCDHCPOptionsAssociationInvokeArgs()
        {
        }
        public static new GetVPCDHCPOptionsAssociationInvokeArgs Empty => new GetVPCDHCPOptionsAssociationInvokeArgs();
    }


    [OutputType]
    public sealed class GetVPCDHCPOptionsAssociationResult
    {
        /// <summary>
        /// The ID of the VPC DHCPOptions Association.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private GetVPCDHCPOptionsAssociationResult(string? id)
        {
            Id = id;
        }
    }
}
