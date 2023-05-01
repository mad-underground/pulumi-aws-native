// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync
{
    public static class GetStorageSystem
    {
        /// <summary>
        /// Resource schema for AWS::DataSync::StorageSystem.
        /// </summary>
        public static Task<GetStorageSystemResult> InvokeAsync(GetStorageSystemArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetStorageSystemResult>("aws-native:datasync:getStorageSystem", args ?? new GetStorageSystemArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::DataSync::StorageSystem.
        /// </summary>
        public static Output<GetStorageSystemResult> Invoke(GetStorageSystemInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetStorageSystemResult>("aws-native:datasync:getStorageSystem", args ?? new GetStorageSystemInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetStorageSystemArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the on-premises storage system added to DataSync Discovery.
        /// </summary>
        [Input("storageSystemArn", required: true)]
        public string StorageSystemArn { get; set; } = null!;

        public GetStorageSystemArgs()
        {
        }
        public static new GetStorageSystemArgs Empty => new GetStorageSystemArgs();
    }

    public sealed class GetStorageSystemInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ARN of the on-premises storage system added to DataSync Discovery.
        /// </summary>
        [Input("storageSystemArn", required: true)]
        public Input<string> StorageSystemArn { get; set; } = null!;

        public GetStorageSystemInvokeArgs()
        {
        }
        public static new GetStorageSystemInvokeArgs Empty => new GetStorageSystemInvokeArgs();
    }


    [OutputType]
    public sealed class GetStorageSystemResult
    {
        /// <summary>
        /// The ARN of the DataSync agent that connects to and reads from the on-premises storage system's management interface.
        /// </summary>
        public readonly ImmutableArray<string> AgentArns;
        /// <summary>
        /// The ARN of the Amazon CloudWatch log group used to monitor and log discovery job events.
        /// </summary>
        public readonly string? CloudWatchLogGroupArn;
        /// <summary>
        /// Indicates whether the DataSync agent can access the on-premises storage system.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.StorageSystemConnectivityStatus? ConnectivityStatus;
        /// <summary>
        /// A familiar name for the on-premises storage system.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The ARN of a secret stored by AWS Secrets Manager.
        /// </summary>
        public readonly string? SecretsManagerArn;
        public readonly Outputs.StorageSystemServerConfiguration? ServerConfiguration;
        /// <summary>
        /// The ARN of the on-premises storage system added to DataSync Discovery.
        /// </summary>
        public readonly string? StorageSystemArn;
        /// <summary>
        /// The type of on-premises storage system that DataSync Discovery will analyze.
        /// </summary>
        public readonly Pulumi.AwsNative.DataSync.StorageSystemSystemType? SystemType;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageSystemTag> Tags;

        [OutputConstructor]
        private GetStorageSystemResult(
            ImmutableArray<string> agentArns,

            string? cloudWatchLogGroupArn,

            Pulumi.AwsNative.DataSync.StorageSystemConnectivityStatus? connectivityStatus,

            string? name,

            string? secretsManagerArn,

            Outputs.StorageSystemServerConfiguration? serverConfiguration,

            string? storageSystemArn,

            Pulumi.AwsNative.DataSync.StorageSystemSystemType? systemType,

            ImmutableArray<Outputs.StorageSystemTag> tags)
        {
            AgentArns = agentArns;
            CloudWatchLogGroupArn = cloudWatchLogGroupArn;
            ConnectivityStatus = connectivityStatus;
            Name = name;
            SecretsManagerArn = secretsManagerArn;
            ServerConfiguration = serverConfiguration;
            StorageSystemArn = storageSystemArn;
            SystemType = systemType;
            Tags = tags;
        }
    }
}
