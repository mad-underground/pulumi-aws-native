// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class ComponentActionParametersArgs : global::Pulumi.ResourceArgs
    {
        [Input("anchor")]
        public Input<Inputs.ComponentPropertyArgs>? Anchor { get; set; }

        [Input("fields")]
        private InputMap<Inputs.ComponentPropertyArgs>? _fields;
        public InputMap<Inputs.ComponentPropertyArgs> Fields
        {
            get => _fields ?? (_fields = new InputMap<Inputs.ComponentPropertyArgs>());
            set => _fields = value;
        }

        [Input("global")]
        public Input<Inputs.ComponentPropertyArgs>? Global { get; set; }

        [Input("id")]
        public Input<Inputs.ComponentPropertyArgs>? Id { get; set; }

        [Input("model")]
        public Input<string>? Model { get; set; }

        [Input("state")]
        public Input<Inputs.ComponentMutationActionSetStateParameterArgs>? State { get; set; }

        [Input("target")]
        public Input<Inputs.ComponentPropertyArgs>? Target { get; set; }

        [Input("type")]
        public Input<Inputs.ComponentPropertyArgs>? Type { get; set; }

        [Input("url")]
        public Input<Inputs.ComponentPropertyArgs>? Url { get; set; }

        public ComponentActionParametersArgs()
        {
        }
        public static new ComponentActionParametersArgs Empty => new ComponentActionParametersArgs();
    }
}
