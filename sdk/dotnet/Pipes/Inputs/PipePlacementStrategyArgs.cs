// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipePlacementStrategyArgs : global::Pulumi.ResourceArgs
    {
        [Input("field")]
        public Input<string>? Field { get; set; }

        [Input("type")]
        public Input<Pulumi.AwsNative.Pipes.PipePlacementStrategyType>? Type { get; set; }

        public PipePlacementStrategyArgs()
        {
        }
        public static new PipePlacementStrategyArgs Empty => new PipePlacementStrategyArgs();
    }
}
