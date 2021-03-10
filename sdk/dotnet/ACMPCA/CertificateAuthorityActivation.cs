// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ACMPCA
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html
    /// </summary>
    [AwsNativeResourceType("aws-native:ACMPCA:CertificateAuthorityActivation")]
    public partial class CertificateAuthorityActivation : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificate
        /// </summary>
        [Output("Certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificateauthorityarn
        /// </summary>
        [Output("CertificateAuthorityArn")]
        public Output<string> CertificateAuthorityArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificatechain
        /// </summary>
        [Output("CertificateChain")]
        public Output<string?> CertificateChain { get; private set; } = null!;

        [Output("CompleteCertificateChain")]
        public Output<string> CompleteCertificateChain { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-status
        /// </summary>
        [Output("Status")]
        public Output<string?> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateAuthorityActivation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateAuthorityActivation(string name, CertificateAuthorityActivationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ACMPCA:CertificateAuthorityActivation", name, args ?? new CertificateAuthorityActivationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateAuthorityActivation(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ACMPCA:CertificateAuthorityActivation", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CertificateAuthorityActivation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateAuthorityActivation Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CertificateAuthorityActivation(name, id, options);
        }
    }

    public sealed class CertificateAuthorityActivationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificate
        /// </summary>
        [Input("Certificate", required: true)]
        public Input<string> Certificate { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificateauthorityarn
        /// </summary>
        [Input("CertificateAuthorityArn", required: true)]
        public Input<string> CertificateAuthorityArn { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-certificatechain
        /// </summary>
        [Input("CertificateChain")]
        public Input<string>? CertificateChain { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-acmpca-certificateauthorityactivation.html#cfn-acmpca-certificateauthorityactivation-status
        /// </summary>
        [Input("Status")]
        public Input<string>? Status { get; set; }

        public CertificateAuthorityActivationArgs()
        {
        }
    }
}
