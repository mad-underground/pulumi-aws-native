// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Outputs
{

    [OutputType]
    public sealed class FormFieldInputConfig
    {
        public readonly bool? DefaultChecked;
        public readonly string? DefaultCountryCode;
        public readonly string? DefaultValue;
        public readonly string? DescriptiveText;
        public readonly Outputs.FormFileUploaderFieldConfig? FileUploaderConfig;
        public readonly bool? IsArray;
        public readonly double? MaxValue;
        public readonly double? MinValue;
        public readonly string? Name;
        public readonly string? Placeholder;
        public readonly bool? ReadOnly;
        public readonly bool? Required;
        public readonly double? Step;
        public readonly string Type;
        public readonly string? Value;
        public readonly Outputs.FormValueMappings? ValueMappings;

        [OutputConstructor]
        private FormFieldInputConfig(
            bool? defaultChecked,

            string? defaultCountryCode,

            string? defaultValue,

            string? descriptiveText,

            Outputs.FormFileUploaderFieldConfig? fileUploaderConfig,

            bool? isArray,

            double? maxValue,

            double? minValue,

            string? name,

            string? placeholder,

            bool? readOnly,

            bool? required,

            double? step,

            string type,

            string? value,

            Outputs.FormValueMappings? valueMappings)
        {
            DefaultChecked = defaultChecked;
            DefaultCountryCode = defaultCountryCode;
            DefaultValue = defaultValue;
            DescriptiveText = descriptiveText;
            FileUploaderConfig = fileUploaderConfig;
            IsArray = isArray;
            MaxValue = maxValue;
            MinValue = minValue;
            Name = name;
            Placeholder = placeholder;
            ReadOnly = readOnly;
            Required = required;
            Step = step;
            Type = type;
            Value = value;
            ValueMappings = valueMappings;
        }
    }
}
