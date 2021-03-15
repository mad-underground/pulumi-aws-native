// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html
    /// </summary>
    public sealed class TopicRuleStepFunctionsActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html#cfn-iot-topicrule-stepfunctionsaction-executionnameprefix
        /// </summary>
        [Input("executionNamePrefix")]
        public Input<string>? ExecutionNamePrefix { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html#cfn-iot-topicrule-stepfunctionsaction-rolearn
        /// </summary>
        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-stepfunctionsaction.html#cfn-iot-topicrule-stepfunctionsaction-statemachinename
        /// </summary>
        [Input("stateMachineName", required: true)]
        public Input<string> StateMachineName { get; set; } = null!;

        public TopicRuleStepFunctionsActionArgs()
        {
        }
    }
}
