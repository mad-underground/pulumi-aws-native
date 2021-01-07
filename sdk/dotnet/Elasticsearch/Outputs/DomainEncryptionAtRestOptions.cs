// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.Elasticsearch.Outputs
{

    [OutputType]
    public sealed class DomainEncryptionAtRestOptions
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-encryptionatrestoptions.html#cfn-elasticsearch-domain-encryptionatrestoptions-enabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-encryptionatrestoptions.html#cfn-elasticsearch-domain-encryptionatrestoptions-kmskeyid
        /// </summary>
        public readonly string? KmsKeyId;

        [OutputConstructor]
        private DomainEncryptionAtRestOptions(
            bool? Enabled,

            string? KmsKeyId)
        {
            this.Enabled = Enabled;
            this.KmsKeyId = KmsKeyId;
        }
    }
}