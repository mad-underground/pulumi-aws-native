// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    public sealed class TaskDefinitionFirelensConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("options")]
        private InputMap<string>? _options;
        public InputMap<string> Options
        {
            get => _options ?? (_options = new InputMap<string>());
            set => _options = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public TaskDefinitionFirelensConfigurationArgs()
        {
        }
        public static new TaskDefinitionFirelensConfigurationArgs Empty => new TaskDefinitionFirelensConfigurationArgs();
    }
}
