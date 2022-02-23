// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFormation
{
    public static class GetCustomResource
    {
        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::CustomResource
        /// </summary>
        public static Task<GetCustomResourceResult> InvokeAsync(GetCustomResourceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCustomResourceResult>("aws-native:cloudformation:getCustomResource", args ?? new GetCustomResourceArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CloudFormation::CustomResource
        /// </summary>
        public static Output<GetCustomResourceResult> Invoke(GetCustomResourceInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCustomResourceResult>("aws-native:cloudformation:getCustomResource", args ?? new GetCustomResourceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCustomResourceArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetCustomResourceArgs()
        {
        }
    }

    public sealed class GetCustomResourceInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetCustomResourceInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCustomResourceResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetCustomResourceResult(string? id)
        {
            Id = id;
        }
    }
}