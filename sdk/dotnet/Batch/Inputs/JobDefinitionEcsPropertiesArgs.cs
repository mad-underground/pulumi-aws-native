// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionEcsPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("taskProperties", required: true)]
        private InputList<Inputs.JobDefinitionEcsTaskPropertiesArgs>? _taskProperties;
        public InputList<Inputs.JobDefinitionEcsTaskPropertiesArgs> TaskProperties
        {
            get => _taskProperties ?? (_taskProperties = new InputList<Inputs.JobDefinitionEcsTaskPropertiesArgs>());
            set => _taskProperties = value;
        }

        public JobDefinitionEcsPropertiesArgs()
        {
        }
        public static new JobDefinitionEcsPropertiesArgs Empty => new JobDefinitionEcsPropertiesArgs();
    }
}
