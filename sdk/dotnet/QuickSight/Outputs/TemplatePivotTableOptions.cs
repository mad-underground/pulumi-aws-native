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
    public sealed class TemplatePivotTableOptions
    {
        public readonly Outputs.TemplateTableCellStyle? CellStyle;
        public readonly Pulumi.AwsNative.QuickSight.TemplateVisibility? CollapsedRowDimensionsVisibility;
        public readonly Outputs.TemplateTableCellStyle? ColumnHeaderStyle;
        public readonly Pulumi.AwsNative.QuickSight.TemplateVisibility? ColumnNamesVisibility;
        /// <summary>
        /// String based length that is composed of value and unit in px
        /// </summary>
        public readonly string? DefaultCellWidth;
        public readonly Pulumi.AwsNative.QuickSight.TemplatePivotTableMetricPlacement? MetricPlacement;
        public readonly Outputs.TemplateRowAlternateColorOptions? RowAlternateColorOptions;
        public readonly Outputs.TemplateTableCellStyle? RowFieldNamesStyle;
        public readonly Outputs.TemplateTableCellStyle? RowHeaderStyle;
        public readonly Outputs.TemplatePivotTableRowsLabelOptions? RowsLabelOptions;
        public readonly Pulumi.AwsNative.QuickSight.TemplatePivotTableRowsLayout? RowsLayout;
        public readonly Pulumi.AwsNative.QuickSight.TemplateVisibility? SingleMetricVisibility;
        public readonly Pulumi.AwsNative.QuickSight.TemplateVisibility? ToggleButtonsVisibility;

        [OutputConstructor]
        private TemplatePivotTableOptions(
            Outputs.TemplateTableCellStyle? cellStyle,

            Pulumi.AwsNative.QuickSight.TemplateVisibility? collapsedRowDimensionsVisibility,

            Outputs.TemplateTableCellStyle? columnHeaderStyle,

            Pulumi.AwsNative.QuickSight.TemplateVisibility? columnNamesVisibility,

            string? defaultCellWidth,

            Pulumi.AwsNative.QuickSight.TemplatePivotTableMetricPlacement? metricPlacement,

            Outputs.TemplateRowAlternateColorOptions? rowAlternateColorOptions,

            Outputs.TemplateTableCellStyle? rowFieldNamesStyle,

            Outputs.TemplateTableCellStyle? rowHeaderStyle,

            Outputs.TemplatePivotTableRowsLabelOptions? rowsLabelOptions,

            Pulumi.AwsNative.QuickSight.TemplatePivotTableRowsLayout? rowsLayout,

            Pulumi.AwsNative.QuickSight.TemplateVisibility? singleMetricVisibility,

            Pulumi.AwsNative.QuickSight.TemplateVisibility? toggleButtonsVisibility)
        {
            CellStyle = cellStyle;
            CollapsedRowDimensionsVisibility = collapsedRowDimensionsVisibility;
            ColumnHeaderStyle = columnHeaderStyle;
            ColumnNamesVisibility = columnNamesVisibility;
            DefaultCellWidth = defaultCellWidth;
            MetricPlacement = metricPlacement;
            RowAlternateColorOptions = rowAlternateColorOptions;
            RowFieldNamesStyle = rowFieldNamesStyle;
            RowHeaderStyle = rowHeaderStyle;
            RowsLabelOptions = rowsLabelOptions;
            RowsLayout = rowsLayout;
            SingleMetricVisibility = singleMetricVisibility;
            ToggleButtonsVisibility = toggleButtonsVisibility;
        }
    }
}
