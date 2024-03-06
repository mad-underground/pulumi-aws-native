// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CertificateManager
{
    /// <summary>
    /// Resource Type definition for AWS::CertificateManager::Certificate
    /// </summary>
    [Obsolete(@"Certificate is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:certificatemanager:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("certificateAuthorityArn")]
        public Output<string?> CertificateAuthorityArn { get; private set; } = null!;

        [Output("certificateTransparencyLoggingPreference")]
        public Output<string?> CertificateTransparencyLoggingPreference { get; private set; } = null!;

        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        [Output("domainValidationOptions")]
        public Output<ImmutableArray<Outputs.CertificateDomainValidationOption>> DomainValidationOptions { get; private set; } = null!;

        [Output("keyAlgorithm")]
        public Output<string?> KeyAlgorithm { get; private set; } = null!;

        [Output("subjectAlternativeNames")]
        public Output<ImmutableArray<string>> SubjectAlternativeNames { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("validationMethod")]
        public Output<string?> ValidationMethod { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:certificatemanager:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:certificatemanager:Certificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "certificateAuthorityArn",
                    "domainName",
                    "domainValidationOptions[*]",
                    "keyAlgorithm",
                    "subjectAlternativeNames[*]",
                    "validationMethod",
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
        [Input("certificateAuthorityArn")]
        public Input<string>? CertificateAuthorityArn { get; set; }

        [Input("certificateTransparencyLoggingPreference")]
        public Input<string>? CertificateTransparencyLoggingPreference { get; set; }

        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        [Input("domainValidationOptions")]
        private InputList<Inputs.CertificateDomainValidationOptionArgs>? _domainValidationOptions;
        public InputList<Inputs.CertificateDomainValidationOptionArgs> DomainValidationOptions
        {
            get => _domainValidationOptions ?? (_domainValidationOptions = new InputList<Inputs.CertificateDomainValidationOptionArgs>());
            set => _domainValidationOptions = value;
        }

        [Input("keyAlgorithm")]
        public Input<string>? KeyAlgorithm { get; set; }

        [Input("subjectAlternativeNames")]
        private InputList<string>? _subjectAlternativeNames;
        public InputList<string> SubjectAlternativeNames
        {
            get => _subjectAlternativeNames ?? (_subjectAlternativeNames = new InputList<string>());
            set => _subjectAlternativeNames = value;
        }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("validationMethod")]
        public Input<string>? ValidationMethod { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }
}
