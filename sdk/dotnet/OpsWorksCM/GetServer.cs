// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorksCM
{
    public static class GetServer
    {
        /// <summary>
        /// Resource Type definition for AWS::OpsWorksCM::Server
        /// </summary>
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("aws-native:opsworkscm:getServer", args ?? new GetServerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::OpsWorksCM::Server
        /// </summary>
        public static Output<GetServerResult> Invoke(GetServerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServerResult>("aws-native:opsworkscm:getServer", args ?? new GetServerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServerArgs : global::Pulumi.InvokeArgs
    {
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetServerArgs()
        {
        }
        public static new GetServerArgs Empty => new GetServerArgs();
    }

    public sealed class GetServerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        public GetServerInvokeArgs()
        {
        }
        public static new GetServerInvokeArgs Empty => new GetServerInvokeArgs();
    }


    [OutputType]
    public sealed class GetServerResult
    {
        public readonly string? Arn;
        public readonly int? BackupRetentionCount;
        public readonly bool? DisableAutomatedBackup;
        public readonly string? Endpoint;
        public readonly string? Id;
        public readonly string? PreferredBackupWindow;
        public readonly string? PreferredMaintenanceWindow;
        public readonly ImmutableArray<Outputs.ServerTag> Tags;

        [OutputConstructor]
        private GetServerResult(
            string? arn,

            int? backupRetentionCount,

            bool? disableAutomatedBackup,

            string? endpoint,

            string? id,

            string? preferredBackupWindow,

            string? preferredMaintenanceWindow,

            ImmutableArray<Outputs.ServerTag> tags)
        {
            Arn = arn;
            BackupRetentionCount = backupRetentionCount;
            DisableAutomatedBackup = disableAutomatedBackup;
            Endpoint = endpoint;
            Id = id;
            PreferredBackupWindow = preferredBackupWindow;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            Tags = tags;
        }
    }
}
