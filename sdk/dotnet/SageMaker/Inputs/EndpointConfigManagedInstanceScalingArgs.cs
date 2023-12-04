// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class EndpointConfigManagedInstanceScalingArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxInstanceCount")]
        public Input<int>? MaxInstanceCount { get; set; }

        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public EndpointConfigManagedInstanceScalingArgs()
        {
        }
        public static new EndpointConfigManagedInstanceScalingArgs Empty => new EndpointConfigManagedInstanceScalingArgs();
    }
}
