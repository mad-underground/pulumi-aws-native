// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html
    /// </summary>
    [AwsNativeResourceType("aws-native:SageMaker:ModelPackageGroup")]
    public partial class ModelPackageGroup : Pulumi.CustomResource
    {
        [Output("CreationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        [Output("ModelPackageGroupArn")]
        public Output<string> ModelPackageGroupArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegroupdescription
        /// </summary>
        [Output("ModelPackageGroupDescription")]
        public Output<string?> ModelPackageGroupDescription { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegroupname
        /// </summary>
        [Output("ModelPackageGroupName")]
        public Output<string> ModelPackageGroupName { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegrouppolicy
        /// </summary>
        [Output("ModelPackageGroupPolicy")]
        public Output<object?> ModelPackageGroupPolicy { get; private set; } = null!;

        [Output("ModelPackageGroupStatus")]
        public Output<string> ModelPackageGroupStatus { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-tags
        /// </summary>
        [Output("Tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ModelPackageGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModelPackageGroup(string name, ModelPackageGroupArgs args, CustomResourceOptions? options = null)
            : base("aws-native:SageMaker:ModelPackageGroup", name, args ?? new ModelPackageGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModelPackageGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:SageMaker:ModelPackageGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ModelPackageGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModelPackageGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ModelPackageGroup(name, id, options);
        }
    }

    public sealed class ModelPackageGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegroupdescription
        /// </summary>
        [Input("ModelPackageGroupDescription")]
        public Input<string>? ModelPackageGroupDescription { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegroupname
        /// </summary>
        [Input("ModelPackageGroupName", required: true)]
        public Input<string> ModelPackageGroupName { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-modelpackagegrouppolicy
        /// </summary>
        [Input("ModelPackageGroupPolicy")]
        public Input<object>? ModelPackageGroupPolicy { get; set; }

        [Input("Tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _Tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-modelpackagegroup.html#cfn-sagemaker-modelpackagegroup-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _Tags ?? (_Tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _Tags = value;
        }

        public ModelPackageGroupArgs()
        {
        }
    }
}
