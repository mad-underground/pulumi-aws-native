// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetZonalAutoshiftConfigurationArgs, GetZonalAutoshiftConfigurationResult, GetZonalAutoshiftConfigurationOutputArgs } from "./getZonalAutoshiftConfiguration";
export const getZonalAutoshiftConfiguration: typeof import("./getZonalAutoshiftConfiguration").getZonalAutoshiftConfiguration = null as any;
export const getZonalAutoshiftConfigurationOutput: typeof import("./getZonalAutoshiftConfiguration").getZonalAutoshiftConfigurationOutput = null as any;
utilities.lazyLoad(exports, ["getZonalAutoshiftConfiguration","getZonalAutoshiftConfigurationOutput"], () => require("./getZonalAutoshiftConfiguration"));

export { ZonalAutoshiftConfigurationArgs } from "./zonalAutoshiftConfiguration";
export type ZonalAutoshiftConfiguration = import("./zonalAutoshiftConfiguration").ZonalAutoshiftConfiguration;
export const ZonalAutoshiftConfiguration: typeof import("./zonalAutoshiftConfiguration").ZonalAutoshiftConfiguration = null as any;
utilities.lazyLoad(exports, ["ZonalAutoshiftConfiguration"], () => require("./zonalAutoshiftConfiguration"));


// Export enums:
export * from "../types/enums/arczonalshift";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:arczonalshift:ZonalAutoshiftConfiguration":
                return new ZonalAutoshiftConfiguration(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "arczonalshift", _module)
