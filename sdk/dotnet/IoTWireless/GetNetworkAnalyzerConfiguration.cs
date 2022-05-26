// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless
{
    public static class GetNetworkAnalyzerConfiguration
    {
        /// <summary>
        /// Create and manage NetworkAnalyzerConfiguration resource.
        /// </summary>
        public static Task<GetNetworkAnalyzerConfigurationResult> InvokeAsync(GetNetworkAnalyzerConfigurationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkAnalyzerConfigurationResult>("aws-native:iotwireless:getNetworkAnalyzerConfiguration", args ?? new GetNetworkAnalyzerConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Create and manage NetworkAnalyzerConfiguration resource.
        /// </summary>
        public static Output<GetNetworkAnalyzerConfigurationResult> Invoke(GetNetworkAnalyzerConfigurationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkAnalyzerConfigurationResult>("aws-native:iotwireless:getNetworkAnalyzerConfiguration", args ?? new GetNetworkAnalyzerConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkAnalyzerConfigurationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the network analyzer configuration
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetNetworkAnalyzerConfigurationArgs()
        {
        }
    }

    public sealed class GetNetworkAnalyzerConfigurationInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the network analyzer configuration
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetNetworkAnalyzerConfigurationInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkAnalyzerConfigurationResult
    {
        /// <summary>
        /// Arn for network analyzer configuration, Returned upon successful create.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// The description of the new resource
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Trace content for your wireless gateway and wireless device resources
        /// </summary>
        public readonly Outputs.TraceContentProperties? TraceContent;
        /// <summary>
        /// List of wireless gateway resources that have been added to the network analyzer configuration
        /// </summary>
        public readonly ImmutableArray<string> WirelessDevices;
        /// <summary>
        /// List of wireless gateway resources that have been added to the network analyzer configuration
        /// </summary>
        public readonly ImmutableArray<string> WirelessGateways;

        [OutputConstructor]
        private GetNetworkAnalyzerConfigurationResult(
            string? arn,

            string? description,

            Outputs.TraceContentProperties? traceContent,

            ImmutableArray<string> wirelessDevices,

            ImmutableArray<string> wirelessGateways)
        {
            Arn = arn;
            Description = description;
            TraceContent = traceContent;
            WirelessDevices = wirelessDevices;
            WirelessGateways = wirelessGateways;
        }
    }
}
