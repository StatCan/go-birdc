[(Français)](#le-nom-du-projet)

## go-birdc

This is a GO library for working with the BIRD internet routing daemon through the provided unix domain socket. It is primarily an abstraction layer allowing users to communicate with a BIRD daemon without needing to directly talk with the BIRD daemon or spawn goroutines to issue `birdc` commands and parsing the output.

## Usage

Install the library as a dependency

```bash
go get -u github.com/StatCan/go-birdc@latest
```

Then start using it in your application

```golang
import "github.com/StatCan/go-birdc"

func main() {
	b := New() // optionally you can specify a path to the unix socket endpoint (defaults to: /run/bird/bird.ctl)
	resp, replyCode, err := b.ShowStatus()
	if err != nil {
		log.Fatalf("something went wrong. %s", replyCode)
	}

	resp, replyCode, err := b.RestartProtocol("bgp")
	if err != nil {
		log.Fatalf("somthing went wrong. %s", replyCode)
	}

	fmt.Println(string(resp))
}

```

## Related Documentation

- [BIRD](https://gitlab.nic.cz/labs/bird/-/tree/v2.14/doc)

### How to Contribute

See [CONTRIBUTING.md](CONTRIBUTING.md)

### License

Unless otherwise noted, the source code of this project is covered under Crown Copyright, Government of Canada, and is distributed under the [MIT License](LICENSE).

The Canada wordmark and related graphics associated with this distribution are protected under trademark law and copyright law. No permission is granted to use them outside the parameters of the Government of Canada's corporate identity program. For more information, see [Federal identity requirements](https://www.canada.ca/en/treasury-board-secretariat/topics/government-communications/federal-identity-requirements.html).

______________________

## go-birdc

Voici une bibliothèque GO pour travailler avec le démon de routage internet BIRD via le socket de domaine Unix fourni. Il s'agit principalement d'une couche d'abstraction permettant aux utilisateurs de communiquer avec un démon BIRD sans avoir besoin de dialoguer directement avec le démon BIRD ou de lancer des goroutines pour émettre des commandes birdc et analyser la sortie.

## Utilisation

Installer la bibliothèque en tant que dépendance.

```bash
go get -u github.com/StatCan/go-birdc@latest
```

Ensuite, commencez à l'utiliser dans votre application.

```golang
import "github.com/StatCan/go-birdc"

func main() {
	b := New() // optionally you can specify a path to the unix socket endpoint (defaults to: /run/bird/bird.ctl)
	resp, replyCode, err := b.ShowStatus()
	if err != nil {
		log.Fatalf("something went wrong. %s", replyCode)
	}

	resp, replyCode, err := b.RestartProtocol("bgp")
	if err != nil {
		log.Fatalf("somthing went wrong. %s", replyCode)
	}

	fmt.Println(string(resp))
}

```

## Documentation Connexe

- [BIRD](https://gitlab.nic.cz/labs/bird/-/tree/v2.14/doc)

### Comment contribuer

Voir [CONTRIBUTING.md](CONTRIBUTING.md)

### Licence

Sauf indication contraire, le code source de ce projet est protégé par le droit d'auteur de la Couronne du gouvernement du Canada et distribué sous la [licence MIT](LICENSE).

Le mot-symbole « Canada » et les éléments graphiques connexes liés à cette distribution sont protégés en vertu des lois portant sur les marques de commerce et le droit d'auteur. Aucune autorisation n'est accordée pour leur utilisation à l'extérieur des paramètres du programme de coordination de l'image de marque du gouvernement du Canada. Pour obtenir davantage de renseignements à ce sujet, veuillez consulter les [Exigences pour l'image de marque](https://www.canada.ca/fr/secretariat-conseil-tresor/sujets/communications-gouvernementales/exigences-image-marque.html).
