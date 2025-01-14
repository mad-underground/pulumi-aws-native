// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Macie.Outputs
{

    [OutputType]
    public sealed class FindingsFilterCriterionAdditionalProperties
    {
        public readonly ImmutableArray<string> Eq;
        public readonly int? Gt;
        public readonly int? Gte;
        public readonly int? Lt;
        public readonly int? Lte;
        public readonly ImmutableArray<string> Neq;

        [OutputConstructor]
        private FindingsFilterCriterionAdditionalProperties(
            ImmutableArray<string> eq,

            int? gt,

            int? gte,

            int? lt,

            int? lte,

            ImmutableArray<string> neq)
        {
            Eq = eq;
            Gt = gt;
            Gte = gte;
            Lt = lt;
            Lte = lte;
            Neq = neq;
        }
    }
}
