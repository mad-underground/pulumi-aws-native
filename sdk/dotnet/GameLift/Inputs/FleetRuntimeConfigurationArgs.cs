// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html
    /// </summary>
    public sealed class FleetRuntimeConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-gamesessionactivationtimeoutseconds
        /// </summary>
        [Input("gameSessionActivationTimeoutSeconds")]
        public Input<int>? GameSessionActivationTimeoutSeconds { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-maxconcurrentgamesessionactivations
        /// </summary>
        [Input("maxConcurrentGameSessionActivations")]
        public Input<int>? MaxConcurrentGameSessionActivations { get; set; }

        [Input("serverProcesses")]
        private InputList<Inputs.FleetServerProcessArgs>? _serverProcesses;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-fleet-runtimeconfiguration.html#cfn-gamelift-fleet-runtimeconfiguration-serverprocesses
        /// </summary>
        public InputList<Inputs.FleetServerProcessArgs> ServerProcesses
        {
            get => _serverProcesses ?? (_serverProcesses = new InputList<Inputs.FleetServerProcessArgs>());
            set => _serverProcesses = value;
        }

        public FleetRuntimeConfigurationArgs()
        {
        }
    }
}