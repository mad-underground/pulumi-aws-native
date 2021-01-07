// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation
{
    public static class GetPartition
    {
        public static Task<GetPartitionResult> InvokeAsync(InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPartitionResult>("cloudformation:index:getPartition", InvokeArgs.Empty, options.WithVersion());
    }


    [OutputType]
    public sealed class GetPartitionResult
    {
        public readonly string Partition;

        [OutputConstructor]
        private GetPartitionResult(string partition)
        {
            Partition = partition;
        }
    }
}