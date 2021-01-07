// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.EMR.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html
    /// </summary>
    public sealed class ClusterInstanceFleetProvisioningSpecificationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-cluster-instancefleetprovisioningspecifications-ondemandspecification
        /// </summary>
        [Input("OnDemandSpecification")]
        public Input<Inputs.ClusterOnDemandProvisioningSpecificationArgs>? OnDemandSpecification { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-cluster-instancefleetprovisioningspecifications-spotspecification
        /// </summary>
        [Input("SpotSpecification")]
        public Input<Inputs.ClusterSpotProvisioningSpecificationArgs>? SpotSpecification { get; set; }

        public ClusterInstanceFleetProvisioningSpecificationsArgs()
        {
        }
    }
}
