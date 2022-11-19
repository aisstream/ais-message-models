/**
 * Ais-Stream WebsocketObjects
 * A sample API to illustrate OpenAPI concepts
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class ShipStaticDataDimension {
    'A': number;
    'B': number;
    'C': number;
    'D': number;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "A",
            "baseName": "A",
            "type": "number",
            "format": ""
        },
        {
            "name": "B",
            "baseName": "B",
            "type": "number",
            "format": ""
        },
        {
            "name": "C",
            "baseName": "C",
            "type": "number",
            "format": ""
        },
        {
            "name": "D",
            "baseName": "D",
            "type": "number",
            "format": ""
        }    ];

    static getAttributeTypeMap() {
        return ShipStaticDataDimension.attributeTypeMap;
    }

    public constructor() {
    }
}
