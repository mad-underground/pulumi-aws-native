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
    public sealed class FargateProfileSelector
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html#cfn-eks-fargateprofile-selector-labels
        /// </summary>
        public readonly ImmutableArray<Outputs.FargateProfileLabel> Labels;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-fargateprofile-selector.html#cfn-eks-fargateprofile-selector-namespace
        /// </summary>
        public readonly string Namespace;

        [OutputConstructor]
        private FargateProfileSelector(
            ImmutableArray<Outputs.FargateProfileLabel> Labels,

            string Namespace)
        {
            this.Labels = Labels;
            this.Namespace = Namespace;
        }
    }
}