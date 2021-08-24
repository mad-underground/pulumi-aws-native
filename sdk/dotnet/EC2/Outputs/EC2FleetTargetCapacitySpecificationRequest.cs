// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html
    /// </summary>
    [OutputType]
    public sealed class EC2FleetTargetCapacitySpecificationRequest
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-defaulttargetcapacitytype
        /// </summary>
        public readonly string? DefaultTargetCapacityType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-ondemandtargetcapacity
        /// </summary>
        public readonly int? OnDemandTargetCapacity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-spottargetcapacity
        /// </summary>
        public readonly int? SpotTargetCapacity;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ec2fleet-targetcapacityspecificationrequest.html#cfn-ec2-ec2fleet-targetcapacityspecificationrequest-totaltargetcapacity
        /// </summary>
        public readonly int TotalTargetCapacity;

        [OutputConstructor]
        private EC2FleetTargetCapacitySpecificationRequest(
            string? defaultTargetCapacityType,

            int? onDemandTargetCapacity,

            int? spotTargetCapacity,

            int totalTargetCapacity)
        {
            DefaultTargetCapacityType = defaultTargetCapacityType;
            OnDemandTargetCapacity = onDemandTargetCapacity;
            SpotTargetCapacity = spotTargetCapacity;
            TotalTargetCapacity = totalTargetCapacity;
        }
    }
}