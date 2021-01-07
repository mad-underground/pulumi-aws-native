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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html
    /// </summary>
    public sealed class InstanceFleetConfigInstanceFleetProvisioningSpecificationsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications-ondemandspecification
        /// </summary>
        [Input("OnDemandSpecification")]
        public Input<Inputs.InstanceFleetConfigOnDemandProvisioningSpecificationArgs>? OnDemandSpecification { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications.html#cfn-elasticmapreduce-instancefleetconfig-instancefleetprovisioningspecifications-spotspecification
        /// </summary>
        [Input("SpotSpecification")]
        public Input<Inputs.InstanceFleetConfigSpotProvisioningSpecificationArgs>? SpotSpecification { get; set; }

        public InstanceFleetConfigInstanceFleetProvisioningSpecificationsArgs()
        {
        }
    }
}
