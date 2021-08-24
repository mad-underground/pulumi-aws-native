// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EKS
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html
    /// </summary>
    [AwsNativeResourceType("aws-native:eks:Addon")]
    public partial class Addon : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-addonname
        /// </summary>
        [Output("addonName")]
        public Output<string> AddonName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-addonversion
        /// </summary>
        [Output("addonVersion")]
        public Output<string?> AddonVersion { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-clustername
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-resolveconflicts
        /// </summary>
        [Output("resolveConflicts")]
        public Output<string?> ResolveConflicts { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-serviceaccountrolearn
        /// </summary>
        [Output("serviceAccountRoleArn")]
        public Output<string?> ServiceAccountRoleArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Addon resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Addon(string name, AddonArgs args, CustomResourceOptions? options = null)
            : base("aws-native:eks:Addon", name, args ?? new AddonArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Addon(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:eks:Addon", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Addon resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Addon Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Addon(name, id, options);
        }
    }

    public sealed class AddonArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-addonname
        /// </summary>
        [Input("addonName", required: true)]
        public Input<string> AddonName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-addonversion
        /// </summary>
        [Input("addonVersion")]
        public Input<string>? AddonVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-clustername
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-resolveconflicts
        /// </summary>
        [Input("resolveConflicts")]
        public Input<string>? ResolveConflicts { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-serviceaccountrolearn
        /// </summary>
        [Input("serviceAccountRoleArn")]
        public Input<string>? ServiceAccountRoleArn { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-addon.html#cfn-eks-addon-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public AddonArgs()
        {
        }
    }
}