// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation
{
    /// <summary>
    /// Resource Type definition for AWS::LakeFormation::Resource
    /// </summary>
    [Obsolete(@"Resource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:lakeformation:Resource")]
    public partial class Resource : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("hybridAccessEnabled")]
        public Output<bool?> HybridAccessEnabled { get; private set; } = null!;

        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        [Output("useServiceLinkedRole")]
        public Output<bool> UseServiceLinkedRole { get; private set; } = null!;

        [Output("withFederation")]
        public Output<bool?> WithFederation { get; private set; } = null!;


        /// <summary>
        /// Create a Resource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resource(string name, ResourceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:lakeformation:Resource", name, args ?? new ResourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:lakeformation:Resource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "resourceArn",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Resource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Resource(name, id, options);
        }
    }

    public sealed class ResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("hybridAccessEnabled")]
        public Input<bool>? HybridAccessEnabled { get; set; }

        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("useServiceLinkedRole", required: true)]
        public Input<bool> UseServiceLinkedRole { get; set; } = null!;

        [Input("withFederation")]
        public Input<bool>? WithFederation { get; set; }

        public ResourceArgs()
        {
        }
        public static new ResourceArgs Empty => new ResourceArgs();
    }
}
