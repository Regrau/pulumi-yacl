// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class VpcRouteTable extends pulumi.CustomResource {
    /**
     * Get an existing VpcRouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcRouteTableState, opts?: pulumi.CustomResourceOptions): VpcRouteTable {
        return new VpcRouteTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'yandex:index/vpcRouteTable:VpcRouteTable';

    /**
     * Returns true if the given object is an instance of VpcRouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcRouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcRouteTable.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly folderId!: pulumi.Output<string>;
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly name!: pulumi.Output<string>;
    public readonly networkId!: pulumi.Output<string>;
    public readonly staticRoutes!: pulumi.Output<outputs.VpcRouteTableStaticRoute[] | undefined>;

    /**
     * Create a VpcRouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcRouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcRouteTableArgs | VpcRouteTableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcRouteTableState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["folderId"] = state ? state.folderId : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["staticRoutes"] = state ? state.staticRoutes : undefined;
        } else {
            const args = argsOrState as VpcRouteTableArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["folderId"] = args ? args.folderId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["staticRoutes"] = args ? args.staticRoutes : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcRouteTable.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcRouteTable resources.
 */
export interface VpcRouteTableState {
    createdAt?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    folderId?: pulumi.Input<string>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    networkId?: pulumi.Input<string>;
    staticRoutes?: pulumi.Input<pulumi.Input<inputs.VpcRouteTableStaticRoute>[]>;
}

/**
 * The set of arguments for constructing a VpcRouteTable resource.
 */
export interface VpcRouteTableArgs {
    description?: pulumi.Input<string>;
    folderId?: pulumi.Input<string>;
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name?: pulumi.Input<string>;
    networkId: pulumi.Input<string>;
    staticRoutes?: pulumi.Input<pulumi.Input<inputs.VpcRouteTableStaticRoute>[]>;
}