// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EFS.Outputs
{

    [OutputType]
    public sealed class AccessPointRootDirectory
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-creationinfo
        /// </summary>
        public readonly Outputs.AccessPointCreationInfo? CreationInfo;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-efs-accesspoint-rootdirectory.html#cfn-efs-accesspoint-rootdirectory-path
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private AccessPointRootDirectory(
            Outputs.AccessPointCreationInfo? creationInfo,

            string? path)
        {
            CreationInfo = creationInfo;
            Path = path;
        }
    }
}
