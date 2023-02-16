// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFRegional.Outputs
{

    [OutputType]
    public sealed class RateBasedRulePredicate
    {
        public readonly string DataId;
        public readonly bool Negated;
        public readonly string Type;

        [OutputConstructor]
        private RateBasedRulePredicate(
            string dataId,

            bool negated,

            string type)
        {
            DataId = dataId;
            Negated = negated;
            Type = type;
        }
    }
}
