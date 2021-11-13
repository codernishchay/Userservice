package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	randomState = "random"
	// TODO: randomize it
)

func HandleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(randomState)
	http.Redirect(c.Writer, c.Request, url, http.StatusTemporaryRedirect)
}

func HandleCallback(c *gin.Context) {
	if c.Request.FormValue("state") != randomState {
		fmt.Println("state is not valid")
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, c.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("Could not get token ", err.Error())
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return

	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get  request", err.Error())
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response ", err.Error())
		http.Redirect(c.Writer, c.Request, "/", http.StatusTemporaryRedirect)
		return

	}
	fmt.Fprintf(c.Writer, "Response: %s", content)
}
