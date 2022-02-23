// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    public static class GetTagOption
    {
        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::TagOption
        /// </summary>
        public static Task<GetTagOptionResult> InvokeAsync(GetTagOptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagOptionResult>("aws-native:servicecatalog:getTagOption", args ?? new GetTagOptionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::TagOption
        /// </summary>
        public static Output<GetTagOptionResult> Invoke(GetTagOptionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTagOptionResult>("aws-native:servicecatalog:getTagOption", args ?? new GetTagOptionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTagOptionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetTagOptionArgs()
        {
        }
    }

    public sealed class GetTagOptionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetTagOptionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagOptionResult
    {
        public readonly bool? Active;
        public readonly string? Id;

        [OutputConstructor]
        private GetTagOptionResult(
            bool? active,

            string? id)
        {
            Active = active;
            Id = id;
        }
    }
}