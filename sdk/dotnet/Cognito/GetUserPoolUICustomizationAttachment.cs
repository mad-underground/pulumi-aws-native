// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    public static class GetUserPoolUICustomizationAttachment
    {
        /// <summary>
        /// Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
        /// </summary>
        public static Task<GetUserPoolUICustomizationAttachmentResult> InvokeAsync(GetUserPoolUICustomizationAttachmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserPoolUICustomizationAttachmentResult>("aws-native:cognito:getUserPoolUICustomizationAttachment", args ?? new GetUserPoolUICustomizationAttachmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment
        /// </summary>
        public static Output<GetUserPoolUICustomizationAttachmentResult> Invoke(GetUserPoolUICustomizationAttachmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUserPoolUICustomizationAttachmentResult>("aws-native:cognito:getUserPoolUICustomizationAttachment", args ?? new GetUserPoolUICustomizationAttachmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserPoolUICustomizationAttachmentArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetUserPoolUICustomizationAttachmentArgs()
        {
        }
    }

    public sealed class GetUserPoolUICustomizationAttachmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetUserPoolUICustomizationAttachmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserPoolUICustomizationAttachmentResult
    {
        public readonly string? CSS;
        public readonly string? Id;

        [OutputConstructor]
        private GetUserPoolUICustomizationAttachmentResult(
            string? cSS,

            string? id)
        {
            CSS = cSS;
            Id = id;
        }
    }
}