// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cloudformation.AppFlow.Outputs
{

    [OutputType]
    public sealed class FlowZendeskSourceProperties
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appflow-flow-zendesksourceproperties.html#cfn-appflow-flow-zendesksourceproperties-object
        /// </summary>
        public readonly string Object;

        [OutputConstructor]
        private FlowZendeskSourceProperties(string Object)
        {
            this.Object = Object;
        }
    }
}
