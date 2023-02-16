// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ConnectorArgs } from "./connector";
export type Connector = import("./connector").Connector;
export const Connector: typeof import("./connector").Connector = null as any;
utilities.lazyLoad(exports, ["Connector"], () => require("./connector"));

export { ConnectorProfileArgs } from "./connectorProfile";
export type ConnectorProfile = import("./connectorProfile").ConnectorProfile;
export const ConnectorProfile: typeof import("./connectorProfile").ConnectorProfile = null as any;
utilities.lazyLoad(exports, ["ConnectorProfile"], () => require("./connectorProfile"));

export { FlowArgs } from "./flow";
export type Flow = import("./flow").Flow;
export const Flow: typeof import("./flow").Flow = null as any;
utilities.lazyLoad(exports, ["Flow"], () => require("./flow"));

export { GetConnectorArgs, GetConnectorResult, GetConnectorOutputArgs } from "./getConnector";
export const getConnector: typeof import("./getConnector").getConnector = null as any;
export const getConnectorOutput: typeof import("./getConnector").getConnectorOutput = null as any;
utilities.lazyLoad(exports, ["getConnector","getConnectorOutput"], () => require("./getConnector"));

export { GetConnectorProfileArgs, GetConnectorProfileResult, GetConnectorProfileOutputArgs } from "./getConnectorProfile";
export const getConnectorProfile: typeof import("./getConnectorProfile").getConnectorProfile = null as any;
export const getConnectorProfileOutput: typeof import("./getConnectorProfile").getConnectorProfileOutput = null as any;
utilities.lazyLoad(exports, ["getConnectorProfile","getConnectorProfileOutput"], () => require("./getConnectorProfile"));

export { GetFlowArgs, GetFlowResult, GetFlowOutputArgs } from "./getFlow";
export const getFlow: typeof import("./getFlow").getFlow = null as any;
export const getFlowOutput: typeof import("./getFlow").getFlowOutput = null as any;
utilities.lazyLoad(exports, ["getFlow","getFlowOutput"], () => require("./getFlow"));


// Export enums:
export * from "../types/enums/appflow";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:appflow:Connector":
                return new Connector(name, <any>undefined, { urn })
            case "aws-native:appflow:ConnectorProfile":
                return new ConnectorProfile(name, <any>undefined, { urn })
            case "aws-native:appflow:Flow":
                return new Flow(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "appflow", _module)
