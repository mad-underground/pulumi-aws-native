// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT
{
    public static class GetCACertificate
    {
        /// <summary>
        /// Registers a CA Certificate in IoT.
        /// </summary>
        public static Task<GetCACertificateResult> InvokeAsync(GetCACertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCACertificateResult>("aws-native:iot:getCACertificate", args ?? new GetCACertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Registers a CA Certificate in IoT.
        /// </summary>
        public static Output<GetCACertificateResult> Invoke(GetCACertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCACertificateResult>("aws-native:iot:getCACertificate", args ?? new GetCACertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCACertificateArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCACertificateArgs()
        {
        }
        public static new GetCACertificateArgs Empty => new GetCACertificateArgs();
    }

    public sealed class GetCACertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCACertificateInvokeArgs()
        {
        }
        public static new GetCACertificateInvokeArgs Empty => new GetCACertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCACertificateResult
    {
        public readonly string? Arn;
        public readonly Pulumi.AwsNative.IoT.CACertificateAutoRegistrationStatus? AutoRegistrationStatus;
        public readonly string? Id;
        public readonly Outputs.CACertificateRegistrationConfig? RegistrationConfig;
        public readonly bool? RemoveAutoRegistration;
        public readonly Pulumi.AwsNative.IoT.CACertificateStatus? Status;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.CACertificateTag> Tags;

        [OutputConstructor]
        private GetCACertificateResult(
            string? arn,

            Pulumi.AwsNative.IoT.CACertificateAutoRegistrationStatus? autoRegistrationStatus,

            string? id,

            Outputs.CACertificateRegistrationConfig? registrationConfig,

            bool? removeAutoRegistration,

            Pulumi.AwsNative.IoT.CACertificateStatus? status,

            ImmutableArray<Outputs.CACertificateTag> tags)
        {
            Arn = arn;
            AutoRegistrationStatus = autoRegistrationStatus;
            Id = id;
            RegistrationConfig = registrationConfig;
            RemoveAutoRegistration = removeAutoRegistration;
            Status = status;
            Tags = tags;
        }
    }
}
