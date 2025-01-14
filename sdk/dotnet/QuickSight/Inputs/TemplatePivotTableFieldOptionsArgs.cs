// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class TemplatePivotTableFieldOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("collapseStateOptions")]
        private InputList<Inputs.TemplatePivotTableFieldCollapseStateOptionArgs>? _collapseStateOptions;
        public InputList<Inputs.TemplatePivotTableFieldCollapseStateOptionArgs> CollapseStateOptions
        {
            get => _collapseStateOptions ?? (_collapseStateOptions = new InputList<Inputs.TemplatePivotTableFieldCollapseStateOptionArgs>());
            set => _collapseStateOptions = value;
        }

        [Input("dataPathOptions")]
        private InputList<Inputs.TemplatePivotTableDataPathOptionArgs>? _dataPathOptions;
        public InputList<Inputs.TemplatePivotTableDataPathOptionArgs> DataPathOptions
        {
            get => _dataPathOptions ?? (_dataPathOptions = new InputList<Inputs.TemplatePivotTableDataPathOptionArgs>());
            set => _dataPathOptions = value;
        }

        [Input("selectedFieldOptions")]
        private InputList<Inputs.TemplatePivotTableFieldOptionArgs>? _selectedFieldOptions;
        public InputList<Inputs.TemplatePivotTableFieldOptionArgs> SelectedFieldOptions
        {
            get => _selectedFieldOptions ?? (_selectedFieldOptions = new InputList<Inputs.TemplatePivotTableFieldOptionArgs>());
            set => _selectedFieldOptions = value;
        }

        public TemplatePivotTableFieldOptionsArgs()
        {
        }
        public static new TemplatePivotTableFieldOptionsArgs Empty => new TemplatePivotTableFieldOptionsArgs();
    }
}
