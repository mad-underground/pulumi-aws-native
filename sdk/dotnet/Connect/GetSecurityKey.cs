// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect
{
    public static class GetSecurityKey
    {
        /// <summary>
        /// Resource Type definition for AWS::Connect::SecurityKey
        /// </summary>
        public static Task<GetSecurityKeyResult> InvokeAsync(GetSecurityKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityKeyResult>("aws-native:connect:getSecurityKey", args ?? new GetSecurityKeyArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Connect::SecurityKey
        /// </summary>
        public static Output<GetSecurityKeyResult> Invoke(GetSecurityKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityKeyResult>("aws-native:connect:getSecurityKey", args ?? new GetSecurityKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityKeyArgs : global::Pulumi.InvokeArgs
    {
        [Input("associationId", required: true)]
        public string AssociationId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        public GetSecurityKeyArgs()
        {
        }
        public static new GetSecurityKeyArgs Empty => new GetSecurityKeyArgs();
    }

    public sealed class GetSecurityKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("associationId", required: true)]
        public Input<string> AssociationId { get; set; } = null!;

        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public GetSecurityKeyInvokeArgs()
        {
        }
        public static new GetSecurityKeyInvokeArgs Empty => new GetSecurityKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityKeyResult
    {
        public readonly string? AssociationId;

        [OutputConstructor]
        private GetSecurityKeyResult(string? associationId)
        {
            AssociationId = associationId;
        }
    }
}
