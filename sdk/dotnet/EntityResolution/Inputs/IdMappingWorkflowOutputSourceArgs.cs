// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EntityResolution.Inputs
{

    public sealed class IdMappingWorkflowOutputSourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("kmsArn")]
        public Input<string>? KmsArn { get; set; }

        /// <summary>
        /// The S3 path to which Entity Resolution will write the output table
        /// </summary>
        [Input("outputS3Path", required: true)]
        public Input<string> OutputS3Path { get; set; } = null!;

        public IdMappingWorkflowOutputSourceArgs()
        {
        }
        public static new IdMappingWorkflowOutputSourceArgs Empty => new IdMappingWorkflowOutputSourceArgs();
    }
}
