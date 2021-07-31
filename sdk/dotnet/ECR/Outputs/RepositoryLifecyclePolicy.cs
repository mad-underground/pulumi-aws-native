// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECR.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html
    /// </summary>
    [OutputType]
    public sealed class RepositoryLifecyclePolicy
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-lifecyclepolicytext
        /// </summary>
        public readonly string? LifecyclePolicyText;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-registryid
        /// </summary>
        public readonly string? RegistryId;

        [OutputConstructor]
        private RepositoryLifecyclePolicy(
            string? lifecyclePolicyText,

            string? registryId)
        {
            LifecyclePolicyText = lifecyclePolicyText;
            RegistryId = registryId;
        }
    }
}
