// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EKS.Outputs
{

    [OutputType]
    public sealed class FargateProfileProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-clustername
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-fargateprofilename
        /// </summary>
        public readonly string? FargateProfileName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-podexecutionrolearn
        /// </summary>
        public readonly string PodExecutionRoleArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-selectors
        /// </summary>
        public readonly ImmutableArray<Outputs.FargateProfileSelector> Selectors;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-subnets
        /// </summary>
        public readonly ImmutableArray<string> Subnets;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html#cfn-eks-fargateprofile-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private FargateProfileProperties(
            string ClusterName,

            string? FargateProfileName,

            string PodExecutionRoleArn,

            ImmutableArray<Outputs.FargateProfileSelector> Selectors,

            ImmutableArray<string> Subnets,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.ClusterName = ClusterName;
            this.FargateProfileName = FargateProfileName;
            this.PodExecutionRoleArn = PodExecutionRoleArn;
            this.Selectors = Selectors;
            this.Subnets = Subnets;
            this.Tags = Tags;
        }
    }
}
