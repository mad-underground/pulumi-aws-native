// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class ComponentVersionLambdaVolumeMount
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-addgroupowner
        /// </summary>
        public readonly bool? AddGroupOwner;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-destinationpath
        /// </summary>
        public readonly string? DestinationPath;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-permission
        /// </summary>
        public readonly string? Permission;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdavolumemount.html#cfn-greengrassv2-componentversion-lambdavolumemount-sourcepath
        /// </summary>
        public readonly string? SourcePath;

        [OutputConstructor]
        private ComponentVersionLambdaVolumeMount(
            bool? AddGroupOwner,

            string? DestinationPath,

            string? Permission,

            string? SourcePath)
        {
            this.AddGroupOwner = AddGroupOwner;
            this.DestinationPath = DestinationPath;
            this.Permission = Permission;
            this.SourcePath = SourcePath;
        }
    }
}
