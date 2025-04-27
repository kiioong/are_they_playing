import { BaseService } from "./base.service";
import { AuthenticationClient } from "../../gen/ts/kiioong/authentication/authentication_service.client";
import { Preferences } from "@capacitor/preferences";
import router from "@/router";
import { RpcError } from "@protobuf-ts/runtime-rpc";

export class AuthService extends BaseService {
  private authClient: AuthenticationClient;

  constructor() {
    super();
    this.authClient =
      this.grpcClient.createClient<AuthenticationClient>(AuthenticationClient);
  }

  public async login(username: string, password: string): Promise<boolean> {
    try {
      const result = await this.authClient.login({
        username: username,
        password: password,
      });

      if (result.response.jwtToken === "") {
        console.error("Username not found or wrong password");
        return false;
      }

      await Preferences.set({
        key: "authToken",
        value: result.response.jwtToken,
      });

      this.setToken(result.response.jwtToken);
      return true;
    } catch (error) {
      await this.handleGrpcError(error);
      return false;
    }
  }

  public async validateToken(token: string): Promise<boolean> {
    if (!token) return false;

    try {
      const validationClient =
        this.grpcClient.createClient<AuthenticationClient>(
          AuthenticationClient,
        );

      await validationClient.validateToken({
        jwtToken: token,
      });

      return true;
    } catch (error) {
      if (error instanceof RpcError && error.code === "UNAUTHENTICATED") {
        await this.logout();
        return false;
      }
      await this.handleGrpcError(error);
      return false;
    }
  }

  public async logout(): Promise<void> {
    try {
      await Preferences.remove({ key: "authToken" });
      this.setToken("");
      await router.push("/login");
    } catch (error) {
      await this.handleGrpcError(error);
    }
  }
}
