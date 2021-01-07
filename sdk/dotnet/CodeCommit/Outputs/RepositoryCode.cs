// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.CodeCommit.Outputs
{

    [OutputType]
    public sealed class RepositoryCode
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-branchname
        /// </summary>
        public readonly string? BranchName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codecommit-repository-code.html#cfn-codecommit-repository-code-s3
        /// </summary>
        public readonly Outputs.RepositoryS3 S3;

        [OutputConstructor]
        private RepositoryCode(
            string? BranchName,

            Outputs.RepositoryS3 S3)
        {
            this.BranchName = BranchName;
            this.S3 = S3;
        }
    }
}
