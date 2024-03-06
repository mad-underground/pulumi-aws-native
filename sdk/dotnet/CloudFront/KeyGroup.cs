// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    /// <summary>
    /// Resource Type definition for AWS::CloudFront::KeyGroup
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudfront:KeyGroup")]
    public partial class KeyGroup : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("keyGroupConfig")]
        public Output<Outputs.KeyGroupConfig> KeyGroupConfig { get; private set; } = null!;

        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;


        /// <summary>
        /// Create a KeyGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KeyGroup(string name, KeyGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:KeyGroup", name, args ?? new KeyGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KeyGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:KeyGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing KeyGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KeyGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new KeyGroup(name, id, options);
        }
    }

    public sealed class KeyGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("keyGroupConfig", required: true)]
        public Input<Inputs.KeyGroupConfigArgs> KeyGroupConfig { get; set; } = null!;

        public KeyGroupArgs()
        {
        }
        public static new KeyGroupArgs Empty => new KeyGroupArgs();
    }
}
