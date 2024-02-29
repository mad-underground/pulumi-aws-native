// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetEip
    {
        /// <summary>
        /// Specifies an Elastic IP (EIP) address and can, optionally, associate it with an Amazon EC2 instance.
        ///  You can allocate an Elastic IP address from an address pool owned by AWS or from an address pool created from a public IPv4 address range that you have brought to AWS for use with your AWS resources using bring your own IP addresses (BYOIP). For more information, see [Bring Your Own IP Addresses (BYOIP)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the *Amazon EC2 User Guide*.
        ///  For more information, see [Elastic IP Addresses](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) in the *Amazon EC2 User Guide*.
        /// </summary>
        public static Task<GetEipResult> InvokeAsync(GetEipArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEipResult>("aws-native:ec2:getEip", args ?? new GetEipArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies an Elastic IP (EIP) address and can, optionally, associate it with an Amazon EC2 instance.
        ///  You can allocate an Elastic IP address from an address pool owned by AWS or from an address pool created from a public IPv4 address range that you have brought to AWS for use with your AWS resources using bring your own IP addresses (BYOIP). For more information, see [Bring Your Own IP Addresses (BYOIP)](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html) in the *Amazon EC2 User Guide*.
        ///  For more information, see [Elastic IP Addresses](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html) in the *Amazon EC2 User Guide*.
        /// </summary>
        public static Output<GetEipResult> Invoke(GetEipInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEipResult>("aws-native:ec2:getEip", args ?? new GetEipInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEipArgs : global::Pulumi.InvokeArgs
    {
        [Input("allocationId", required: true)]
        public string AllocationId { get; set; } = null!;

        [Input("publicIp", required: true)]
        public string PublicIp { get; set; } = null!;

        public GetEipArgs()
        {
        }
        public static new GetEipArgs Empty => new GetEipArgs();
    }

    public sealed class GetEipInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("allocationId", required: true)]
        public Input<string> AllocationId { get; set; } = null!;

        [Input("publicIp", required: true)]
        public Input<string> PublicIp { get; set; } = null!;

        public GetEipInvokeArgs()
        {
        }
        public static new GetEipInvokeArgs Empty => new GetEipInvokeArgs();
    }


    [OutputType]
    public sealed class GetEipResult
    {
        public readonly string? AllocationId;
        /// <summary>
        /// The ID of the instance.
        ///   Updates to the ``InstanceId`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
        /// </summary>
        public readonly string? InstanceId;
        public readonly string? PublicIp;
        /// <summary>
        /// The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool.
        ///   Updates to the ``PublicIpv4Pool`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
        /// </summary>
        public readonly string? PublicIpv4Pool;
        /// <summary>
        /// Any tags assigned to the Elastic IP address.
        ///   Updates to the ``Tags`` property may require *some interruptions*. Updates on an EIP reassociates the address on its associated resource.
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetEipResult(
            string? allocationId,

            string? instanceId,

            string? publicIp,

            string? publicIpv4Pool,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            AllocationId = allocationId;
            InstanceId = instanceId;
            PublicIp = publicIp;
            PublicIpv4Pool = publicIpv4Pool;
            Tags = tags;
        }
    }
}
