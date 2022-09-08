// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CertificateArgs } from "./certificate";
export type Certificate = import("./certificate").Certificate;
export const Certificate: typeof import("./certificate").Certificate = null as any;

export { CertificateAuthorityArgs } from "./certificateAuthority";
export type CertificateAuthority = import("./certificateAuthority").CertificateAuthority;
export const CertificateAuthority: typeof import("./certificateAuthority").CertificateAuthority = null as any;

export { CertificateAuthorityActivationArgs } from "./certificateAuthorityActivation";
export type CertificateAuthorityActivation = import("./certificateAuthorityActivation").CertificateAuthorityActivation;
export const CertificateAuthorityActivation: typeof import("./certificateAuthorityActivation").CertificateAuthorityActivation = null as any;

export { GetCertificateArgs, GetCertificateResult, GetCertificateOutputArgs } from "./getCertificate";
export const getCertificate: typeof import("./getCertificate").getCertificate = null as any;
export const getCertificateOutput: typeof import("./getCertificate").getCertificateOutput = null as any;

export { GetCertificateAuthorityArgs, GetCertificateAuthorityResult, GetCertificateAuthorityOutputArgs } from "./getCertificateAuthority";
export const getCertificateAuthority: typeof import("./getCertificateAuthority").getCertificateAuthority = null as any;
export const getCertificateAuthorityOutput: typeof import("./getCertificateAuthority").getCertificateAuthorityOutput = null as any;

export { GetCertificateAuthorityActivationArgs, GetCertificateAuthorityActivationResult, GetCertificateAuthorityActivationOutputArgs } from "./getCertificateAuthorityActivation";
export const getCertificateAuthorityActivation: typeof import("./getCertificateAuthorityActivation").getCertificateAuthorityActivation = null as any;
export const getCertificateAuthorityActivationOutput: typeof import("./getCertificateAuthorityActivation").getCertificateAuthorityActivationOutput = null as any;

export { PermissionArgs } from "./permission";
export type Permission = import("./permission").Permission;
export const Permission: typeof import("./permission").Permission = null as any;

utilities.lazyLoad(exports, ["Certificate"], () => require("./certificate"));
utilities.lazyLoad(exports, ["CertificateAuthority"], () => require("./certificateAuthority"));
utilities.lazyLoad(exports, ["CertificateAuthorityActivation"], () => require("./certificateAuthorityActivation"));
utilities.lazyLoad(exports, ["getCertificate","getCertificateOutput"], () => require("./getCertificate"));
utilities.lazyLoad(exports, ["getCertificateAuthority","getCertificateAuthorityOutput"], () => require("./getCertificateAuthority"));
utilities.lazyLoad(exports, ["getCertificateAuthorityActivation","getCertificateAuthorityActivationOutput"], () => require("./getCertificateAuthorityActivation"));
utilities.lazyLoad(exports, ["Permission"], () => require("./permission"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:acmpca:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "aws-native:acmpca:CertificateAuthority":
                return new CertificateAuthority(name, <any>undefined, { urn })
            case "aws-native:acmpca:CertificateAuthorityActivation":
                return new CertificateAuthorityActivation(name, <any>undefined, { urn })
            case "aws-native:acmpca:Permission":
                return new Permission(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "acmpca", _module)
