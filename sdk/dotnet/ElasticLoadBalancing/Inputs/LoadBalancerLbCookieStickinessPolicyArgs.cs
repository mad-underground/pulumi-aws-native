// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancing.Inputs
{

    public sealed class LoadBalancerLbCookieStickinessPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("cookieExpirationPeriod")]
        public Input<string>? CookieExpirationPeriod { get; set; }

        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        public LoadBalancerLbCookieStickinessPolicyArgs()
        {
        }
        public static new LoadBalancerLbCookieStickinessPolicyArgs Empty => new LoadBalancerLbCookieStickinessPolicyArgs();
    }
}
