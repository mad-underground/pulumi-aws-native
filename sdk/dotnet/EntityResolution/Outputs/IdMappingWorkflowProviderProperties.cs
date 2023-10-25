// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Outputs
{

    [OutputType]
    public sealed class IdMappingWorkflowProviderProperties
    {
        public readonly Outputs.IdMappingWorkflowIntermediateSourceConfiguration? IntermediateSourceConfiguration;
        /// <summary>
        /// Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format
        /// </summary>
        public readonly object? ProviderConfiguration;
        /// <summary>
        /// Arn of the Provider Service being used.
        /// </summary>
        public readonly string ProviderServiceArn;

        [OutputConstructor]
        private IdMappingWorkflowProviderProperties(
            Outputs.IdMappingWorkflowIntermediateSourceConfiguration? intermediateSourceConfiguration,

            object? providerConfiguration,

            string providerServiceArn)
        {
            IntermediateSourceConfiguration = intermediateSourceConfiguration;
            ProviderConfiguration = providerConfiguration;
            ProviderServiceArn = providerServiceArn;
        }
    }
}
