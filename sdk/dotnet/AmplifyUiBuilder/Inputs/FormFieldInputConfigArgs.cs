// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormFieldInputConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultChecked")]
        public Input<bool>? DefaultChecked { get; set; }

        [Input("defaultCountryCode")]
        public Input<string>? DefaultCountryCode { get; set; }

        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        [Input("descriptiveText")]
        public Input<string>? DescriptiveText { get; set; }

        [Input("fileUploaderConfig")]
        public Input<Inputs.FormFileUploaderFieldConfigArgs>? FileUploaderConfig { get; set; }

        [Input("isArray")]
        public Input<bool>? IsArray { get; set; }

        [Input("maxValue")]
        public Input<double>? MaxValue { get; set; }

        [Input("minValue")]
        public Input<double>? MinValue { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("placeholder")]
        public Input<string>? Placeholder { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("required")]
        public Input<bool>? Required { get; set; }

        [Input("step")]
        public Input<double>? Step { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        [Input("valueMappings")]
        public Input<Inputs.FormValueMappingsArgs>? ValueMappings { get; set; }

        public FormFieldInputConfigArgs()
        {
        }
        public static new FormFieldInputConfigArgs Empty => new FormFieldInputConfigArgs();
    }
}
