// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Outputs
{

    /// <summary>
    /// An access policy to associate with the current access entry.
    /// </summary>
    [OutputType]
    public sealed class AccessEntryAccessPolicy
    {
        public readonly Outputs.AccessEntryAccessScope AccessScope;
        /// <summary>
        /// The ARN of the access policy to add to the access entry.
        /// </summary>
        public readonly string PolicyArn;

        [OutputConstructor]
        private AccessEntryAccessPolicy(
            Outputs.AccessEntryAccessScope accessScope,

            string policyArn)
        {
            AccessScope = accessScope;
            PolicyArn = policyArn;
        }
    }
}
