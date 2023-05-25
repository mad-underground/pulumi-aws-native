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
    /// The evaluation form base item.
    /// </summary>
    public sealed class EvaluationFormBaseItemArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The evaluation form section item
        /// </summary>
        [Input("section", required: true)]
        public Input<Inputs.EvaluationFormSectionArgs> Section { get; set; } = null!;

        public EvaluationFormBaseItemArgs()
        {
        }
        public static new EvaluationFormBaseItemArgs Empty => new EvaluationFormBaseItemArgs();
    }
}