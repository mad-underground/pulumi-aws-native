// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass
{
    public static class GetConnectorDefinitionVersion
    {
        /// <summary>
        /// Resource Type definition for AWS::Greengrass::ConnectorDefinitionVersion
        /// </summary>
        public static Task<GetConnectorDefinitionVersionResult> InvokeAsync(GetConnectorDefinitionVersionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectorDefinitionVersionResult>("aws-native:greengrass:getConnectorDefinitionVersion", args ?? new GetConnectorDefinitionVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Greengrass::ConnectorDefinitionVersion
        /// </summary>
        public static Output<GetConnectorDefinitionVersionResult> Invoke(GetConnectorDefinitionVersionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConnectorDefinitionVersionResult>("aws-native:greengrass:getConnectorDefinitionVersion", args ?? new GetConnectorDefinitionVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectorDefinitionVersionArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetConnectorDefinitionVersionArgs()
        {
        }
    }

    public sealed class GetConnectorDefinitionVersionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetConnectorDefinitionVersionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectorDefinitionVersionResult
    {
        public readonly string? Id;

        [OutputConstructor]
        private GetConnectorDefinitionVersionResult(string? id)
        {
            Id = id;
        }
    }
}