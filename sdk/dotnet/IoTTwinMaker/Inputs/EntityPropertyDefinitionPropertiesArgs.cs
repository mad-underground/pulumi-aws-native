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
    /// An object that specifies information about a property.
    /// </summary>
    public sealed class EntityPropertyDefinitionPropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An object that specifies information about a property.
        /// </summary>
        [Input("configuration")]
        public Input<Inputs.EntityPropertyDefinitionConfigurationArgs>? Configuration { get; set; }

        /// <summary>
        /// An object that contains information about the data type.
        /// </summary>
        [Input("dataType")]
        public Input<Inputs.EntityDataTypeArgs>? DataType { get; set; }

        /// <summary>
        /// An object that contains the default value.
        /// </summary>
        [Input("defaultValue")]
        public Input<Inputs.EntityDataValueArgs>? DefaultValue { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property ID comes from an external data store.
        /// </summary>
        [Input("isExternalId")]
        public Input<bool>? IsExternalId { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property definition can be updated.
        /// </summary>
        [Input("isFinal")]
        public Input<bool>? IsFinal { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property definition is imported from an external data store.
        /// </summary>
        [Input("isImported")]
        public Input<bool>? IsImported { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property definition is inherited from a parent entity.
        /// </summary>
        [Input("isInherited")]
        public Input<bool>? IsInherited { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property is required.
        /// </summary>
        [Input("isRequiredInEntity")]
        public Input<bool>? IsRequiredInEntity { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property is stored externally.
        /// </summary>
        [Input("isStoredExternally")]
        public Input<bool>? IsStoredExternally { get; set; }

        /// <summary>
        /// A Boolean value that specifies whether the property consists of time series data.
        /// </summary>
        [Input("isTimeSeries")]
        public Input<bool>? IsTimeSeries { get; set; }

        public EntityPropertyDefinitionPropertiesArgs()
        {
        }
        public static new EntityPropertyDefinitionPropertiesArgs Empty => new EntityPropertyDefinitionPropertiesArgs();
    }
}