package main

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Crie uma nova sessão do DiscordGo.
	sess, err := discordgo.New("Bot MTA5OTcyOTYzMjM4MDYwMDMyMQ.GKbg4f.To2kuzjkSdg3m37ircV83bF7Hrm1SHSvE1Oe-Y")
	if err != nil {
		fmt.Println("Erro ao criar a sessão do DiscordGo:", err)
		return
	}

	// Adicione um ouvinte para a mensagem criada (MessageCreate).
	sess.AddHandler(messageCreate)

	// Abra a conexão da sessão com o Discord.
	err = sess.Open()
	if err != nil {
		fmt.Println("Erro ao abrir a conexão da sessão:", err)
		return
	}

	// Mantenha o programa em execução.
	fmt.Println("Bot está em execução. Pressione CTRL-C para sair.")
	<-make(chan struct{})
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignorar mensagens do próprio bot e mensagens que não começam com "^".
	if m.Author.ID == s.State.User.ID || m.Content[0] != '^' {
		return
	}

	// Processar comandos.
	switch m.Content[1:] {
	case "oi":
		helloCommand(s, m)
	case "penis":
		penisCommand(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "Comando inválido. Tente ^oi ou ^penis.")
	}
}

func helloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignorar mensagens de bots.
	if m.Author.Bot {
		response := fmt.Sprintf("Vá se fuder %s seu bot de merda", m.Author.Username)
		s.ChannelMessageSend(m.ChannelID, response)
		return
	}

	// Selecionar uma saudação aleatória.
	hellos := []string{
		fmt.Sprintf("Olá %s meu lindo", m.Author.Username),
		fmt.Sprintf("Olá %s seu FEDIDO", m.Author.Username),
		fmt.Sprintf("Olá %s seu BOIOLA", m.Author.Username),
		fmt.Sprintf("Que foi o %s seu esquisito?", m.Author.Username),
	}
	chooser := rand.Intn(len(hellos))
	response := hellos[chooser]

	s.ChannelMessageSend(m.ChannelID, response)
}

func penisCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignorar mensagens de bots.
	if m.Author.Bot {
		response := "```\n _ __    ___  _ __  (_) ___\n| '_ \\  / _ \\| '_ \\ | |/ __|\n| |_) ||  __/| | | || |\\__ \\\n| .__/  \\___||_| |_||_||___/\n|_|```"
		s.ChannelMessageSend(m.ChannelID, response)
		return
	}

	response := fmt.Sprintf("Tá querendo pika? %s seu bot de merda", s.State.User.Username)
	s.ChannelMessageSend(m.ChannelID, response)
}
