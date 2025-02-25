// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam
{
    public static class GetSamlProvider
    {
        /// <summary>
        /// Resource Type definition for AWS::IAM::SAMLProvider
        /// </summary>
        public static Task<GetSamlProviderResult> InvokeAsync(GetSamlProviderArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSamlProviderResult>("aws-native:iam:getSamlProvider", args ?? new GetSamlProviderArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::IAM::SAMLProvider
        /// </summary>
        public static Output<GetSamlProviderResult> Invoke(GetSamlProviderInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSamlProviderResult>("aws-native:iam:getSamlProvider", args ?? new GetSamlProviderInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSamlProviderArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the SAML provider
        /// </summary>
        [Input("arn", required: true)]
        public string Arn { get; set; } = null!;

        public GetSamlProviderArgs()
        {
        }
        public static new GetSamlProviderArgs Empty => new GetSamlProviderArgs();
    }

    public sealed class GetSamlProviderInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the SAML provider
        /// </summary>
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        public GetSamlProviderInvokeArgs()
        {
        }
        public static new GetSamlProviderInvokeArgs Empty => new GetSamlProviderInvokeArgs();
    }


    [OutputType]
    public sealed class GetSamlProviderResult
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of the SAML provider
        /// </summary>
        public readonly string? Arn;
        public readonly string? SamlMetadataDocument;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetSamlProviderResult(
            string? arn,

            string? samlMetadataDocument,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            Arn = arn;
            SamlMetadataDocument = samlMetadataDocument;
            Tags = tags;
        }
    }
}
