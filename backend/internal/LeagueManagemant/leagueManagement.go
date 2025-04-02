package leaguemanagement

import (
	"context"
	"fmt"
	"strconv"
	"time"

	lm "github.com/kiioong/are_they_playing/gen/go/kiioong/league_management"
	authenticationService "github.com/kiioong/are_they_playing/internal/AuthenticationService"
	"github.com/kiioong/are_they_playing/internal/Database"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type LeagueManagementServer struct {
	lm.UnimplementedLeagueManagementServer
}

func (s *LeagueManagementServer) AddTeam(ctx context.Context, in *lm.Team) (*lm.Team, error) {
	var team Database.Team

	result := Database.DB.Where("name = ?", in.Name).First(&team)

	if result.Error == gorm.ErrRecordNotFound {
		team = Database.Team{Name: in.Name}

		result := Database.DB.Create(&team)

		if result.Error == nil {
			return &lm.Team{
				Id:         team.ID,
				Name:       team.Name,
				PathToLogo: team.PathToLogo,
			}, nil
		}

		fmt.Println(result.Error)
		return nil, nil
	}

	return &lm.Team{
		Id:         team.ID,
		Name:       team.Name,
		PathToLogo: team.PathToLogo,
	}, nil
}

func (s *LeagueManagementServer) AddTeamToLeague(ctx context.Context, in *lm.TeamLeague) (*lm.MutationResult, error) {
	var team Database.Team
	var league Database.League

	result := Database.DB.Where("id = ?", in.Team.Id).First(&team)

	if result.Error != nil {
		fmt.Println(result.Error)
		return &lm.MutationResult{
			Success: false,
		}, nil
	}

	result = Database.DB.Where("id = ?", in.League.Id).First(&league)

	if result.Error != nil {
		fmt.Println(result.Error)
		return &lm.MutationResult{
			Success: false,
		}, nil
	}

	team.Leagues = append(team.Leagues, league)
	Database.DB.Save(&team)

	return &lm.MutationResult{
		Success: true,
	}, nil
}

func (s *LeagueManagementServer) AddGame(ctx context.Context, in_game *lm.Game) (*lm.MutationResult, error) {
	var homeTeam Database.Team
	var awayTeam Database.Team
	var league Database.League

	result := Database.DB.Where("id = ?", in_game.League.Id).First(&league)

	if result.Error == gorm.ErrRecordNotFound {
		return &lm.MutationResult{
			Success: false,
		}, nil
	}

	result = Database.DB.Where("name = ?", in_game.HomeTeam.Name).First(&homeTeam)

	if result.Error == gorm.ErrRecordNotFound {
		return &lm.MutationResult{
			Success: false,
		}, nil
	}

	result = Database.DB.Where("name = ?", in_game.AwayTeam.Name).First(&awayTeam)

	if result.Error == gorm.ErrRecordNotFound {
		return &lm.MutationResult{
			Success: false,
		}, nil
	}

	game := Database.Game{HomeTeamID: uint(homeTeam.ID), AwayTeamID: uint(awayTeam.ID), StartTime: time.Unix(int64(in_game.StartTimestamp), 0), LeagueID: uint(league.ID)}

	result = Database.DB.Create(&game)

	if result.Error == nil {
		return &lm.MutationResult{
			Success: true,
		}, nil
	}

	return &lm.MutationResult{
		Success: false,
	}, nil

}

func (s *LeagueManagementServer) GetSports(empty *emptypb.Empty, stream lm.LeagueManagement_GetSportsServer) error {
	var sports []Database.Sport
	result := Database.DB.Find(&sports)

	if result.Error != nil {
		return result.Error
	}

	for _, sport := range sports {
		if err := stream.Send(&lm.Sport{Id: uint64(sport.ID), Name: sport.Name}); err != nil {
			return err
		}
	}

	return nil
}

func (s *LeagueManagementServer) GetLeagues(sport *lm.Sport, stream lm.LeagueManagement_GetLeaguesServer) error {
	var leagues []Database.League
	result := Database.DB.Find(&leagues, "sport_id =  ?", sport.Id)

	if result.Error != nil {
		return result.Error
	}

	for _, league := range leagues {
		if err := stream.Send(&lm.League{Id: league.ID, Name: league.Name}); err != nil {
			return err
		}
	}

	return nil
}

func (s *LeagueManagementServer) GetTeams(league *lm.League, stream lm.LeagueManagement_GetTeamsServer) error {
	var teams []Database.Team

	result := Database.DB.Joins("JOIN league_teams ON league_teams.team_id = teams.id").Where("league_teams.league_id = ?", league.Id).Find(&teams)

	if result.Error != nil {
		return result.Error
	}

	for _, team := range teams {
		if err := stream.Send(&lm.Team{Id: team.ID, Name: team.Name, PathToLogo: team.PathToLogo}); err != nil {
			return err
		}
	}

	return nil
}

func (s *LeagueManagementServer) AddTeamToFavourites(ctx context.Context, in_team *lm.Team) (*lm.MutationResult, error) {
	var user Database.User
	var team Database.Team

	user_id := ctx.Value(authenticationService.User{})

	result := Database.DB.Where("id = ?", user_id).First(&user)

	if result.Error != nil {
		return &lm.MutationResult{
			Success: false,
		}, result.Error
	}

	result = Database.DB.Where("id = ?", in_team.Id).First(&team)

	if result.Error != nil {
		return &lm.MutationResult{
			Success: false,
		}, result.Error
	}

	userTeam := Database.UserTeam{
		UserID: uint(user.ID),
		TeamID: team.ID,
	}

	result = Database.DB.Create(userTeam)

	if result.Error != nil {
		return &lm.MutationResult{
			Success: false,
		}, result.Error
	}

	return &lm.MutationResult{
		Success: true,
	}, nil

}

func (s *LeagueManagementServer) GetGames(game_request *lm.GameRequest, stream lm.LeagueManagement_GetGamesServer) error {
	var games []Database.Game
	var home_team Database.Team
	var away_team Database.Team
	var clear_team Database.Team
	var league Database.League

	ctx := stream.Context()

	user_id, err := strconv.ParseUint(ctx.Value(authenticationService.User{}).(string), 10, 64)

	if err != nil {
		return err
	}

	teams, err := getFavouriteTeams(user_id)

	if err != nil {
		return err
	}

	var team_ids []uint32

	for _, team := range teams {
		team_ids = append(team_ids, team.ID)
	}

	result := Database.DB.Find(&games, "date(start_time) = ? AND (home_team_id in (?) OR games.away_team_id in (?))", time.Unix(game_request.TimestampOfDay, 0).Format("2006-01-02"), team_ids, team_ids)

	if result.Error != nil {
		return result.Error
	}

	for _, game := range games {
		home_team = clear_team
		away_team = clear_team

		result = Database.DB.Where("id = ?", game.HomeTeamID).First(&home_team)

		if result.Error != nil {
			continue
		}

		result = Database.DB.Where("id = ?", game.AwayTeamID).First(&away_team)

		if result.Error != nil {
			continue
		}

		result = Database.DB.Where("id = ?", game.LeagueID).First(&league)

		if result.Error != nil {
			continue
		}

		if err := stream.Send(&lm.Game{HomeTeam: &lm.Team{Id: home_team.ID, Name: home_team.Name}, AwayTeam: &lm.Team{Id: away_team.ID, Name: away_team.Name}, StartTimestamp: game.StartTime.Unix(), League: &lm.League{Id: league.ID, Name: league.Name}}); err != nil {
			return err
		}
	}

	return nil
}

func getFavouriteTeams(user_id uint64) ([]Database.Team, error) {
	var teams []Database.Team

	result := Database.DB.Joins("JOIN user_teams ON user_teams.team_id = teams.id").Where("user_teams.user_id = ?", user_id).Find(&teams)

	if result.Error != nil {
		return nil, result.Error
	}

	return teams, nil
}

func NewServer() *LeagueManagementServer {
	s := &LeagueManagementServer{}
	return s
}
