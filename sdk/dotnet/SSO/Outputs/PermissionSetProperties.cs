// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SSO.Outputs
{

    [OutputType]
    public sealed class PermissionSetProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-inlinepolicy
        /// </summary>
        public readonly string? InlinePolicy;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-instancearn
        /// </summary>
        public readonly string InstanceArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-managedpolicies
        /// </summary>
        public readonly ImmutableArray<string> ManagedPolicies;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-relaystatetype
        /// </summary>
        public readonly string? RelayStateType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-sessionduration
        /// </summary>
        public readonly string? SessionDuration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-permissionset.html#cfn-sso-permissionset-tags
        /// </summary>
        public readonly ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags;

        [OutputConstructor]
        private PermissionSetProperties(
            string? Description,

            string? InlinePolicy,

            string InstanceArn,

            ImmutableArray<string> ManagedPolicies,

            string Name,

            string? RelayStateType,

            string? SessionDuration,

            ImmutableArray<Pulumi.Cloudformation.Outputs.Tag> Tags)
        {
            this.Description = Description;
            this.InlinePolicy = InlinePolicy;
            this.InstanceArn = InstanceArn;
            this.ManagedPolicies = ManagedPolicies;
            this.Name = Name;
            this.RelayStateType = RelayStateType;
            this.SessionDuration = SessionDuration;
            this.Tags = Tags;
        }
    }
}
