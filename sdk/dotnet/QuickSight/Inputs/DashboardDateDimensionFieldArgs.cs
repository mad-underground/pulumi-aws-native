// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class DashboardDateDimensionFieldArgs : global::Pulumi.ResourceArgs
    {
        [Input("column", required: true)]
        public Input<Inputs.DashboardColumnIdentifierArgs> Column { get; set; } = null!;

        [Input("dateGranularity")]
        public Input<Pulumi.AwsNative.QuickSight.DashboardTimeGranularity>? DateGranularity { get; set; }

        [Input("fieldId", required: true)]
        public Input<string> FieldId { get; set; } = null!;

        [Input("formatConfiguration")]
        public Input<Inputs.DashboardDateTimeFormatConfigurationArgs>? FormatConfiguration { get; set; }

        [Input("hierarchyId")]
        public Input<string>? HierarchyId { get; set; }

        public DashboardDateDimensionFieldArgs()
        {
        }
        public static new DashboardDateDimensionFieldArgs Empty => new DashboardDateDimensionFieldArgs();
    }
}
