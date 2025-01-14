// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceCatalog
{
    public static class GetPortfolioShare
    {
        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::PortfolioShare
        /// </summary>
        public static Task<GetPortfolioShareResult> InvokeAsync(GetPortfolioShareArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPortfolioShareResult>("aws-native:servicecatalog:getPortfolioShare", args ?? new GetPortfolioShareArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ServiceCatalog::PortfolioShare
        /// </summary>
        public static Output<GetPortfolioShareResult> Invoke(GetPortfolioShareInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPortfolioShareResult>("aws-native:servicecatalog:getPortfolioShare", args ?? new GetPortfolioShareInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPortfolioShareArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetPortfolioShareArgs()
        {
        }
        public static new GetPortfolioShareArgs Empty => new GetPortfolioShareArgs();
    }

    public sealed class GetPortfolioShareInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetPortfolioShareInvokeArgs()
        {
        }
        public static new GetPortfolioShareInvokeArgs Empty => new GetPortfolioShareInvokeArgs();
    }


    [OutputType]
    public sealed class GetPortfolioShareResult
    {
        public readonly string? Id;
        public readonly bool? ShareTagOptions;

        [OutputConstructor]
        private GetPortfolioShareResult(
            string? id,

            bool? shareTagOptions)
        {
            Id = id;
            ShareTagOptions = shareTagOptions;
        }
    }
}
