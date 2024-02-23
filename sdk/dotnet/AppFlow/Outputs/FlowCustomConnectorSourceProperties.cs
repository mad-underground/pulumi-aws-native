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
    public sealed class FlowCustomConnectorSourceProperties
    {
        public readonly ImmutableDictionary<string, string>? CustomProperties;
        public readonly Outputs.FlowCustomConnectorSourcePropertiesDataTransferApiProperties? DataTransferApi;
        public readonly string EntityName;

        [OutputConstructor]
        private FlowCustomConnectorSourceProperties(
            ImmutableDictionary<string, string>? customProperties,

            Outputs.FlowCustomConnectorSourcePropertiesDataTransferApiProperties? dataTransferApi,

            string entityName)
        {
            CustomProperties = customProperties;
            DataTransferApi = dataTransferApi;
            EntityName = entityName;
        }
    }
}
