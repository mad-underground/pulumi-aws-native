// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream
{
    /// <summary>
    /// Resource Type definition for AWS::AppStream::ImageBuilder
    /// </summary>
    [AwsNativeResourceType("aws-native:appstream:ImageBuilder")]
    public partial class ImageBuilder : global::Pulumi.CustomResource
    {
        [Output("accessEndpoints")]
        public Output<ImmutableArray<Outputs.ImageBuilderAccessEndpoint>> AccessEndpoints { get; private set; } = null!;

        [Output("appstreamAgentVersion")]
        public Output<string?> AppstreamAgentVersion { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        [Output("domainJoinInfo")]
        public Output<Outputs.ImageBuilderDomainJoinInfo?> DomainJoinInfo { get; private set; } = null!;

        [Output("enableDefaultInternetAccess")]
        public Output<bool?> EnableDefaultInternetAccess { get; private set; } = null!;

        [Output("iamRoleArn")]
        public Output<string?> IamRoleArn { get; private set; } = null!;

        [Output("imageArn")]
        public Output<string?> ImageArn { get; private set; } = null!;

        [Output("imageName")]
        public Output<string?> ImageName { get; private set; } = null!;

        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("streamingUrl")]
        public Output<string> StreamingUrl { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        [Output("vpcConfig")]
        public Output<Outputs.ImageBuilderVpcConfig?> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a ImageBuilder resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageBuilder(string name, ImageBuilderArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appstream:ImageBuilder", name, args ?? new ImageBuilderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageBuilder(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appstream:ImageBuilder", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ImageBuilder resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageBuilder Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ImageBuilder(name, id, options);
        }
    }

    public sealed class ImageBuilderArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessEndpoints")]
        private InputList<Inputs.ImageBuilderAccessEndpointArgs>? _accessEndpoints;
        public InputList<Inputs.ImageBuilderAccessEndpointArgs> AccessEndpoints
        {
            get => _accessEndpoints ?? (_accessEndpoints = new InputList<Inputs.ImageBuilderAccessEndpointArgs>());
            set => _accessEndpoints = value;
        }

        [Input("appstreamAgentVersion")]
        public Input<string>? AppstreamAgentVersion { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("domainJoinInfo")]
        public Input<Inputs.ImageBuilderDomainJoinInfoArgs>? DomainJoinInfo { get; set; }

        [Input("enableDefaultInternetAccess")]
        public Input<bool>? EnableDefaultInternetAccess { get; set; }

        [Input("iamRoleArn")]
        public Input<string>? IamRoleArn { get; set; }

        [Input("imageArn")]
        public Input<string>? ImageArn { get; set; }

        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        [Input("vpcConfig")]
        public Input<Inputs.ImageBuilderVpcConfigArgs>? VpcConfig { get; set; }

        public ImageBuilderArgs()
        {
        }
        public static new ImageBuilderArgs Empty => new ImageBuilderArgs();
    }
}
