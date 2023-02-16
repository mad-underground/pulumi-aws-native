// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetIPAMResourceDiscovery
    {
        /// <summary>
        /// Resource Schema of AWS::EC2::IPAMResourceDiscovery Type
        /// </summary>
        public static Task<GetIPAMResourceDiscoveryResult> InvokeAsync(GetIPAMResourceDiscoveryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIPAMResourceDiscoveryResult>("aws-native:ec2:getIPAMResourceDiscovery", args ?? new GetIPAMResourceDiscoveryArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Schema of AWS::EC2::IPAMResourceDiscovery Type
        /// </summary>
        public static Output<GetIPAMResourceDiscoveryResult> Invoke(GetIPAMResourceDiscoveryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIPAMResourceDiscoveryResult>("aws-native:ec2:getIPAMResourceDiscovery", args ?? new GetIPAMResourceDiscoveryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIPAMResourceDiscoveryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the IPAM Pool.
        /// </summary>
        [Input("ipamResourceDiscoveryId", required: true)]
        public string IpamResourceDiscoveryId { get; set; } = null!;

        public GetIPAMResourceDiscoveryArgs()
        {
        }
        public static new GetIPAMResourceDiscoveryArgs Empty => new GetIPAMResourceDiscoveryArgs();
    }

    public sealed class GetIPAMResourceDiscoveryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the IPAM Pool.
        /// </summary>
        [Input("ipamResourceDiscoveryId", required: true)]
        public Input<string> IpamResourceDiscoveryId { get; set; } = null!;

        public GetIPAMResourceDiscoveryInvokeArgs()
        {
        }
        public static new GetIPAMResourceDiscoveryInvokeArgs Empty => new GetIPAMResourceDiscoveryInvokeArgs();
    }


    [OutputType]
    public sealed class GetIPAMResourceDiscoveryResult
    {
        public readonly string? Description;
        /// <summary>
        /// Amazon Resource Name (Arn) for the Resource Discovery.
        /// </summary>
        public readonly string? IpamResourceDiscoveryArn;
        /// <summary>
        /// Id of the IPAM Pool.
        /// </summary>
        public readonly string? IpamResourceDiscoveryId;
        /// <summary>
        /// The region the resource discovery is setup in. 
        /// </summary>
        public readonly string? IpamResourceDiscoveryRegion;
        /// <summary>
        /// Determines whether or not address space from this pool is publicly advertised. Must be set if and only if the pool is IPv6.
        /// </summary>
        public readonly bool? IsDefault;
        /// <summary>
        /// The regions Resource Discovery is enabled for. Allows resource discoveries to be created in these regions, as well as enabling monitoring
        /// </summary>
        public readonly ImmutableArray<Outputs.IPAMResourceDiscoveryIpamOperatingRegion> OperatingRegions;
        /// <summary>
        /// Owner Account ID of the Resource Discovery
        /// </summary>
        public readonly string? OwnerId;
        /// <summary>
        /// The state of this Resource Discovery.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.IPAMResourceDiscoveryTag> Tags;

        [OutputConstructor]
        private GetIPAMResourceDiscoveryResult(
            string? description,

            string? ipamResourceDiscoveryArn,

            string? ipamResourceDiscoveryId,

            string? ipamResourceDiscoveryRegion,

            bool? isDefault,

            ImmutableArray<Outputs.IPAMResourceDiscoveryIpamOperatingRegion> operatingRegions,

            string? ownerId,

            string? state,

            ImmutableArray<Outputs.IPAMResourceDiscoveryTag> tags)
        {
            Description = description;
            IpamResourceDiscoveryArn = ipamResourceDiscoveryArn;
            IpamResourceDiscoveryId = ipamResourceDiscoveryId;
            IpamResourceDiscoveryRegion = ipamResourceDiscoveryRegion;
            IsDefault = isDefault;
            OperatingRegions = operatingRegions;
            OwnerId = ownerId;
            State = state;
            Tags = tags;
        }
    }
}
