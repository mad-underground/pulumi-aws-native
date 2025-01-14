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
    /// Defines the input needed to run a transform job using the inference specification specified in the algorithm.
    /// </summary>
    [OutputType]
    public sealed class ModelPackageTransformJobDefinition
    {
        /// <summary>
        /// A string that determines the number of records included in a single mini-batch.
        /// </summary>
        public readonly Pulumi.AwsNative.SageMaker.ModelPackageTransformJobDefinitionBatchStrategy? BatchStrategy;
        public readonly Outputs.ModelPackageEnvironment? Environment;
        /// <summary>
        /// The maximum number of parallel requests that can be sent to each instance in a transform job. The default value is 1.
        /// </summary>
        public readonly int? MaxConcurrentTransforms;
        /// <summary>
        /// The maximum payload size allowed, in MB. A payload is the data portion of a record (without metadata).
        /// </summary>
        public readonly int? MaxPayloadInMb;
        public readonly Outputs.ModelPackageTransformInput TransformInput;
        public readonly Outputs.ModelPackageTransformOutput TransformOutput;
        public readonly Outputs.ModelPackageTransformResources TransformResources;

        [OutputConstructor]
        private ModelPackageTransformJobDefinition(
            Pulumi.AwsNative.SageMaker.ModelPackageTransformJobDefinitionBatchStrategy? batchStrategy,

            Outputs.ModelPackageEnvironment? environment,

            int? maxConcurrentTransforms,

            int? maxPayloadInMb,

            Outputs.ModelPackageTransformInput transformInput,

            Outputs.ModelPackageTransformOutput transformOutput,

            Outputs.ModelPackageTransformResources transformResources)
        {
            BatchStrategy = batchStrategy;
            Environment = environment;
            MaxConcurrentTransforms = maxConcurrentTransforms;
            MaxPayloadInMb = maxPayloadInMb;
            TransformInput = transformInput;
            TransformOutput = transformOutput;
            TransformResources = transformResources;
        }
    }
}
