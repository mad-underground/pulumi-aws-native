// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    /// <summary>
    /// The EC2 launch template that contains configuration settings and game server code to be deployed to all instances in the game server group.
    /// </summary>
    public sealed class GameServerGroupLaunchTemplateArgs : global::Pulumi.ResourceArgs
    {
        [Input("launchTemplateId")]
        public Input<string>? LaunchTemplateId { get; set; }

        [Input("launchTemplateName")]
        public Input<string>? LaunchTemplateName { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public GameServerGroupLaunchTemplateArgs()
        {
        }
        public static new GameServerGroupLaunchTemplateArgs Empty => new GameServerGroupLaunchTemplateArgs();
    }
}
