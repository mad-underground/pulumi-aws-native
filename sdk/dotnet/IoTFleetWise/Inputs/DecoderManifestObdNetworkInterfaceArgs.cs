// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTFleetWise.Inputs
{

    public sealed class DecoderManifestObdNetworkInterfaceArgs : global::Pulumi.ResourceArgs
    {
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        [Input("obdInterface", required: true)]
        public Input<Inputs.DecoderManifestObdInterfaceArgs> ObdInterface { get; set; } = null!;

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.IoTFleetWise.DecoderManifestObdNetworkInterfaceType> Type { get; set; } = null!;

        public DecoderManifestObdNetworkInterfaceArgs()
        {
        }
        public static new DecoderManifestObdNetworkInterfaceArgs Empty => new DecoderManifestObdNetworkInterfaceArgs();
    }
}
