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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html
    /// </summary>
    public sealed class MatchmakingConfigurationPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-acceptancerequired
        /// </summary>
        [Input("AcceptanceRequired", required: true)]
        public Input<bool> AcceptanceRequired { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-acceptancetimeoutseconds
        /// </summary>
        [Input("AcceptanceTimeoutSeconds")]
        public Input<int>? AcceptanceTimeoutSeconds { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-additionalplayercount
        /// </summary>
        [Input("AdditionalPlayerCount")]
        public Input<int>? AdditionalPlayerCount { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-backfillmode
        /// </summary>
        [Input("BackfillMode")]
        public Input<string>? BackfillMode { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-customeventdata
        /// </summary>
        [Input("CustomEventData")]
        public Input<string>? CustomEventData { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-description
        /// </summary>
        [Input("Description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-flexmatchmode
        /// </summary>
        [Input("FlexMatchMode")]
        public Input<string>? FlexMatchMode { get; set; }

        [Input("GameProperties")]
        private InputList<Inputs.MatchmakingConfigurationGamePropertyArgs>? _GameProperties;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-gameproperties
        /// </summary>
        public InputList<Inputs.MatchmakingConfigurationGamePropertyArgs> GameProperties
        {
            get => _GameProperties ?? (_GameProperties = new InputList<Inputs.MatchmakingConfigurationGamePropertyArgs>());
            set => _GameProperties = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-gamesessiondata
        /// </summary>
        [Input("GameSessionData")]
        public Input<string>? GameSessionData { get; set; }

        [Input("GameSessionQueueArns")]
        private InputList<string>? _GameSessionQueueArns;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-gamesessionqueuearns
        /// </summary>
        public InputList<string> GameSessionQueueArns
        {
            get => _GameSessionQueueArns ?? (_GameSessionQueueArns = new InputList<string>());
            set => _GameSessionQueueArns = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-name
        /// </summary>
        [Input("Name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-notificationtarget
        /// </summary>
        [Input("NotificationTarget")]
        public Input<string>? NotificationTarget { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-requesttimeoutseconds
        /// </summary>
        [Input("RequestTimeoutSeconds", required: true)]
        public Input<int> RequestTimeoutSeconds { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-matchmakingconfiguration.html#cfn-gamelift-matchmakingconfiguration-rulesetname
        /// </summary>
        [Input("RuleSetName", required: true)]
        public Input<string> RuleSetName { get; set; } = null!;

        public MatchmakingConfigurationPropertiesArgs()
        {
        }
    }
}
