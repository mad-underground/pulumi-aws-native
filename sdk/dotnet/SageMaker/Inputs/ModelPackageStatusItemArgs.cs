// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// Represents the overall status of a model package.
    /// </summary>
    public sealed class ModelPackageStatusItemArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the overall status is Failed, the reason for the failure.
        /// </summary>
        [Input("failureReason")]
        public Input<string>? FailureReason { get; set; }

        /// <summary>
        /// The name of the model package for which the overall status is being reported.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The current status.
        /// </summary>
        [Input("status", required: true)]
        public Input<Pulumi.AwsNative.SageMaker.ModelPackageStatusItemStatus> Status { get; set; } = null!;

        public ModelPackageStatusItemArgs()
        {
        }
    }
}