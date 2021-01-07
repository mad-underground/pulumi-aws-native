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
    public sealed class ClientVpnEndpointClientAuthenticationRequest
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-activedirectory
        /// </summary>
        public readonly Outputs.ClientVpnEndpointDirectoryServiceAuthenticationRequest? ActiveDirectory;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-federatedauthentication
        /// </summary>
        public readonly Outputs.ClientVpnEndpointFederatedAuthenticationRequest? FederatedAuthentication;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-mutualauthentication
        /// </summary>
        public readonly Outputs.ClientVpnEndpointCertificateAuthenticationRequest? MutualAuthentication;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-clientvpnendpoint-clientauthenticationrequest.html#cfn-ec2-clientvpnendpoint-clientauthenticationrequest-type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ClientVpnEndpointClientAuthenticationRequest(
            Outputs.ClientVpnEndpointDirectoryServiceAuthenticationRequest? ActiveDirectory,

            Outputs.ClientVpnEndpointFederatedAuthenticationRequest? FederatedAuthentication,

            Outputs.ClientVpnEndpointCertificateAuthenticationRequest? MutualAuthentication,

            string Type)
        {
            this.ActiveDirectory = ActiveDirectory;
            this.FederatedAuthentication = FederatedAuthentication;
            this.MutualAuthentication = MutualAuthentication;
            this.Type = Type;
        }
    }
}