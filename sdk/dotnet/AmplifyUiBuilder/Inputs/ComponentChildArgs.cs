// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class ComponentChildArgs : global::Pulumi.ResourceArgs
    {
        [Input("children")]
        private InputList<Inputs.ComponentChildArgs>? _children;
        public InputList<Inputs.ComponentChildArgs> Children
        {
            get => _children ?? (_children = new InputList<Inputs.ComponentChildArgs>());
            set => _children = value;
        }

        [Input("componentType", required: true)]
        public Input<string> ComponentType { get; set; } = null!;

        [Input("events")]
        private InputMap<Inputs.ComponentEventArgs>? _events;
        public InputMap<Inputs.ComponentEventArgs> Events
        {
            get => _events ?? (_events = new InputMap<Inputs.ComponentEventArgs>());
            set => _events = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("properties", required: true)]
        private InputMap<Inputs.ComponentPropertyArgs>? _properties;
        public InputMap<Inputs.ComponentPropertyArgs> Properties
        {
            get => _properties ?? (_properties = new InputMap<Inputs.ComponentPropertyArgs>());
            set => _properties = value;
        }

        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        public ComponentChildArgs()
        {
        }
        public static new ComponentChildArgs Empty => new ComponentChildArgs();
    }
}
