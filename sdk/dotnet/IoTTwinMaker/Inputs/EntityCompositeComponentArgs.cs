// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Inputs
{

    public sealed class EntityCompositeComponentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the component.
        /// </summary>
        [Input("componentName")]
        public Input<string>? ComponentName { get; set; }

        /// <summary>
        /// The path of the component.
        /// </summary>
        [Input("componentPath")]
        public Input<string>? ComponentPath { get; set; }

        /// <summary>
        /// The ID of the component type.
        /// </summary>
        [Input("componentTypeId")]
        public Input<string>? ComponentTypeId { get; set; }

        /// <summary>
        /// The description of the component.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("properties")]
        private InputMap<Inputs.EntityPropertyArgs>? _properties;

        /// <summary>
        /// An object that maps strings to the properties to set in the component type. Each string in the mapping must be unique to this object.
        /// </summary>
        public InputMap<Inputs.EntityPropertyArgs> Properties
        {
            get => _properties ?? (_properties = new InputMap<Inputs.EntityPropertyArgs>());
            set => _properties = value;
        }

        [Input("propertyGroups")]
        private InputMap<Inputs.EntityPropertyGroupArgs>? _propertyGroups;

        /// <summary>
        /// An object that maps strings to the property groups to set in the component type. Each string in the mapping must be unique to this object.
        /// </summary>
        public InputMap<Inputs.EntityPropertyGroupArgs> PropertyGroups
        {
            get => _propertyGroups ?? (_propertyGroups = new InputMap<Inputs.EntityPropertyGroupArgs>());
            set => _propertyGroups = value;
        }

        /// <summary>
        /// The current status of the component.
        /// </summary>
        [Input("status")]
        public Input<Inputs.EntityStatusArgs>? Status { get; set; }

        public EntityCompositeComponentArgs()
        {
        }
        public static new EntityCompositeComponentArgs Empty => new EntityCompositeComponentArgs();
    }
}
