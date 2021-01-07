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
    public sealed class FunctionConfigurationSyncConfig
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-syncconfig.html#cfn-appsync-functionconfiguration-syncconfig-conflictdetection
        /// </summary>
        public readonly string ConflictDetection;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-syncconfig.html#cfn-appsync-functionconfiguration-syncconfig-conflicthandler
        /// </summary>
        public readonly string? ConflictHandler;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-functionconfiguration-syncconfig.html#cfn-appsync-functionconfiguration-syncconfig-lambdaconflicthandlerconfig
        /// </summary>
        public readonly Outputs.FunctionConfigurationLambdaConflictHandlerConfig? LambdaConflictHandlerConfig;

        [OutputConstructor]
        private FunctionConfigurationSyncConfig(
            string ConflictDetection,

            string? ConflictHandler,

            Outputs.FunctionConfigurationLambdaConflictHandlerConfig? LambdaConflictHandlerConfig)
        {
            this.ConflictDetection = ConflictDetection;
            this.ConflictHandler = ConflictHandler;
            this.LambdaConflictHandlerConfig = LambdaConflictHandlerConfig;
        }
    }
}