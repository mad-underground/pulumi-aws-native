// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MemoryDB
{
    public static class GetACL
    {
        /// <summary>
        /// Resource Type definition for AWS::MemoryDB::ACL
        /// </summary>
        public static Task<GetACLResult> InvokeAsync(GetACLArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetACLResult>("aws-native:memorydb:getACL", args ?? new GetACLArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::MemoryDB::ACL
        /// </summary>
        public static Output<GetACLResult> Invoke(GetACLInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetACLResult>("aws-native:memorydb:getACL", args ?? new GetACLInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetACLArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the acl.
        /// </summary>
        [Input("aCLName", required: true)]
        public string ACLName { get; set; } = null!;

        public GetACLArgs()
        {
        }
    }

    public sealed class GetACLInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the acl.
        /// </summary>
        [Input("aCLName", required: true)]
        public Input<string> ACLName { get; set; } = null!;

        public GetACLInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetACLResult
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) of the acl.
        /// </summary>
        public readonly string? Arn;
        /// <summary>
        /// Indicates acl status. Can be "creating", "active", "modifying", "deleting".
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// An array of key-value pairs to apply to this cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.ACLTag> Tags;
        /// <summary>
        /// List of users associated to this acl.
        /// </summary>
        public readonly ImmutableArray<string> UserNames;

        [OutputConstructor]
        private GetACLResult(
            string? arn,

            string? status,

            ImmutableArray<Outputs.ACLTag> tags,

            ImmutableArray<string> userNames)
        {
            Arn = arn;
            Status = status;
            Tags = tags;
            UserNames = userNames;
        }
    }
}
