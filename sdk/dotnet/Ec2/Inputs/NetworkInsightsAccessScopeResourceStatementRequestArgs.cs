// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Inputs
{

    public sealed class NetworkInsightsAccessScopeResourceStatementRequestArgs : global::Pulumi.ResourceArgs
    {
        [Input("resourceTypes")]
        private InputList<string>? _resourceTypes;
        public InputList<string> ResourceTypes
        {
            get => _resourceTypes ?? (_resourceTypes = new InputList<string>());
            set => _resourceTypes = value;
        }

        [Input("resources")]
        private InputList<string>? _resources;
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        public NetworkInsightsAccessScopeResourceStatementRequestArgs()
        {
        }
        public static new NetworkInsightsAccessScopeResourceStatementRequestArgs Empty => new NetworkInsightsAccessScopeResourceStatementRequestArgs();
    }
}
