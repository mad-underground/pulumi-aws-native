// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeTargetEventBridgeEventBusParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("detailType")]
        public Input<string>? DetailType { get; set; }

        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        [Input("resources")]
        private InputList<string>? _resources;
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("time")]
        public Input<string>? Time { get; set; }

        public PipeTargetEventBridgeEventBusParametersArgs()
        {
        }
        public static new PipeTargetEventBridgeEventBusParametersArgs Empty => new PipeTargetEventBridgeEventBusParametersArgs();
    }
}
