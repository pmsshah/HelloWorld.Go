package main

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "os"
	"strings"
)

type Person struct {
        Name string
        Phone string
}

func main() {
	arg := len(os.Args)
	
	if(arg != 4) {
		fmt.Printf("USAGE: cmd[add,find,del] name phone\n")
		return
	}
	
	cmd:=os.Args[1];
	name:=os.Args[2];
	phone:=os.Args[3];
	
	fmt.Printf("cmd:%s with name:%s, phone:%s\n", cmd,name,phone)
	
        session, err := mgo.Dial("127.0.0.1")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB("test").C("people")
        err = c.Insert(
			&Person{"Al1", "+55 53 8116 9631"},
			&Person{"Al2", "+55 53 8116 9632"},
			&Person{"Al3", "+55 53 8116 9633"},
			&Person{"Al4", "+55 53 8116 9634"},
			&Person{"Al5", "+55 53 8116 9635"},
			&Person{"Cl1", "+55 53 8402 8511"},
			&Person{"Cl2", "+55 53 8402 8512"},
			&Person{"Cl3", "+55 53 8402 8513"},
			&Person{"Cl4", "+55 53 8402 8514"},
			&Person{"Cl5", "+55 53 8402 8515"})
        if err != nil {
                log.Fatal(err)
        }

        if(strings.EqualFold(cmd,"add")) {
		fmt.Printf("Performing ADD with name:%s, phone:%s\n", name,phone)
        }

        if(strings.EqualFold(cmd,"del")) {
		fmt.Printf("Performing DEL for name:%s and/or phone:%s\n", name,phone)
        }

        if(strings.EqualFold(cmd,"find")) {
        
		fmt.Printf("Performing FIND for name:%s and/or phone:%s\n", name,phone)

		result := Person{}
		err = c.Find(bson.M{"name": name}).One(&result)
		if err != nil {
			err = c.Find(bson.M{"phone": phone}).One(&result)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Phone:%s own by Name:%s", phone, result.Name)
			}
		} else {
	        	fmt.Printf("Name:%s has Phone:%s", name, result.Phone)
	        }
        }
        if(strings.EqualFold(cmd,"all")) {
        
		fmt.Printf("Performing ALL \n")

		result := Person{}
		err = c.Find(nil).All(&result)
		if err != nil {
			log.Fatal(err)
		} else {
	        	fmt.Println("All records: ", result)
	        }
        }

}
