// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TargetGroupMatcher
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-matcher.html#cfn-elasticloadbalancingv2-targetgroup-matcher-httpcode
        /// </summary>
        public readonly string? HttpCode;

        [OutputConstructor]
        private TargetGroupMatcher(string? HttpCode)
        {
            this.HttpCode = HttpCode;
        }
    }
}
