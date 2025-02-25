// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplateSankeyDiagramChartConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("dataLabels")]
        public Input<Inputs.TemplateDataLabelOptionsArgs>? DataLabels { get; set; }

        [Input("fieldWells")]
        public Input<Inputs.TemplateSankeyDiagramFieldWellsArgs>? FieldWells { get; set; }

        [Input("sortConfiguration")]
        public Input<Inputs.TemplateSankeyDiagramSortConfigurationArgs>? SortConfiguration { get; set; }

        public TemplateSankeyDiagramChartConfigurationArgs()
        {
        }
        public static new TemplateSankeyDiagramChartConfigurationArgs Empty => new TemplateSankeyDiagramChartConfigurationArgs();
    }
}
