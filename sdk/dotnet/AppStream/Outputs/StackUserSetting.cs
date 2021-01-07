// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppStream.Outputs
{

    [OutputType]
    public sealed class StackUserSetting
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-action
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-usersetting.html#cfn-appstream-stack-usersetting-permission
        /// </summary>
        public readonly string Permission;

        [OutputConstructor]
        private StackUserSetting(
            string Action,

            string Permission)
        {
            this.Action = Action;
            this.Permission = Permission;
        }
    }
}