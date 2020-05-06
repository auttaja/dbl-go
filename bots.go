package dbl_go

import (
	"time"
)

// Bot represents a single bot on top.gg
type Bot struct {
	DefAvatar        string    `json:"defAvatar"`
	Invite           string    `json:"invite"`
	Website          string    `json:"website"`
	Support          string    `json:"support"`
	Github           string    `json:"github"`
	Longdesc         string    `json:"longdesc"`
	Shortdesc        string    `json:"shortdesc"`
	Prefix           string    `json:"prefix"`
	Lib              string    `json:"lib"`
	Clientid         string    `json:"clientid"`
	Avatar           string    `json:"avatar"`
	ID               string    `json:"id"`
	Discriminator    string    `json:"discriminator"`
	Username         string    `json:"username"`
	Date             time.Time `json:"date"`
	ServerCount      int       `json:"server_count"`
	ShardCount       int       `json:"shard_count"`
	Guilds           []string  `json:"guilds"`
	Shards           []int     `json:"shards"`
	MonthlyPoints    int       `json:"monthlyPoints"`
	Points           int       `json:"points"`
	CertifiedBot     bool      `json:"certifiedBot"`
	Owners           []string  `json:"owners"`
	Tags             []string  `json:"tags"`
	Donatebotguildid string    `json:"donatebotguildid"`
}

// Stats contains the stats for a bot.
type Stats struct {
	ServerCount int   `json:"server_count,omitempty"`
	Shards      []int `json:"shards,omitempty"`
	ShardCount  int   `json:"shard_count,omitempty"`
	ShardID     int   `json:"shard_id,omitempty"`
}

// GetBot returns a single Bot from top.gg with the given id
func (c *DBL) GetBot(id string) (*Bot, error) {
	req, err := c.newRequest("GET", "/bots/"+id, nil)
	if err != nil {
		return nil, err
	}

	bot := new(Bot)
	_, err = c.do(req, bot)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

// CheckUserVote checks whether or not the given user has voted for the given bot
func (c *DBL) CheckUserVote(bot, user string) (bool, error) {
	req, err := c.newRequest("GET", "/bots/"+bot+"/check?userId="+user, nil)
	if err != nil {
		return false, err
	}

	status := make(map[string]string)
	_, err = c.do(req, &status)
	if err != nil {
		return false, err
	}

	if val, ok := status["voted"]; ok {
		return val == "1", nil
	}

	return false, nil
}

// GetBotStats gets the Stats for the given bot from top.gg
func (c *DBL) GetBotStats(bot string) (*Stats, error) {
	req, err := c.newRequest("GET", "/bots/"+bot+"/stats", nil)
	if err != nil {
		return nil, err
	}

	stats := new(Stats)
	_, err = c.do(req, stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

// PostBotStats updates the given bot with the Stats provided
func (c *DBL) PostBotStats(bot string, stats *Stats) error {
	req, err := c.newRequest("POST", "/bots/"+bot+"/stats", stats)
	if err != nil {
		return err
	}

	var resp interface{}
	_, err = c.do(req, &resp)
	if err != nil {
		return err
	}

	return nil
}
