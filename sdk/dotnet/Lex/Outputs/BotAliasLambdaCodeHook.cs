// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// Contains information about code hooks that Amazon Lex calls during a conversation.
    /// </summary>
    [OutputType]
    public sealed class BotAliasLambdaCodeHook
    {
        /// <summary>
        /// The version of the request-response that you want Amazon Lex to use to invoke your Lambda function.
        /// </summary>
        public readonly string CodeHookInterfaceVersion;
        /// <summary>
        /// The Amazon Resource Name (ARN) of the Lambda function.
        /// </summary>
        public readonly string LambdaArn;

        [OutputConstructor]
        private BotAliasLambdaCodeHook(
            string codeHookInterfaceVersion,

            string lambdaArn)
        {
            CodeHookInterfaceVersion = codeHookInterfaceVersion;
            LambdaArn = lambdaArn;
        }
    }
}