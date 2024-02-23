// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder.Inputs
{

    public sealed class FormValueMappingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("bindingProperties")]
        private InputMap<Inputs.FormInputBindingPropertiesValueArgs>? _bindingProperties;
        public InputMap<Inputs.FormInputBindingPropertiesValueArgs> BindingProperties
        {
            get => _bindingProperties ?? (_bindingProperties = new InputMap<Inputs.FormInputBindingPropertiesValueArgs>());
            set => _bindingProperties = value;
        }

        [Input("values", required: true)]
        private InputList<Inputs.FormValueMappingArgs>? _values;
        public InputList<Inputs.FormValueMappingArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.FormValueMappingArgs>());
            set => _values = value;
        }

        public FormValueMappingsArgs()
        {
        }
        public static new FormValueMappingsArgs Empty => new FormValueMappingsArgs();
    }
}
