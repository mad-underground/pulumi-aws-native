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
    public sealed class StackUserAssociationProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-authenticationtype
        /// </summary>
        public readonly string AuthenticationType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-sendemailnotification
        /// </summary>
        public readonly bool? SendEmailNotification;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-stackname
        /// </summary>
        public readonly string StackName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-username
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private StackUserAssociationProperties(
            string AuthenticationType,

            bool? SendEmailNotification,

            string StackName,

            string UserName)
        {
            this.AuthenticationType = AuthenticationType;
            this.SendEmailNotification = SendEmailNotification;
            this.StackName = StackName;
            this.UserName = UserName;
        }
    }
}