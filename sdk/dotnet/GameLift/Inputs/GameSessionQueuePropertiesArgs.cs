// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.GameLift.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html
    /// </summary>
    public sealed class GameSessionQueuePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("Destinations")]
        private InputList<Inputs.GameSessionQueueDestinationArgs>? _Destinations;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-destinations
        /// </summary>
        public InputList<Inputs.GameSessionQueueDestinationArgs> Destinations
        {
            get => _Destinations ?? (_Destinations = new InputList<Inputs.GameSessionQueueDestinationArgs>());
            set => _Destinations = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("PlayerLatencyPolicies")]
        private InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs>? _PlayerLatencyPolicies;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-playerlatencypolicies
        /// </summary>
        public InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs> PlayerLatencyPolicies
        {
            get => _PlayerLatencyPolicies ?? (_PlayerLatencyPolicies = new InputList<Inputs.GameSessionQueuePlayerLatencyPolicyArgs>());
            set => _PlayerLatencyPolicies = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-gamesessionqueue.html#cfn-gamelift-gamesessionqueue-timeoutinseconds
        /// </summary>
        [Input("TimeoutInSeconds")]
        public Input<int>? TimeoutInSeconds { get; set; }

        public GameSessionQueuePropertiesArgs()
        {
        }
    }
}