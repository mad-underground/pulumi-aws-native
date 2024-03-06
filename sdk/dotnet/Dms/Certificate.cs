// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Dms
{
    /// <summary>
    /// Resource Type definition for AWS::DMS::Certificate
    /// </summary>
    [Obsolete(@"Certificate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:dms:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("certificateIdentifier")]
        public Output<string?> CertificateIdentifier { get; private set; } = null!;

        [Output("certificatePem")]
        public Output<string?> CertificatePem { get; private set; } = null!;

        [Output("certificateWallet")]
        public Output<string?> CertificateWallet { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:dms:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:dms:Certificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "certificateIdentifier",
                    "certificatePem",
                    "certificateWallet",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, options);
        }
    }

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificateIdentifier")]
        public Input<string>? CertificateIdentifier { get; set; }

        [Input("certificatePem")]
        public Input<string>? CertificatePem { get; set; }

        [Input("certificateWallet")]
        public Input<string>? CertificateWallet { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}
