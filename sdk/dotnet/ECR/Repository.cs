// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECR
{
    /// <summary>
    /// The AWS::ECR::Repository resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html
    /// </summary>
    [AwsNativeResourceType("aws-native:ecr:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("encryptionConfiguration")]
        public Output<Outputs.RepositoryEncryptionConfiguration?> EncryptionConfiguration { get; private set; } = null!;

        [Output("imageScanningConfiguration")]
        public Output<Outputs.RepositoryImageScanningConfiguration?> ImageScanningConfiguration { get; private set; } = null!;

        /// <summary>
        /// The image tag mutability setting for the repository.
        /// </summary>
        [Output("imageTagMutability")]
        public Output<Pulumi.AwsNative.ECR.RepositoryImageTagMutability?> ImageTagMutability { get; private set; } = null!;

        [Output("lifecyclePolicy")]
        public Output<Outputs.RepositoryLifecyclePolicy?> LifecyclePolicy { get; private set; } = null!;

        /// <summary>
        /// The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
        /// </summary>
        [Output("repositoryName")]
        public Output<string?> RepositoryName { get; private set; } = null!;

        /// <summary>
        /// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. 
        /// </summary>
        [Output("repositoryPolicyText")]
        public Output<object?> RepositoryPolicyText { get; private set; } = null!;

        [Output("repositoryUri")]
        public Output<string> RepositoryUri { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.RepositoryTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ecr:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ecr:Repository", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, options);
        }
    }

    public sealed class RepositoryArgs : global::Pulumi.ResourceArgs
    {
        [Input("encryptionConfiguration")]
        public Input<Inputs.RepositoryEncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        [Input("imageScanningConfiguration")]
        public Input<Inputs.RepositoryImageScanningConfigurationArgs>? ImageScanningConfiguration { get; set; }

        /// <summary>
        /// The image tag mutability setting for the repository.
        /// </summary>
        [Input("imageTagMutability")]
        public Input<Pulumi.AwsNative.ECR.RepositoryImageTagMutability>? ImageTagMutability { get; set; }

        [Input("lifecyclePolicy")]
        public Input<Inputs.RepositoryLifecyclePolicyArgs>? LifecyclePolicy { get; set; }

        /// <summary>
        /// The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
        /// </summary>
        [Input("repositoryName")]
        public Input<string>? RepositoryName { get; set; }

        /// <summary>
        /// The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. 
        /// </summary>
        [Input("repositoryPolicyText")]
        public Input<object>? RepositoryPolicyText { get; set; }

        [Input("tags")]
        private InputList<Inputs.RepositoryTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.RepositoryTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.RepositoryTagArgs>());
            set => _tags = value;
        }

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }
}
