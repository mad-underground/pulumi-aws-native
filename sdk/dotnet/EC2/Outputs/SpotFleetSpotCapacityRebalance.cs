// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EC2.Outputs
{

    [OutputType]
    public sealed class SpotFleetSpotCapacityRebalance
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-spotcapacityrebalance.html#cfn-ec2-spotfleet-spotcapacityrebalance-replacementstrategy
        /// </summary>
        public readonly string? ReplacementStrategy;

        [OutputConstructor]
        private SpotFleetSpotCapacityRebalance(string? ReplacementStrategy)
        {
            this.ReplacementStrategy = ReplacementStrategy;
        }
    }
}