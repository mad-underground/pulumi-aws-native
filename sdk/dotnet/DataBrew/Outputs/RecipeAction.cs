// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.DataBrew.Outputs
{

    [OutputType]
    public sealed class RecipeAction
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-action.html#cfn-databrew-recipe-action-operation
        /// </summary>
        public readonly string Operation;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-action.html#cfn-databrew-recipe-action-parameters
        /// </summary>
        public readonly object? Parameters;

        [OutputConstructor]
        private RecipeAction(
            string Operation,

            object? Parameters)
        {
            this.Operation = Operation;
            this.Parameters = Parameters;
        }
    }
}