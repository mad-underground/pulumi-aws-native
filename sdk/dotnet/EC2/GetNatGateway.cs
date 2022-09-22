// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetNatGateway
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::NatGateway
        /// </summary>
        public static Task<GetNatGatewayResult> InvokeAsync(GetNatGatewayArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNatGatewayResult>("aws-native:ec2:getNatGateway", args ?? new GetNatGatewayArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::NatGateway
        /// </summary>
        public static Output<GetNatGatewayResult> Invoke(GetNatGatewayInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNatGatewayResult>("aws-native:ec2:getNatGateway", args ?? new GetNatGatewayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNatGatewayArgs : global::Pulumi.InvokeArgs
    {
        [Input("natGatewayId", required: true)]
        public string NatGatewayId { get; set; } = null!;

        public GetNatGatewayArgs()
        {
        }
        public static new GetNatGatewayArgs Empty => new GetNatGatewayArgs();
    }

    public sealed class GetNatGatewayInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("natGatewayId", required: true)]
        public Input<string> NatGatewayId { get; set; } = null!;

        public GetNatGatewayInvokeArgs()
        {
        }
        public static new GetNatGatewayInvokeArgs Empty => new GetNatGatewayInvokeArgs();
    }


    [OutputType]
    public sealed class GetNatGatewayResult
    {
        public readonly string? NatGatewayId;
        public readonly ImmutableArray<Outputs.NatGatewayTag> Tags;

        [OutputConstructor]
        private GetNatGatewayResult(
            string? natGatewayId,

            ImmutableArray<Outputs.NatGatewayTag> tags)
        {
            NatGatewayId = natGatewayId;
            Tags = tags;
        }
    }
}
