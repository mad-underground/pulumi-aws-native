// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EMR.Outputs
{

    [OutputType]
    public sealed class ClusterManagedScalingPolicy
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-managedscalingpolicy.html#cfn-elasticmapreduce-cluster-managedscalingpolicy-computelimits
        /// </summary>
        public readonly Outputs.ClusterComputeLimits? ComputeLimits;

        [OutputConstructor]
        private ClusterManagedScalingPolicy(Outputs.ClusterComputeLimits? ComputeLimits)
        {
            this.ComputeLimits = ComputeLimits;
        }
    }
}