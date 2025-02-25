// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Events.Inputs
{

    public sealed class RuleInputTransformerArgs : global::Pulumi.ResourceArgs
    {
        [Input("inputPathsMap")]
        private InputMap<string>? _inputPathsMap;
        public InputMap<string> InputPathsMap
        {
            get => _inputPathsMap ?? (_inputPathsMap = new InputMap<string>());
            set => _inputPathsMap = value;
        }

        [Input("inputTemplate", required: true)]
        public Input<string> InputTemplate { get; set; } = null!;

        public RuleInputTransformerArgs()
        {
        }
        public static new RuleInputTransformerArgs Empty => new RuleInputTransformerArgs();
    }
}
