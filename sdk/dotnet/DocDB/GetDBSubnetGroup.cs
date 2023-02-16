// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DocDB
{
    public static class GetDBSubnetGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBSubnetGroup
        /// </summary>
        public static Task<GetDBSubnetGroupResult> InvokeAsync(GetDBSubnetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDBSubnetGroupResult>("aws-native:docdb:getDBSubnetGroup", args ?? new GetDBSubnetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::DocDB::DBSubnetGroup
        /// </summary>
        public static Output<GetDBSubnetGroupResult> Invoke(GetDBSubnetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDBSubnetGroupResult>("aws-native:docdb:getDBSubnetGroup", args ?? new GetDBSubnetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBSubnetGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDBSubnetGroupArgs()
        {
        }
        public static new GetDBSubnetGroupArgs Empty => new GetDBSubnetGroupArgs();
    }

    public sealed class GetDBSubnetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDBSubnetGroupInvokeArgs()
        {
        }
        public static new GetDBSubnetGroupInvokeArgs Empty => new GetDBSubnetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetDBSubnetGroupResult
    {
        public readonly string? DBSubnetGroupDescription;
        public readonly string? Id;
        public readonly ImmutableArray<string> SubnetIds;
        public readonly ImmutableArray<Outputs.DBSubnetGroupTag> Tags;

        [OutputConstructor]
        private GetDBSubnetGroupResult(
            string? dBSubnetGroupDescription,

            string? id,

            ImmutableArray<string> subnetIds,

            ImmutableArray<Outputs.DBSubnetGroupTag> tags)
        {
            DBSubnetGroupDescription = dBSubnetGroupDescription;
            Id = id;
            SubnetIds = subnetIds;
            Tags = tags;
        }
    }
}
