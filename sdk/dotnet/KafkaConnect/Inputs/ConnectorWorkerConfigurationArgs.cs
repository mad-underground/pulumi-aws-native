// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KafkaConnect.Inputs
{

    /// <summary>
    /// Specifies the worker configuration to use with the connector.
    /// </summary>
    public sealed class ConnectorWorkerConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The revision of the worker configuration to use.
        /// </summary>
        [Input("revision", required: true)]
        public Input<int> Revision { get; set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the worker configuration to use.
        /// </summary>
        [Input("workerConfigurationArn", required: true)]
        public Input<string> WorkerConfigurationArn { get; set; } = null!;

        public ConnectorWorkerConfigurationArgs()
        {
        }
    }
}