// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeEcsEnvironmentVariableArgs : global::Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public PipeEcsEnvironmentVariableArgs()
        {
        }
        public static new PipeEcsEnvironmentVariableArgs Empty => new PipeEcsEnvironmentVariableArgs();
    }
}
