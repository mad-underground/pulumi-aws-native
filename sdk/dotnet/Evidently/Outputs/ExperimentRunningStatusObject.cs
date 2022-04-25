// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Evidently.Outputs
{

    [OutputType]
    public sealed class ExperimentRunningStatusObject
    {
        /// <summary>
        /// Provide the analysis Completion time for an experiment
        /// </summary>
        public readonly string? AnalysisCompleteTime;
        /// <summary>
        /// Provide CANCELLED or COMPLETED desired state when stopping an experiment
        /// </summary>
        public readonly string? DesiredState;
        /// <summary>
        /// Reason is a required input for stopping the experiment
        /// </summary>
        public readonly string? Reason;
        /// <summary>
        /// Provide START or STOP action to apply on an experiment
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private ExperimentRunningStatusObject(
            string? analysisCompleteTime,

            string? desiredState,

            string? reason,

            string? status)
        {
            AnalysisCompleteTime = analysisCompleteTime;
            DesiredState = desiredState;
            Reason = reason;
            Status = status;
        }
    }
}