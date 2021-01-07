// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualGatewayAttributes
    {
        public readonly string Arn;
        public readonly string MeshName;
        public readonly string MeshOwner;
        public readonly string ResourceOwner;
        public readonly string Uid;
        public readonly string VirtualGatewayName;

        [OutputConstructor]
        private VirtualGatewayAttributes(
            string Arn,

            string MeshName,

            string MeshOwner,

            string ResourceOwner,

            string Uid,

            string VirtualGatewayName)
        {
            this.Arn = Arn;
            this.MeshName = MeshName;
            this.MeshOwner = MeshOwner;
            this.ResourceOwner = ResourceOwner;
            this.Uid = Uid;
            this.VirtualGatewayName = VirtualGatewayName;
        }
    }
}
