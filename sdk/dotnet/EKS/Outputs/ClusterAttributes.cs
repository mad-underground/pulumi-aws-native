// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EKS.Outputs
{

    [OutputType]
    public sealed class ClusterAttributes
    {
        public readonly string Arn;
        public readonly string CertificateAuthorityData;
        public readonly string ClusterSecurityGroupId;
        public readonly string EncryptionConfigKeyArn;
        public readonly string Endpoint;

        [OutputConstructor]
        private ClusterAttributes(
            string Arn,

            string CertificateAuthorityData,

            string ClusterSecurityGroupId,

            string EncryptionConfigKeyArn,

            string Endpoint)
        {
            this.Arn = Arn;
            this.CertificateAuthorityData = CertificateAuthorityData;
            this.ClusterSecurityGroupId = ClusterSecurityGroupId;
            this.EncryptionConfigKeyArn = EncryptionConfigKeyArn;
            this.Endpoint = Endpoint;
        }
    }
}