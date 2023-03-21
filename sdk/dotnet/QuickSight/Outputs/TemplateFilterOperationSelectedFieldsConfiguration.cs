// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class TemplateFilterOperationSelectedFieldsConfiguration
    {
        public readonly Pulumi.AwsNative.QuickSight.TemplateSelectedFieldOptions? SelectedFieldOptions;
        public readonly ImmutableArray<string> SelectedFields;

        [OutputConstructor]
        private TemplateFilterOperationSelectedFieldsConfiguration(
            Pulumi.AwsNative.QuickSight.TemplateSelectedFieldOptions? selectedFieldOptions,

            ImmutableArray<string> selectedFields)
        {
            SelectedFieldOptions = selectedFieldOptions;
            SelectedFields = selectedFields;
        }
    }
}
