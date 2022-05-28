// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito
{
    public static class GetUserPool
    {
        /// <summary>
        /// Resource Type definition for AWS::Cognito::UserPool
        /// </summary>
        public static Task<GetUserPoolResult> InvokeAsync(GetUserPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserPoolResult>("aws-native:cognito:getUserPool", args ?? new GetUserPoolArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::Cognito::UserPool
        /// </summary>
        public static Output<GetUserPoolResult> Invoke(GetUserPoolInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetUserPoolResult>("aws-native:cognito:getUserPool", args ?? new GetUserPoolInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserPoolArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetUserPoolArgs()
        {
        }
    }

    public sealed class GetUserPoolInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetUserPoolInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserPoolResult
    {
        public readonly Outputs.UserPoolAccountRecoverySetting? AccountRecoverySetting;
        public readonly Outputs.UserPoolAdminCreateUserConfig? AdminCreateUserConfig;
        public readonly ImmutableArray<string> AliasAttributes;
        public readonly string? Arn;
        public readonly ImmutableArray<string> AutoVerifiedAttributes;
        public readonly Outputs.UserPoolDeviceConfiguration? DeviceConfiguration;
        public readonly Outputs.UserPoolEmailConfiguration? EmailConfiguration;
        public readonly string? EmailVerificationMessage;
        public readonly string? EmailVerificationSubject;
        public readonly ImmutableArray<string> EnabledMfas;
        public readonly string? Id;
        public readonly Outputs.UserPoolLambdaConfig? LambdaConfig;
        public readonly string? MfaConfiguration;
        public readonly Outputs.UserPoolPolicies? Policies;
        public readonly string? ProviderName;
        public readonly string? ProviderURL;
        public readonly ImmutableArray<Outputs.UserPoolSchemaAttribute> Schema;
        public readonly string? SmsAuthenticationMessage;
        public readonly Outputs.UserPoolSmsConfiguration? SmsConfiguration;
        public readonly string? SmsVerificationMessage;
        public readonly Outputs.UserPoolUserAttributeUpdateSettings? UserAttributeUpdateSettings;
        public readonly Outputs.UserPoolAddOns? UserPoolAddOns;
        public readonly string? UserPoolName;
        public readonly object? UserPoolTags;
        public readonly ImmutableArray<string> UsernameAttributes;
        public readonly Outputs.UserPoolUsernameConfiguration? UsernameConfiguration;
        public readonly Outputs.UserPoolVerificationMessageTemplate? VerificationMessageTemplate;

        [OutputConstructor]
        private GetUserPoolResult(
            Outputs.UserPoolAccountRecoverySetting? accountRecoverySetting,

            Outputs.UserPoolAdminCreateUserConfig? adminCreateUserConfig,

            ImmutableArray<string> aliasAttributes,

            string? arn,

            ImmutableArray<string> autoVerifiedAttributes,

            Outputs.UserPoolDeviceConfiguration? deviceConfiguration,

            Outputs.UserPoolEmailConfiguration? emailConfiguration,

            string? emailVerificationMessage,

            string? emailVerificationSubject,

            ImmutableArray<string> enabledMfas,

            string? id,

            Outputs.UserPoolLambdaConfig? lambdaConfig,

            string? mfaConfiguration,

            Outputs.UserPoolPolicies? policies,

            string? providerName,

            string? providerURL,

            ImmutableArray<Outputs.UserPoolSchemaAttribute> schema,

            string? smsAuthenticationMessage,

            Outputs.UserPoolSmsConfiguration? smsConfiguration,

            string? smsVerificationMessage,

            Outputs.UserPoolUserAttributeUpdateSettings? userAttributeUpdateSettings,

            Outputs.UserPoolAddOns? userPoolAddOns,

            string? userPoolName,

            object? userPoolTags,

            ImmutableArray<string> usernameAttributes,

            Outputs.UserPoolUsernameConfiguration? usernameConfiguration,

            Outputs.UserPoolVerificationMessageTemplate? verificationMessageTemplate)
        {
            AccountRecoverySetting = accountRecoverySetting;
            AdminCreateUserConfig = adminCreateUserConfig;
            AliasAttributes = aliasAttributes;
            Arn = arn;
            AutoVerifiedAttributes = autoVerifiedAttributes;
            DeviceConfiguration = deviceConfiguration;
            EmailConfiguration = emailConfiguration;
            EmailVerificationMessage = emailVerificationMessage;
            EmailVerificationSubject = emailVerificationSubject;
            EnabledMfas = enabledMfas;
            Id = id;
            LambdaConfig = lambdaConfig;
            MfaConfiguration = mfaConfiguration;
            Policies = policies;
            ProviderName = providerName;
            ProviderURL = providerURL;
            Schema = schema;
            SmsAuthenticationMessage = smsAuthenticationMessage;
            SmsConfiguration = smsConfiguration;
            SmsVerificationMessage = smsVerificationMessage;
            UserAttributeUpdateSettings = userAttributeUpdateSettings;
            UserPoolAddOns = userPoolAddOns;
            UserPoolName = userPoolName;
            UserPoolTags = userPoolTags;
            UsernameAttributes = usernameAttributes;
            UsernameConfiguration = usernameConfiguration;
            VerificationMessageTemplate = verificationMessageTemplate;
        }
    }
}
