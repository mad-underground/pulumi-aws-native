// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Amplify
{
    /// <summary>
    /// The AWS::Amplify::Domain resource allows you to connect a custom domain to your app.
    /// </summary>
    [AwsNativeResourceType("aws-native:amplify:Domain")]
    public partial class Domain : global::Pulumi.CustomResource
    {
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("autoSubDomainCreationPatterns")]
        public Output<ImmutableArray<string>> AutoSubDomainCreationPatterns { get; private set; } = null!;

        [Output("autoSubDomainIamRole")]
        public Output<string?> AutoSubDomainIamRole { get; private set; } = null!;

        [Output("certificate")]
        public Output<Outputs.DomainCertificate?> Certificate { get; private set; } = null!;

        [Output("certificateRecord")]
        public Output<string> CertificateRecord { get; private set; } = null!;

        [Output("certificateSettings")]
        public Output<Outputs.DomainCertificateSettings?> CertificateSettings { get; private set; } = null!;

        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        [Output("domainStatus")]
        public Output<string> DomainStatus { get; private set; } = null!;

        [Output("enableAutoSubDomain")]
        public Output<bool?> EnableAutoSubDomain { get; private set; } = null!;

        [Output("statusReason")]
        public Output<string> StatusReason { get; private set; } = null!;

        [Output("subDomainSettings")]
        public Output<ImmutableArray<Outputs.DomainSubDomainSetting>> SubDomainSettings { get; private set; } = null!;

        [Output("updateStatus")]
        public Output<string?> UpdateStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("aws-native:amplify:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:amplify:Domain", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "appId",
                    "domainName",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, options);
        }
    }

    public sealed class DomainArgs : global::Pulumi.ResourceArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("autoSubDomainCreationPatterns")]
        private InputList<string>? _autoSubDomainCreationPatterns;
        public InputList<string> AutoSubDomainCreationPatterns
        {
            get => _autoSubDomainCreationPatterns ?? (_autoSubDomainCreationPatterns = new InputList<string>());
            set => _autoSubDomainCreationPatterns = value;
        }

        [Input("autoSubDomainIamRole")]
        public Input<string>? AutoSubDomainIamRole { get; set; }

        [Input("certificate")]
        public Input<Inputs.DomainCertificateArgs>? Certificate { get; set; }

        [Input("certificateSettings")]
        public Input<Inputs.DomainCertificateSettingsArgs>? CertificateSettings { get; set; }

        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        [Input("enableAutoSubDomain")]
        public Input<bool>? EnableAutoSubDomain { get; set; }

        [Input("subDomainSettings", required: true)]
        private InputList<Inputs.DomainSubDomainSettingArgs>? _subDomainSettings;
        public InputList<Inputs.DomainSubDomainSettingArgs> SubDomainSettings
        {
            get => _subDomainSettings ?? (_subDomainSettings = new InputList<Inputs.DomainSubDomainSettingArgs>());
            set => _subDomainSettings = value;
        }

        [Input("updateStatus")]
        public Input<string>? UpdateStatus { get; set; }

        public DomainArgs()
        {
        }
        public static new DomainArgs Empty => new DomainArgs();
    }
}
