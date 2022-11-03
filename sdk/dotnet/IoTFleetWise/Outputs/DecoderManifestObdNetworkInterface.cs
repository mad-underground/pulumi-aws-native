// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTFleetWise.Outputs
{

    [OutputType]
    public sealed class DecoderManifestObdNetworkInterface
    {
        public readonly string InterfaceId;
        public readonly Outputs.DecoderManifestObdInterface ObdInterface;
        public readonly Pulumi.AwsNative.IoTFleetWise.DecoderManifestObdNetworkInterfaceType Type;

        [OutputConstructor]
        private DecoderManifestObdNetworkInterface(
            string interfaceId,

            Outputs.DecoderManifestObdInterface obdInterface,

            Pulumi.AwsNative.IoTFleetWise.DecoderManifestObdNetworkInterfaceType type)
        {
            InterfaceId = interfaceId;
            ObdInterface = obdInterface;
            Type = type;
        }
    }
}