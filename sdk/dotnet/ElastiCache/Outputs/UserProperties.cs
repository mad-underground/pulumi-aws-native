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
    public sealed class UserProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-accessstring
        /// </summary>
        public readonly string? AccessString;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-engine
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-nopasswordrequired
        /// </summary>
        public readonly bool? NoPasswordRequired;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-passwords
        /// </summary>
        public readonly Outputs.UserPasswordList? Passwords;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-userid
        /// </summary>
        public readonly string UserId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-user.html#cfn-elasticache-user-username
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private UserProperties(
            string? AccessString,

            string Engine,

            bool? NoPasswordRequired,

            Outputs.UserPasswordList? Passwords,

            string UserId,

            string UserName)
        {
            this.AccessString = AccessString;
            this.Engine = Engine;
            this.NoPasswordRequired = NoPasswordRequired;
            this.Passwords = Passwords;
            this.UserId = UserId;
            this.UserName = UserName;
        }
    }
}
