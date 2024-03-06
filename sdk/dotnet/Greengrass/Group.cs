// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    /// <summary>
    /// Resource Type definition for AWS::Greengrass::Group
    /// </summary>
    [Obsolete(@"Group is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:greengrass:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("initialVersion")]
        public Output<Outputs.GroupVersion?> InitialVersion { get; private set; } = null!;

        [Output("latestVersionArn")]
        public Output<string> LatestVersionArn { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string?> RoleArn { get; private set; } = null!;

        [Output("roleAttachedAt")]
        public Output<string> RoleAttachedAt { get; private set; } = null!;

        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Greengrass::Group` for more information about the expected schema for this property.
        /// </summary>
        [Output("tags")]
        public Output<object?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:greengrass:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:greengrass:Group", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "initialVersion",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Group(name, id, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("initialVersion")]
        public Input<Inputs.GroupVersionArgs>? InitialVersion { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        /// <summary>
        /// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Greengrass::Group` for more information about the expected schema for this property.
        /// </summary>
        [Input("tags")]
        public Input<object>? Tags { get; set; }

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }
}
