import { BaseService } from "./base.service";
import { AuthenticationClient } from "../../gen/ts/kiioong/authentication/authentication_service.client";
import { Preferences } from "@capacitor/preferences";

export class AuthService extends BaseService {
  private authClient: AuthenticationClient;

  constructor() {
    super();
    this.authClient =
      this.grpcClient.createClient<AuthenticationClient>(AuthenticationClient);
  }

  public async login(username: string, password: string) {
    const result = await this.authClient.login({
      username: username,
      password: password,
    });

    if (result.response.jwtToken === "") {
      // validation
      console.error("Username not found or wrong password");
      return false;
    }

    await Preferences.set({
      key: "authToken",
      value: result.response.jwtToken,
    });

    this.setToken(result.response.jwtToken);

    return true;
  }
}
