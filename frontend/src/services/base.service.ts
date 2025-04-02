import { GrpcClientService } from "./grpc/grpc-client.service";

export class BaseService {
  protected grpcClient: GrpcClientService;

  constructor() {
    this.grpcClient = GrpcClientService.getInstance();
  }

  protected getTransport() {
    return this.grpcClient.getTransport();
  }
}
