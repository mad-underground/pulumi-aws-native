// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue
{
    public static class GetConnection
    {
        /// <summary>
        /// Resource Type definition for AWS::Glue::Connection
        /// </summary>
        public static Task<GetConnectionResult> InvokeAsync(GetConnectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectionResult>("aws-native:glue:getConnection", args ?? new GetConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Glue::Connection
        /// </summary>
        public static Output<GetConnectionResult> Invoke(GetConnectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectionResult>("aws-native:glue:getConnection", args ?? new GetConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectionArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetConnectionArgs()
        {
        }
        public static new GetConnectionArgs Empty => new GetConnectionArgs();
    }

    public sealed class GetConnectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetConnectionInvokeArgs()
        {
        }
        public static new GetConnectionInvokeArgs Empty => new GetConnectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectionResult
    {
        public readonly Outputs.ConnectionInput? ConnectionInput;
        public readonly string? Id;

        [OutputConstructor]
        private GetConnectionResult(
            Outputs.ConnectionInput? connectionInput,

            string? id)
        {
            ConnectionInput = connectionInput;
            Id = id;
        }
    }
}
