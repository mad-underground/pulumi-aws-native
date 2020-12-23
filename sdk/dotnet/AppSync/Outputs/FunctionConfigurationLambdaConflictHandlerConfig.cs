// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppSync.Outputs
{

    [OutputType]
    public sealed class FunctionConfigurationLambdaConflictHandlerConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-lambdaconflicthandlerconfig.html#cfn-appsync-functionconfiguration-lambdaconflicthandlerconfig-lambdaconflicthandlerarn
        /// </summary>
        public readonly string? LambdaConflictHandlerArn;

        [OutputConstructor]
        private FunctionConfigurationLambdaConflictHandlerConfig(string? LambdaConflictHandlerArn)
        {
            this.LambdaConflictHandlerArn = LambdaConflictHandlerArn;
        }
    }
}
