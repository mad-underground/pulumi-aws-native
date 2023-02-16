// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetNamespaceArgs, GetNamespaceResult, GetNamespaceOutputArgs } from "./getNamespace";
export const getNamespace: typeof import("./getNamespace").getNamespace = null as any;
export const getNamespaceOutput: typeof import("./getNamespace").getNamespaceOutput = null as any;
utilities.lazyLoad(exports, ["getNamespace","getNamespaceOutput"], () => require("./getNamespace"));

export { GetWorkgroupArgs, GetWorkgroupResult, GetWorkgroupOutputArgs } from "./getWorkgroup";
export const getWorkgroup: typeof import("./getWorkgroup").getWorkgroup = null as any;
export const getWorkgroupOutput: typeof import("./getWorkgroup").getWorkgroupOutput = null as any;
utilities.lazyLoad(exports, ["getWorkgroup","getWorkgroupOutput"], () => require("./getWorkgroup"));

export { NamespaceArgs } from "./namespace";
export type Namespace = import("./namespace").Namespace;
export const Namespace: typeof import("./namespace").Namespace = null as any;
utilities.lazyLoad(exports, ["Namespace"], () => require("./namespace"));

export { WorkgroupArgs } from "./workgroup";
export type Workgroup = import("./workgroup").Workgroup;
export const Workgroup: typeof import("./workgroup").Workgroup = null as any;
utilities.lazyLoad(exports, ["Workgroup"], () => require("./workgroup"));


// Export enums:
export * from "../types/enums/redshiftserverless";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:redshiftserverless:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            case "aws-native:redshiftserverless:Workgroup":
                return new Workgroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "redshiftserverless", _module)
