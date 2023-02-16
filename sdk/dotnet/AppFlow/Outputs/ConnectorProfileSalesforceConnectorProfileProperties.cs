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
    public sealed class ConnectorProfileSalesforceConnectorProfileProperties
    {
        /// <summary>
        /// The location of the Salesforce resource
        /// </summary>
        public readonly string? InstanceUrl;
        /// <summary>
        /// Indicates whether the connector profile applies to a sandbox or production environment
        /// </summary>
        public readonly bool? IsSandboxEnvironment;

        [OutputConstructor]
        private ConnectorProfileSalesforceConnectorProfileProperties(
            string? instanceUrl,

            bool? isSandboxEnvironment)
        {
            InstanceUrl = instanceUrl;
            IsSandboxEnvironment = isSandboxEnvironment;
        }
    }
}
