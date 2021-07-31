// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html
    /// </summary>
    [OutputType]
    public sealed class DetectorModelState
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-onenter
        /// </summary>
        public readonly Outputs.DetectorModelOnEnter? OnEnter;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-onexit
        /// </summary>
        public readonly Outputs.DetectorModelOnExit? OnExit;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-oninput
        /// </summary>
        public readonly Outputs.DetectorModelOnInput? OnInput;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-state.html#cfn-iotevents-detectormodel-state-statename
        /// </summary>
        public readonly string? StateName;

        [OutputConstructor]
        private DetectorModelState(
            Outputs.DetectorModelOnEnter? onEnter,

            Outputs.DetectorModelOnExit? onExit,

            Outputs.DetectorModelOnInput? onInput,

            string? stateName)
        {
            OnEnter = onEnter;
            OnExit = onExit;
            OnInput = onInput;
            StateName = stateName;
        }
    }
}
