// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    /// <summary>
    /// Resource Type definition for AWS::EC2::NetworkInterfaceAttachment
    /// </summary>
    [AwsNativeResourceType("aws-native:ec2:NetworkInterfaceAttachment")]
    public partial class NetworkInterfaceAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the network interface attachment.
        /// </summary>
        [Output("attachmentId")]
        public Output<string> AttachmentId { get; private set; } = null!;

        /// <summary>
        /// Whether to delete the network interface when the instance terminates. By default, this value is set to true.
        /// </summary>
        [Output("deleteOnTermination")]
        public Output<bool?> DeleteOnTermination { get; private set; } = null!;

        /// <summary>
        /// The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.
        /// </summary>
        [Output("deviceIndex")]
        public Output<string> DeviceIndex { get; private set; } = null!;

        [Output("enaSrdSpecification")]
        public Output<Outputs.NetworkInterfaceAttachmentEnaSrdSpecification?> EnaSrdSpecification { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance to which you will attach the ENI.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The ID of the ENI that you want to attach.
        /// </summary>
        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfaceAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfaceAttachment(string name, NetworkInterfaceAttachmentArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInterfaceAttachment", name, args ?? new NetworkInterfaceAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfaceAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInterfaceAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "deviceIndex",
                    "instanceId",
                    "networkInterfaceId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkInterfaceAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfaceAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkInterfaceAttachment(name, id, options);
        }
    }

    public sealed class NetworkInterfaceAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to delete the network interface when the instance terminates. By default, this value is set to true.
        /// </summary>
        [Input("deleteOnTermination")]
        public Input<bool>? DeleteOnTermination { get; set; }

        /// <summary>
        /// The network interface's position in the attachment order. For example, the first attached network interface has a DeviceIndex of 0.
        /// </summary>
        [Input("deviceIndex", required: true)]
        public Input<string> DeviceIndex { get; set; } = null!;

        [Input("enaSrdSpecification")]
        public Input<Inputs.NetworkInterfaceAttachmentEnaSrdSpecificationArgs>? EnaSrdSpecification { get; set; }

        /// <summary>
        /// The ID of the instance to which you will attach the ENI.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The ID of the ENI that you want to attach.
        /// </summary>
        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        public NetworkInterfaceAttachmentArgs()
        {
        }
        public static new NetworkInterfaceAttachmentArgs Empty => new NetworkInterfaceAttachmentArgs();
    }
}
