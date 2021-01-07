// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.ServiceCatalog.Outputs
{

    [OutputType]
    public sealed class CloudFormationProvisionedProductAttributes
    {
        public readonly string CloudformationStackArn;
        public readonly ImmutableDictionary<string, string> Outputs;
        public readonly string ProvisionedProductId;
        public readonly string RecordId;

        [OutputConstructor]
        private CloudFormationProvisionedProductAttributes(
            string CloudformationStackArn,

            ImmutableDictionary<string, string> Outputs,

            string ProvisionedProductId,

            string RecordId)
        {
            this.CloudformationStackArn = CloudformationStackArn;
            this.Outputs = Outputs;
            this.ProvisionedProductId = ProvisionedProductId;
            this.RecordId = RecordId;
        }
    }
}