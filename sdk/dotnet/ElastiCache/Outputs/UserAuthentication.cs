// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ElastiCache.Outputs
{

    [OutputType]
    public sealed class UserAuthentication
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html#cfn-elasticache-user-authentication-passwordcount
        /// </summary>
        public readonly int? PasswordCount;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-authentication.html#cfn-elasticache-user-authentication-type
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private UserAuthentication(
            int? PasswordCount,

            string? Type)
        {
            this.PasswordCount = PasswordCount;
            this.Type = Type;
        }
    }
}
