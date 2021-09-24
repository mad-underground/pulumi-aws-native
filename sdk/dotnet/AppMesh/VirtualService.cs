// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh
{
    /// <summary>
    /// Resource Type definition for AWS::AppMesh::VirtualService
    /// </summary>
    [Obsolete(@"VirtualService is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:appmesh:VirtualService")]
    public partial class VirtualService : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("meshName")]
        public Output<string> MeshName { get; private set; } = null!;

        [Output("meshOwner")]
        public Output<string?> MeshOwner { get; private set; } = null!;

        [Output("resourceOwner")]
        public Output<string> ResourceOwner { get; private set; } = null!;

        [Output("spec")]
        public Output<Outputs.VirtualServiceVirtualServiceSpec> Spec { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.VirtualServiceTag>> Tags { get; private set; } = null!;

        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        [Output("virtualServiceName")]
        public Output<string> VirtualServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualService(string name, VirtualServiceArgs args, CustomResourceOptions? options = null)
            : base("aws-native:appmesh:VirtualService", name, args ?? new VirtualServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualService(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:appmesh:VirtualService", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing VirtualService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualService Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualService(name, id, options);
        }
    }

    public sealed class VirtualServiceArgs : Pulumi.ResourceArgs
    {
        [Input("meshName", required: true)]
        public Input<string> MeshName { get; set; } = null!;

        [Input("meshOwner")]
        public Input<string>? MeshOwner { get; set; }

        [Input("spec", required: true)]
        public Input<Inputs.VirtualServiceVirtualServiceSpecArgs> Spec { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.VirtualServiceTagArgs>? _tags;
        public InputList<Inputs.VirtualServiceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.VirtualServiceTagArgs>());
            set => _tags = value;
        }

        [Input("virtualServiceName", required: true)]
        public Input<string> VirtualServiceName { get; set; } = null!;

        public VirtualServiceArgs()
        {
        }
    }
}