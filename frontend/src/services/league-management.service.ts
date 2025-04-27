import { BaseService } from "./base.service";
import { LeagueManagementClient } from "../../gen/ts/kiioong/league_management/league_management.client";
import {
  Game,
  League,
  Sport,
  Team,
} from "../../gen/ts/kiioong/league_management/league_management";

export class LeagueManagementService extends BaseService {
  private leagueManagementClient: LeagueManagementClient;

  constructor() {
    super();
    this.leagueManagementClient =
      this.grpcClient.createClient<LeagueManagementClient>(
        LeagueManagementClient,
      );
  }

  public async getSports(): Promise<Sport[]> {
    try {
      const sports: Sport[] = [];
      const sportsStream = this.leagueManagementClient.getSports({});

      for await (const sport of sportsStream.responses) {
        sports.push(sport);
      }

      return sports;
    } catch (error) {
      await this.handleGrpcError(error);
      return [];
    }
  }

  public async getLeagues(sport: Sport): Promise<League[]> {
    try {
      const leagues: League[] = [];
      const leaguesStream = this.leagueManagementClient.getLeagues(sport);

      for await (const league of leaguesStream.responses) {
        leagues.push(league);
      }

      return leagues;
    } catch (error) {
      await this.handleGrpcError(error);
      return [];
    }
  }

  public async getTeams(league: League): Promise<Team[]> {
    try {
      const teams: Team[] = [];
      const teamsStream = this.leagueManagementClient.getTeams(league);

      for await (const team of teamsStream.responses) {
        teams.push(team);
      }

      return teams;
    } catch (error) {
      await this.handleGrpcError(error);
      return [];
    }
  }

  public async addTeamToFavorites(team: Team): Promise<boolean> {
    try {
      const result = await this.leagueManagementClient.addTeamToFavourites(team);
      return result.response.success;
    } catch (error) {
      await this.handleGrpcError(error);
      return false;
    }
  }

  public async removeTeamFromFavorites(team: Team): Promise<boolean> {
    try {
      const result = await this.leagueManagementClient.removeTeamFromFavourites(team);
      return result.response.success;
    } catch (error) {
      await this.handleGrpcError(error);
      return false;
    }
  }

  public async getFavouriteTeams(): Promise<Team[]> {
    try {
      const teams: Team[] = [];
      const teamsStream = this.leagueManagementClient.getFavouriteTeams({});

      for await (const team of teamsStream.responses) {
        teams.push(team);
      }

      return teams;
    } catch (error) {
      await this.handleGrpcError(error);
      return [];
    }
  }

  public async getGames(pickedDay: Date): Promise<Game[]> {
    try {
      const games: Game[] = [];
      const gamesStream = this.leagueManagementClient.getGames({
        timestampOfDay: BigInt((pickedDay.valueOf() / 1000).toFixed(0)),
      });

      for await (const game of gamesStream.responses) {
        games.push(game);
      }

      return games;
    } catch (error) {
      await this.handleGrpcError(error);
      return [];
    }
  }
}
