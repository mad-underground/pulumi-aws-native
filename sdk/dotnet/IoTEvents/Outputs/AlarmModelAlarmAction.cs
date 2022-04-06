// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// The actions to be performed.
    /// </summary>
    [OutputType]
    public sealed class AlarmModelAlarmAction
    {
        public readonly Outputs.AlarmModelDynamoDB? DynamoDB;
        public readonly Outputs.AlarmModelDynamoDBv2? DynamoDBv2;
        public readonly Outputs.AlarmModelFirehose? Firehose;
        public readonly Outputs.AlarmModelIotEvents? IotEvents;
        public readonly Outputs.AlarmModelIotSiteWise? IotSiteWise;
        public readonly Outputs.AlarmModelIotTopicPublish? IotTopicPublish;
        public readonly Outputs.AlarmModelLambda? Lambda;
        public readonly Outputs.AlarmModelSns? Sns;
        public readonly Outputs.AlarmModelSqs? Sqs;

        [OutputConstructor]
        private AlarmModelAlarmAction(
            Outputs.AlarmModelDynamoDB? dynamoDB,

            Outputs.AlarmModelDynamoDBv2? dynamoDBv2,

            Outputs.AlarmModelFirehose? firehose,

            Outputs.AlarmModelIotEvents? iotEvents,

            Outputs.AlarmModelIotSiteWise? iotSiteWise,

            Outputs.AlarmModelIotTopicPublish? iotTopicPublish,

            Outputs.AlarmModelLambda? lambda,

            Outputs.AlarmModelSns? sns,

            Outputs.AlarmModelSqs? sqs)
        {
            DynamoDB = dynamoDB;
            DynamoDBv2 = dynamoDBv2;
            Firehose = firehose;
            IotEvents = iotEvents;
            IotSiteWise = iotSiteWise;
            IotTopicPublish = iotTopicPublish;
            Lambda = lambda;
            Sns = sns;
            Sqs = sqs;
        }
    }
}