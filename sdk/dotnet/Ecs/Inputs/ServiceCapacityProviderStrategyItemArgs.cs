// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class ServiceCapacityProviderStrategyItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("base")]
        public Input<int>? Base { get; set; }

        [Input("capacityProvider")]
        public Input<string>? CapacityProvider { get; set; }

        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public ServiceCapacityProviderStrategyItemArgs()
        {
        }
        public static new ServiceCapacityProviderStrategyItemArgs Empty => new ServiceCapacityProviderStrategyItemArgs();
    }
}
