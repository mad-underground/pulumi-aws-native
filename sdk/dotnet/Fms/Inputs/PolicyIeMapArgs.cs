// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Fms.Inputs
{

    /// <summary>
    /// An FMS includeMap or excludeMap.
    /// </summary>
    public sealed class PolicyIeMapArgs : global::Pulumi.ResourceArgs
    {
        [Input("account")]
        private InputList<string>? _account;
        public InputList<string> Account
        {
            get => _account ?? (_account = new InputList<string>());
            set => _account = value;
        }

        [Input("orgunit")]
        private InputList<string>? _orgunit;
        public InputList<string> Orgunit
        {
            get => _orgunit ?? (_orgunit = new InputList<string>());
            set => _orgunit = value;
        }

        public PolicyIeMapArgs()
        {
        }
        public static new PolicyIeMapArgs Empty => new PolicyIeMapArgs();
    }
}
