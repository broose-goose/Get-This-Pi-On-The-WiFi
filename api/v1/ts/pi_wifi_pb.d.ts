import * as jspb from 'google-protobuf'



export class EmptyRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmptyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EmptyRequest): EmptyRequest.AsObject;
  static serializeBinaryToWriter(message: EmptyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmptyRequest;
  static deserializeBinaryFromReader(message: EmptyRequest, reader: jspb.BinaryReader): EmptyRequest;
}

export namespace EmptyRequest {
  export type AsObject = {
  }
}

export class SetWiFiDetails extends jspb.Message {
  getSsid(): string;
  setSsid(value: string): SetWiFiDetails;

  getPassword(): string;
  setPassword(value: string): SetWiFiDetails;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetWiFiDetails.AsObject;
  static toObject(includeInstance: boolean, msg: SetWiFiDetails): SetWiFiDetails.AsObject;
  static serializeBinaryToWriter(message: SetWiFiDetails, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetWiFiDetails;
  static deserializeBinaryFromReader(message: SetWiFiDetails, reader: jspb.BinaryReader): SetWiFiDetails;
}

export namespace SetWiFiDetails {
  export type AsObject = {
    ssid: string,
    password: string,
  }
}

export class PiWiFiRequest extends jspb.Message {
  getRequestId(): number;
  setRequestId(value: number): PiWiFiRequest;

  getGetdevicedescription(): EmptyRequest | undefined;
  setGetdevicedescription(value?: EmptyRequest): PiWiFiRequest;
  hasGetdevicedescription(): boolean;
  clearGetdevicedescription(): PiWiFiRequest;

  getSetwifidetails(): SetWiFiDetails | undefined;
  setSetwifidetails(value?: SetWiFiDetails): PiWiFiRequest;
  hasSetwifidetails(): boolean;
  clearSetwifidetails(): PiWiFiRequest;

  getRequestCase(): PiWiFiRequest.RequestCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PiWiFiRequest.AsObject;
  static toObject(includeInstance: boolean, msg: PiWiFiRequest): PiWiFiRequest.AsObject;
  static serializeBinaryToWriter(message: PiWiFiRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PiWiFiRequest;
  static deserializeBinaryFromReader(message: PiWiFiRequest, reader: jspb.BinaryReader): PiWiFiRequest;
}

export namespace PiWiFiRequest {
  export type AsObject = {
    requestId: number,
    getdevicedescription?: EmptyRequest.AsObject,
    setwifidetails?: SetWiFiDetails.AsObject,
  }

  export enum RequestCase { 
    REQUEST_NOT_SET = 0,
    GETDEVICEDESCRIPTION = 2,
    SETWIFIDETAILS = 3,
  }
}

export class EmptyResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EmptyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EmptyResponse): EmptyResponse.AsObject;
  static serializeBinaryToWriter(message: EmptyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EmptyResponse;
  static deserializeBinaryFromReader(message: EmptyResponse, reader: jspb.BinaryReader): EmptyResponse;
}

export namespace EmptyResponse {
  export type AsObject = {
  }
}

export class SuccessFailure extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): SuccessFailure;

  getError(): string;
  setError(value: string): SuccessFailure;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SuccessFailure.AsObject;
  static toObject(includeInstance: boolean, msg: SuccessFailure): SuccessFailure.AsObject;
  static serializeBinaryToWriter(message: SuccessFailure, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SuccessFailure;
  static deserializeBinaryFromReader(message: SuccessFailure, reader: jspb.BinaryReader): SuccessFailure;
}

export namespace SuccessFailure {
  export type AsObject = {
    success: boolean,
    error: string,
  }
}

export class DeviceDescription extends jspb.Message {
  getDevicename(): string;
  setDevicename(value: string): DeviceDescription;

  getDevicewifiname(): string;
  setDevicewifiname(value: string): DeviceDescription;

  getWifistatus(): WiFiStatus;
  setWifistatus(value: WiFiStatus): DeviceDescription;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeviceDescription.AsObject;
  static toObject(includeInstance: boolean, msg: DeviceDescription): DeviceDescription.AsObject;
  static serializeBinaryToWriter(message: DeviceDescription, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeviceDescription;
  static deserializeBinaryFromReader(message: DeviceDescription, reader: jspb.BinaryReader): DeviceDescription;
}

export namespace DeviceDescription {
  export type AsObject = {
    devicename: string,
    devicewifiname: string,
    wifistatus: WiFiStatus,
  }
}

export class PiWiFiResponse extends jspb.Message {
  getResponseId(): number;
  setResponseId(value: number): PiWiFiResponse;

  getStatus(): SuccessFailure | undefined;
  setStatus(value?: SuccessFailure): PiWiFiResponse;
  hasStatus(): boolean;
  clearStatus(): PiWiFiResponse;

  getDevicedescription(): DeviceDescription | undefined;
  setDevicedescription(value?: DeviceDescription): PiWiFiResponse;
  hasDevicedescription(): boolean;
  clearDevicedescription(): PiWiFiResponse;

  getSetwifidetailsresponse(): EmptyResponse | undefined;
  setSetwifidetailsresponse(value?: EmptyResponse): PiWiFiResponse;
  hasSetwifidetailsresponse(): boolean;
  clearSetwifidetailsresponse(): PiWiFiResponse;

  getResponseCase(): PiWiFiResponse.ResponseCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PiWiFiResponse.AsObject;
  static toObject(includeInstance: boolean, msg: PiWiFiResponse): PiWiFiResponse.AsObject;
  static serializeBinaryToWriter(message: PiWiFiResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PiWiFiResponse;
  static deserializeBinaryFromReader(message: PiWiFiResponse, reader: jspb.BinaryReader): PiWiFiResponse;
}

export namespace PiWiFiResponse {
  export type AsObject = {
    responseId: number,
    status?: SuccessFailure.AsObject,
    devicedescription?: DeviceDescription.AsObject,
    setwifidetailsresponse?: EmptyResponse.AsObject,
  }

  export enum ResponseCase { 
    RESPONSE_NOT_SET = 0,
    DEVICEDESCRIPTION = 3,
    SETWIFIDETAILSRESPONSE = 4,
  }
}

export enum WiFiStatus { 
  DISCONNECTING = 0,
  SETTING = 1,
  RESTARTING = 2,
  CONNECTING = 3,
  CONNECTED = 4,
}
