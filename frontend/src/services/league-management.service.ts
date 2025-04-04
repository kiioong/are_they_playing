import { BaseService } from "./base.service";
import { LeagueManagementClient } from "../../gen/ts/kiioong/league_management/league_management.client";
import {
  Game,
  League,
  Sport,
  Team,
} from "../../gen/ts/kiioong/league_management/league_management";
import { AuthenticationClient } from "../../gen/ts/kiioong/authentication/authentication_service.client";
import { Ref } from "vue";

export class LeagueManagementService extends BaseService {
  private leagueManagementClient: LeagueManagementClient;

  constructor() {
    super();
    this.leagueManagementClient =
      this.grpcClient.createClient<LeagueManagementClient>(
        LeagueManagementClient,
      );
  }

  public async getSports() {
    let sports: Sport[] = [];

    const sportsStream = this.leagueManagementClient.getSports({});

    for await (let sport of sportsStream.responses) {
      sports.push(sport);
    }

    return sports;
  }

  public async getLeagues(sport: Sport) {
    let leagues: League[] = [];

    const leaguesStream = this.leagueManagementClient.getLeagues(sport);

    for await (let league of leaguesStream.responses) {
      leagues.push(league);
    }

    return leagues;
  }

  public async getTeams(league: League) {
    let teams: Team[] = [];

    const teamsStream = this.leagueManagementClient.getTeams(league);

    for await (let team of teamsStream.responses) {
      teams.push(team);
    }

    return teams;
  }

  public async addTeamToFavorites(team: Team) {
    const result = await this.leagueManagementClient.addTeamToFavourites(team);

    return result.response.success;
  }

  public async getGames(pickedDay: Date) {
    let games: Game[] = [];

    const gamesStream = this.leagueManagementClient.getGames({
      timestampOfDay: BigInt((pickedDay.valueOf() / 1000).toFixed(0)),
    });

    for await (let game of gamesStream.responses) {
      games.push(game);
    }

    return games;
  }
}
