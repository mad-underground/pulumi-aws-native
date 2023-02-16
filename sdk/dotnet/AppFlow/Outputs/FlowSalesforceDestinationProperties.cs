// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowSalesforceDestinationProperties
    {
        public readonly Pulumi.AwsNative.AppFlow.FlowDataTransferApi? DataTransferApi;
        public readonly Outputs.FlowErrorHandlingConfig? ErrorHandlingConfig;
        /// <summary>
        /// List of fields used as ID when performing a write operation.
        /// </summary>
        public readonly ImmutableArray<string> IdFieldNames;
        public readonly string Object;
        public readonly Pulumi.AwsNative.AppFlow.FlowWriteOperationType? WriteOperationType;

        [OutputConstructor]
        private FlowSalesforceDestinationProperties(
            Pulumi.AwsNative.AppFlow.FlowDataTransferApi? dataTransferApi,

            Outputs.FlowErrorHandlingConfig? errorHandlingConfig,

            ImmutableArray<string> idFieldNames,

            string @object,

            Pulumi.AwsNative.AppFlow.FlowWriteOperationType? writeOperationType)
        {
            DataTransferApi = dataTransferApi;
            ErrorHandlingConfig = errorHandlingConfig;
            IdFieldNames = idFieldNames;
            Object = @object;
            WriteOperationType = writeOperationType;
        }
    }
}
