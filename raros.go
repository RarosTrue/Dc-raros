package main

import(
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"

)

func main(){
	


	if discordgo.MessageCreate == "^oi"{
		helloComand()
	}

	if discordgo.MessageCreate == "^penis"{
		penisCommand()
	}





}

func helloComand(s *discordgo.Session m *discordgo.MessageCreate){
	

	if m.Author.Bot{	
		reponse := fmt.Sprintf("Vá se fuder %s seu bot de  merda", bot.Username)
		s.ChannelMessageSend(m.ChannelID, reponse)
	return
	}

	

	if m.Author.User{
		hellos = [4]String{("Olá %s meu lindo", user.Username), 
		("Olá %s seu FEDIDO", user.Username), 
		("Olá %s seu BOIOLA", user.Username), 
		("Que foi o %s seu esquisito?", user.Username)}

		defer chooser := rand.Intn(len(hellos))

		response := hellos[chooser]

		s.ChannelMessageSend(m.ChannelID, reponse)


	}


}




func penisCommand(){

		if m.Author.Bot {

			response := fmt.println("

				```
				 _ __    ___  _ __  (_) ___
				| '_ \  / _ \| '_ \ | |/ __|
				| |_) ||  __/| | | || |\__ \
				| .__/  \___||_| |_||_||___/
				|_|

				```

		")


		} else m.Author.Bot{


			response := fmt.println("tá querendo pika? %s seu bot de merda", s.Username)
		
		}

		m.ChannelMessageSend(m.ChannelID, response)

		
}