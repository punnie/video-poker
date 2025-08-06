package pkg

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
)

const (
	maxHighScores = 10
	configDir     = "video-poker"
	highScoreFile = "highscores.json"
)

type HighScore struct {
	Initials string `json:"initials"`
	Score    int    `json:"score"`
}

type HighScores []HighScore

func (hs HighScores) Len() int           { return len(hs) }
func (hs HighScores) Less(i, j int) bool { return hs[i].Score > hs[j].Score }
func (hs HighScores) Swap(i, j int)      { hs[i], hs[j] = hs[j], hs[i] }

func getHighScoreFilePath() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configPath, configDir, highScoreFile), nil
}

func LoadHighScores() (HighScores, error) {
	filePath, err := getHighScoreFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return HighScores{}, nil
		}
		return nil, err
	}

	var highScores HighScores
	err = json.Unmarshal(data, &highScores)
	if err != nil {
		return nil, err
	}
	return highScores, nil
}

func (hs HighScores) Save() error {
	filePath, err := getHighScoreFilePath()
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(hs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

func (hs HighScores) Add(score HighScore) HighScores {
	hs = append(hs, score)
	sort.Sort(hs)
	if len(hs) > maxHighScores {
		hs = hs[:maxHighScores]
	}
	return hs
}

func (hs HighScores) IsHighScore(score int) bool {
	if len(hs) < maxHighScores {
		return true
	}
	return score > hs[len(hs)-1].Score
}
