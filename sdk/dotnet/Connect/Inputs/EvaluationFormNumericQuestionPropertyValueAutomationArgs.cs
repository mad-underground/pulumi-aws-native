// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// The automation property name of the question.
    /// </summary>
    public sealed class EvaluationFormNumericQuestionPropertyValueAutomationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The automation property label.
        /// </summary>
        [Input("label", required: true)]
        public Input<Pulumi.AwsNative.Connect.EvaluationFormNumericQuestionPropertyValueAutomationLabel> Label { get; set; } = null!;

        public EvaluationFormNumericQuestionPropertyValueAutomationArgs()
        {
        }
        public static new EvaluationFormNumericQuestionPropertyValueAutomationArgs Empty => new EvaluationFormNumericQuestionPropertyValueAutomationArgs();
    }
}