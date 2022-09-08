// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CampaignArgs } from "./campaign";
export type Campaign = import("./campaign").Campaign;
export const Campaign: typeof import("./campaign").Campaign = null as any;

export { GetCampaignArgs, GetCampaignResult, GetCampaignOutputArgs } from "./getCampaign";
export const getCampaign: typeof import("./getCampaign").getCampaign = null as any;
export const getCampaignOutput: typeof import("./getCampaign").getCampaignOutput = null as any;

utilities.lazyLoad(exports, ["Campaign"], () => require("./campaign"));
utilities.lazyLoad(exports, ["getCampaign","getCampaignOutput"], () => require("./getCampaign"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:connectcampaigns:Campaign":
                return new Campaign(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "connectcampaigns", _module)
