// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Outputs
{

    [OutputType]
    public sealed class TableOpenTableFormatInput
    {
        public readonly Outputs.TableIcebergInput? IcebergInput;

        [OutputConstructor]
        private TableOpenTableFormatInput(Outputs.TableIcebergInput? icebergInput)
        {
            IcebergInput = icebergInput;
        }
    }
}
