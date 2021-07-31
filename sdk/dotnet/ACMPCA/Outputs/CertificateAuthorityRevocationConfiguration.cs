// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-revocationconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class CertificateAuthorityRevocationConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-acmpca-certificateauthority-revocationconfiguration.html#cfn-acmpca-certificateauthority-revocationconfiguration-crlconfiguration
        /// </summary>
        public readonly Outputs.CertificateAuthorityCrlConfiguration? CrlConfiguration;

        [OutputConstructor]
        private CertificateAuthorityRevocationConfiguration(Outputs.CertificateAuthorityCrlConfiguration? crlConfiguration)
        {
            CrlConfiguration = crlConfiguration;
        }
    }
}
