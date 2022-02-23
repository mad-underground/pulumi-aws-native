// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DocDB
{
    public static class GetDBInstance
    {
        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBInstance
        /// </summary>
        public static Task<GetDBInstanceResult> InvokeAsync(GetDBInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDBInstanceResult>("aws-native:docdb:getDBInstance", args ?? new GetDBInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBInstance
        /// </summary>
        public static Output<GetDBInstanceResult> Invoke(GetDBInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDBInstanceResult>("aws-native:docdb:getDBInstance", args ?? new GetDBInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDBInstanceArgs()
        {
        }
    }

    public sealed class GetDBInstanceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDBInstanceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDBInstanceResult
    {
        public readonly bool? AutoMinorVersionUpgrade;
        public readonly string? DBInstanceClass;
        public readonly string? Endpoint;
        public readonly string? Id;
        public readonly string? Port;
        public readonly string? PreferredMaintenanceWindow;
        public readonly ImmutableArray<Outputs.DBInstanceTag> Tags;

        [OutputConstructor]
        private GetDBInstanceResult(
            bool? autoMinorVersionUpgrade,

            string? dBInstanceClass,

            string? endpoint,

            string? id,

            string? port,

            string? preferredMaintenanceWindow,

            ImmutableArray<Outputs.DBInstanceTag> tags)
        {
            AutoMinorVersionUpgrade = autoMinorVersionUpgrade;
            DBInstanceClass = dBInstanceClass;
            Endpoint = endpoint;
            Id = id;
            Port = port;
            PreferredMaintenanceWindow = preferredMaintenanceWindow;
            Tags = tags;
        }
    }
}