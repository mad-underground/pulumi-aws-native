// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RDS
{
    public static class GetDBParameterGroup
    {
        /// <summary>
        /// The AWS::RDS::DBParameterGroup resource creates a custom parameter group for an RDS database family
        /// </summary>
        public static Task<GetDBParameterGroupResult> InvokeAsync(GetDBParameterGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDBParameterGroupResult>("aws-native:rds:getDBParameterGroup", args ?? new GetDBParameterGroupArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::RDS::DBParameterGroup resource creates a custom parameter group for an RDS database family
        /// </summary>
        public static Output<GetDBParameterGroupResult> Invoke(GetDBParameterGroupInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDBParameterGroupResult>("aws-native:rds:getDBParameterGroup", args ?? new GetDBParameterGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDBParameterGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the DB parameter group
        /// </summary>
        [Input("dBParameterGroupName", required: true)]
        public string DBParameterGroupName { get; set; } = null!;

        public GetDBParameterGroupArgs()
        {
        }
    }

    public sealed class GetDBParameterGroupInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the name of the DB parameter group
        /// </summary>
        [Input("dBParameterGroupName", required: true)]
        public Input<string> DBParameterGroupName { get; set; } = null!;

        public GetDBParameterGroupInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDBParameterGroupResult
    {
        /// <summary>
        /// Specifies the name of the DB parameter group
        /// </summary>
        public readonly string? DBParameterGroupName;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.DBParameterGroupTag> Tags;

        [OutputConstructor]
        private GetDBParameterGroupResult(
            string? dBParameterGroupName,

            ImmutableArray<Outputs.DBParameterGroupTag> tags)
        {
            DBParameterGroupName = dBParameterGroupName;
            Tags = tags;
        }
    }
}
