// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoT.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html
    /// </summary>
    public sealed class TopicRuleTopicRulePayloadArgs : Pulumi.ResourceArgs
    {
        [Input("Actions", required: true)]
        private InputList<Inputs.TopicRuleActionArgs>? _Actions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-actions
        /// </summary>
        public InputList<Inputs.TopicRuleActionArgs> Actions
        {
            get => _Actions ?? (_Actions = new InputList<Inputs.TopicRuleActionArgs>());
            set => _Actions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-awsiotsqlversion
        /// </summary>
        [Input("AwsIotSqlVersion")]
        public Input<string>? AwsIotSqlVersion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-description
        /// </summary>
        [Input("Description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-erroraction
        /// </summary>
        [Input("ErrorAction")]
        public Input<Inputs.TopicRuleActionArgs>? ErrorAction { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-ruledisabled
        /// </summary>
        [Input("RuleDisabled", required: true)]
        public Input<bool> RuleDisabled { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-topicrulepayload.html#cfn-iot-topicrule-topicrulepayload-sql
        /// </summary>
        [Input("Sql", required: true)]
        public Input<string> Sql { get; set; } = null!;

        public TopicRuleTopicRulePayloadArgs()
        {
        }
    }
}