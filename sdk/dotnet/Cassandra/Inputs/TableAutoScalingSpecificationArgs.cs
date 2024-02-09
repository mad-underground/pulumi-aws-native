// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra.Inputs
{

    /// <summary>
    /// Represents the read and write settings used for AutoScaling.
    /// </summary>
    public sealed class TableAutoScalingSpecificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("readCapacityAutoScaling")]
        public Input<Inputs.TableAutoScalingSettingArgs>? ReadCapacityAutoScaling { get; set; }

        [Input("writeCapacityAutoScaling")]
        public Input<Inputs.TableAutoScalingSettingArgs>? WriteCapacityAutoScaling { get; set; }

        public TableAutoScalingSpecificationArgs()
        {
        }
        public static new TableAutoScalingSpecificationArgs Empty => new TableAutoScalingSpecificationArgs();
    }
}
