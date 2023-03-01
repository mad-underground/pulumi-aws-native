// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Outputs
{

    /// <summary>
    /// Response headers that indicate success or failure of a login request
    /// </summary>
    [OutputType]
    public sealed class WebACLResponseInspectionHeader
    {
        public readonly ImmutableArray<string> FailureValues;
        public readonly string Name;
        public readonly ImmutableArray<string> SuccessValues;

        [OutputConstructor]
        private WebACLResponseInspectionHeader(
            ImmutableArray<string> failureValues,

            string name,

            ImmutableArray<string> successValues)
        {
            FailureValues = failureValues;
            Name = name;
            SuccessValues = successValues;
        }
    }
}
