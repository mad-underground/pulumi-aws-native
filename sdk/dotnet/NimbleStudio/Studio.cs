// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio
{
    /// <summary>
    /// Represents a studio that contains other Nimble Studio resources
    /// </summary>
    [AwsNativeResourceType("aws-native:nimblestudio:Studio")]
    public partial class Studio : global::Pulumi.CustomResource
    {
        /// <summary>
        /// &lt;p&gt;The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.&lt;/p&gt;
        /// </summary>
        [Output("adminRoleArn")]
        public Output<string> AdminRoleArn { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;A friendly name for the studio.&lt;/p&gt;
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The Amazon Web Services Region where the studio resource is located.&lt;/p&gt;
        /// </summary>
        [Output("homeRegion")]
        public Output<string> HomeRegion { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The Amazon Web Services SSO application client ID used to integrate with Amazon Web Services SSO to enable Amazon Web Services SSO users to log in to Nimble Studio portal.&lt;/p&gt;
        /// </summary>
        [Output("ssoClientId")]
        public Output<string> SsoClientId { get; private set; } = null!;

        [Output("studioEncryptionConfiguration")]
        public Output<Outputs.StudioEncryptionConfiguration?> StudioEncryptionConfiguration { get; private set; } = null!;

        [Output("studioId")]
        public Output<string> StudioId { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.&lt;/p&gt;
        /// </summary>
        [Output("studioName")]
        public Output<string> StudioName { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The address of the web page for the studio.&lt;/p&gt;
        /// </summary>
        [Output("studioUrl")]
        public Output<string> StudioUrl { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.&lt;/p&gt;
        /// </summary>
        [Output("userRoleArn")]
        public Output<string> UserRoleArn { get; private set; } = null!;


        /// <summary>
        /// Create a Studio resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Studio(string name, StudioArgs args, CustomResourceOptions? options = null)
            : base("aws-native:nimblestudio:Studio", name, args ?? new StudioArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Studio(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:nimblestudio:Studio", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "studioName",
                    "tags.*",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Studio resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Studio Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Studio(name, id, options);
        }
    }

    public sealed class StudioArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The IAM role that Studio Admins will assume when logging in to the Nimble Studio portal.&lt;/p&gt;
        /// </summary>
        [Input("adminRoleArn", required: true)]
        public Input<string> AdminRoleArn { get; set; } = null!;

        /// <summary>
        /// &lt;p&gt;A friendly name for the studio.&lt;/p&gt;
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        [Input("studioEncryptionConfiguration")]
        public Input<Inputs.StudioEncryptionConfigurationArgs>? StudioEncryptionConfiguration { get; set; }

        /// <summary>
        /// &lt;p&gt;The studio name that is used in the URL of the Nimble Studio portal when accessed by Nimble Studio users.&lt;/p&gt;
        /// </summary>
        [Input("studioName")]
        public Input<string>? StudioName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// &lt;p&gt;The IAM role that Studio Users will assume when logging in to the Nimble Studio portal.&lt;/p&gt;
        /// </summary>
        [Input("userRoleArn", required: true)]
        public Input<string> UserRoleArn { get; set; } = null!;

        public StudioArgs()
        {
        }
        public static new StudioArgs Empty => new StudioArgs();
    }
}
