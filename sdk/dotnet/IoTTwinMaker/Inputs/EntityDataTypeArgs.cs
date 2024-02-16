// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTTwinMaker.Inputs
{

    /// <summary>
    /// An object that specifies the data type of a property.
    /// </summary>
    public sealed class EntityDataTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedValues")]
        private InputList<Inputs.EntityDataValueArgs>? _allowedValues;

        /// <summary>
        /// The allowed values for this data type.
        /// </summary>
        public InputList<Inputs.EntityDataValueArgs> AllowedValues
        {
            get => _allowedValues ?? (_allowedValues = new InputList<Inputs.EntityDataValueArgs>());
            set => _allowedValues = value;
        }

        /// <summary>
        /// The nested type in the data type.
        /// </summary>
        [Input("nestedType")]
        public Input<Inputs.EntityDataTypeArgs>? NestedType { get; set; }

        /// <summary>
        /// A relationship that associates a component with another component.
        /// </summary>
        [Input("relationship")]
        public Input<Inputs.EntityRelationshipArgs>? Relationship { get; set; }

        /// <summary>
        /// The underlying type of the data type.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AwsNative.IoTTwinMaker.EntityDataTypeType>? Type { get; set; }

        /// <summary>
        /// The unit of measure used in this data type.
        /// </summary>
        [Input("unitOfMeasure")]
        public Input<string>? UnitOfMeasure { get; set; }

        public EntityDataTypeArgs()
        {
        }
        public static new EntityDataTypeArgs Empty => new EntityDataTypeArgs();
    }
}
