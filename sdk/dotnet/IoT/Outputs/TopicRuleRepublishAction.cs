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
    public sealed class TopicRuleRepublishAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-qos
        /// </summary>
        public readonly int? Qos;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-rolearn
        /// </summary>
        public readonly string RoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-topic
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private TopicRuleRepublishAction(
            int? Qos,

            string RoleArn,

            string Topic)
        {
            this.Qos = Qos;
            this.RoleArn = RoleArn;
            this.Topic = Topic;
        }
    }
}