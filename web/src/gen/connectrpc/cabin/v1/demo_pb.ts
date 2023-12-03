// @generated by protoc-gen-es v1.5.0 with parameter "target=ts"
// @generated from file cabin/v1/demo.proto (package cabin.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum cabin.v1.DemoStatus
 */
export enum DemoStatus {
  /**
   * @generated from enum value: DEMO_STATUS_UNSPECIFIED = 0;
   */
  DEMO_STATUS_UNSPECIFIED = 0,

  /**
   * @generated from enum value: AWAITING = 1;
   */
  AWAITING = 1,

  /**
   * @generated from enum value: PROCESSED = 2;
   */
  PROCESSED = 2,

  /**
   * @generated from enum value: ERROR = 3;
   */
  ERROR = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(DemoStatus)
proto3.util.setEnumType(DemoStatus, "cabin.v1.DemoStatus", [
  { no: 0, name: "DEMO_STATUS_UNSPECIFIED" },
  { no: 1, name: "AWAITING" },
  { no: 2, name: "PROCESSED" },
  { no: 3, name: "ERROR" },
]);

/**
 * @generated from message cabin.v1.Demo
 */
export class Demo extends Message<Demo> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string identity_id = 2;
   */
  identityId = "";

  /**
   * @generated from field: cabin.v1.DemoStatus status = 3;
   */
  status = DemoStatus.DEMO_STATUS_UNSPECIFIED;

  /**
   * @generated from field: string reason = 4;
   */
  reason = "";

  /**
   * @generated from field: google.protobuf.Timestamp uploaded_at = 5;
   */
  uploadedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp processed_at = 6;
   */
  processedAt?: Timestamp;

  constructor(data?: PartialMessage<Demo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cabin.v1.Demo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "identity_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(DemoStatus) },
    { no: 4, name: "reason", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "uploaded_at", kind: "message", T: Timestamp },
    { no: 6, name: "processed_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Demo {
    return new Demo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Demo {
    return new Demo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Demo {
    return new Demo().fromJsonString(jsonString, options);
  }

  static equals(a: Demo | PlainMessage<Demo> | undefined, b: Demo | PlainMessage<Demo> | undefined): boolean {
    return proto3.util.equals(Demo, a, b);
  }
}

/**
 * @generated from message cabin.v1.GetDemoRequest
 */
export class GetDemoRequest extends Message<GetDemoRequest> {
  /**
   * @generated from field: string demo_id = 1;
   */
  demoId = "";

  constructor(data?: PartialMessage<GetDemoRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cabin.v1.GetDemoRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "demo_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDemoRequest {
    return new GetDemoRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDemoRequest {
    return new GetDemoRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDemoRequest {
    return new GetDemoRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetDemoRequest | PlainMessage<GetDemoRequest> | undefined, b: GetDemoRequest | PlainMessage<GetDemoRequest> | undefined): boolean {
    return proto3.util.equals(GetDemoRequest, a, b);
  }
}

/**
 * @generated from message cabin.v1.GetDemoResponse
 */
export class GetDemoResponse extends Message<GetDemoResponse> {
  /**
   * @generated from field: cabin.v1.Demo demo = 1;
   */
  demo?: Demo;

  constructor(data?: PartialMessage<GetDemoResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cabin.v1.GetDemoResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "demo", kind: "message", T: Demo },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetDemoResponse {
    return new GetDemoResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetDemoResponse {
    return new GetDemoResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetDemoResponse {
    return new GetDemoResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetDemoResponse | PlainMessage<GetDemoResponse> | undefined, b: GetDemoResponse | PlainMessage<GetDemoResponse> | undefined): boolean {
    return proto3.util.equals(GetDemoResponse, a, b);
  }
}

/**
 * @generated from message cabin.v1.ListDemosRequest
 */
export class ListDemosRequest extends Message<ListDemosRequest> {
  constructor(data?: PartialMessage<ListDemosRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cabin.v1.ListDemosRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListDemosRequest {
    return new ListDemosRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListDemosRequest {
    return new ListDemosRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListDemosRequest {
    return new ListDemosRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListDemosRequest | PlainMessage<ListDemosRequest> | undefined, b: ListDemosRequest | PlainMessage<ListDemosRequest> | undefined): boolean {
    return proto3.util.equals(ListDemosRequest, a, b);
  }
}

/**
 * @generated from message cabin.v1.ListDemosResponse
 */
export class ListDemosResponse extends Message<ListDemosResponse> {
  constructor(data?: PartialMessage<ListDemosResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "cabin.v1.ListDemosResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListDemosResponse {
    return new ListDemosResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListDemosResponse {
    return new ListDemosResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListDemosResponse {
    return new ListDemosResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListDemosResponse | PlainMessage<ListDemosResponse> | undefined, b: ListDemosResponse | PlainMessage<ListDemosResponse> | undefined): boolean {
    return proto3.util.equals(ListDemosResponse, a, b);
  }
}

