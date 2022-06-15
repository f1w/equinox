<div align="center">
	<h1>Equinox</h1>
	<p>Equinox is an API client written in golang with the goal of providing an easy to use interface to interact with all of the <a href="https://developer.riotgames.com/apis">Riot Games API</a> endpoints.
	</p>
	<p>
		<a href="https://github.com/Kyagara/equinox/releases">
			<img src="https://img.shields.io/github/v/tag/Kyagara/equinox?label=Version"/>
		</a>
		<a href="https://github.com/Kyagara/equinox">
			<img src="https://img.shields.io/github/go-mod/go-version/Kyagara/equinox"/>
		</a>    
		<a href="https://pkg.go.dev/github.com/Kyagara/equinox">
			<img src="https://img.shields.io/static/v1?label=Godoc&message=reference&color=blue"/>
		</a>
		<a href="https://github.com/Kyagara/equinox/actions?query=workflow%3Atests">
			<img src="https://img.shields.io/github/workflow/status/Kyagara/equinox/Tests?label=Tests"/>
		</a>
		<a href="https://codecov.io/gh/Kyagara/equinox">
			<img src="https://codecov.io/gh/Kyagara/equinox/branch/master/graph/badge.svg"/>
		</a>
	</p>
	<p>
		<a href="#features">Features</a> •
		<a href="#todo">Todo</a> •
		<a href="#installation">Installation</a> •
		<a href="#usage">Usage</a> •
		<a href="#example">Example</a> •
		<a href="#disclaimer">Disclaimer</a>
	</p>
</div>

#### Features

-   All Riot APIs implemented
    -   Riot Account
    -   League of Legends
    -   Teamfight Tactics
    -   Valorant
    -   Legends of Runeterra
-   Cache
-   Rate Limiting

> Caching and Rate limiting is in an early stage, they are in-memory only for now.

#### Todo

-   Add integration tests.
-   Review the cache and rate limiting implementation
-   Add Data Dragon support
-   Add Helper Methods (e.g.: GetChampion or GetSummoner inside a match response)

## Installation

To install equinox you can either import it in a package:

```go
import "github.com/Kyagara/equinox"
```

and run `go get` without any parameters, or, if you want to use the latest version from this repo, you can use the following command:

```bash
go get github.com/Kyagara/equinox@master
```

## Usage

To access the diferent parts of the Riot Games API, create a new instance of the Equinox client:

```go
client, err := equinox.NewClient("RIOT_API_KEY")
```

A client without a configuration struct comes with the default options:

```go
config := &api.EquinoxConfig{
	Key: "RIOT_API_KEY", // The API Key provided as a parameter.
	Cluster:  api.AmericasCluster, // Riot API cluster, use the nearest cluster. Options: AmericasCluster, EuropeCluster, AsiaCluster.
	LogLevel: api.FatalLevel, // The logging level, the FatalLevel provided effectively disables logging.
	Timeout:  10, // http.Client timeout in seconds.
	TTL:  120, // TTL for cache in seconds, 0 disables caching.
	Retry:    true, // Retry if the API returns a 429 response.
	RateLimit: true // If rate limit is enabled or not
}
```

A configuration struct can be used to create a custom Client using `equinox.NewClientWithConfig()`.

Now you can access different games endpoints by their abbreviations, provided you have access to them. For example:

```go
// This method uses a lol.Region. Can be accessed with a Development key.
summoner, err := client.LOL.Summoner.ByName(lol.BR1, "Loveable Senpai")

// The client.Cluster will be used as the region. Can be accessed with a Development key.
account, err := client.Riot.Account.ByPUUID("puuid")

// This method uses a val.Shard. May not be available in your policy.
recentMatches, err := client.VAL.Match.Recent(val.BR, val.CompetitiveQueue)
```

## Example

```go
package main

import (
	"fmt"

	"github.com/Kyagara/equinox"
	"github.com/Kyagara/equinox/lol"
)

func main() {
	// For custom configurations, you can use NewClientWithConfig(),
	// you will need to provide an api.EquinoxConfig{} object.
	client, err := equinox.NewClient("RIOT_API_KEY")

	if err != nil {
		fmt.Println("error creating client: ", err)
		return
	}

	// Get this week's champion rotations.
	rotation, err := client.LOL.Champion.Rotations(lol.BR1)

	if err != nil {
		fmt.Println("error retrieving champion rotations: ", err)
		return
	}

	fmt.Printf("%+v\n", rotation)
	// &{FreeChampionIDs:[17 43 56 62 67 79 85 90 133 145 147 157 201 203 245 518]
	// FreeChampionIDsForNewPlayers:[222 254 427 82 131 147 54 17 18 37]
	// MaxNewPlayerLevel:10}
}
```

## Disclaimer

Equinox isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing Riot Games properties. Riot Games, and all associated properties are trademarks or registered trademarks of Riot Games, Inc.
