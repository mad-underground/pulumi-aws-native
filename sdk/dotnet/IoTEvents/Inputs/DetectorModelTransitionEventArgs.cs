// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.IoTEvents.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html
    /// </summary>
    public sealed class DetectorModelTransitionEventArgs : Pulumi.ResourceArgs
    {
        [Input("Actions")]
        private InputList<Inputs.DetectorModelActionArgs>? _Actions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-actions
        /// </summary>
        public InputList<Inputs.DetectorModelActionArgs> Actions
        {
            get => _Actions ?? (_Actions = new InputList<Inputs.DetectorModelActionArgs>());
            set => _Actions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-condition
        /// </summary>
        [Input("Condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-eventname
        /// </summary>
        [Input("EventName")]
        public Input<string>? EventName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotevents-detectormodel-transitionevent.html#cfn-iotevents-detectormodel-transitionevent-nextstate
        /// </summary>
        [Input("NextState")]
        public Input<string>? NextState { get; set; }

        public DetectorModelTransitionEventArgs()
        {
        }
    }
}