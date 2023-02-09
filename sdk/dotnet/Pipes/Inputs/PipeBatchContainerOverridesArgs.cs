// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pipes.Inputs
{

    public sealed class PipeBatchContainerOverridesArgs : global::Pulumi.ResourceArgs
    {
        [Input("command")]
        private InputList<string>? _command;
        public InputList<string> Command
        {
            get => _command ?? (_command = new InputList<string>());
            set => _command = value;
        }

        [Input("environment")]
        private InputList<Inputs.PipeBatchEnvironmentVariableArgs>? _environment;
        public InputList<Inputs.PipeBatchEnvironmentVariableArgs> Environment
        {
            get => _environment ?? (_environment = new InputList<Inputs.PipeBatchEnvironmentVariableArgs>());
            set => _environment = value;
        }

        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        [Input("resourceRequirements")]
        private InputList<Inputs.PipeBatchResourceRequirementArgs>? _resourceRequirements;
        public InputList<Inputs.PipeBatchResourceRequirementArgs> ResourceRequirements
        {
            get => _resourceRequirements ?? (_resourceRequirements = new InputList<Inputs.PipeBatchResourceRequirementArgs>());
            set => _resourceRequirements = value;
        }

        public PipeBatchContainerOverridesArgs()
        {
        }
        public static new PipeBatchContainerOverridesArgs Empty => new PipeBatchContainerOverridesArgs();
    }
}