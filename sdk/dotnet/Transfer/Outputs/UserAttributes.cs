// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Transfer.Outputs
{

    [OutputType]
    public sealed class UserAttributes
    {
        public readonly string Arn;
        public readonly string ServerId;
        public readonly string UserName;

        [OutputConstructor]
        private UserAttributes(
            string Arn,

            string ServerId,

            string UserName)
        {
            this.Arn = Arn;
            this.ServerId = ServerId;
            this.UserName = UserName;
        }
    }
}