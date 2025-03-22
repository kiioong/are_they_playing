package leaguemanagement

import (
	"context"
	"fmt"

	lm "github.com/kiioong/are_they_playing/gen/go/kiioong/league_management"
	"github.com/kiioong/are_they_playing/internal/Database"
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
				Id:         int32(team.ID),
				Name:       team.Name,
				PathToLogo: team.PathToLogo,
			}, nil
		}

		fmt.Println(result.Error)
		return nil, nil
	}

	return &lm.Team{
		Id:         int32(team.ID),
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

func NewServer() *LeagueManagementServer {
	s := &LeagueManagementServer{}
	return s
}
