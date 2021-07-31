// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-vpcsubnets.html
    /// </summary>
    [OutputType]
    public sealed class GameServerGroupVpcSubnets
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-gameservergroup-vpcsubnets.html#cfn-gamelift-gameservergroup-vpcsubnets-vpcsubnets
        /// </summary>
        public readonly ImmutableArray<string> VpcSubnets;

        [OutputConstructor]
        private GameServerGroupVpcSubnets(ImmutableArray<string> vpcSubnets)
        {
            VpcSubnets = vpcSubnets;
        }
    }
}
