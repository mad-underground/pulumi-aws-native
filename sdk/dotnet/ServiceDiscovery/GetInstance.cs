// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ServiceDiscovery
{
    public static class GetInstance
    {
        /// <summary>
        /// Resource Type definition for AWS::ServiceDiscovery::Instance
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("aws-native:servicediscovery:getInstance", args ?? new GetInstanceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::ServiceDiscovery::Instance
        /// </summary>
        public static Output<GetInstanceResult> Invoke(GetInstanceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetInstanceResult>("aws-native:servicediscovery:getInstance", args ?? new GetInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public string InstanceId { get; set; } = null!;

        public GetInstanceArgs()
        {
        }
    }

    public sealed class GetInstanceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public GetInstanceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly object? InstanceAttributes;

        [OutputConstructor]
        private GetInstanceResult(object? instanceAttributes)
        {
            InstanceAttributes = instanceAttributes;
        }
    }
}