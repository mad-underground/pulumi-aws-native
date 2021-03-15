// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Synthetics.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html
    /// </summary>
    public sealed class CanaryRunConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-activetracing
        /// </summary>
        [Input("activeTracing")]
        public Input<bool>? ActiveTracing { get; set; }

        [Input("environmentVariables")]
        private InputMap<string>? _environmentVariables;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-environmentvariables
        /// </summary>
        public InputMap<string> EnvironmentVariables
        {
            get => _environmentVariables ?? (_environmentVariables = new InputMap<string>());
            set => _environmentVariables = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-memoryinmb
        /// </summary>
        [Input("memoryInMB")]
        public Input<int>? MemoryInMB { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-synthetics-canary-runconfig.html#cfn-synthetics-canary-runconfig-timeoutinseconds
        /// </summary>
        [Input("timeoutInSeconds", required: true)]
        public Input<int> TimeoutInSeconds { get; set; } = null!;

        public CanaryRunConfigArgs()
        {
        }
    }
}
