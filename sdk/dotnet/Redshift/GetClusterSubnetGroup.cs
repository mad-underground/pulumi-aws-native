// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift
{
    public static class GetClusterSubnetGroup
    {
        /// <summary>
        /// Specifies an Amazon Redshift subnet group.
        /// </summary>
        public static Task<GetClusterSubnetGroupResult> InvokeAsync(GetClusterSubnetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetClusterSubnetGroupResult>("aws-native:redshift:getClusterSubnetGroup", args ?? new GetClusterSubnetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies an Amazon Redshift subnet group.
        /// </summary>
        public static Output<GetClusterSubnetGroupResult> Invoke(GetClusterSubnetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetClusterSubnetGroupResult>("aws-native:redshift:getClusterSubnetGroup", args ?? new GetClusterSubnetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetClusterSubnetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default". 
        /// </summary>
        [Input("clusterSubnetGroupName", required: true)]
        public string ClusterSubnetGroupName { get; set; } = null!;

        public GetClusterSubnetGroupArgs()
        {
        }
        public static new GetClusterSubnetGroupArgs Empty => new GetClusterSubnetGroupArgs();
    }

    public sealed class GetClusterSubnetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default". 
        /// </summary>
        [Input("clusterSubnetGroupName", required: true)]
        public Input<string> ClusterSubnetGroupName { get; set; } = null!;

        public GetClusterSubnetGroupInvokeArgs()
        {
        }
        public static new GetClusterSubnetGroupInvokeArgs Empty => new GetClusterSubnetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetClusterSubnetGroupResult
    {
        /// <summary>
        /// This name must be unique for all subnet groups that are created by your AWS account. If costumer do not provide it, cloudformation will generate it. Must not be "Default". 
        /// </summary>
        public readonly string? ClusterSubnetGroupName;
        /// <summary>
        /// The description of the parameter group.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The list of VPC subnet IDs
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The list of tags for the cluster parameter group.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterSubnetGroupTag> Tags;

        [OutputConstructor]
        private GetClusterSubnetGroupResult(
            string? clusterSubnetGroupName,

            string? description,

            ImmutableArray<string> subnetIds,

            ImmutableArray<Outputs.ClusterSubnetGroupTag> tags)
        {
            ClusterSubnetGroupName = clusterSubnetGroupName;
            Description = description;
            SubnetIds = subnetIds;
            Tags = tags;
        }
    }
}
