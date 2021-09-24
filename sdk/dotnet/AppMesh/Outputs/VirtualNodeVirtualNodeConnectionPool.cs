// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualNodeVirtualNodeConnectionPool
    {
        public readonly Outputs.VirtualNodeVirtualNodeGrpcConnectionPool? GRPC;
        public readonly Outputs.VirtualNodeVirtualNodeHttpConnectionPool? HTTP;
        public readonly Outputs.VirtualNodeVirtualNodeHttp2ConnectionPool? HTTP2;
        public readonly Outputs.VirtualNodeVirtualNodeTcpConnectionPool? TCP;

        [OutputConstructor]
        private VirtualNodeVirtualNodeConnectionPool(
            Outputs.VirtualNodeVirtualNodeGrpcConnectionPool? gRPC,

            Outputs.VirtualNodeVirtualNodeHttpConnectionPool? hTTP,

            Outputs.VirtualNodeVirtualNodeHttp2ConnectionPool? hTTP2,

            Outputs.VirtualNodeVirtualNodeTcpConnectionPool? tCP)
        {
            GRPC = gRPC;
            HTTP = hTTP;
            HTTP2 = hTTP2;
            TCP = tCP;
        }
    }
}