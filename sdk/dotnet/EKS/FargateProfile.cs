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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html
    /// </summary>
    [AwsNativeResourceType("aws-native:EKS:FargateProfile")]
    public partial class FargateProfile : Pulumi.CustomResource
    {
        [Output("Arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-clustername
        /// </summary>
        [Output("ClusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-fargateprofilename
        /// </summary>
        [Output("FargateProfileName")]
        public Output<string?> FargateProfileName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-podexecutionrolearn
        /// </summary>
        [Output("PodExecutionRoleArn")]
        public Output<string> PodExecutionRoleArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-selectors
        /// </summary>
        [Output("Selectors")]
        public Output<ImmutableArray<Outputs.FargateProfileSelector>> Selectors { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-subnets
        /// </summary>
        [Output("Subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-tags
        /// </summary>
        [Output("Tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a FargateProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FargateProfile(string name, FargateProfileArgs args, CustomResourceOptions? options = null)
            : base("aws-native:EKS:FargateProfile", name, args ?? new FargateProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FargateProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:EKS:FargateProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FargateProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FargateProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FargateProfile(name, id, options);
        }
    }

    public sealed class FargateProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-clustername
        /// </summary>
        [Input("ClusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-fargateprofilename
        /// </summary>
        [Input("FargateProfileName")]
        public Input<string>? FargateProfileName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-podexecutionrolearn
        /// </summary>
        [Input("PodExecutionRoleArn", required: true)]
        public Input<string> PodExecutionRoleArn { get; set; } = null!;

        [Input("Selectors", required: true)]
        private InputList<Inputs.FargateProfileSelectorArgs>? _Selectors;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-selectors
        /// </summary>
        public InputList<Inputs.FargateProfileSelectorArgs> Selectors
        {
            get => _Selectors ?? (_Selectors = new InputList<Inputs.FargateProfileSelectorArgs>());
            set => _Selectors = value;
        }

        [Input("Subnets")]
        private InputList<string>? _Subnets;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-subnets
        /// </summary>
        public InputList<string> Subnets
        {
            get => _Subnets ?? (_Subnets = new InputList<string>());
            set => _Subnets = value;
        }

        [Input("Tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _Tags = value;
        }

        public FargateProfileArgs()
        {
        }
    }
}
