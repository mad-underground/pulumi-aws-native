// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Inputs
{

    public sealed class CertificatePolicyQualifierInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("policyQualifierId", required: true)]
        public Input<string> PolicyQualifierId { get; set; } = null!;

        [Input("qualifier", required: true)]
        public Input<Inputs.CertificateQualifierArgs> Qualifier { get; set; } = null!;

        public CertificatePolicyQualifierInfoArgs()
        {
        }
        public static new CertificatePolicyQualifierInfoArgs Empty => new CertificatePolicyQualifierInfoArgs();
    }
}
