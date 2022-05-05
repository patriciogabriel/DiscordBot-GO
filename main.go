package main

import (
	"flag"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	prefixChar string
)

func init() {
	flag.StringVar(&prefixChar, "pfx", ".", "Prefix that the bot uses for commands")
}

func main() {
	token := ""

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating discord session", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection,", err)
		return
	}

	for {
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	c := strings.ToLower(m.Content)
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == ".help" {
		s.ChannelMessageSendReply(m.ChannelID, "\n.random\n.time\n.calc", m.Reference())
	} else if m.Content == ".random" {
		rand.Seed(time.Now().Unix())
		random := []string{
			"A persistência é o caminho do êxito.",
			"Que venham dias melhores",
			"Hoje vai dar tudo certo.",
			"Jamais desista de ser feliz.",
		}
		n := rand.Int() % len(random)
		s.ChannelMessageSendReply(m.ChannelID, random[n], m.Reference())
	} else if m.Content == ".time" {
		timenow := time.Now()
		s.ChannelMessageSendReply(m.ChannelID, timenow.Format(time.Kitchen), m.Reference())
	} else if strings.HasPrefix(c, prefixChar+"calc") {
		if m.Content == ".calc" {
			s.ChannelMessageSendReply(m.ChannelID, "Digite um número válido", m.Reference())
		} else {
			re := regexp.MustCompile("[0-9]+")
			re2 := re.FindAllString(m.Content, -1)
			op := regexp.MustCompile(`[()+\-*/.]`)
			op2 := op.FindAllString(m.Content, -1)
			mathop := op2[len(op2)-1]

			n1string := re2[0]
			n2string := re2[1]
			n1int, err := strconv.Atoi(n1string)
			n2int, err := strconv.Atoi(n2string)
			if err != nil {
				fmt.Println(err)
			}

			if mathop == "*" {
				calcInt := n1int * n2int
				calcStr := strconv.Itoa(calcInt)
				s.ChannelMessageSendReply(m.ChannelID, calcStr, m.Reference())
			} else if mathop == "+" {
				calcInt := n1int + n2int
				calcStr := strconv.Itoa(calcInt)
				s.ChannelMessageSendReply(m.ChannelID, calcStr, m.Reference())
			} else if mathop == "-" {
				calcInt := n1int - n2int
				calcStr := strconv.Itoa(calcInt)
				s.ChannelMessageSendReply(m.ChannelID, calcStr, m.Reference())
			} else if mathop == "/" {
				n1float := float64(n1int)
				n2float := float64(n2int)
				calcInt := n1float / n2float
				calcStr := fmt.Sprintf("%v", calcInt)
				s.ChannelMessageSendReply(m.ChannelID, calcStr, m.Reference())
			} else {
				s.ChannelMessageSendReply(m.ChannelID, "Operador matemático inválido.\n\n +  =  Adição\n -   =  Subtração\n /  =  Divisão\n *   =  Multiplicação", m.Reference())
			}
		}

	}

}
