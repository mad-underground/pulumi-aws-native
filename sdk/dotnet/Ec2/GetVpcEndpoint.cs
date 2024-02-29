// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetVpcEndpoint
    {
        /// <summary>
        /// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the [User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/).
        ///  An endpoint of type ``Interface`` establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
        ///  An endpoint of type ``gateway`` serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see [W
        /// </summary>
        public static Task<GetVpcEndpointResult> InvokeAsync(GetVpcEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpcEndpointResult>("aws-native:ec2:getVpcEndpoint", args ?? new GetVpcEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Specifies a VPC endpoint. A VPC endpoint provides a private connection between your VPC and an endpoint service. You can use an endpoint service provided by AWS, an MKT Partner, or another AWS accounts in your organization. For more information, see the [User Guide](https://docs.aws.amazon.com/vpc/latest/privatelink/).
        ///  An endpoint of type ``Interface`` establishes connections between the subnets in your VPC and an AWS-service, your own service, or a service hosted by another AWS-account. With an interface VPC endpoint, you specify the subnets in which to create the endpoint and the security groups to associate with the endpoint network interfaces.
        ///  An endpoint of type ``gateway`` serves as a target for a route in your route table for traffic destined for S3 or DDB. You can specify an endpoint policy for the endpoint, which controls access to the service from your VPC. You can also specify the VPC route tables that use the endpoint. For more information about connectivity to S3, see [W
        /// </summary>
        public static Output<GetVpcEndpointResult> Invoke(GetVpcEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpcEndpointResult>("aws-native:ec2:getVpcEndpoint", args ?? new GetVpcEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpcEndpointArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVpcEndpointArgs()
        {
        }
        public static new GetVpcEndpointArgs Empty => new GetVpcEndpointArgs();
    }

    public sealed class GetVpcEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVpcEndpointInvokeArgs()
        {
        }
        public static new GetVpcEndpointInvokeArgs Empty => new GetVpcEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpcEndpointResult
    {
        public readonly string? CreationTimestamp;
        public readonly ImmutableArray<string> DnsEntries;
        public readonly string? Id;
        public readonly ImmutableArray<string> NetworkInterfaceIds;
        /// <summary>
        /// An endpoint policy, which controls access to the service from the VPC. The default endpoint policy allows full access to the service. Endpoint policies are supported only for gateway and interface endpoints.
        ///  For CloudFormation templates in YAML, you can provide the policy in JSON or YAML format. CFNlong converts YAML policies to JSON format before calling the API to create or modify the VPC endpoint.
        /// </summary>
        public readonly string? PolicyDocument;
        /// <summary>
        /// Indicate whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, ``kinesis.us-east-1.amazonaws.com``), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.
        ///  To use a private hosted zone, you must set the following VPC attributes to ``true``: ``enableDnsHostnames`` and ``enableDnsSupport``.
        ///  This property is supported only for interface endpoints.
        ///  Default: ``false``
        /// </summary>
        public readonly bool? PrivateDnsEnabled;
        /// <summary>
        /// The IDs of the route tables. Routing is supported only for gateway endpoints.
        /// </summary>
        public readonly ImmutableArray<string> RouteTableIds;
        /// <summary>
        /// The IDs of the security groups to associate with the endpoint network interfaces. If this parameter is not specified, we use the default security group for the VPC. Security groups are supported only for interface endpoints.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The IDs of the subnets in which to create endpoint network interfaces. You must specify this property for an interface endpoint or a Gateway Load Balancer endpoint. You can't specify this property for a gateway endpoint. For a Gateway Load Balancer endpoint, you can specify only one subnet.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;

        [OutputConstructor]
        private GetVpcEndpointResult(
            string? creationTimestamp,

            ImmutableArray<string> dnsEntries,

            string? id,

            ImmutableArray<string> networkInterfaceIds,

            string? policyDocument,

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
