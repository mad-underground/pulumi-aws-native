// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    public sealed class EventInvokeConfigOnSuccessArgs : global::Pulumi.ResourceArgs
    {
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        public EventInvokeConfigOnSuccessArgs()
        {
        }
        public static new EventInvokeConfigOnSuccessArgs Empty => new EventInvokeConfigOnSuccessArgs();
    }
}
