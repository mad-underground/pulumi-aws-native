// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        public PipeFilterArgs()
        {
        }
        public static new PipeFilterArgs Empty => new PipeFilterArgs();
    }
}
