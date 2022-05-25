// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkManager.Outputs
{

    [OutputType]
    public sealed class ConnectPeerConfiguration
    {
        public readonly ImmutableArray<Outputs.ConnectPeerBgpConfiguration> BgpConfigurations;
        public readonly string? CoreNetworkAddress;
        public readonly ImmutableArray<string> InsideCidrBlocks;
        public readonly string? PeerAddress;
        public readonly string? Protocol;

        [OutputConstructor]
        private ConnectPeerConfiguration(
            ImmutableArray<Outputs.ConnectPeerBgpConfiguration> bgpConfigurations,

            string? coreNetworkAddress,

            ImmutableArray<string> insideCidrBlocks,

            string? peerAddress,

            string? protocol)
        {
            BgpConfigurations = bgpConfigurations;
            CoreNetworkAddress = coreNetworkAddress;
            InsideCidrBlocks = insideCidrBlocks;
            PeerAddress = peerAddress;
            Protocol = protocol;
        }
    }
}