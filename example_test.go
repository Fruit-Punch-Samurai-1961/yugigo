package yugigo

import (
	"fmt"
	"log"
)

func ExampleGetCards() {
	//define the search query
	search := CardSearch{
		FName:     "shark",
		Level:     "4",
		Attribute: "water",
		Sort:      "Atk",
	}

	//call the GetCards Method
	cards, _, err := GetCards(search)
	if err != nil {
		log.Fatal(err)
	}

	//Print out the Name and Attack of the cards
	for _, card := range cards {
		fmt.Printf(" Name: %v\n Attack: %v\n", card.Name, card.Atk)
	}
	//OutPut://Name: Double Fin Shark
	//Attack: 1000
	//Name: Double Shark
	//Attack: 1200
	//Name: Lantern Shark
	//Attack: 1500
	//Name: Right-Hand Shark
	//Attack: 1500
	//Name: Wind-Up Shark
	//Attack: 1500
	//Name: Buzzsaw Shark
	//Attack: 1600
	//Name: Saber Shark
	//Attack: 1600
	//Name: Spear Shark
	//Attack: 1600
	//Name: Abyssal Kingshark
	//Attack: 1700
	//Name: Hammer Shark
	//Attack: 1700
	//Name: Metabo-Shark
	//Attack: 1800
	//Name: Bahamut Shark
	//Attack: 2600
	//Name: Number 37: Hope Woven Dragon Spider Shark
	//Attack: 2600
	//Name: Number 32: Shark Drake
	//Attack: 2800
	//Name: Number C32: Shark Drake Veiss
	//Attack: 2800
}

func ExampleGetCardFromCardSet() {
	//Call the GetCardFromCardSet method
	data, err := GetCardFromCardSet("SDY-045")
	if err != nil {
		log.Fatal(err)
	}
	//Print out the info received
	fmt.Printf(" Set Name: %v\n Set Code: %v\n Name: %v\n ID: %v\n Set Price: %v\n Set Rarity: %v\n",
		data.SetName, data.SetCode, data.Name, data.ID, data.SetPrice, data.SetRarity)

	//Output:  Set Name: Starter Deck: Yugi
	// Set Code: SDY-045
	// Name: Yami
	// ID: 59197169
	// Set Price: 1.18
	// Set Rarity: Common
}

func ExampleGetAllArchetypes() {
	//Call the GetAllArchetypes method
	data, err := GetAllArchetypes()
	if err != nil {
		log.Fatal(err)
	}
	//Print out the first five names
	for i := 0; i <= 5; i++ {
		fmt.Println(data[i].ArchetypeName)
	}

	//Output:@Ignister
	//ABC
	//Abyss Actor
	//Adamancipator
	//Aesir
	//Aether
}

func ExampleGetAllCardSets() {
	//Call the GetAllCardSets Method
	data, err := GetAllCardSets()
	if err != nil {
		log.Fatal(data)
	}
	//Print the information for the first three card sets
	for i := 0; i <= 3; i++ {
		cardset := data[i]
		fmt.Printf(" Set Name: %v\n Set Code: %v\n Number of Cards: %v\n ", cardset.SetName, cardset.SetCode, cardset.NumOfCards)
	}

	//Output: Set Name: 2-Player Starter Deck: Yuya & Declan
	// Set Code: YS15
	// Number of Cards: 42
	//  Set Name: 2013 Collectible Tins Wave 1
	// Set Code: CT10
	// Number of Cards: 9
	//  Set Name: 2013 Collectible Tins Wave 2
	// Set Code: CT10
	// Number of Cards: 9
	//  Set Name: 2014 Mega-Tin Mega Pack
	// Set Code: MP14
	// Number of Cards: 247
}

func ExampleGetDataBaseInfo() {
	//Get simple information on the database
	info, err := GetDataBaseInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" Database Version: %v\n Last Update: %v\n", info.DatabaseVersion, info.LastUpdate)

	//Output: Database Version: 5.86
	// Last Update: 2020-12-29 19:31:19
}

func ExampleGetCardWithFullUrl() {
	//Define the search query
	//NOTE: Since we didn't specify the Nums query, the default 50 will be used. We can check that by calling cards.Nums.
	search := CardSearch{
		CardSet: "metal raiders",
		Offset:  "10",
	}
	//Call the GetCards Method initially
	cards, meta, err := GetCards(search)
	if err != nil {
		log.Fatal(err)
	}
	//Print information for the first three cards for the first three pages returned
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			card := cards[j]
			fmt.Printf(" Name: %v\n Race: %v\n", card.Name, card.Race)
		}
		cards, meta, err = GetCardWithFullUrl(meta.NextPage)
	}
	//Output:Name: Big Eye
	// Race: Fiend
	// Name: Black Skull Dragon
	// Race: Dragon
	// Name: Blackland Fire Dragon
	// Race: Dragon
	// Name: Jirai Gumo
	// Race: Insect
	// Name: Kaminari Attack
	// Race: Thunder
	// Name: Kazejin
	// Race: Spellcaster
	// Name: Seven Tools of the Bandit
	// Race: Counter
	// Name: Shadow Ghoul
	// Race: Zombie
	// Name: Share the Pain
	// Race: Normal
}
