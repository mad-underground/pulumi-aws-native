// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    /// <summary>
    /// Resource schema for AWS::EC2::NetworkInsightsAccessScope
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:NetworkInsightsAccessScope")]
    public partial class NetworkInsightsAccessScope : global::Pulumi.CustomResource
    {
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        [Output("excludePaths")]
        public Output<ImmutableArray<Outputs.NetworkInsightsAccessScopeAccessScopePathRequest>> ExcludePaths { get; private set; } = null!;

        [Output("matchPaths")]
        public Output<ImmutableArray<Outputs.NetworkInsightsAccessScopeAccessScopePathRequest>> MatchPaths { get; private set; } = null!;

        [Output("networkInsightsAccessScopeArn")]
        public Output<string> NetworkInsightsAccessScopeArn { get; private set; } = null!;

        [Output("networkInsightsAccessScopeId")]
        public Output<string> NetworkInsightsAccessScopeId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.NetworkInsightsAccessScopeTag>> Tags { get; private set; } = null!;

        [Output("updatedDate")]
        public Output<string> UpdatedDate { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInsightsAccessScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInsightsAccessScope(string name, NetworkInsightsAccessScopeArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInsightsAccessScope", name, args ?? new NetworkInsightsAccessScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInsightsAccessScope(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInsightsAccessScope", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInsightsAccessScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInsightsAccessScope Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkInsightsAccessScope(name, id, options);
        }
    }

    public sealed class NetworkInsightsAccessScopeArgs : global::Pulumi.ResourceArgs
    {
        [Input("excludePaths")]
        private InputList<Inputs.NetworkInsightsAccessScopeAccessScopePathRequestArgs>? _excludePaths;
        public InputList<Inputs.NetworkInsightsAccessScopeAccessScopePathRequestArgs> ExcludePaths
        {
            get => _excludePaths ?? (_excludePaths = new InputList<Inputs.NetworkInsightsAccessScopeAccessScopePathRequestArgs>());
            set => _excludePaths = value;
        }

        [Input("matchPaths")]
        private InputList<Inputs.NetworkInsightsAccessScopeAccessScopePathRequestArgs>? _matchPaths;
        public InputList<Inputs.NetworkInsightsAccessScopeAccessScopePathRequestArgs> MatchPaths
        {
            get => _matchPaths ?? (_matchPaths = new InputList<Inputs.NetworkInsightsAccessScopeAccessScopePathRequestArgs>());
            set => _matchPaths = value;
        }

        [Input("tags")]
        private InputList<Inputs.NetworkInsightsAccessScopeTagArgs>? _tags;
        public InputList<Inputs.NetworkInsightsAccessScopeTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.NetworkInsightsAccessScopeTagArgs>());
            set => _tags = value;
        }

        public NetworkInsightsAccessScopeArgs()
        {
        }
        public static new NetworkInsightsAccessScopeArgs Empty => new NetworkInsightsAccessScopeArgs();
    }
}
