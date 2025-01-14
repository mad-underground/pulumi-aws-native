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
    public sealed class DashboardSmallMultiplesOptions
    {
        public readonly double? MaxVisibleColumns;
        public readonly double? MaxVisibleRows;
        public readonly Outputs.DashboardPanelConfiguration? PanelConfiguration;
        public readonly Outputs.DashboardSmallMultiplesAxisProperties? XAxis;
        public readonly Outputs.DashboardSmallMultiplesAxisProperties? YAxis;

        [OutputConstructor]
        private DashboardSmallMultiplesOptions(
            double? maxVisibleColumns,

            double? maxVisibleRows,

            Outputs.DashboardPanelConfiguration? panelConfiguration,

            Outputs.DashboardSmallMultiplesAxisProperties? xAxis,

            Outputs.DashboardSmallMultiplesAxisProperties? yAxis)
        {
            MaxVisibleColumns = maxVisibleColumns;
            MaxVisibleRows = maxVisibleRows;
            PanelConfiguration = panelConfiguration;
            XAxis = xAxis;
            YAxis = yAxis;
        }
    }
}
