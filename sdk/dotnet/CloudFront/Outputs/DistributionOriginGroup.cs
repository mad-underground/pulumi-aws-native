// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CloudFront.Outputs
{

    [OutputType]
    public sealed class DistributionOriginGroup
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-failovercriteria
        /// </summary>
        public readonly Outputs.DistributionOriginGroupFailoverCriteria FailoverCriteria;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-origingroup.html#cfn-cloudfront-distribution-origingroup-members
        /// </summary>
        public readonly Outputs.DistributionOriginGroupMembers Members;

        [OutputConstructor]
        private DistributionOriginGroup(
            Outputs.DistributionOriginGroupFailoverCriteria FailoverCriteria,

            string Id,

            Outputs.DistributionOriginGroupMembers Members)
        {
            this.FailoverCriteria = FailoverCriteria;
            this.Id = Id;
            this.Members = Members;
        }
    }
}