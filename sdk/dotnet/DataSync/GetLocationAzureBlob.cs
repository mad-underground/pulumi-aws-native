// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    public static class GetLocationAzureBlob
    {
        /// <summary>
        /// Resource schema for AWS::DataSync::LocationAzureBlob.
        /// </summary>
        public static Task<GetLocationAzureBlobResult> InvokeAsync(GetLocationAzureBlobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLocationAzureBlobResult>("aws-native:datasync:getLocationAzureBlob", args ?? new GetLocationAzureBlobArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::LocationAzureBlob.
        /// </summary>
        public static Output<GetLocationAzureBlobResult> Invoke(GetLocationAzureBlobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLocationAzureBlobResult>("aws-native:datasync:getLocationAzureBlob", args ?? new GetLocationAzureBlobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLocationAzureBlobArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Azure Blob Location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public string LocationArn { get; set; } = null!;

        public GetLocationAzureBlobArgs()
        {
        }
        public static new GetLocationAzureBlobArgs Empty => new GetLocationAzureBlobArgs();
    }

    public sealed class GetLocationAzureBlobInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Azure Blob Location that is created.
        /// </summary>
        [Input("locationArn", required: true)]
        public Input<string> LocationArn { get; set; } = null!;

        public GetLocationAzureBlobInvokeArgs()
        {
        }
        public static new GetLocationAzureBlobInvokeArgs Empty => new GetLocationAzureBlobInvokeArgs();
    }


    [OutputType]
    public sealed class GetLocationAzureBlobResult
    {
        /// <summary>
        /// The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.
        /// </summary>
        public readonly ImmutableArray<string> AgentArns;
        /// <summary>
        /// Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.LocationAzureBlobAzureAccessTier? AzureAccessTier;
        /// <summary>
        /// The specific authentication type that you want DataSync to use to access your Azure Blob Container.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.LocationAzureBlobAzureBlobAuthenticationType? AzureBlobAuthenticationType;
        /// <summary>
        /// Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.LocationAzureBlobAzureBlobType? AzureBlobType;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Azure Blob Location that is created.
        /// </summary>
        public readonly string? LocationArn;
        /// <summary>
        /// The URL of the Azure Blob Location that was described.
        /// </summary>
        public readonly string? LocationUri;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.LocationAzureBlobTag> Tags;

        [OutputConstructor]
        private GetLocationAzureBlobResult(
            ImmutableArray<string> agentArns,

            Pulumi.AwsNative.DataSync.LocationAzureBlobAzureAccessTier? azureAccessTier,

            Pulumi.AwsNative.DataSync.LocationAzureBlobAzureBlobAuthenticationType? azureBlobAuthenticationType,

            Pulumi.AwsNative.DataSync.LocationAzureBlobAzureBlobType? azureBlobType,

            string? locationArn,

            string? locationUri,

            ImmutableArray<Outputs.LocationAzureBlobTag> tags)
        {
            AgentArns = agentArns;
            AzureAccessTier = azureAccessTier;
            AzureBlobAuthenticationType = azureBlobAuthenticationType;
            AzureBlobType = azureBlobType;
            LocationArn = locationArn;
            LocationUri = locationUri;
            Tags = tags;
        }
    }
}
