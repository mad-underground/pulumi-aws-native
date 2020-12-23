// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualGatewayVirtualGatewayListenerTlsFileCertificate
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate-certificatechain
        /// </summary>
        public readonly string CertificateChain;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate.html#cfn-appmesh-virtualgateway-virtualgatewaylistenertlsfilecertificate-privatekey
        /// </summary>
        public readonly string PrivateKey;

        [OutputConstructor]
        private VirtualGatewayVirtualGatewayListenerTlsFileCertificate(
            string CertificateChain,

            string PrivateKey)
        {
            this.CertificateChain = CertificateChain;
            this.PrivateKey = PrivateKey;
        }
    }
}
