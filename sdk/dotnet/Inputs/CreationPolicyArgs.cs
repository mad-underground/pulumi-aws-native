// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Inputs
{

    /// <summary>
    /// The CreationPolicy for a resource.
    /// </summary>
    public sealed class CreationPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For an Auto Scaling group replacement update, specifies how many instances must
        /// signal success for the update to succeed.
        /// </summary>
        [Input("AutoScalingCreationPolicy")]
        public Input<Inputs.AutoscalingCreationPolicyArgs>? AutoScalingCreationPolicy { get; set; }

        /// <summary>
        /// When AWS CloudFormation creates the associated resource, configures the number of
        /// required success signals and the length of time that AWS CloudFormation waits for those signals.
        /// </summary>
        [Input("ResourceSignal")]
        public Input<Inputs.ResourceSignalArgs>? ResourceSignal { get; set; }

        public CreationPolicyArgs()
        {
        }
    }
}