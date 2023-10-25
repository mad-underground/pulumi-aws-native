// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetUsagePlanKey
    {
        /// <summary>
        /// The ``AWS::ApiGateway::UsagePlanKey`` resource associates an API key with a usage plan. This association determines which users the usage plan is applied to.
        /// </summary>
        public static Task<GetUsagePlanKeyResult> InvokeAsync(GetUsagePlanKeyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsagePlanKeyResult>("aws-native:apigateway:getUsagePlanKey", args ?? new GetUsagePlanKeyArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::ApiGateway::UsagePlanKey`` resource associates an API key with a usage plan. This association determines which users the usage plan is applied to.
        /// </summary>
        public static Output<GetUsagePlanKeyResult> Invoke(GetUsagePlanKeyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsagePlanKeyResult>("aws-native:apigateway:getUsagePlanKey", args ?? new GetUsagePlanKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsagePlanKeyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetUsagePlanKeyArgs()
        {
        }
        public static new GetUsagePlanKeyArgs Empty => new GetUsagePlanKeyArgs();
    }

    public sealed class GetUsagePlanKeyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetUsagePlanKeyInvokeArgs()
        {
        }
        public static new GetUsagePlanKeyInvokeArgs Empty => new GetUsagePlanKeyInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsagePlanKeyResult
    {
        /// <summary>
        /// An autogenerated ID which is a combination of the ID of the key and ID of the usage plan combined with a : such as 123abcdef:abc123.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private GetUsagePlanKeyResult(string? id)
        {
            Id = id;
        }
    }
}
