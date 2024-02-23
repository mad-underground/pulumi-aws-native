// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowCustomConnectorDestinationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("customProperties")]
        private InputMap<string>? _customProperties;
        public InputMap<string> CustomProperties
        {
            get => _customProperties ?? (_customProperties = new InputMap<string>());
            set => _customProperties = value;
        }

        [Input("entityName", required: true)]
        public Input<string> EntityName { get; set; } = null!;

        [Input("errorHandlingConfig")]
        public Input<Inputs.FlowErrorHandlingConfigArgs>? ErrorHandlingConfig { get; set; }

        [Input("idFieldNames")]
        private InputList<string>? _idFieldNames;

        /// <summary>
        /// List of fields used as ID when performing a write operation.
        /// </summary>
        public InputList<string> IdFieldNames
        {
            get => _idFieldNames ?? (_idFieldNames = new InputList<string>());
            set => _idFieldNames = value;
        }

        [Input("writeOperationType")]
        public Input<Pulumi.AwsNative.AppFlow.FlowWriteOperationType>? WriteOperationType { get; set; }

        public FlowCustomConnectorDestinationPropertiesArgs()
        {
        }
        public static new FlowCustomConnectorDestinationPropertiesArgs Empty => new FlowCustomConnectorDestinationPropertiesArgs();
    }
}
