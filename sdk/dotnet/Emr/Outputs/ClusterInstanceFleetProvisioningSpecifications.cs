// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Emr.Outputs
{

    [OutputType]
    public sealed class ClusterInstanceFleetProvisioningSpecifications
    {
        public readonly Outputs.ClusterOnDemandProvisioningSpecification? OnDemandSpecification;
        public readonly Outputs.ClusterSpotProvisioningSpecification? SpotSpecification;

        [OutputConstructor]
        private ClusterInstanceFleetProvisioningSpecifications(
            Outputs.ClusterOnDemandProvisioningSpecification? onDemandSpecification,

            Outputs.ClusterSpotProvisioningSpecification? spotSpecification)
        {
            OnDemandSpecification = onDemandSpecification;
            SpotSpecification = spotSpecification;
        }
    }
}
