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
    public sealed class LaunchTemplateSpotOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions-blockdurationminutes
        /// </summary>
        public readonly int? BlockDurationMinutes;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions-instanceinterruptionbehavior
        /// </summary>
        public readonly string? InstanceInterruptionBehavior;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions-maxprice
        /// </summary>
        public readonly string? MaxPrice;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions-spotinstancetype
        /// </summary>
        public readonly string? SpotInstanceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions.html#cfn-ec2-launchtemplate-launchtemplatedata-instancemarketoptions-spotoptions-validuntil
        /// </summary>
        public readonly string? ValidUntil;

        [OutputConstructor]
        private LaunchTemplateSpotOptions(
            int? BlockDurationMinutes,

            string? InstanceInterruptionBehavior,

            string? MaxPrice,

            string? SpotInstanceType,

            string? ValidUntil)
        {
            this.BlockDurationMinutes = BlockDurationMinutes;
            this.InstanceInterruptionBehavior = InstanceInterruptionBehavior;
            this.MaxPrice = MaxPrice;
            this.SpotInstanceType = SpotInstanceType;
            this.ValidUntil = ValidUntil;
        }
    }
}