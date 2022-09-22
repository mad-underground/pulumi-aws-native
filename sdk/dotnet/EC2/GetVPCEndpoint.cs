// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetVPCEndpoint
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::VPCEndpoint
        /// </summary>
        public static Task<GetVPCEndpointResult> InvokeAsync(GetVPCEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVPCEndpointResult>("aws-native:ec2:getVPCEndpoint", args ?? new GetVPCEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::VPCEndpoint
        /// </summary>
        public static Output<GetVPCEndpointResult> Invoke(GetVPCEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVPCEndpointResult>("aws-native:ec2:getVPCEndpoint", args ?? new GetVPCEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVPCEndpointArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVPCEndpointArgs()
        {
        }
        public static new GetVPCEndpointArgs Empty => new GetVPCEndpointArgs();
    }

    public sealed class GetVPCEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVPCEndpointInvokeArgs()
        {
        }
        public static new GetVPCEndpointInvokeArgs Empty => new GetVPCEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetVPCEndpointResult
    {
        public readonly string? CreationTimestamp;
        public readonly ImmutableArray<string> DnsEntries;
        public readonly string? Id;
        public readonly ImmutableArray<string> NetworkInterfaceIds;
        public readonly object? PolicyDocument;
        public readonly bool? PrivateDnsEnabled;
        public readonly ImmutableArray<string> RouteTableIds;
        public readonly ImmutableArray<string> SecurityGroupIds;
        public readonly ImmutableArray<string> SubnetIds;

        [OutputConstructor]
        private GetVPCEndpointResult(
            string? creationTimestamp,

            ImmutableArray<string> dnsEntries,

            string? id,

            ImmutableArray<string> networkInterfaceIds,

            object? policyDocument,

            bool? privateDnsEnabled,

            ImmutableArray<string> routeTableIds,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> subnetIds)
        {
            CreationTimestamp = creationTimestamp;
            DnsEntries = dnsEntries;
            Id = id;
            NetworkInterfaceIds = networkInterfaceIds;
            PolicyDocument = policyDocument;
            PrivateDnsEnabled = privateDnsEnabled;
            RouteTableIds = routeTableIds;
            SecurityGroupIds = securityGroupIds;
            SubnetIds = subnetIds;
        }
    }
}
