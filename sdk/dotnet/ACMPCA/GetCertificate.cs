// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA
{
    public static class GetCertificate
    {
        /// <summary>
        /// A certificate issued via a private certificate authority
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("aws-native:acmpca:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// A certificate issued via a private certificate authority
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("aws-native:acmpca:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the issued certificate.
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the private CA to issue the certificate.
        /// </summary>
        [Input("certificateAuthorityArn", required: true)]
        public string CertificateAuthorityArn { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
        public static new GetCertificateArgs Empty => new GetCertificateArgs();
    }

    public sealed class GetCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the issued certificate.
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the private CA to issue the certificate.
        /// </summary>
        [Input("certificateAuthorityArn", required: true)]
        public Input<string> CertificateAuthorityArn { get; set; } = null!;

        public GetCertificateInvokeArgs()
        {
        }
        public static new GetCertificateInvokeArgs Empty => new GetCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// The ARN of the issued certificate.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The issued certificate in base 64 PEM-encoded format.
        /// </summary>
        public readonly string? CertificateValue;

        [OutputConstructor]
        private GetCertificateResult(
            string? arn,

            string? certificate)
        {
            Arn = arn;
            CertificateValue = certificate;
        }
    }
}
