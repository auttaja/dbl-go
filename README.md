# dbl_go
--
    import "github.com/auttaja/dbl-go"

Package provides a Golang interface to the top.gg API

## Usage

```go
const (
	// The current version of the package
	Version = "0.0.1"
	// The top.gg base URL
	BaseURL = "https://top.gg"
)
```

#### type Bot

```go
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
```

Bot represents a single bot on top.gg

#### type DBL

```go
type DBL struct {
}
```


#### func  New

```go
func New(token string) *DBL
```
New creates a new top.gg client using the given token and the default http
client

#### func  NewDBL

```go
func NewDBL(token string, client *http.Client) *DBL
```
NewDBL creates a new top.gg client with the given token and HTTP client

#### func (*DBL) CheckUserVote

```go
func (c *DBL) CheckUserVote(bot, user string) (bool, error)
```
CheckUserVote checks whether or not the given user has voted for the given bot

#### func (*DBL) GetBot

```go
func (c *DBL) GetBot(id string) (*Bot, error)
```
GetBot returns a single Bot from top.gg with the given id

#### func (*DBL) GetBotStats

```go
func (c *DBL) GetBotStats(bot string) (*Stats, error)
```
GetBotStats gets the Stats for the given bot from top.gg

#### func (*DBL) GetUser

```go
func (c *DBL) GetUser(id string) (*User, error)
```
Get user gets a User for the given id from top.gg

#### func (*DBL) PostBotStats

```go
func (c *DBL) PostBotStats(bot string, stats *Stats) error
```
PostBotStats updates the given bot with the Stats provided

#### type Social

```go
type Social struct {
	Github    string `json:"github"`
	Instagram string `json:"instagram"`
	Reddit    string `json:"reddit"`
	Twitter   string `json:"twitter"`
	Youtube   string `json:"youtube"`
}
```

Social contains a users social connections on top.gg

#### type Stats

```go
type Stats struct {
	ServerCount int   `json:"server_count,omitempty"`
	Shards      []int `json:"shards,omitempty"`
	ShardCount  int   `json:"shard_count,omitempty"`
	ShardID     int   `json:"shard_id,omitempty"`
}
```

Stats contains the stats for a bot.

#### type User

```go
type User struct {
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	ID            string `json:"id"`
	Username      string `json:"username"`
	DefAvatar     string `json:"defAvatar"`
	Banner        string `json:"banner"`
	Bio           string `json:"bio"`
	Color         string `json:"color"`
	Admin         bool   `json:"admin"`
	WebMod        bool   `json:"webMod"`
	Mod           bool   `json:"mod"`
	CertifiedDev  bool   `json:"certifiedDev"`
	Supporter     bool   `json:"supporter"`
	Social        Social `json:"social"`
}
```

User represents a user on top.gg
