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
    /// Specifies configuration for replication between a source and target Kafka cluster.
    /// </summary>
    public sealed class ReplicatorReplicationInfoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration relating to consumer group replication.
        /// </summary>
        [Input("consumerGroupReplication", required: true)]
        public Input<Inputs.ReplicatorConsumerGroupReplicationArgs> ConsumerGroupReplication { get; set; } = null!;

        /// <summary>
        /// Amazon Resource Name of the source Kafka cluster.
        /// </summary>
        [Input("sourceKafkaClusterArn", required: true)]
        public Input<string> SourceKafkaClusterArn { get; set; } = null!;

        /// <summary>
        /// The type of compression to use writing records to target Kafka cluster.
        /// </summary>
        [Input("targetCompressionType", required: true)]
        public Input<Pulumi.AwsNative.Msk.ReplicatorReplicationInfoTargetCompressionType> TargetCompressionType { get; set; } = null!;

        /// <summary>
        /// Amazon Resource Name of the target Kafka cluster.
        /// </summary>
        [Input("targetKafkaClusterArn", required: true)]
        public Input<string> TargetKafkaClusterArn { get; set; } = null!;

        /// <summary>
        /// Configuration relating to topic replication.
        /// </summary>
        [Input("topicReplication", required: true)]
        public Input<Inputs.ReplicatorTopicReplicationArgs> TopicReplication { get; set; } = null!;

        public ReplicatorReplicationInfoArgs()
        {
        }
        public static new ReplicatorReplicationInfoArgs Empty => new ReplicatorReplicationInfoArgs();
    }
}
