// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation
{
    public static class Cidr
    {
        public static Task<CidrResult> InvokeAsync(CidrArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<CidrResult>("cloudformation:index:cidr", args ?? new CidrArgs(), options.WithVersion());
    }


    public sealed class CidrArgs : Pulumi.InvokeArgs
    {
        [Input("cidrBits", required: true)]
        public int CidrBits { get; set; }

        [Input("count", required: true)]
        public int Count { get; set; }

        [Input("ipBlock", required: true)]
        public string IpBlock { get; set; } = null!;

        public CidrArgs()
        {
        }
    }


    [OutputType]
    public sealed class CidrResult
    {
        public readonly ImmutableArray<string> Subnets;

        [OutputConstructor]
        private CidrResult(ImmutableArray<string> subnets)
        {
            Subnets = subnets;
        }
    }
}