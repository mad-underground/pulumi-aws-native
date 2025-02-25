// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Qldb
{
    public static class GetLedger
    {
        /// <summary>
        /// Resource Type definition for AWS::QLDB::Ledger
        /// </summary>
        public static Task<GetLedgerResult> InvokeAsync(GetLedgerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLedgerResult>("aws-native:qldb:getLedger", args ?? new GetLedgerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::QLDB::Ledger
        /// </summary>
        public static Output<GetLedgerResult> Invoke(GetLedgerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLedgerResult>("aws-native:qldb:getLedger", args ?? new GetLedgerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLedgerArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetLedgerArgs()
        {
        }
        public static new GetLedgerArgs Empty => new GetLedgerArgs();
    }

    public sealed class GetLedgerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetLedgerInvokeArgs()
        {
        }
        public static new GetLedgerInvokeArgs Empty => new GetLedgerInvokeArgs();
    }


    [OutputType]
    public sealed class GetLedgerResult
    {
        public readonly bool? DeletionProtection;
        public readonly string? Id;
        public readonly string? KmsKey;
        public readonly string? PermissionsMode;
        public readonly ImmutableArray<Pulumi.AwsNative.Outputs.Tag> Tags;

        [OutputConstructor]
        private GetLedgerResult(
            bool? deletionProtection,

            string? id,

            string? kmsKey,

            string? permissionsMode,

            ImmutableArray<Pulumi.AwsNative.Outputs.Tag> tags)
        {
            DeletionProtection = deletionProtection;
            Id = id;
            KmsKey = kmsKey;
            PermissionsMode = permissionsMode;
            Tags = tags;
        }
    }
}
