// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    /// <summary>
    /// Resource Type definition for AWS::Greengrass::DeviceDefinitionVersion
    /// </summary>
    [Obsolete(@"DeviceDefinitionVersion is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:greengrass:DeviceDefinitionVersion")]
    public partial class DeviceDefinitionVersion : global::Pulumi.CustomResource
    {
        [Output("awsId")]
        public Output<string> AwsId { get; private set; } = null!;

        [Output("deviceDefinitionId")]
        public Output<string> DeviceDefinitionId { get; private set; } = null!;

        [Output("devices")]
        public Output<ImmutableArray<Outputs.DeviceDefinitionVersionDevice>> Devices { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceDefinitionVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceDefinitionVersion(string name, DeviceDefinitionVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:greengrass:DeviceDefinitionVersion", name, args ?? new DeviceDefinitionVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceDefinitionVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:greengrass:DeviceDefinitionVersion", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "deviceDefinitionId",
                    "devices[*]",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DeviceDefinitionVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceDefinitionVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeviceDefinitionVersion(name, id, options);
        }
    }

    public sealed class DeviceDefinitionVersionArgs : global::Pulumi.ResourceArgs
    {
        [Input("deviceDefinitionId", required: true)]
        public Input<string> DeviceDefinitionId { get; set; } = null!;

        [Input("devices", required: true)]
        private InputList<Inputs.DeviceDefinitionVersionDeviceArgs>? _devices;
        public InputList<Inputs.DeviceDefinitionVersionDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.DeviceDefinitionVersionDeviceArgs>());
            set => _devices = value;
        }

        public DeviceDefinitionVersionArgs()
        {
        }
        public static new DeviceDefinitionVersionArgs Empty => new DeviceDefinitionVersionArgs();
    }
}
