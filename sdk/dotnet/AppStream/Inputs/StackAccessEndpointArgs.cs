// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppStream.Inputs
{

    public sealed class StackAccessEndpointArgs : global::Pulumi.ResourceArgs
    {
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        [Input("vpceId", required: true)]
        public Input<string> VpceId { get; set; } = null!;

        public StackAccessEndpointArgs()
        {
        }
        public static new StackAccessEndpointArgs Empty => new StackAccessEndpointArgs();
    }
}
