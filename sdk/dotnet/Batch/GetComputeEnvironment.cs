// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch
{
    public static class GetComputeEnvironment
    {
        /// <summary>
        /// Resource Type definition for AWS::Batch::ComputeEnvironment
        /// </summary>
        public static Task<GetComputeEnvironmentResult> InvokeAsync(GetComputeEnvironmentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetComputeEnvironmentResult>("aws-native:batch:getComputeEnvironment", args ?? new GetComputeEnvironmentArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Batch::ComputeEnvironment
        /// </summary>
        public static Output<GetComputeEnvironmentResult> Invoke(GetComputeEnvironmentInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetComputeEnvironmentResult>("aws-native:batch:getComputeEnvironment", args ?? new GetComputeEnvironmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetComputeEnvironmentArgs : Pulumi.InvokeArgs
    {
        [Input("computeEnvironmentArn", required: true)]
        public string ComputeEnvironmentArn { get; set; } = null!;

        public GetComputeEnvironmentArgs()
        {
        }
    }

    public sealed class GetComputeEnvironmentInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("computeEnvironmentArn", required: true)]
        public Input<string> ComputeEnvironmentArn { get; set; } = null!;

        public GetComputeEnvironmentInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetComputeEnvironmentResult
    {
        public readonly string? ComputeEnvironmentArn;
        public readonly Outputs.ComputeEnvironmentComputeResources? ComputeResources;
        public readonly string? ServiceRole;
        public readonly string? State;
        public readonly int? UnmanagedvCpus;

        [OutputConstructor]
        private GetComputeEnvironmentResult(
            string? computeEnvironmentArn,

            Outputs.ComputeEnvironmentComputeResources? computeResources,

            string? serviceRole,

            string? state,

            int? unmanagedvCpus)
        {
            ComputeEnvironmentArn = computeEnvironmentArn;
            ComputeResources = computeResources;
            ServiceRole = serviceRole;
            State = state;
            UnmanagedvCpus = unmanagedvCpus;
        }
    }
}