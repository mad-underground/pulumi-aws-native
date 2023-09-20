// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.BillingConductor.Inputs
{

    public sealed class CustomLineItemLineItemFilterArgs : global::Pulumi.ResourceArgs
    {
        [Input("attribute", required: true)]
        public Input<Pulumi.AwsNative.BillingConductor.CustomLineItemLineItemFilterAttribute> Attribute { get; set; } = null!;

        [Input("matchOption", required: true)]
        public Input<Pulumi.AwsNative.BillingConductor.CustomLineItemLineItemFilterMatchOption> MatchOption { get; set; } = null!;

        [Input("values", required: true)]
        private InputList<Pulumi.AwsNative.BillingConductor.CustomLineItemLineItemFilterValue>? _values;
        public InputList<Pulumi.AwsNative.BillingConductor.CustomLineItemLineItemFilterValue> Values
        {
            get => _values ?? (_values = new InputList<Pulumi.AwsNative.BillingConductor.CustomLineItemLineItemFilterValue>());
            set => _values = value;
        }

        public CustomLineItemLineItemFilterArgs()
        {
        }
        public static new CustomLineItemLineItemFilterArgs Empty => new CustomLineItemLineItemFilterArgs();
    }
}