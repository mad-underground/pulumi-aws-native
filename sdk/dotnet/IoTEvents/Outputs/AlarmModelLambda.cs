// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    [OutputType]
    public sealed class AlarmModelLambda
    {
        /// <summary>
        /// The ARN of the Lambda function that is executed.
        /// </summary>
        public readonly string FunctionArn;
        public readonly Outputs.AlarmModelPayload? Payload;

        [OutputConstructor]
        private AlarmModelLambda(
            string functionArn,

            Outputs.AlarmModelPayload? payload)
        {
            FunctionArn = functionArn;
            Payload = payload;
        }
    }
}
