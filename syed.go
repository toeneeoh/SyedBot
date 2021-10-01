package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/bwmarrin/discordgo"
	"github.com/machinebox/graphql"
)

type Config struct {
	DiscordToken string
	Twitter      Twitter
}

type Twitter struct {
	Token       string
	TokenSecret string
	Key         string
	KeySecret   string
}

type AniData struct {
	Media struct {
		ID    int `json:"id"`
		Title struct {
			English string `json:"english"`
			Romaji  string `json:"romaji"`
		} `json:"title"`
		Type       string   `json:"type"`
		Genres     []string `json:"genres"`
		CoverImage struct {
			Large string `json:"large"`
			Color string `json:"color"`
		} `json:"coverImage"`
		Status            string `json:"status"`
		Season            string `json:"season"`
		SeasonYear        int    `json:"seasonYear"`
		Episodes          int    `json:"episodes"`
		AverageScore      int    `json:"averageScore"`
		MeanScore         int    `json:"meanScore"`
		Format            string `json:"format"`
		Description       string `json:"description"`
		NextAiringEpisode struct {
			AiringAt int `json:"airingAt"`
			Episode  int `json:"episode"`
		} `json:"nextAiringEpisode"`
	} `json:"Media"`
}

type AniStaffData struct {
	Staff struct {
		ID                 int      `json:"id"`
		Gender             string   `json:"gender"`
		Age                int      `json:"age"`
		PrimaryOccupations []string `json:"primaryOccupations"`
		DateOfBirth        struct {
			Year  int `json:"year"`
			Month int `json:"month"`
			Day   int `json:"day"`
		} `json:"dateOfBirth"`
		Name struct {
			Full string `json:"full"`
		} `json:"name"`
		Image struct {
			Large string `json:"large"`
		} `json:"image"`
		Characters struct {
			Nodes []struct {
				ID   int `json:"id"`
				Name struct {
					Full string `json:"full"`
				} `json:"name"`
				Media struct {
					Nodes []struct {
						ID    int `json:"id"`
						Title struct {
							Romaji  string `json:"romaji"`
							English string `json:"english"`
						} `json:"title"`
					} `json:"nodes"`
				} `json:"media"`
			} `json:"nodes"`
		} `json:"characters"`
	} `json:"Staff"`
}

var config Config

func init() {
	readin, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("Error: " + err.Error())
	} //put some error handling here later
	_ = json.Unmarshal(readin, &config)
}

