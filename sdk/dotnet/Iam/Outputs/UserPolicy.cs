// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Iam.Outputs
{

    /// <summary>
    /// Contains information about an attached policy.
    ///  An attached policy is a managed policy that has been attached to a user, group, or role.
    ///  For more information about managed policies, refer to [Managed Policies and Inline Policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the *User Guide*.
    /// </summary>
    [OutputType]
    public sealed class UserPolicy
    {
        /// <summary>
        /// The entire contents of the policy that defines permissions. For more information, see [Overview of JSON policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json).
        /// </summary>
        public readonly object PolicyDocument;
        /// <summary>
        /// The friendly name (not ARN) identifying the policy.
        /// </summary>
        public readonly string PolicyName;

        [OutputConstructor]
        private UserPolicy(
            object policyDocument,

            string policyName)
        {
            PolicyDocument = policyDocument;
            PolicyName = policyName;
        }
    }
}
