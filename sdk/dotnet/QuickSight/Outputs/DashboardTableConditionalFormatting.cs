// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class DashboardTableConditionalFormatting
    {
        public readonly ImmutableArray<Outputs.DashboardTableConditionalFormattingOption> ConditionalFormattingOptions;

        [OutputConstructor]
        private DashboardTableConditionalFormatting(ImmutableArray<Outputs.DashboardTableConditionalFormattingOption> conditionalFormattingOptions)
        {
            ConditionalFormattingOptions = conditionalFormattingOptions;
        }
    }
}
