// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ArcZonalShift.Inputs
{

    public sealed class ZonalAutoshiftConfigurationPracticeRunConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("blockedDates")]
        private InputList<string>? _blockedDates;
        public InputList<string> BlockedDates
        {
            get => _blockedDates ?? (_blockedDates = new InputList<string>());
            set => _blockedDates = value;
        }

        [Input("blockedWindows")]
        private InputList<string>? _blockedWindows;
        public InputList<string> BlockedWindows
        {
            get => _blockedWindows ?? (_blockedWindows = new InputList<string>());
            set => _blockedWindows = value;
        }

        [Input("blockingAlarms")]
        private InputList<Inputs.ZonalAutoshiftConfigurationControlConditionArgs>? _blockingAlarms;
        public InputList<Inputs.ZonalAutoshiftConfigurationControlConditionArgs> BlockingAlarms
        {
            get => _blockingAlarms ?? (_blockingAlarms = new InputList<Inputs.ZonalAutoshiftConfigurationControlConditionArgs>());
            set => _blockingAlarms = value;
        }

        [Input("outcomeAlarms", required: true)]
        private InputList<Inputs.ZonalAutoshiftConfigurationControlConditionArgs>? _outcomeAlarms;
        public InputList<Inputs.ZonalAutoshiftConfigurationControlConditionArgs> OutcomeAlarms
        {
            get => _outcomeAlarms ?? (_outcomeAlarms = new InputList<Inputs.ZonalAutoshiftConfigurationControlConditionArgs>());
            set => _outcomeAlarms = value;
        }

        public ZonalAutoshiftConfigurationPracticeRunConfigurationArgs()
        {
        }
        public static new ZonalAutoshiftConfigurationPracticeRunConfigurationArgs Empty => new ZonalAutoshiftConfigurationPracticeRunConfigurationArgs();
    }
}
