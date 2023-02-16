// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataBrew.Inputs
{

    /// <summary>
    /// A key-value pair to associate expression's substitution variable names with their values
    /// </summary>
    public sealed class RulesetSubstitutionValueArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value or column name
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        /// <summary>
        /// Variable name
        /// </summary>
        [Input("valueReference", required: true)]
        public Input<string> ValueReference { get; set; } = null!;

        public RulesetSubstitutionValueArgs()
        {
        }
        public static new RulesetSubstitutionValueArgs Empty => new RulesetSubstitutionValueArgs();
    }
}
