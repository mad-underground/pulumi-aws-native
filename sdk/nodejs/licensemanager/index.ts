// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetGrantArgs, GetGrantResult, GetGrantOutputArgs } from "./getGrant";
export const getGrant: typeof import("./getGrant").getGrant = null as any;
export const getGrantOutput: typeof import("./getGrant").getGrantOutput = null as any;

export { GetLicenseArgs, GetLicenseResult, GetLicenseOutputArgs } from "./getLicense";
export const getLicense: typeof import("./getLicense").getLicense = null as any;
export const getLicenseOutput: typeof import("./getLicense").getLicenseOutput = null as any;

export { GrantArgs } from "./grant";
export type Grant = import("./grant").Grant;
export const Grant: typeof import("./grant").Grant = null as any;

export { LicenseArgs } from "./license";
export type License = import("./license").License;
export const License: typeof import("./license").License = null as any;

utilities.lazyLoad(exports, ["getGrant","getGrantOutput"], () => require("./getGrant"));
utilities.lazyLoad(exports, ["getLicense","getLicenseOutput"], () => require("./getLicense"));
utilities.lazyLoad(exports, ["Grant"], () => require("./grant"));
utilities.lazyLoad(exports, ["License"], () => require("./license"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:licensemanager:Grant":
                return new Grant(name, <any>undefined, { urn })
            case "aws-native:licensemanager:License":
                return new License(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "licensemanager", _module)
