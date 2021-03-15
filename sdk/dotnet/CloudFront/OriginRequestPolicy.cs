// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html
    /// </summary>
    [AwsNativeResourceType("aws-native:CloudFront:OriginRequestPolicy")]
    public partial class OriginRequestPolicy : Pulumi.CustomResource
    {
        [Output("id")]
        public Output<string> Id { get; private set; } = null!;

        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig
        /// </summary>
        [Output("originRequestPolicyConfig")]
        public Output<Outputs.OriginRequestPolicyOriginRequestPolicyConfig> OriginRequestPolicyConfig { get; private set; } = null!;


        /// <summary>
        /// Create a OriginRequestPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OriginRequestPolicy(string name, OriginRequestPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:CloudFront:OriginRequestPolicy", name, args ?? new OriginRequestPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OriginRequestPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:CloudFront:OriginRequestPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing OriginRequestPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OriginRequestPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OriginRequestPolicy(name, id, options);
        }
    }

    public sealed class OriginRequestPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originrequestpolicy.html#cfn-cloudfront-originrequestpolicy-originrequestpolicyconfig
        /// </summary>
        [Input("originRequestPolicyConfig", required: true)]
        public Input<Inputs.OriginRequestPolicyOriginRequestPolicyConfigArgs> OriginRequestPolicyConfig { get; set; } = null!;

        public OriginRequestPolicyArgs()
        {
        }
    }
}
