// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS
{
    public static class GetDBProxyTargetGroup
    {
        /// <summary>
        /// Resource schema for AWS::RDS::DBProxyTargetGroup
        /// </summary>
        public static Task<GetDBProxyTargetGroupResult> InvokeAsync(GetDBProxyTargetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDBProxyTargetGroupResult>("aws-native:rds:getDBProxyTargetGroup", args ?? new GetDBProxyTargetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource schema for AWS::RDS::DBProxyTargetGroup
        /// </summary>
        public static Output<GetDBProxyTargetGroupResult> Invoke(GetDBProxyTargetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDBProxyTargetGroupResult>("aws-native:rds:getDBProxyTargetGroup", args ?? new GetDBProxyTargetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBProxyTargetGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the proxy.
        /// </summary>
        [Input("dBProxyName", required: true)]
        public string DBProxyName { get; set; } = null!;

        public GetDBProxyTargetGroupArgs()
        {
        }
        public static new GetDBProxyTargetGroupArgs Empty => new GetDBProxyTargetGroupArgs();
    }

    public sealed class GetDBProxyTargetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The identifier for the proxy.
        /// </summary>
        [Input("dBProxyName", required: true)]
        public Input<string> DBProxyName { get; set; } = null!;

        public GetDBProxyTargetGroupInvokeArgs()
        {
        }
        public static new GetDBProxyTargetGroupInvokeArgs Empty => new GetDBProxyTargetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetDBProxyTargetGroupResult
    {
        public readonly Outputs.DBProxyTargetGroupConnectionPoolConfigurationInfoFormat? ConnectionPoolConfigurationInfo;
        public readonly ImmutableArray<string> DBClusterIdentifiers;
        public readonly ImmutableArray<string> DBInstanceIdentifiers;
        /// <summary>
        /// The identifier for the proxy.
        /// </summary>
        public readonly string? DBProxyName;
        /// <summary>
        /// The Amazon Resource Name (ARN) representing the target group.
        /// </summary>
        public readonly string? TargetGroupArn;

        [OutputConstructor]
        private GetDBProxyTargetGroupResult(
            Outputs.DBProxyTargetGroupConnectionPoolConfigurationInfoFormat? connectionPoolConfigurationInfo,

            ImmutableArray<string> dBClusterIdentifiers,

            ImmutableArray<string> dBInstanceIdentifiers,

            string? dBProxyName,

            string? targetGroupArn)
        {
            ConnectionPoolConfigurationInfo = connectionPoolConfigurationInfo;
            DBClusterIdentifiers = dBClusterIdentifiers;
            DBInstanceIdentifiers = dBInstanceIdentifiers;
            DBProxyName = dBProxyName;
            TargetGroupArn = targetGroupArn;
        }
    }
}
