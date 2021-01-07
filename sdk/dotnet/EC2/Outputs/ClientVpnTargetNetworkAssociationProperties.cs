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
    public sealed class ClientVpnTargetNetworkAssociationProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html#cfn-ec2-clientvpntargetnetworkassociation-clientvpnendpointid
        /// </summary>
        public readonly string ClientVpnEndpointId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html#cfn-ec2-clientvpntargetnetworkassociation-subnetid
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private ClientVpnTargetNetworkAssociationProperties(
            string ClientVpnEndpointId,

            string SubnetId)
        {
            this.ClientVpnEndpointId = ClientVpnEndpointId;
            this.SubnetId = SubnetId;
        }
    }
}