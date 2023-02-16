// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpenSearchServerless
{
    /// <summary>
    /// Amazon OpenSearchServerless access policy resource
    /// </summary>
    [AwsNativeResourceType("aws-native:opensearchserverless:AccessPolicy")]
    public partial class AccessPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the policy
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The JSON policy document that is the content for the policy
        /// </summary>
        [Output("policy")]
        public Output<string?> Policy { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.AwsNative.OpenSearchServerless.AccessPolicyType?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPolicy(string name, AccessPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:opensearchserverless:AccessPolicy", name, args ?? new AccessPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:opensearchserverless:AccessPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPolicy(name, id, options);
        }
    }

    public sealed class AccessPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The JSON policy document that is the content for the policy
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("type")]
        public Input<Pulumi.AwsNative.OpenSearchServerless.AccessPolicyType>? Type { get; set; }

        public AccessPolicyArgs()
        {
        }
        public static new AccessPolicyArgs Empty => new AccessPolicyArgs();
    }
}
