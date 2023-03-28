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
    public sealed class TemplateDefaultNewSheetConfiguration
    {
        public readonly Outputs.TemplateDefaultInteractiveLayoutConfiguration? InteractiveLayoutConfiguration;
        public readonly Outputs.TemplateDefaultPaginatedLayoutConfiguration? PaginatedLayoutConfiguration;
        public readonly Pulumi.AwsNative.QuickSight.TemplateSheetContentType? SheetContentType;

        [OutputConstructor]
        private TemplateDefaultNewSheetConfiguration(
            Outputs.TemplateDefaultInteractiveLayoutConfiguration? interactiveLayoutConfiguration,

            Outputs.TemplateDefaultPaginatedLayoutConfiguration? paginatedLayoutConfiguration,

            Pulumi.AwsNative.QuickSight.TemplateSheetContentType? sheetContentType)
        {
            InteractiveLayoutConfiguration = interactiveLayoutConfiguration;
            PaginatedLayoutConfiguration = paginatedLayoutConfiguration;
            SheetContentType = sheetContentType;
        }
    }
}