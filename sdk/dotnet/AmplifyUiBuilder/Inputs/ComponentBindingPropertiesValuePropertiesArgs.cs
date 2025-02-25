// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class ComponentBindingPropertiesValuePropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        [Input("field")]
        public Input<string>? Field { get; set; }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("model")]
        public Input<string>? Model { get; set; }

        [Input("predicates")]
        private InputList<Inputs.ComponentPredicateArgs>? _predicates;
        public InputList<Inputs.ComponentPredicateArgs> Predicates
        {
            get => _predicates ?? (_predicates = new InputList<Inputs.ComponentPredicateArgs>());
            set => _predicates = value;
        }

        [Input("slotName")]
        public Input<string>? SlotName { get; set; }

        [Input("userAttribute")]
        public Input<string>? UserAttribute { get; set; }

        public ComponentBindingPropertiesValuePropertiesArgs()
        {
        }
        public static new ComponentBindingPropertiesValuePropertiesArgs Empty => new ComponentBindingPropertiesValuePropertiesArgs();
    }
}
