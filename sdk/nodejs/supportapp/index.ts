// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccountAliasArgs } from "./accountAlias";
export type AccountAlias = import("./accountAlias").AccountAlias;
export const AccountAlias: typeof import("./accountAlias").AccountAlias = null as any;

export { GetAccountAliasArgs, GetAccountAliasResult, GetAccountAliasOutputArgs } from "./getAccountAlias";
export const getAccountAlias: typeof import("./getAccountAlias").getAccountAlias = null as any;
export const getAccountAliasOutput: typeof import("./getAccountAlias").getAccountAliasOutput = null as any;

export { GetSlackChannelConfigurationArgs, GetSlackChannelConfigurationResult, GetSlackChannelConfigurationOutputArgs } from "./getSlackChannelConfiguration";
export const getSlackChannelConfiguration: typeof import("./getSlackChannelConfiguration").getSlackChannelConfiguration = null as any;
export const getSlackChannelConfigurationOutput: typeof import("./getSlackChannelConfiguration").getSlackChannelConfigurationOutput = null as any;

export { SlackChannelConfigurationArgs } from "./slackChannelConfiguration";
export type SlackChannelConfiguration = import("./slackChannelConfiguration").SlackChannelConfiguration;
export const SlackChannelConfiguration: typeof import("./slackChannelConfiguration").SlackChannelConfiguration = null as any;

utilities.lazyLoad(exports, ["AccountAlias"], () => require("./accountAlias"));
utilities.lazyLoad(exports, ["getAccountAlias","getAccountAliasOutput"], () => require("./getAccountAlias"));
utilities.lazyLoad(exports, ["getSlackChannelConfiguration","getSlackChannelConfigurationOutput"], () => require("./getSlackChannelConfiguration"));
utilities.lazyLoad(exports, ["SlackChannelConfiguration"], () => require("./slackChannelConfiguration"));

// Export enums:
export * from "../types/enums/supportapp";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:supportapp:AccountAlias":
                return new AccountAlias(name, <any>undefined, { urn })
            case "aws-native:supportapp:SlackChannelConfiguration":
                return new SlackChannelConfiguration(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "supportapp", _module)
