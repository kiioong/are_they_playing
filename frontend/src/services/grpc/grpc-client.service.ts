import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { AuthInterceptor } from "@/services/grpc/auth.interceptor";

export class GrpcClientService {
  private static instance: GrpcClientService;
  private token: string = "";
  private authInterceptor: AuthInterceptor;

  private constructor() {
    this.authInterceptor = new AuthInterceptor(() => this.token);
  }

  public static getInstance(): GrpcClientService {
    if (!GrpcClientService.instance) {
      GrpcClientService.instance = new GrpcClientService();
    }
    return GrpcClientService.instance;
  }

  public setToken(token: string) {
    this.token = token;
  }

  public getToken(): string {
    return this.token;
  }

  public createTransport(): GrpcWebFetchTransport {
    return new GrpcWebFetchTransport({
      baseUrl: "http://localhost:10000",
      interceptors: [this.authInterceptor],
    });
  }

  // Create a client with the interceptor
  public createClient<T>(ClientConstructor: any): T {
    const transport = this.createTransport();
    return new ClientConstructor(transport);
  }
}
