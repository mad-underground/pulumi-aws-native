// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SES
{
    public static class GetVdmAttributes
    {
        /// <summary>
        /// Resource Type definition for AWS::SES::VdmAttributes
        /// </summary>
        public static Task<GetVdmAttributesResult> InvokeAsync(GetVdmAttributesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVdmAttributesResult>("aws-native:ses:getVdmAttributes", args ?? new GetVdmAttributesArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::SES::VdmAttributes
        /// </summary>
        public static Output<GetVdmAttributesResult> Invoke(GetVdmAttributesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVdmAttributesResult>("aws-native:ses:getVdmAttributes", args ?? new GetVdmAttributesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVdmAttributesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for this resource
        /// </summary>
        [Input("vdmAttributesResourceId", required: true)]
        public string VdmAttributesResourceId { get; set; } = null!;

        public GetVdmAttributesArgs()
        {
        }
        public static new GetVdmAttributesArgs Empty => new GetVdmAttributesArgs();
    }

    public sealed class GetVdmAttributesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Unique identifier for this resource
        /// </summary>
        [Input("vdmAttributesResourceId", required: true)]
        public Input<string> VdmAttributesResourceId { get; set; } = null!;

        public GetVdmAttributesInvokeArgs()
        {
        }
        public static new GetVdmAttributesInvokeArgs Empty => new GetVdmAttributesInvokeArgs();
    }


    [OutputType]
    public sealed class GetVdmAttributesResult
    {
        public readonly Outputs.VdmAttributesDashboardAttributes? DashboardAttributes;
        public readonly Outputs.VdmAttributesGuardianAttributes? GuardianAttributes;
        /// <summary>
        /// Unique identifier for this resource
        /// </summary>
        public readonly string? VdmAttributesResourceId;

        [OutputConstructor]
        private GetVdmAttributesResult(
            Outputs.VdmAttributesDashboardAttributes? dashboardAttributes,

            Outputs.VdmAttributesGuardianAttributes? guardianAttributes,

            string? vdmAttributesResourceId)
        {
            DashboardAttributes = dashboardAttributes;
            GuardianAttributes = guardianAttributes;
            VdmAttributesResourceId = vdmAttributesResourceId;
        }
    }
}