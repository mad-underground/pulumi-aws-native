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
    public sealed class VirtualGatewayVirtualGatewayTlsValidationContextTrust
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust-acm
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayTlsValidationContextAcmTrust? ACM;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust.html#cfn-appmesh-virtualgateway-virtualgatewaytlsvalidationcontexttrust-file
        /// </summary>
        public readonly Outputs.VirtualGatewayVirtualGatewayTlsValidationContextFileTrust? File;

        [OutputConstructor]
        private VirtualGatewayVirtualGatewayTlsValidationContextTrust(
            Outputs.VirtualGatewayVirtualGatewayTlsValidationContextAcmTrust? ACM,

            Outputs.VirtualGatewayVirtualGatewayTlsValidationContextFileTrust? File)
        {
            this.ACM = ACM;
            this.File = File;
        }
    }
}