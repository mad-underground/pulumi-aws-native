// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cassandra.Outputs
{

    /// <summary>
    /// Represents the read and write settings used for AutoScaling.
    /// </summary>
    [OutputType]
    public sealed class TableAutoScalingSpecification
    {
        public readonly Outputs.TableAutoScalingSetting? ReadCapacityAutoScaling;
        public readonly Outputs.TableAutoScalingSetting? WriteCapacityAutoScaling;

        [OutputConstructor]
        private TableAutoScalingSpecification(
            Outputs.TableAutoScalingSetting? readCapacityAutoScaling,

            Outputs.TableAutoScalingSetting? writeCapacityAutoScaling)
        {
            ReadCapacityAutoScaling = readCapacityAutoScaling;
            WriteCapacityAutoScaling = writeCapacityAutoScaling;
        }
    }
}