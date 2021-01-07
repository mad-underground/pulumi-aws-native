// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./association";
export * from "./document";
export * from "./maintenanceWindow";
export * from "./maintenanceWindowTarget";
export * from "./maintenanceWindowTask";
export * from "./parameter";
export * from "./patchBaseline";
export * from "./resourceDataSync";

// Import resources to register:
import { Association } from "./association";
import { Document } from "./document";
import { MaintenanceWindow } from "./maintenanceWindow";
import { MaintenanceWindowTarget } from "./maintenanceWindowTarget";
import { MaintenanceWindowTask } from "./maintenanceWindowTask";
import { Parameter } from "./parameter";
import { PatchBaseline } from "./patchBaseline";
import { ResourceDataSync } from "./resourceDataSync";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:SSM:Association":
                return new Association(name, <any>undefined, { urn })
            case "cloudformation:SSM:Document":
                return new Document(name, <any>undefined, { urn })
            case "cloudformation:SSM:MaintenanceWindow":
                return new MaintenanceWindow(name, <any>undefined, { urn })
            case "cloudformation:SSM:MaintenanceWindowTarget":
                return new MaintenanceWindowTarget(name, <any>undefined, { urn })
            case "cloudformation:SSM:MaintenanceWindowTask":
                return new MaintenanceWindowTask(name, <any>undefined, { urn })
            case "cloudformation:SSM:Parameter":
                return new Parameter(name, <any>undefined, { urn })
            case "cloudformation:SSM:PatchBaseline":
                return new PatchBaseline(name, <any>undefined, { urn })
            case "cloudformation:SSM:ResourceDataSync":
                return new ResourceDataSync(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "SSM", _module)