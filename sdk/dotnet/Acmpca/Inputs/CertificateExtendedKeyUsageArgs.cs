// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Inputs
{

    public sealed class CertificateExtendedKeyUsageArgs : global::Pulumi.ResourceArgs
    {
        [Input("extendedKeyUsageObjectIdentifier")]
        public Input<string>? ExtendedKeyUsageObjectIdentifier { get; set; }

        [Input("extendedKeyUsageType")]
        public Input<string>? ExtendedKeyUsageType { get; set; }

        public CertificateExtendedKeyUsageArgs()
        {
        }
        public static new CertificateExtendedKeyUsageArgs Empty => new CertificateExtendedKeyUsageArgs();
    }
}
