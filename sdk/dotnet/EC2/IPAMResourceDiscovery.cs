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
    /// Resource Schema of AWS::EC2::IPAMResourceDiscovery Type
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:IPAMResourceDiscovery")]
    public partial class IPAMResourceDiscovery : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (Arn) for the Resource Discovery.
        /// </summary>
        [Output("ipamResourceDiscoveryArn")]
        public Output<string> IpamResourceDiscoveryArn { get; private set; } = null!;

        /// <summary>
        /// Id of the IPAM Pool.
        /// </summary>
        [Output("ipamResourceDiscoveryId")]
        public Output<string> IpamResourceDiscoveryId { get; private set; } = null!;

        /// <summary>
        /// The region the resource discovery is setup in. 
        /// </summary>
        [Output("ipamResourceDiscoveryRegion")]
        public Output<string> IpamResourceDiscoveryRegion { get; private set; } = null!;

        /// <summary>
        /// Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// The regions Resource Discovery is enabled for. Allows resource discoveries to be created in these regions, as well as enabling monitoring
        /// </summary>
        [Output("operatingRegions")]
        public Output<ImmutableArray<Outputs.IPAMResourceDiscoveryIpamOperatingRegion>> OperatingRegions { get; private set; } = null!;

        /// <summary>
        /// Owner Account ID of the Resource Discovery
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The state of this Resource Discovery.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.IPAMResourceDiscoveryTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a IPAMResourceDiscovery resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IPAMResourceDiscovery(string name, IPAMResourceDiscoveryArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ec2:IPAMResourceDiscovery", name, args ?? new IPAMResourceDiscoveryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IPAMResourceDiscovery(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:IPAMResourceDiscovery", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing IPAMResourceDiscovery resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IPAMResourceDiscovery Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new IPAMResourceDiscovery(name, id, options);
        }
    }

    public sealed class IPAMResourceDiscoveryArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("operatingRegions")]
        private InputList<Inputs.IPAMResourceDiscoveryIpamOperatingRegionArgs>? _operatingRegions;

        /// <summary>
        /// The regions Resource Discovery is enabled for. Allows resource discoveries to be created in these regions, as well as enabling monitoring
        /// </summary>
        public InputList<Inputs.IPAMResourceDiscoveryIpamOperatingRegionArgs> OperatingRegions
        {
            get => _operatingRegions ?? (_operatingRegions = new InputList<Inputs.IPAMResourceDiscoveryIpamOperatingRegionArgs>());
            set => _operatingRegions = value;
        }

        [Input("tags")]
        private InputList<Inputs.IPAMResourceDiscoveryTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.IPAMResourceDiscoveryTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.IPAMResourceDiscoveryTagArgs>());
            set => _tags = value;
        }

        public IPAMResourceDiscoveryArgs()
        {
        }
        public static new IPAMResourceDiscoveryArgs Empty => new IPAMResourceDiscoveryArgs();
    }
}
