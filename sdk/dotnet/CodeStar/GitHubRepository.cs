// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeStar
{
    /// <summary>
    /// Resource Type definition for AWS::CodeStar::GitHubRepository
    /// </summary>
    [Obsolete(@"GitHubRepository is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:codestar:GitHubRepository")]
    public partial class GitHubRepository : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("code")]
        public Output<Outputs.GitHubRepositoryCode?> Code { get; private set; } = null!;

        [Output("connectionArn")]
        public Output<string?> ConnectionArn { get; private set; } = null!;

        [Output("enableIssues")]
        public Output<bool?> EnableIssues { get; private set; } = null!;

        [Output("isPrivate")]
        public Output<bool?> IsPrivate { get; private set; } = null!;

        [Output("repositoryAccessToken")]
        public Output<string?> RepositoryAccessToken { get; private set; } = null!;

        [Output("repositoryDescription")]
        public Output<string?> RepositoryDescription { get; private set; } = null!;

        [Output("repositoryName")]
        public Output<string> RepositoryName { get; private set; } = null!;

        [Output("repositoryOwner")]
        public Output<string> RepositoryOwner { get; private set; } = null!;


        /// <summary>
        /// Create a GitHubRepository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GitHubRepository(string name, GitHubRepositoryArgs args, CustomResourceOptions? options = null)
            : base("aws-native:codestar:GitHubRepository", name, args ?? new GitHubRepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GitHubRepository(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:codestar:GitHubRepository", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing GitHubRepository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GitHubRepository Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GitHubRepository(name, id, options);
        }
    }

    public sealed class GitHubRepositoryArgs : global::Pulumi.ResourceArgs
    {
        [Input("code")]
        public Input<Inputs.GitHubRepositoryCodeArgs>? Code { get; set; }

        [Input("connectionArn")]
        public Input<string>? ConnectionArn { get; set; }

        [Input("enableIssues")]
        public Input<bool>? EnableIssues { get; set; }

        [Input("isPrivate")]
        public Input<bool>? IsPrivate { get; set; }

        [Input("repositoryAccessToken")]
        public Input<string>? RepositoryAccessToken { get; set; }

        [Input("repositoryDescription")]
        public Input<string>? RepositoryDescription { get; set; }

        [Input("repositoryName", required: true)]
        public Input<string> RepositoryName { get; set; } = null!;

        [Input("repositoryOwner", required: true)]
        public Input<string> RepositoryOwner { get; set; } = null!;

        public GitHubRepositoryArgs()
        {
        }
        public static new GitHubRepositoryArgs Empty => new GitHubRepositoryArgs();
    }
}
