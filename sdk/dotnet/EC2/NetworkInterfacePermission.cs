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
    /// Resource Type definition for AWS::EC2::NetworkInterfacePermission
    /// </summary>
    [Obsolete(@"NetworkInterfacePermission is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:ec2:NetworkInterfacePermission")]
    public partial class NetworkInterfacePermission : Pulumi.CustomResource
    {
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        [Output("permission")]
        public Output<string> Permission { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterfacePermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterfacePermission(string name, NetworkInterfacePermissionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInterfacePermission", name, args ?? new NetworkInterfacePermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterfacePermission(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ec2:NetworkInterfacePermission", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkInterfacePermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterfacePermission Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkInterfacePermission(name, id, options);
        }
    }

    public sealed class NetworkInterfacePermissionArgs : Pulumi.ResourceArgs
    {
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        public NetworkInterfacePermissionArgs()
        {
        }
    }
}