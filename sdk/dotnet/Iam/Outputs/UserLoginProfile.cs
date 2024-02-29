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
    /// Creates a password for the specified user, giving the user the ability to access AWS services through the console. For more information about managing passwords, see [Managing Passwords](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html) in the *User Guide*.
    /// </summary>
    [OutputType]
    public sealed class UserLoginProfile
    {
        /// <summary>
        /// The user's password.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Specifies whether the user is required to set a new password on next sign-in.
        /// </summary>
        public readonly bool? PasswordResetRequired;

        [OutputConstructor]
        private UserLoginProfile(
            string password,

            bool? passwordResetRequired)
        {
            Password = password;
            PasswordResetRequired = passwordResetRequired;
        }
    }
}
