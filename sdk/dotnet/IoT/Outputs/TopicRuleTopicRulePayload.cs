// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoT.Outputs
{

    [OutputType]
    public sealed class TopicRuleTopicRulePayload
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-actions
        /// </summary>
        public readonly ImmutableArray<Outputs.TopicRuleAction> Actions;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-awsiotsqlversion
        /// </summary>
        public readonly string? AwsIotSqlVersion;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-erroraction
        /// </summary>
        public readonly Outputs.TopicRuleAction? ErrorAction;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-ruledisabled
        /// </summary>
        public readonly bool RuleDisabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-sql
        /// </summary>
        public readonly string Sql;

        [OutputConstructor]
        private TopicRuleTopicRulePayload(
            ImmutableArray<Outputs.TopicRuleAction> Actions,

            string? AwsIotSqlVersion,

            string? Description,

            Outputs.TopicRuleAction? ErrorAction,

            bool RuleDisabled,

            string Sql)
        {
            this.Actions = Actions;
            this.AwsIotSqlVersion = AwsIotSqlVersion;
            this.Description = Description;
            this.ErrorAction = ErrorAction;
            this.RuleDisabled = RuleDisabled;
            this.Sql = Sql;
        }
    }
}