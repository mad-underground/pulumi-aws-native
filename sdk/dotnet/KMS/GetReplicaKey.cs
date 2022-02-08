// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KMS
{
    public static class GetReplicaKey
    {
        /// <summary>
        /// The AWS::KMS::ReplicaKey resource specifies a multi-region replica customer master key (CMK) in AWS Key Management Service (AWS KMS).
        /// </summary>
        public static Task<GetReplicaKeyResult> InvokeAsync(GetReplicaKeyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetReplicaKeyResult>("aws-native:kms:getReplicaKey", args ?? new GetReplicaKeyArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::KMS::ReplicaKey resource specifies a multi-region replica customer master key (CMK) in AWS Key Management Service (AWS KMS).
        /// </summary>
        public static Output<GetReplicaKeyResult> Invoke(GetReplicaKeyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetReplicaKeyResult>("aws-native:kms:getReplicaKey", args ?? new GetReplicaKeyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetReplicaKeyArgs : Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public string KeyId { get; set; } = null!;

        public GetReplicaKeyArgs()
        {
        }
    }

    public sealed class GetReplicaKeyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("keyId", required: true)]
        public Input<string> KeyId { get; set; } = null!;

        public GetReplicaKeyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetReplicaKeyResult
    {
        public readonly string? Arn;
        /// <summary>
        /// A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.
        /// </summary>
        public readonly bool? Enabled;
        public readonly string? KeyId;
        /// <summary>
        /// The key policy that authorizes use of the CMK. The key policy must observe the following rules.
        /// </summary>
        public readonly object? KeyPolicy;
        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReplicaKeyTag> Tags;

        [OutputConstructor]
        private GetReplicaKeyResult(
            string? arn,

            string? description,

            bool? enabled,

            string? keyId,

            object? keyPolicy,

            ImmutableArray<Outputs.ReplicaKeyTag> tags)
        {
            Arn = arn;
            Description = description;
            Enabled = enabled;
            KeyId = keyId;
            KeyPolicy = keyPolicy;
            Tags = tags;
        }
    }
}
