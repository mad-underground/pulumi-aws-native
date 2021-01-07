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
    public sealed class NetworkInsightsAnalysisAlternatePathHint
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-alternatepathhint.html#cfn-ec2-networkinsightsanalysis-alternatepathhint-componentarn
        /// </summary>
        public readonly string? ComponentArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinsightsanalysis-alternatepathhint.html#cfn-ec2-networkinsightsanalysis-alternatepathhint-componentid
        /// </summary>
        public readonly string? ComponentId;

        [OutputConstructor]
        private NetworkInsightsAnalysisAlternatePathHint(
            string? ComponentArn,

            string? ComponentId)
        {
            this.ComponentArn = ComponentArn;
            this.ComponentId = ComponentId;
        }
    }
}