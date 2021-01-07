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
    public sealed class RecipeRecipeStep
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipestep.html#cfn-databrew-recipe-recipestep-action
        /// </summary>
        public readonly Outputs.RecipeAction Action;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-recipe-recipestep.html#cfn-databrew-recipe-recipestep-conditionexpressions
        /// </summary>
        public readonly ImmutableArray<Outputs.RecipeConditionExpression> ConditionExpressions;

        [OutputConstructor]
        private RecipeRecipeStep(
            Outputs.RecipeAction Action,

            ImmutableArray<Outputs.RecipeConditionExpression> ConditionExpressions)
        {
            this.Action = Action;
            this.ConditionExpressions = ConditionExpressions;
        }
    }
}
