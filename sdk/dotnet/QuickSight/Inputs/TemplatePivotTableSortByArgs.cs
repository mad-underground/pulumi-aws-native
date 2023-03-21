// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePivotTableSortByArgs : global::Pulumi.ResourceArgs
    {
        [Input("column")]
        public Input<Inputs.TemplateColumnSortArgs>? Column { get; set; }

        [Input("dataPath")]
        public Input<Inputs.TemplateDataPathSortArgs>? DataPath { get; set; }

        [Input("field")]
        public Input<Inputs.TemplateFieldSortArgs>? Field { get; set; }

        public TemplatePivotTableSortByArgs()
        {
        }
        public static new TemplatePivotTableSortByArgs Empty => new TemplatePivotTableSortByArgs();
    }
}
