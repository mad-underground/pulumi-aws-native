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
    public sealed class DashboardFilledMapShapeConditionalFormatting
    {
        public readonly string FieldId;
        public readonly Outputs.DashboardShapeConditionalFormat? Format;

        [OutputConstructor]
        private DashboardFilledMapShapeConditionalFormatting(
            string fieldId,

            Outputs.DashboardShapeConditionalFormat? format)
        {
            FieldId = fieldId;
            Format = format;
        }
    }
}
