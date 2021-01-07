// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./instanceFleetConfig";
export * from "./instanceGroupConfig";
export * from "./securityConfiguration";
export * from "./step";

// Import resources to register:
import { Cluster } from "./cluster";
import { InstanceFleetConfig } from "./instanceFleetConfig";
import { InstanceGroupConfig } from "./instanceGroupConfig";
import { SecurityConfiguration } from "./securityConfiguration";
import { Step } from "./step";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:EMR:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "cloudformation:EMR:InstanceFleetConfig":
                return new InstanceFleetConfig(name, <any>undefined, { urn })
            case "cloudformation:EMR:InstanceGroupConfig":
                return new InstanceGroupConfig(name, <any>undefined, { urn })
            case "cloudformation:EMR:SecurityConfiguration":
                return new SecurityConfiguration(name, <any>undefined, { urn })
            case "cloudformation:EMR:Step":
                return new Step(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "EMR", _module)