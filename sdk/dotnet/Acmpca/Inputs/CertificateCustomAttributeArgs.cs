// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Acmpca.Inputs
{

    public sealed class CertificateCustomAttributeArgs : global::Pulumi.ResourceArgs
    {
        [Input("objectIdentifier", required: true)]
        public Input<string> ObjectIdentifier { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public CertificateCustomAttributeArgs()
        {
        }
        public static new CertificateCustomAttributeArgs Empty => new CertificateCustomAttributeArgs();
    }
}
