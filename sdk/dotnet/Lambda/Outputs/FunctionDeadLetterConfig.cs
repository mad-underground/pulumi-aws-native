// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Outputs
{

    /// <summary>
    /// The [dead-letter queue](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq) for failed asynchronous invocations.
    /// </summary>
    [OutputType]
    public sealed class FunctionDeadLetterConfig
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
        /// </summary>
        public readonly string? TargetArn;

        [OutputConstructor]
        private FunctionDeadLetterConfig(string? targetArn)
        {
            TargetArn = targetArn;
        }
    }
}
