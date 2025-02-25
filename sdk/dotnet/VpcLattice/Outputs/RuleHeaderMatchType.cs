// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VpcLattice.Outputs
{

    [OutputType]
    public sealed class RuleHeaderMatchType
    {
        public readonly string? Contains;
        public readonly string? Exact;
        public readonly string? Prefix;

        [OutputConstructor]
        private RuleHeaderMatchType(
            string? contains,

            string? exact,

            string? prefix)
        {
            Contains = contains;
            Exact = exact;
            Prefix = prefix;
        }
    }
}
