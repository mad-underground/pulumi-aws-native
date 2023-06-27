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
    /// Resource Type definition for AWS::EC2::Host
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:Host")]
    public partial class Host : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Outpost hardware asset.
        /// </summary>
        [Output("assetId")]
        public Output<string?> AssetId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
        /// </summary>
        [Output("autoPlacement")]
        public Output<string?> AutoPlacement { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone in which to allocate the Dedicated Host.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// ID of the host created.
        /// </summary>
        [Output("hostId")]
        public Output<string> HostId { get; private set; } = null!;

        /// <summary>
        /// Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host.
        /// </summary>
        [Output("hostMaintenance")]
        public Output<string?> HostMaintenance { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
        /// </summary>
        [Output("hostRecovery")]
        public Output<string?> HostRecovery { get; private set; } = null!;

        /// <summary>
        /// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.
        /// </summary>
        [Output("instanceFamily")]
        public Output<string?> InstanceFamily { get; private set; } = null!;

        /// <summary>
        /// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
        /// </summary>
        [Output("instanceType")]
        public Output<string?> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.
        /// </summary>
        [Output("outpostArn")]
        public Output<string?> OutpostArn { get; private set; } = null!;


        /// <summary>
        /// Create a Host resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Host(string name, HostArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:Host", name, args ?? new HostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Host(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:Host", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Host resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Host Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Host(name, id, options);
        }
    }

    public sealed class HostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Outpost hardware asset.
        /// </summary>
        [Input("assetId")]
        public Input<string>? AssetId { get; set; }

        /// <summary>
        /// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
        /// </summary>
        [Input("autoPlacement")]
        public Input<string>? AutoPlacement { get; set; }

        /// <summary>
        /// The Availability Zone in which to allocate the Dedicated Host.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// Automatically allocates a new dedicated host and moves your instances on to it if a degradation is detected on your current host.
        /// </summary>
        [Input("hostMaintenance")]
        public Input<string>? HostMaintenance { get; set; }

        /// <summary>
        /// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
        /// </summary>
        [Input("hostRecovery")]
        public Input<string>? HostRecovery { get; set; }

        /// <summary>
        /// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family.
        /// </summary>
        [Input("instanceFamily")]
        public Input<string>? InstanceFamily { get; set; }

        /// <summary>
        /// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Amazon Web Services Outpost on which to allocate the Dedicated Host.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        public HostArgs()
        {
        }
        public static new HostArgs Empty => new HostArgs();
    }
}
