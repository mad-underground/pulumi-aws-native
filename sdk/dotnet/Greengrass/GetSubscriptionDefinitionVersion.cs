// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetSubscriptionDefinitionVersion
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::SubscriptionDefinitionVersion
        /// </summary>
        public static Task<GetSubscriptionDefinitionVersionResult> InvokeAsync(GetSubscriptionDefinitionVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSubscriptionDefinitionVersionResult>("aws-native:greengrass:getSubscriptionDefinitionVersion", args ?? new GetSubscriptionDefinitionVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::SubscriptionDefinitionVersion
        /// </summary>
        public static Output<GetSubscriptionDefinitionVersionResult> Invoke(GetSubscriptionDefinitionVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetSubscriptionDefinitionVersionResult>("aws-native:greengrass:getSubscriptionDefinitionVersion", args ?? new GetSubscriptionDefinitionVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSubscriptionDefinitionVersionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetSubscriptionDefinitionVersionArgs()
        {
        }
    }

    public sealed class GetSubscriptionDefinitionVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetSubscriptionDefinitionVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSubscriptionDefinitionVersionResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetSubscriptionDefinitionVersionResult(string? id)
        {
            Id = id;
        }
    }
}