func main() {

	DiscordToken := config.DiscordToken

	DiscordSession, err := discordgo.New("Bot " + DiscordToken)
	if err != nil {
		log.Fatalln("Error creating Discord session" + err.Error())
	}

	rand.Seed(time.Now().UnixNano())
	DiscordSession.AddHandler(MessageHandler)
	DiscordSession.AddHandler(ReactHandler)
	DiscordSession.Identify.Intents = discordgo.IntentsGuildMessages

	err = DiscordSession.Open()
	if err != nil {
		log.Fatalln("Error opening Discord connection" + err.Error())
	}

	log.Println("Bot started")

	//Run until term signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Kill)
	<-sc

	//Close the bot
	DiscordSession.Close()
}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	log.Println(m.Content)

	if m.Content == "piss" {
		s.ChannelMessageSend(m.ChannelID, "shid")
	}

	if m.Content == "syed" {
		s.ChannelMessageSend(m.ChannelID, "ji?")
	}

	if m.Content == "salam" {
		s.ChannelMessageSend(m.ChannelID, "salam")
	}

	if m.Content == "?github" {
		s.ChannelMessageSend(m.ChannelID, "https://github.com/Monko2k/SyedBot")
	}

	if strings.HasPrefix(m.Content, "?wholesome ") {
		clipped := strings.Replace(m.Content, "?wholesome ", "", 1)
		wholesomeamt := rand.Intn(101)
		var wholesomestat string
		if wholesomeamt == 0 {
			wholesomestat = " is the least wholesome of them all."
		} else if wholesomeamt < 25 {
			wholesomestat = " is definitively unwholesome."
		} else if wholesomeamt < 50 {
			wholesomestat = " is pretty unwholesome."
		} else if wholesomeamt < 75 {
			wholesomestat = " is pretty wholesome!"
		} else if wholesomeamt < 100 {
			wholesomestat = " is incredibly wholesome!"
		} else {
			wholesomestat = " is super freaking wholesome!"
		}
		message := clipped + " is " + strconv.Itoa(wholesomeamt) + "% wholesome\n" + clipped + wholesomestat

		s.ChannelMessageSend(m.ChannelID, message)
	}

	if strings.HasPrefix(m.Content, "?whitecatify ") {
		clipped := strings.Replace(m.Content, "?whitecatify ", "", 1)
		s.ChannelMessageSend(m.ChannelID, "holy shit guys, "+clipped)
	}

	//surely these are better to do with regex
	if strings.HasPrefix(m.Content, "im ") {
		clipped := strings.Replace(m.Content, "im ", "", 1)
		s.ChannelMessageSend(m.ChannelID, "hi "+clipped)
	}
	if strings.HasPrefix(m.Content, "i'm ") {
		clipped := strings.Replace(m.Content, "i'm ", "", 1)
		s.ChannelMessageSend(m.ChannelID, "hi "+clipped)
	}
	if strings.HasPrefix(m.Content, "IM ") {
		clipped := strings.Replace(m.Content, "IM ", "", 1)
		s.ChannelMessageSend(m.ChannelID, "hi "+clipped)
	}
	if strings.HasPrefix(m.Content, "I'M ") {
		clipped := strings.Replace(m.Content, "I'M ", "", 1)
		s.ChannelMessageSend(m.ChannelID, "hi "+clipped)
	}
	if strings.HasPrefix(m.Content, "I'm ") {
		clipped := strings.Replace(m.Content, "I'm ", "", 1)
		s.ChannelMessageSend(m.ChannelID, "hi "+clipped)
	}

	if m.Content == "dsd" {
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
		time.Sleep(2 * time.Second)
		reactionMessage, _ := s.ChannelMessage(m.ChannelID, m.ID)

		for _, x := range reactionMessage.Reactions {
			log.Println(reactionMessage.Reactions[0].Emoji)
			if x.Emoji.Name == "✅" && x.Count > 1 {
				s.ChannelMessageSend(m.ChannelID, "yeaaah")
			}
			log.Println(x.Emoji.Name)
		}
	}

	if strings.HasPrefix(m.Content, "?tweet ") {
		clipped := strings.Replace(m.Content, "?tweet ", "", 1)
		s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
		s.MessageReactionAdd(m.ChannelID, m.ID, "🖕")
		time.Sleep(10 * time.Second)
		reactionMessage, _ := s.ChannelMessage(m.ChannelID, m.ID)

		upvote := 0
		downvote := 0
		for _, x := range reactionMessage.Reactions {
			if x.Emoji.Name == "✅" {
				upvote = x.Count
			} else if x.Emoji.Name == "🖕" {
				downvote = x.Count
			}
		}

		if upvote > 3 && upvote-downvote > 2 {
			TwitterSession := anaconda.NewTwitterApiWithCredentials(config.Twitter.Token, config.Twitter.TokenSecret, config.Twitter.Key, config.Twitter.KeySecret)
			tweet, err := TwitterSession.PostTweet(clipped, url.Values{})
			if err != nil {
				log.Println("Tweet post failed" + err.Error())
			} else {
				tweeturl := "https://twitter.com/BotSyed/status/" + strconv.Itoa(int(tweet.Id))
				s.ChannelMessageSend(m.ChannelID, tweeturl)
			}
			TwitterSession.Close()
		} else {
			s.ChannelMessageSend(m.ChannelID, "Not enough upvotes! (need at least 3)")
		}

	}

	if strings.HasPrefix(m.Content, "?choose ") {
		clipped := strings.Replace(m.Content, "?choose ", "", 1)
		var divider string
		if strings.Contains(m.Content, ", ") {
			divider = ", "
		} else {
			divider = " "
		}
		options := strings.Split(clipped, divider)
		log.Println(len(options))
		for _, x := range options {
			log.Println(x)
		}
		if len(options) == 0 {
			return
		} else if len(options) == 1 {
			s.ChannelMessageSend(m.ChannelID, options[0])
		} else {
			selection := rand.Intn(len(options))
			s.ChannelMessageSend(m.ChannelID, options[selection])
		}
	}

	if strings.HasPrefix(m.Content, "?anime ") {
		clipped := strings.Replace(m.Content, "?anime ", "", 1)
		graphqlClient := graphql.NewClient("https://graphql.anilist.co")
		graphqlRequest := graphql.NewRequest(`
			{
				Media(search: "` + clipped + `", type: ANIME, sort: SEARCH_MATCH	) {
					id
					title {
						romaji
						english
					}
					type
					genres
					coverImage {
						large
						color
					}
					status
					season
					seasonYear
					episodes
					averageScore
					meanScore
					format
					description (asHtml: false)
					nextAiringEpisode {
						airingAt
						episode
					}
				}
			}
		`)
		var graphqlResponse AniData

		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			log.Println(err.Error())
			s.ChannelMessageSend(m.ChannelID, "Anime not found!")
			return
		}
		color := 0xFFFFFF
		if graphqlResponse.Media.CoverImage.Color != "" {
			colorhexstring := strings.Replace(graphqlResponse.Media.CoverImage.Color, "#", "", 1)
			colorvalue, _ := strconv.ParseInt(colorhexstring, 16, 64)
			color = int(colorvalue)
		}

		var title string
		var subtitle string
		if graphqlResponse.Media.Title.English != "" {
			title = graphqlResponse.Media.Title.English
			if graphqlResponse.Media.Title.English != graphqlResponse.Media.Title.Romaji {
				subtitle = "**" + graphqlResponse.Media.Title.Romaji + "**\n"
			} else {
				subtitle = ""
			}
		} else {
			title = graphqlResponse.Media.Title.Romaji
			subtitle = ""
		}
		var genres string
		for i, s := range graphqlResponse.Media.Genres {
			if i == 0 {
				genres += s
			} else {
				genres += ", " + s
			}
		}
		var airingTime string
		if graphqlResponse.Media.Status == "RELEASING" {
			convtime := time.Unix(int64(graphqlResponse.Media.NextAiringEpisode.AiringAt), 0)
			airingTime = "\n**Next Airing: **Episode " + strconv.Itoa(graphqlResponse.Media.NextAiringEpisode.Episode) + " on " + convtime.Month().String() + " " + strconv.Itoa(convtime.Day()) + " " + strconv.Itoa(convtime.Year())
		} else {
			airingTime = ""
		}
		var episodes string
		if graphqlResponse.Media.Format == "MOVIE" {
			episodes = ""
		} else if graphqlResponse.Media.Episodes != 0 {
			episodes = "\n**Episodes:  **" + strconv.Itoa(graphqlResponse.Media.Episodes)
		} else {
			episodes = "\n**Not Yet Aired**"
		}
		description := strings.Split(graphqlResponse.Media.Description, "<br>")[0] + "\n\n" // only use everything before the first linebreak returned by description

		re, err := regexp.Compile(`(?:<[\/a-z]*>)`)
		if err != nil {
			log.Println(err.Error())
		}
		description = re.ReplaceAllString(description, "")
		var format string
		switch graphqlResponse.Media.Format {
		case "TV":
			format = "*TV Series*\n\n"
		case "TV_SHORT":
			format = "*TV Short*\n\n"
		case "MOVIE":
			format = "*Movie*\n\n"
		case "SPECIAL":
			format = "*Special*\n\n"
		case "MUSIC":
			format = "*Music*\n\n"
		default:
			format = "*" + graphqlResponse.Media.Format + "*\n\n"
		}
		var season string
		if graphqlResponse.Media.Season != "" {
			season = "**Season:  **" + strings.Title(strings.ToLower(graphqlResponse.Media.Season)+" "+strconv.Itoa(graphqlResponse.Media.SeasonYear))
		} else {
			season = ""
		}
		var averageScore string
		if graphqlResponse.Media.AverageScore != 0 {
			averageScore = "\n**Average Score:  **" + strconv.Itoa(graphqlResponse.Media.AverageScore) + "%"
		} else {
			averageScore = "\n**Mean Score:  **" + strconv.Itoa(graphqlResponse.Media.MeanScore) + "%"
		}
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       color,
			Description: subtitle + format + description + season + episodes + averageScore + airingTime,
			URL:         "https://anilist.co/anime/" + strconv.Itoa(graphqlResponse.Media.ID),
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "Genres",
					Value:  genres,
					Inline: false,
				},
			},

			Image: &discordgo.MessageEmbedImage{
				URL: graphqlResponse.Media.CoverImage.Large,
			},
			Title: title,
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	if strings.HasPrefix(m.Content, "?aniperson ") {
		clipped := strings.Replace(m.Content, "?aniperson ", "", 1)
		graphqlClient := graphql.NewClient("https://graphql.anilist.co")
		graphqlRequest := graphql.NewRequest(`
			{
				Staff(search: "` + clipped + `", sort: SEARCH_MATCH	) {
					id
					gender
					age
					primaryOccupations	
					dateOfBirth {
						year
						month
						day
					}
					name {
						full
					}
					image {
						large
					}

					characters(sort: FAVOURITES_DESC, page: 1, perPage: 3 ) {
						nodes {
							id
							name {
								full
							}
							media(sort: POPULARITY_DESC) {
								nodes {
									id
									title {
										romaji
										english
									}
								}
							}
						}
					}
				}
			}
		`)
		var graphqlResponse AniStaffData
		if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
			log.Println(err.Error())
			s.ChannelMessageSend(m.ChannelID, "Person not found!")
			return
		}

		var occupations string
		for i, s := range graphqlResponse.Staff.PrimaryOccupations {
			if i == 0 {
				occupations += s
			} else {
				occupations += ", " + s
			}
		}
		if occupations != "" {
			occupations = "*" + occupations + "*\n"
		}

		var birth string
		if graphqlResponse.Staff.DateOfBirth.Day != 0 {
			birth = "\n**Birth: **" + strconv.Itoa(graphqlResponse.Staff.DateOfBirth.Day) + "/" + strconv.Itoa(graphqlResponse.Staff.DateOfBirth.Month) + "/" + strconv.Itoa(graphqlResponse.Staff.DateOfBirth.Year)
		} else {
			birth = ""
		}

		var age string
		if graphqlResponse.Staff.Age != 0 {
			age = "\n**Age: **" + strconv.Itoa(graphqlResponse.Staff.Age)
		} else {
			age = ""
		}

		var gender string
		if graphqlResponse.Staff.Gender != "" {
			gender = "\n**Gender: **" + graphqlResponse.Staff.Gender
		} else {
			gender = ""
		}

		var roles string
		for _, s := range graphqlResponse.Staff.Characters.Nodes {
			roles += "[" + s.Name.Full + "](https://anilist.co/character/" + strconv.Itoa(s.ID) + ") "
			if s.Media.Nodes[0].Title.English != "" {
				roles += "[(" + s.Media.Nodes[0].Title.English
			} else {
				roles += "[(" + s.Media.Nodes[0].Title.Romaji
			}
			roles += ")](https://anilist.co/anime/" + strconv.Itoa(s.Media.Nodes[0].ID) + ")\n"
		}

		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0xFFFFFF,
			Description: occupations + birth + age + gender,
			URL:         "https://anilist.co/staff/" + strconv.Itoa(graphqlResponse.Staff.ID),
			Image: &discordgo.MessageEmbedImage{
				URL: graphqlResponse.Staff.Image.Large,
			},
			Title: graphqlResponse.Staff.Name.Full,
		}
		if roles != "" {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:  "\n\nCharacter Roles",
				Value: roles,
			})
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}

//Handler doesn't actually detect reactions, not sure why

func ReactHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	log.Println("react")

	if r.UserID == s.State.User.ID {
		return
	}
	log.Println(r.Emoji.Name)
	if r.Emoji.Name == "📌" {

		reactionMessage, _ := s.ChannelMessage(r.ChannelID, r.MessageID)
		for _, x := range reactionMessage.Reactions {
			if x.Emoji.Name == "📌" && x.Count > 1 {
				s.ChannelMessagePin(r.ChannelID, r.MessageID)

			}
		}

	}
}
