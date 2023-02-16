// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowSAPODataDestinationPropertiesArgs : global::Pulumi.ResourceArgs
    {
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

        [Input("objectPath", required: true)]
        public Input<string> ObjectPath { get; set; } = null!;

        [Input("successResponseHandlingConfig")]
        public Input<Inputs.FlowSuccessResponseHandlingConfigArgs>? SuccessResponseHandlingConfig { get; set; }

        [Input("writeOperationType")]
        public Input<Pulumi.AwsNative.AppFlow.FlowWriteOperationType>? WriteOperationType { get; set; }

        public FlowSAPODataDestinationPropertiesArgs()
        {
        }
        public static new FlowSAPODataDestinationPropertiesArgs Empty => new FlowSAPODataDestinationPropertiesArgs();
    }
}
