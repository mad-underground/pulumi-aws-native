// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.SageMaker.Outputs
{

    [OutputType]
    public sealed class ModelExplainabilityJobDefinitionModelExplainabilityJobInput
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput.html#cfn-sagemaker-modelexplainabilityjobdefinition-modelexplainabilityjobinput-endpointinput
        /// </summary>
        public readonly Outputs.ModelExplainabilityJobDefinitionEndpointInput EndpointInput;

        [OutputConstructor]
        private ModelExplainabilityJobDefinitionModelExplainabilityJobInput(Outputs.ModelExplainabilityJobDefinitionEndpointInput EndpointInput)
        {
            this.EndpointInput = EndpointInput;
        }
    }
}