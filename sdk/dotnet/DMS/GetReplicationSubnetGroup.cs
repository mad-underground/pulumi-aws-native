// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DMS
{
    public static class GetReplicationSubnetGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::DMS::ReplicationSubnetGroup
        /// </summary>
        public static Task<GetReplicationSubnetGroupResult> InvokeAsync(GetReplicationSubnetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReplicationSubnetGroupResult>("aws-native:dms:getReplicationSubnetGroup", args ?? new GetReplicationSubnetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DMS::ReplicationSubnetGroup
        /// </summary>
        public static Output<GetReplicationSubnetGroupResult> Invoke(GetReplicationSubnetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetReplicationSubnetGroupResult>("aws-native:dms:getReplicationSubnetGroup", args ?? new GetReplicationSubnetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplicationSubnetGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetReplicationSubnetGroupArgs()
        {
        }
        public static new GetReplicationSubnetGroupArgs Empty => new GetReplicationSubnetGroupArgs();
    }

    public sealed class GetReplicationSubnetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetReplicationSubnetGroupInvokeArgs()
        {
        }
        public static new GetReplicationSubnetGroupInvokeArgs Empty => new GetReplicationSubnetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetReplicationSubnetGroupResult
    {
        public readonly string? Id;
        public readonly string? ReplicationSubnetGroupDescription;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly ImmutableArray<Outputs.ReplicationSubnetGroupTag> Tags;

        [OutputConstructor]
        private GetReplicationSubnetGroupResult(
            string? id,

            string? replicationSubnetGroupDescription,

            ImmutableArray<string> subnetIds,

            ImmutableArray<Outputs.ReplicationSubnetGroupTag> tags)
        {
            Id = id;
            ReplicationSubnetGroupDescription = replicationSubnetGroupDescription;
            SubnetIds = subnetIds;
            Tags = tags;
        }
    }
}
