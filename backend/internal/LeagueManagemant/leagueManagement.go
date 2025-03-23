package leaguemanagement

import (
	"context"
	"fmt"
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

func (s *LeagueManagementServer) GetLeagues(empty *emptypb.Empty, stream lm.LeagueManagement_GetLeaguesServer) error {
	var leagues []Database.League
	result := Database.DB.Find(&leagues)

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
	result := Database.DB.Find(&teams, "league_id = ?", league.Id)

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

	fmt.Println(in_team)

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

func NewServer() *LeagueManagementServer {
	s := &LeagueManagementServer{}
	return s
}
