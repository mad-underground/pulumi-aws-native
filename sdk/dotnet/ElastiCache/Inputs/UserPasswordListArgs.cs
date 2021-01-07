// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ElastiCache.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-passwordlist.html
    /// </summary>
    public sealed class UserPasswordListArgs : Pulumi.ResourceArgs
    {
        [Input("PasswordList")]
        private InputList<string>? _PasswordList;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-user-passwordlist.html#cfn-elasticache-user-passwordlist-passwordlist
        /// </summary>
        public InputList<string> PasswordList
        {
            get => _PasswordList ?? (_PasswordList = new InputList<string>());
            set => _PasswordList = value;
        }

        public UserPasswordListArgs()
        {
        }
    }
}
