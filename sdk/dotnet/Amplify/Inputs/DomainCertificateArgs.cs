// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Amplify.Inputs
{

    public sealed class DomainCertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        [Input("certificateType")]
        public Input<Pulumi.AwsNative.Amplify.DomainCertificateCertificateType>? CertificateType { get; set; }

        [Input("certificateVerificationDnsRecord")]
        public Input<string>? CertificateVerificationDnsRecord { get; set; }

        public DomainCertificateArgs()
        {
        }
        public static new DomainCertificateArgs Empty => new DomainCertificateArgs();
    }
}
