// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// Baseline configuration used to validate that the data conforms to the specified constraints and statistics.
    /// </summary>
    [OutputType]
    public sealed class ModelBiasJobDefinitionModelBiasBaselineConfig
    {
        public readonly string? BaseliningJobName;
        public readonly Outputs.ModelBiasJobDefinitionConstraintsResource? ConstraintsResource;

        [OutputConstructor]
        private ModelBiasJobDefinitionModelBiasBaselineConfig(
            string? baseliningJobName,

            Outputs.ModelBiasJobDefinitionConstraintsResource? constraintsResource)
        {
            BaseliningJobName = baseliningJobName;
            ConstraintsResource = constraintsResource;
        }
    }
}
