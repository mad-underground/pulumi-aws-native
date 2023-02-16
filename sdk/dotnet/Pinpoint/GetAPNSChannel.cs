// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    public static class GetAPNSChannel
    {
        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::APNSChannel
        /// </summary>
        public static Task<GetAPNSChannelResult> InvokeAsync(GetAPNSChannelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAPNSChannelResult>("aws-native:pinpoint:getAPNSChannel", args ?? new GetAPNSChannelArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Pinpoint::APNSChannel
        /// </summary>
        public static Output<GetAPNSChannelResult> Invoke(GetAPNSChannelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAPNSChannelResult>("aws-native:pinpoint:getAPNSChannel", args ?? new GetAPNSChannelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAPNSChannelArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetAPNSChannelArgs()
        {
        }
        public static new GetAPNSChannelArgs Empty => new GetAPNSChannelArgs();
    }

    public sealed class GetAPNSChannelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetAPNSChannelInvokeArgs()
        {
        }
        public static new GetAPNSChannelInvokeArgs Empty => new GetAPNSChannelInvokeArgs();
    }


    [OutputType]
    public sealed class GetAPNSChannelResult
    {
        public readonly string? BundleId;
        public readonly string? Certificate;
        public readonly string? DefaultAuthenticationMethod;
        public readonly bool? Enabled;
        public readonly string? Id;
        public readonly string? PrivateKey;
        public readonly string? TeamId;
        public readonly string? TokenKey;
        public readonly string? TokenKeyId;

        [OutputConstructor]
        private GetAPNSChannelResult(
            string? bundleId,

            string? certificate,

            string? defaultAuthenticationMethod,

            bool? enabled,

            string? id,

            string? privateKey,

            string? teamId,

            string? tokenKey,

            string? tokenKeyId)
        {
            BundleId = bundleId;
            Certificate = certificate;
            DefaultAuthenticationMethod = defaultAuthenticationMethod;
            Enabled = enabled;
            Id = id;
            PrivateKey = privateKey;
            TeamId = teamId;
            TokenKey = tokenKey;
            TokenKeyId = tokenKeyId;
        }
    }
}
