// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2
{
    public static class GetVPNConnection
    {
        /// <summary>
        /// Resource Type definition for AWS::EC2::VPNConnection
        /// </summary>
        public static Task<GetVPNConnectionResult> InvokeAsync(GetVPNConnectionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVPNConnectionResult>("aws-native:ec2:getVPNConnection", args ?? new GetVPNConnectionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::EC2::VPNConnection
        /// </summary>
        public static Output<GetVPNConnectionResult> Invoke(GetVPNConnectionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVPNConnectionResult>("aws-native:ec2:getVPNConnection", args ?? new GetVPNConnectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVPNConnectionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetVPNConnectionArgs()
        {
        }
    }

    public sealed class GetVPNConnectionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetVPNConnectionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVPNConnectionResult
    {
        public readonly string? Id;
        public readonly ImmutableArray<Outputs.VPNConnectionTag> Tags;

        [OutputConstructor]
        private GetVPNConnectionResult(
            string? id,

            ImmutableArray<Outputs.VPNConnectionTag> tags)
        {
            Id = id;
            Tags = tags;
        }
    }
}