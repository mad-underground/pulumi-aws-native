// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EmrServerless.Inputs
{

    public sealed class ApplicationLogTypeMapKeyValuePairArgs : global::Pulumi.ResourceArgs
    {
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("value", required: true)]
        private InputList<string>? _value;
        public InputList<string> Value
        {
            get => _value ?? (_value = new InputList<string>());
            set => _value = value;
        }

        public ApplicationLogTypeMapKeyValuePairArgs()
        {
        }
        public static new ApplicationLogTypeMapKeyValuePairArgs Empty => new ApplicationLogTypeMapKeyValuePairArgs();
    }
}