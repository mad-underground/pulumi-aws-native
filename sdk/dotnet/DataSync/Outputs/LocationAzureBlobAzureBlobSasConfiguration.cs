// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Outputs
{

    /// <summary>
    /// Specifies the shared access signature (SAS) that DataSync uses to access your Azure Blob Storage container.
    /// </summary>
    [OutputType]
    public sealed class LocationAzureBlobAzureBlobSasConfiguration
    {
        /// <summary>
        /// Specifies the shared access signature (SAS) token, which indicates the permissions DataSync needs to access your Azure Blob Storage container.
        /// </summary>
        public readonly string AzureBlobSasToken;

        [OutputConstructor]
        private LocationAzureBlobAzureBlobSasConfiguration(string azureBlobSasToken)
        {
            AzureBlobSasToken = azureBlobSasToken;
        }
    }
}
