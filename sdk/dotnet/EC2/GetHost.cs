// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetHost
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::Host
        /// </summary>
        public static Task<GetHostResult> InvokeAsync(GetHostArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHostResult>("aws-native:ec2:getHost", args ?? new GetHostArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::Host
        /// </summary>
        public static Output<GetHostResult> Invoke(GetHostInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHostResult>("aws-native:ec2:getHost", args ?? new GetHostInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHostArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the host created.
        /// </summary>
        [Input("hostId", required: true)]
        public string HostId { get; set; } = null!;

        public GetHostArgs()
        {
        }
        public static new GetHostArgs Empty => new GetHostArgs();
    }

    public sealed class GetHostInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the host created.
        /// </summary>
        [Input("hostId", required: true)]
        public Input<string> HostId { get; set; } = null!;

        public GetHostInvokeArgs()
        {
        }
        public static new GetHostInvokeArgs Empty => new GetHostInvokeArgs();
    }


    [OutputType]
    public sealed class GetHostResult
    {
        /// <summary>
        /// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
        /// </summary>
        public readonly string? AutoPlacement;
        /// <summary>
        /// Id of the host created.
        /// </summary>
        public readonly string? HostId;
        /// <summary>
        /// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
        /// </summary>
        public readonly string? HostRecovery;

        [OutputConstructor]
        private GetHostResult(
            string? autoPlacement,

            string? hostId,

            string? hostRecovery)
        {
            AutoPlacement = autoPlacement;
            HostId = hostId;
            HostRecovery = hostRecovery;
        }
    }
}
