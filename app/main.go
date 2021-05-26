package main

import (
	"fmt"
	"time"

	"github.com/denautonomepirat/leihs"
)

func main() {
	fmt.Println("morning")
	l := leihs.NewLeihs(&leihs.Config{
		Token:    "2JO6OZTGXHWNPRBF6P75HEWMUNZWZG7C",
		LeihsURL: "https://leihs.hopto.org",
	})
	ts := time.Now()
	/*

		users, err := l.FindUsers()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Found %v users\n", len(*users))

		user, err := l.FindUser("thk@id.aau.dk")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("User ID: \"%s\"\n", user.ID)
	*/

	a, err := l.AuthenticationSystemByName("AAU SSO")
	if err != nil {
		fmt.Println("Auth system not found")
		s := ""
		fmt.Scanln(&s)
		if s == "Y" {
			na := &leihs.AuthenticationSystem{
				Name:                  "AAU SSO",
				ExternalSignInURL:     "https://leihs.hopto.org:8282/login",
				InternalPrivateKey:    "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEIHIudvH7GFCSU6Jzg8bttlqvAEeCho9x+31Ibph8rru6oAoGCCqGSM49\nAwEHoUQDQgAEKTFu0lm7P63LuNSlqF3O7yYmpVot5ft7K3AUGH1OvVHnMPGepLgP\nWC07UIH4RMGtkMcnurzNUdyjTs17/bnkUg==\n-----END EC PRIVATE KEY-----",
				SendEmail:             true,
				ShortcutSignInEnabled: true,
				ExternalPublicKey:     "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEmmCU7vfc55KpPohsaSn2pptM5x1b\naKPccv8zF/8i65Rrv53Ja16s0Z9YNKJ3/Ztiq+CJ6JkBgVCd9cybIoCO7Q==\n-----END PUBLIC KEY-----\n",
				Type:                  "external",
				SignUpEmailMatch:      ".aau.dk",
				InternalPublicKey:     "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEKTFu0lm7P63LuNSlqF3O7yYmpVot\n5ft7K3AUGH1OvVHnMPGepLgPWC07UIH4RMGtkMcnurzNUdyjTs17/bnkUg==\n-----END PUBLIC KEY-----",
				Priority:              3,
				ID:                    "aau",
				ExternalSignOutURL:    "https://leihs.hopto.org:8282/logout",
				Enabled:               true,
				CreatedAt:             time.Now(),
				UpdatedAt:             time.Now(),
			}
			l.AddAuthenticationSystem(na)

		}
	} else {
		fmt.Printf("Authentication System ID: \"%s\"\n", a.ID)
	}
	u := &leihs.User{
		Email:                 "",
		AccountEnabled:        true,
		PasswordSignInEnabled: false,
	}

	u, err = l.AddUser(u)
	if err != nil {
		fmt.Println(err)
	}
	grPtr := &leihs.Group{}

	groups, err := l.FindGroups()
	if err != nil {
		fmt.Println(err)
	}
	for i, group := range *groups {
		if group.Name == "AAU students" {
			fmt.Printf("%d : %s\n", i, group.ID)
			grPtr = &group
		}
	}

	l.AddToGroup(user, grPtr)

	td := time.Now().Sub(ts).String()
	fmt.Printf("det tog %s\n", td)

}
