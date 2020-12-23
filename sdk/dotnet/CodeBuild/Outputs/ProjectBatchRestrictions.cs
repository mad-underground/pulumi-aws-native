// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CodeBuild.Outputs
{

    [OutputType]
    public sealed class ProjectBatchRestrictions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-batchrestrictions.html#cfn-codebuild-project-batchrestrictions-computetypesallowed
        /// </summary>
        public readonly ImmutableArray<string> ComputeTypesAllowed;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-batchrestrictions.html#cfn-codebuild-project-batchrestrictions-maximumbuildsallowed
        /// </summary>
        public readonly int? MaximumBuildsAllowed;

        [OutputConstructor]
        private ProjectBatchRestrictions(
            ImmutableArray<string> ComputeTypesAllowed,

            int? MaximumBuildsAllowed)
        {
            this.ComputeTypesAllowed = ComputeTypesAllowed;
            this.MaximumBuildsAllowed = MaximumBuildsAllowed;
        }
    }
}
