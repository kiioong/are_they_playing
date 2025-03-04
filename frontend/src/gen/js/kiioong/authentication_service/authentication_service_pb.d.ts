import * as jspb from 'google-protobuf'



export class AuthenticationData extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): AuthenticationData;

  getPassword(): string;
  setPassword(value: string): AuthenticationData;

  getSessionId(): number;
  setSessionId(value: number): AuthenticationData;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthenticationData.AsObject;
  static toObject(includeInstance: boolean, msg: AuthenticationData): AuthenticationData.AsObject;
  static serializeBinaryToWriter(message: AuthenticationData, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthenticationData;
  static deserializeBinaryFromReader(message: AuthenticationData, reader: jspb.BinaryReader): AuthenticationData;
}

export namespace AuthenticationData {
  export type AsObject = {
    username: string,
    password: string,
    sessionId: number,
  }
}

export class AuthentificationStatus extends jspb.Message {
  getIsLoggedIn(): boolean;
  setIsLoggedIn(value: boolean): AuthentificationStatus;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AuthentificationStatus.AsObject;
  static toObject(includeInstance: boolean, msg: AuthentificationStatus): AuthentificationStatus.AsObject;
  static serializeBinaryToWriter(message: AuthentificationStatus, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AuthentificationStatus;
  static deserializeBinaryFromReader(message: AuthentificationStatus, reader: jspb.BinaryReader): AuthentificationStatus;
}

export namespace AuthentificationStatus {
  export type AsObject = {
    isLoggedIn: boolean,
  }
}

