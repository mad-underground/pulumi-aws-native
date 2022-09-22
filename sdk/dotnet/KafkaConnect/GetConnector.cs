// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect
{
    public static class GetConnector
    {
        /// <summary>
        /// Resource Type definition for AWS::KafkaConnect::Connector
        /// </summary>
        public static Task<GetConnectorResult> InvokeAsync(GetConnectorArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetConnectorResult>("aws-native:kafkaconnect:getConnector", args ?? new GetConnectorArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::KafkaConnect::Connector
        /// </summary>
        public static Output<GetConnectorResult> Invoke(GetConnectorInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetConnectorResult>("aws-native:kafkaconnect:getConnector", args ?? new GetConnectorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConnectorArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name for the created Connector.
        /// </summary>
        [Input("connectorArn", required: true)]
        public string ConnectorArn { get; set; } = null!;

        public GetConnectorArgs()
        {
        }
        public static new GetConnectorArgs Empty => new GetConnectorArgs();
    }

    public sealed class GetConnectorInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Amazon Resource Name for the created Connector.
        /// </summary>
        [Input("connectorArn", required: true)]
        public Input<string> ConnectorArn { get; set; } = null!;

        public GetConnectorInvokeArgs()
        {
        }
        public static new GetConnectorInvokeArgs Empty => new GetConnectorInvokeArgs();
    }


    [OutputType]
    public sealed class GetConnectorResult
    {
        public readonly Outputs.ConnectorCapacity? Capacity;
        /// <summary>
        /// Amazon Resource Name for the created Connector.
        /// </summary>
        public readonly string? ConnectorArn;

        [OutputConstructor]
        private GetConnectorResult(
            Outputs.ConnectorCapacity? capacity,

            string? connectorArn)
        {
            Capacity = capacity;
            ConnectorArn = connectorArn;
        }
    }
}
