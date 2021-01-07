// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./config";
export * from "./dataflowEndpointGroup";
export * from "./missionProfile";

// Import resources to register:
import { Config } from "./config";
import { DataflowEndpointGroup } from "./dataflowEndpointGroup";
import { MissionProfile } from "./missionProfile";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloudformation:GroundStation:Config":
                return new Config(name, <any>undefined, { urn })
            case "cloudformation:GroundStation:DataflowEndpointGroup":
                return new DataflowEndpointGroup(name, <any>undefined, { urn })
            case "cloudformation:GroundStation:MissionProfile":
                return new MissionProfile(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloudformation", "GroundStation", _module)
