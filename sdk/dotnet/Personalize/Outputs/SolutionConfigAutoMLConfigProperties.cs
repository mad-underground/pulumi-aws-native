// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Personalize.Outputs
{

    /// <summary>
    /// The AutoMLConfig object containing a list of recipes to search when AutoML is performed.
    /// </summary>
    [OutputType]
    public sealed class SolutionConfigAutoMLConfigProperties
    {
        /// <summary>
        /// The metric to optimize.
        /// </summary>
        public readonly string? MetricName;
        /// <summary>
        /// The list of candidate recipes.
        /// </summary>
        public readonly ImmutableArray<string> RecipeList;

        [OutputConstructor]
        private SolutionConfigAutoMLConfigProperties(
            string? metricName,

            ImmutableArray<string> recipeList)
        {
            MetricName = metricName;
            RecipeList = recipeList;
        }
    }
}
