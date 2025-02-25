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
    public sealed class DashboardParameterDropDownControl
    {
        public readonly Outputs.DashboardCascadingControlConfiguration? CascadingControlConfiguration;
        public readonly Outputs.DashboardDropDownControlDisplayOptions? DisplayOptions;
        public readonly string ParameterControlId;
        public readonly Outputs.DashboardParameterSelectableValues? SelectableValues;
        public readonly string SourceParameterName;
        public readonly string Title;
        public readonly Pulumi.AwsNative.QuickSight.DashboardSheetControlListType? Type;

        [OutputConstructor]
        private DashboardParameterDropDownControl(
            Outputs.DashboardCascadingControlConfiguration? cascadingControlConfiguration,

            Outputs.DashboardDropDownControlDisplayOptions? displayOptions,

            string parameterControlId,

            Outputs.DashboardParameterSelectableValues? selectableValues,

            string sourceParameterName,

            string title,

            Pulumi.AwsNative.QuickSight.DashboardSheetControlListType? type)
        {
            CascadingControlConfiguration = cascadingControlConfiguration;
            DisplayOptions = displayOptions;
            ParameterControlId = parameterControlId;
            SelectableValues = selectableValues;
            SourceParameterName = sourceParameterName;
            Title = title;
            Type = type;
        }
    }
}
