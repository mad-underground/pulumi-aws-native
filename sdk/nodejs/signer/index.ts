// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetSigningProfileArgs, GetSigningProfileResult, GetSigningProfileOutputArgs } from "./getSigningProfile";
export const getSigningProfile: typeof import("./getSigningProfile").getSigningProfile = null as any;
export const getSigningProfileOutput: typeof import("./getSigningProfile").getSigningProfileOutput = null as any;
utilities.lazyLoad(exports, ["getSigningProfile","getSigningProfileOutput"], () => require("./getSigningProfile"));

export { ProfilePermissionArgs } from "./profilePermission";
export type ProfilePermission = import("./profilePermission").ProfilePermission;
export const ProfilePermission: typeof import("./profilePermission").ProfilePermission = null as any;
utilities.lazyLoad(exports, ["ProfilePermission"], () => require("./profilePermission"));

export { SigningProfileArgs } from "./signingProfile";
export type SigningProfile = import("./signingProfile").SigningProfile;
export const SigningProfile: typeof import("./signingProfile").SigningProfile = null as any;
utilities.lazyLoad(exports, ["SigningProfile"], () => require("./signingProfile"));


// Export enums:
export * from "../types/enums/signer";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:signer:ProfilePermission":
                return new ProfilePermission(name, <any>undefined, { urn })
            case "aws-native:signer:SigningProfile":
                return new SigningProfile(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "signer", _module)
