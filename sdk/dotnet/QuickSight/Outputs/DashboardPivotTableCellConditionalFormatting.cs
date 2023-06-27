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
    public sealed class DashboardPivotTableCellConditionalFormatting
    {
        public readonly string FieldId;
        public readonly Outputs.DashboardPivotTableConditionalFormattingScope? Scope;
        public readonly ImmutableArray<Outputs.DashboardPivotTableConditionalFormattingScope> Scopes;
        public readonly Outputs.DashboardTextConditionalFormat? TextFormat;

        [OutputConstructor]
        private DashboardPivotTableCellConditionalFormatting(
            string fieldId,

            Outputs.DashboardPivotTableConditionalFormattingScope? scope,

            ImmutableArray<Outputs.DashboardPivotTableConditionalFormattingScope> scopes,

            Outputs.DashboardTextConditionalFormat? textFormat)
        {
            FieldId = fieldId;
            Scope = scope;
            Scopes = scopes;
            TextFormat = textFormat;
        }
    }
}
