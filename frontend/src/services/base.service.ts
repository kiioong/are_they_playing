import { GrpcClientService } from "./grpc/grpc-client.service";

export class BaseService {
  protected grpcClient: GrpcClientService;

  constructor() {
    this.grpcClient = GrpcClientService.getInstance();
  }

  public setToken(token: string) {
    this.grpcClient.setToken(token);
  }
}
