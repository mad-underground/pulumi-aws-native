// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipePlacementConstraintArgs : global::Pulumi.ResourceArgs
    {
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        [Input("type")]
        public Input<Pulumi.AwsNative.Pipes.PipePlacementConstraintType>? Type { get; set; }

        public PipePlacementConstraintArgs()
        {
        }
        public static new PipePlacementConstraintArgs Empty => new PipePlacementConstraintArgs();
    }
}
