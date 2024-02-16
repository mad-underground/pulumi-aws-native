// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGatewayV2
{
    /// <summary>
    /// Resource Type definition for AWS::ApiGatewayV2::VpcLink
    /// </summary>
    [AwsNativeResourceType("aws-native:apigatewayv2:VpcLink")]
    public partial class VpcLink : global::Pulumi.CustomResource
    {
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// This resource type use map for Tags, suggest to use List of Tag
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("vpcLinkId")]
        public Output<string> VpcLinkId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcLink resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcLink(string name, VpcLinkArgs args, CustomResourceOptions? options = null)
            : base("aws-native:apigatewayv2:VpcLink", name, args ?? new VpcLinkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcLink(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:apigatewayv2:VpcLink", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "securityGroupIds[*]",
                    "subnetIds[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpcLink resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcLink Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpcLink(name, id, options);
        }
    }

    public sealed class VpcLinkArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// This resource type use map for Tags, suggest to use List of Tag
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public VpcLinkArgs()
        {
        }
        public static new VpcLinkArgs Empty => new VpcLinkArgs();
    }
}
