// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionTaskContainerDependencyArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition", required: true)]
        public Input<string> Condition { get; set; } = null!;

        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        public JobDefinitionTaskContainerDependencyArgs()
        {
        }
        public static new JobDefinitionTaskContainerDependencyArgs Empty => new JobDefinitionTaskContainerDependencyArgs();
    }
}
