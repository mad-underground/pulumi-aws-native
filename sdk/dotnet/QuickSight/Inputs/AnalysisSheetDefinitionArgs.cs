// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisSheetDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("contentType")]
        public Input<Pulumi.AwsNative.QuickSight.AnalysisSheetContentType>? ContentType { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("filterControls")]
        private InputList<Inputs.AnalysisFilterControlArgs>? _filterControls;
        public InputList<Inputs.AnalysisFilterControlArgs> FilterControls
        {
            get => _filterControls ?? (_filterControls = new InputList<Inputs.AnalysisFilterControlArgs>());
            set => _filterControls = value;
        }

        [Input("layouts")]
        private InputList<Inputs.AnalysisLayoutArgs>? _layouts;
        public InputList<Inputs.AnalysisLayoutArgs> Layouts
        {
            get => _layouts ?? (_layouts = new InputList<Inputs.AnalysisLayoutArgs>());
            set => _layouts = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameterControls")]
        private InputList<Inputs.AnalysisParameterControlArgs>? _parameterControls;
        public InputList<Inputs.AnalysisParameterControlArgs> ParameterControls
        {
            get => _parameterControls ?? (_parameterControls = new InputList<Inputs.AnalysisParameterControlArgs>());
            set => _parameterControls = value;
        }

        [Input("sheetControlLayouts")]
        private InputList<Inputs.AnalysisSheetControlLayoutArgs>? _sheetControlLayouts;
        public InputList<Inputs.AnalysisSheetControlLayoutArgs> SheetControlLayouts
        {
            get => _sheetControlLayouts ?? (_sheetControlLayouts = new InputList<Inputs.AnalysisSheetControlLayoutArgs>());
            set => _sheetControlLayouts = value;
        }

        [Input("sheetId", required: true)]
        public Input<string> SheetId { get; set; } = null!;

        [Input("textBoxes")]
        private InputList<Inputs.AnalysisSheetTextBoxArgs>? _textBoxes;
        public InputList<Inputs.AnalysisSheetTextBoxArgs> TextBoxes
        {
            get => _textBoxes ?? (_textBoxes = new InputList<Inputs.AnalysisSheetTextBoxArgs>());
            set => _textBoxes = value;
        }

        [Input("title")]
        public Input<string>? Title { get; set; }

        [Input("visuals")]
        private InputList<Inputs.AnalysisVisualArgs>? _visuals;
        public InputList<Inputs.AnalysisVisualArgs> Visuals
        {
            get => _visuals ?? (_visuals = new InputList<Inputs.AnalysisVisualArgs>());
            set => _visuals = value;
        }

        public AnalysisSheetDefinitionArgs()
        {
        }
        public static new AnalysisSheetDefinitionArgs Empty => new AnalysisSheetDefinitionArgs();
    }
}