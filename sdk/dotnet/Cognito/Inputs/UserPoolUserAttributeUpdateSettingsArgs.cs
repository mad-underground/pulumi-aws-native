// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolUserAttributeUpdateSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributesRequireVerificationBeforeUpdate", required: true)]
        private InputList<string>? _attributesRequireVerificationBeforeUpdate;
        public InputList<string> AttributesRequireVerificationBeforeUpdate
        {
            get => _attributesRequireVerificationBeforeUpdate ?? (_attributesRequireVerificationBeforeUpdate = new InputList<string>());
            set => _attributesRequireVerificationBeforeUpdate = value;
        }

        public UserPoolUserAttributeUpdateSettingsArgs()
        {
        }
        public static new UserPoolUserAttributeUpdateSettingsArgs Empty => new UserPoolUserAttributeUpdateSettingsArgs();
    }
}
