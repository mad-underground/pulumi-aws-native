// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EnabledBaselineArgs } from "./enabledBaseline";
export type EnabledBaseline = import("./enabledBaseline").EnabledBaseline;
export const EnabledBaseline: typeof import("./enabledBaseline").EnabledBaseline = null as any;
utilities.lazyLoad(exports, ["EnabledBaseline"], () => require("./enabledBaseline"));

export { EnabledControlArgs } from "./enabledControl";
export type EnabledControl = import("./enabledControl").EnabledControl;
export const EnabledControl: typeof import("./enabledControl").EnabledControl = null as any;
utilities.lazyLoad(exports, ["EnabledControl"], () => require("./enabledControl"));

export { GetEnabledBaselineArgs, GetEnabledBaselineResult, GetEnabledBaselineOutputArgs } from "./getEnabledBaseline";
export const getEnabledBaseline: typeof import("./getEnabledBaseline").getEnabledBaseline = null as any;
export const getEnabledBaselineOutput: typeof import("./getEnabledBaseline").getEnabledBaselineOutput = null as any;
utilities.lazyLoad(exports, ["getEnabledBaseline","getEnabledBaselineOutput"], () => require("./getEnabledBaseline"));

export { GetEnabledControlArgs, GetEnabledControlResult, GetEnabledControlOutputArgs } from "./getEnabledControl";
export const getEnabledControl: typeof import("./getEnabledControl").getEnabledControl = null as any;
export const getEnabledControlOutput: typeof import("./getEnabledControl").getEnabledControlOutput = null as any;
utilities.lazyLoad(exports, ["getEnabledControl","getEnabledControlOutput"], () => require("./getEnabledControl"));

export { GetLandingZoneArgs, GetLandingZoneResult, GetLandingZoneOutputArgs } from "./getLandingZone";
export const getLandingZone: typeof import("./getLandingZone").getLandingZone = null as any;
export const getLandingZoneOutput: typeof import("./getLandingZone").getLandingZoneOutput = null as any;
utilities.lazyLoad(exports, ["getLandingZone","getLandingZoneOutput"], () => require("./getLandingZone"));

export { LandingZoneArgs } from "./landingZone";
export type LandingZone = import("./landingZone").LandingZone;
export const LandingZone: typeof import("./landingZone").LandingZone = null as any;
utilities.lazyLoad(exports, ["LandingZone"], () => require("./landingZone"));


// Export enums:
export * from "../types/enums/controltower";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:controltower:EnabledBaseline":
                return new EnabledBaseline(name, <any>undefined, { urn })
            case "aws-native:controltower:EnabledControl":
                return new EnabledControl(name, <any>undefined, { urn })
            case "aws-native:controltower:LandingZone":
                return new LandingZone(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "controltower", _module)
