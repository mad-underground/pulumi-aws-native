// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    /// <summary>
    /// Details of a Kafka cluster for replication.
    /// </summary>
    public sealed class ReplicatorKafkaClusterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of an Amazon MSK cluster. Exactly one of AmazonMskCluster is required.
        /// </summary>
        [Input("amazonMskCluster", required: true)]
        public Input<Inputs.ReplicatorAmazonMskClusterArgs> AmazonMskCluster { get; set; } = null!;

        /// <summary>
        /// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
        /// </summary>
        [Input("vpcConfig", required: true)]
        public Input<Inputs.ReplicatorKafkaClusterClientVpcConfigArgs> VpcConfig { get; set; } = null!;

        public ReplicatorKafkaClusterArgs()
        {
        }
        public static new ReplicatorKafkaClusterArgs Empty => new ReplicatorKafkaClusterArgs();
    }
}
