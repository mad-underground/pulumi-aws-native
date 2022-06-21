// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GlobalAccelerator
{
    public static class GetAccelerator
    {
        /// <summary>
        /// Resource Type definition for AWS::GlobalAccelerator::Accelerator
        /// </summary>
        public static Task<GetAcceleratorResult> InvokeAsync(GetAcceleratorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAcceleratorResult>("aws-native:globalaccelerator:getAccelerator", args ?? new GetAcceleratorArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::GlobalAccelerator::Accelerator
        /// </summary>
        public static Output<GetAcceleratorResult> Invoke(GetAcceleratorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAcceleratorResult>("aws-native:globalaccelerator:getAccelerator", args ?? new GetAcceleratorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAcceleratorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the accelerator.
        /// </summary>
        [Input("acceleratorArn", required: true)]
        public string AcceleratorArn { get; set; } = null!;

        public GetAcceleratorArgs()
        {
        }
    }

    public sealed class GetAcceleratorInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the accelerator.
        /// </summary>
        [Input("acceleratorArn", required: true)]
        public Input<string> AcceleratorArn { get; set; } = null!;

        public GetAcceleratorInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAcceleratorResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the accelerator.
        /// </summary>
        public readonly string? AcceleratorArn;
        /// <summary>
        /// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
        /// </summary>
        public readonly string? DnsName;
        /// <summary>
        /// Indicates whether an accelerator is enabled. The value is true or false.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// IP Address type.
        /// </summary>
        public readonly Pulumi.AwsNative.GlobalAccelerator.AcceleratorIpAddressType? IpAddressType;
        /// <summary>
        /// The IP addresses from BYOIP Prefix pool.
        /// </summary>
        public readonly ImmutableArray<string> IpAddresses;
        /// <summary>
        /// The IPv4 addresses assigned to the accelerator.
        /// </summary>
        public readonly ImmutableArray<string> Ipv4Addresses;
        /// <summary>
        /// Name of accelerator.
        /// </summary>
        public readonly string? Name;
        public readonly ImmutableArray<Outputs.AcceleratorTag> Tags;

        [OutputConstructor]
        private GetAcceleratorResult(
            string? acceleratorArn,

            string? dnsName,

            bool? enabled,

            Pulumi.AwsNative.GlobalAccelerator.AcceleratorIpAddressType? ipAddressType,

            ImmutableArray<string> ipAddresses,

            ImmutableArray<string> ipv4Addresses,

            string? name,

            ImmutableArray<Outputs.AcceleratorTag> tags)
        {
            AcceleratorArn = acceleratorArn;
            DnsName = dnsName;
            Enabled = enabled;
            IpAddressType = ipAddressType;
            IpAddresses = ipAddresses;
            Ipv4Addresses = ipv4Addresses;
            Name = name;
            Tags = tags;
        }
    }
}
