// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Outputs
{

    [OutputType]
    public sealed class TopicRuleAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-cloudwatchalarm
        /// </summary>
        public readonly Outputs.TopicRuleCloudwatchAlarmAction? CloudwatchAlarm;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-cloudwatchmetric
        /// </summary>
        public readonly Outputs.TopicRuleCloudwatchMetricAction? CloudwatchMetric;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-dynamodb
        /// </summary>
        public readonly Outputs.TopicRuleDynamoDBAction? DynamoDB;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-dynamodbv2
        /// </summary>
        public readonly Outputs.TopicRuleDynamoDBv2Action? DynamoDBv2;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-elasticsearch
        /// </summary>
        public readonly Outputs.TopicRuleElasticsearchAction? Elasticsearch;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-firehose
        /// </summary>
        public readonly Outputs.TopicRuleFirehoseAction? Firehose;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-http
        /// </summary>
        public readonly Outputs.TopicRuleHttpAction? Http;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-iotanalytics
        /// </summary>
        public readonly Outputs.TopicRuleIotAnalyticsAction? IotAnalytics;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-iotevents
        /// </summary>
        public readonly Outputs.TopicRuleIotEventsAction? IotEvents;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-iotsitewise
        /// </summary>
        public readonly Outputs.TopicRuleIotSiteWiseAction? IotSiteWise;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-kinesis
        /// </summary>
        public readonly Outputs.TopicRuleKinesisAction? Kinesis;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-lambda
        /// </summary>
        public readonly Outputs.TopicRuleLambdaAction? Lambda;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-republish
        /// </summary>
        public readonly Outputs.TopicRuleRepublishAction? Republish;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-s3
        /// </summary>
        public readonly Outputs.TopicRuleS3Action? S3;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-sns
        /// </summary>
        public readonly Outputs.TopicRuleSnsAction? Sns;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-sqs
        /// </summary>
        public readonly Outputs.TopicRuleSqsAction? Sqs;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-action.html#cfn-iot-topicrule-action-stepfunctions
        /// </summary>
        public readonly Outputs.TopicRuleStepFunctionsAction? StepFunctions;

        [OutputConstructor]
        private TopicRuleAction(
            Outputs.TopicRuleCloudwatchAlarmAction? cloudwatchAlarm,

            Outputs.TopicRuleCloudwatchMetricAction? cloudwatchMetric,

            Outputs.TopicRuleDynamoDBAction? dynamoDB,

            Outputs.TopicRuleDynamoDBv2Action? dynamoDBv2,

            Outputs.TopicRuleElasticsearchAction? elasticsearch,

            Outputs.TopicRuleFirehoseAction? firehose,

            Outputs.TopicRuleHttpAction? http,

            Outputs.TopicRuleIotAnalyticsAction? iotAnalytics,

            Outputs.TopicRuleIotEventsAction? iotEvents,

            Outputs.TopicRuleIotSiteWiseAction? iotSiteWise,

            Outputs.TopicRuleKinesisAction? kinesis,

            Outputs.TopicRuleLambdaAction? lambda,

            Outputs.TopicRuleRepublishAction? republish,

            Outputs.TopicRuleS3Action? s3,

            Outputs.TopicRuleSnsAction? sns,

            Outputs.TopicRuleSqsAction? sqs,

            Outputs.TopicRuleStepFunctionsAction? stepFunctions)
        {
            CloudwatchAlarm = cloudwatchAlarm;
            CloudwatchMetric = cloudwatchMetric;
            DynamoDB = dynamoDB;
            DynamoDBv2 = dynamoDBv2;
            Elasticsearch = elasticsearch;
            Firehose = firehose;
            Http = http;
            IotAnalytics = iotAnalytics;
            IotEvents = iotEvents;
            IotSiteWise = iotSiteWise;
            Kinesis = kinesis;
            Lambda = lambda;
            Republish = republish;
            S3 = s3;
            Sns = sns;
            Sqs = sqs;
            StepFunctions = stepFunctions;
        }
    }
}
