// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Greengrass.Outputs
{

    [OutputType]
    public sealed class DeviceDefinitionVersionProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html#cfn-greengrass-devicedefinitionversion-devicedefinitionid
        /// </summary>
        public readonly string DeviceDefinitionId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinitionversion.html#cfn-greengrass-devicedefinitionversion-devices
        /// </summary>
        public readonly ImmutableArray<Outputs.DeviceDefinitionVersionDevice> Devices;

        [OutputConstructor]
        private DeviceDefinitionVersionProperties(
            string DeviceDefinitionId,

            ImmutableArray<Outputs.DeviceDefinitionVersionDevice> Devices)
        {
            this.DeviceDefinitionId = DeviceDefinitionId;
            this.Devices = Devices;
        }
    }
}