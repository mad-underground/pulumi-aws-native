// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualGatewayListenerTlsCertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("acm")]
        public Input<Inputs.VirtualGatewayListenerTlsAcmCertificateArgs>? Acm { get; set; }

        [Input("file")]
        public Input<Inputs.VirtualGatewayListenerTlsFileCertificateArgs>? File { get; set; }

        [Input("sds")]
        public Input<Inputs.VirtualGatewayListenerTlsSdsCertificateArgs>? Sds { get; set; }

        public VirtualGatewayListenerTlsCertificateArgs()
        {
        }
        public static new VirtualGatewayListenerTlsCertificateArgs Empty => new VirtualGatewayListenerTlsCertificateArgs();
    }
}
