// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EventDataStoreArgs } from "./eventDataStore";
export type EventDataStore = import("./eventDataStore").EventDataStore;
export const EventDataStore: typeof import("./eventDataStore").EventDataStore = null as any;

export { GetEventDataStoreArgs, GetEventDataStoreResult, GetEventDataStoreOutputArgs } from "./getEventDataStore";
export const getEventDataStore: typeof import("./getEventDataStore").getEventDataStore = null as any;
export const getEventDataStoreOutput: typeof import("./getEventDataStore").getEventDataStoreOutput = null as any;

export { GetTrailArgs, GetTrailResult, GetTrailOutputArgs } from "./getTrail";
export const getTrail: typeof import("./getTrail").getTrail = null as any;
export const getTrailOutput: typeof import("./getTrail").getTrailOutput = null as any;

export { TrailArgs } from "./trail";
export type Trail = import("./trail").Trail;
export const Trail: typeof import("./trail").Trail = null as any;

utilities.lazyLoad(exports, ["EventDataStore"], () => require("./eventDataStore"));
utilities.lazyLoad(exports, ["getEventDataStore","getEventDataStoreOutput"], () => require("./getEventDataStore"));
utilities.lazyLoad(exports, ["getTrail","getTrailOutput"], () => require("./getTrail"));
utilities.lazyLoad(exports, ["Trail"], () => require("./trail"));

// Export enums:
export * from "../types/enums/cloudtrail";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:cloudtrail:EventDataStore":
                return new EventDataStore(name, <any>undefined, { urn })
            case "aws-native:cloudtrail:Trail":
                return new Trail(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "cloudtrail", _module)
