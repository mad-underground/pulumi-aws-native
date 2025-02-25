// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTAnalytics.Outputs
{

    [OutputType]
    public sealed class PipelineLambda
    {
        public readonly int BatchSize;
        public readonly string LambdaName;
        public readonly string Name;
        public readonly string? Next;

        [OutputConstructor]
        private PipelineLambda(
            int batchSize,

            string lambdaName,

            string name,

            string? next)
        {
            BatchSize = batchSize;
            LambdaName = lambdaName;
            Name = name;
            Next = next;
        }
    }
}
