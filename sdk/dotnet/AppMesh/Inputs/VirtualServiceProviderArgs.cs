// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualServiceProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("virtualNode")]
        public Input<Inputs.VirtualServiceVirtualNodeServiceProviderArgs>? VirtualNode { get; set; }

        [Input("virtualRouter")]
        public Input<Inputs.VirtualServiceVirtualRouterServiceProviderArgs>? VirtualRouter { get; set; }

        public VirtualServiceProviderArgs()
        {
        }
        public static new VirtualServiceProviderArgs Empty => new VirtualServiceProviderArgs();
    }
}
