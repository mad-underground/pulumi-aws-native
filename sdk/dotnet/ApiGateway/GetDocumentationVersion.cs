// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ApiGateway
{
    public static class GetDocumentationVersion
    {
        /// <summary>
        /// A snapshot of the documentation of an API.
        /// </summary>
        public static Task<GetDocumentationVersionResult> InvokeAsync(GetDocumentationVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDocumentationVersionResult>("aws-native:apigateway:getDocumentationVersion", args ?? new GetDocumentationVersionArgs(), options.WithDefaults());

        /// <summary>
        /// A snapshot of the documentation of an API.
        /// </summary>
        public static Output<GetDocumentationVersionResult> Invoke(GetDocumentationVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDocumentationVersionResult>("aws-native:apigateway:getDocumentationVersion", args ?? new GetDocumentationVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDocumentationVersionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The version identifier of the API documentation snapshot.
        /// </summary>
        [Input("documentationVersion", required: true)]
        public string DocumentationVersionValue { get; set; } = null!;

        /// <summary>
        /// The identifier of the API.
        /// </summary>
        [Input("restApiId", required: true)]
        public string RestApiId { get; set; } = null!;

        public GetDocumentationVersionArgs()
        {
        }
    }

    public sealed class GetDocumentationVersionInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The version identifier of the API documentation snapshot.
        /// </summary>
        [Input("documentationVersion", required: true)]
        public Input<string> DocumentationVersion { get; set; } = null!;

        /// <summary>
        /// The identifier of the API.
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public GetDocumentationVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDocumentationVersionResult
    {
        /// <summary>
        /// The description of the API documentation snapshot.
        /// </summary>
        public readonly string? Description;

        [OutputConstructor]
        private GetDocumentationVersionResult(string? description)
        {
            Description = description;
        }
    }
}