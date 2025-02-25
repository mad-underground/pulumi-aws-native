// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fis.Inputs
{

    /// <summary>
    /// Specifies an action for the experiment template.
    /// </summary>
    public sealed class ExperimentTemplateActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actionId", required: true)]
        public Input<string> ActionId { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// The parameters for the action, if applicable.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("startAfter")]
        private InputList<string>? _startAfter;
        public InputList<string> StartAfter
        {
            get => _startAfter ?? (_startAfter = new InputList<string>());
            set => _startAfter = value;
        }

        [Input("targets")]
        private InputMap<string>? _targets;

        /// <summary>
        /// One or more targets for the action.
        /// </summary>
        public InputMap<string> Targets
        {
            get => _targets ?? (_targets = new InputMap<string>());
            set => _targets = value;
        }

        public ExperimentTemplateActionArgs()
        {
        }
        public static new ExperimentTemplateActionArgs Empty => new ExperimentTemplateActionArgs();
    }
}
