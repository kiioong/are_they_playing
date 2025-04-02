import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";

export class GrpcClientService {
  private static instance: GrpcClientService;
  private transport: GrpcWebFetchTransport | null = null;

  private constructor() {
    this.initializeChannel();
  }

  public static getInstance(): GrpcClientService {
    if (!GrpcClientService.instance) {
      GrpcClientService.instance = new GrpcClientService();
    }
    return GrpcClientService.instance;
  }

  private initializeChannel() {
    this.transport = new GrpcWebFetchTransport({
      baseUrl: "http://localhost:10000",
    });
  }

  public getTransport() {
    return this.transport;
  }
}
