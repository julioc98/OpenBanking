package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/openbanking/cmd/restapi/handlers"
	"github.com/julioc98/openbanking/pkg/env"
	"github.com/julioc98/openbanking/pkg/gateways"
	"github.com/julioc98/openbanking/pkg/middleware"
	"github.com/julioc98/openbanking/pkg/services"
)

func handlerHi(w http.ResponseWriter, r *http.Request) {
	msg := "Ola, Seja bem vindo(a) ao openbankinghacka!!"
	log.Println(msg)
	w.Write([]byte(msg))
}

func main() {
	r := mux.NewRouter()

	// Choose the folder to serve
	staticDir := "/static/"

	// Create the route
	r.PathPrefix("/callback").
		Handler(http.StripPrefix("/callback", http.FileServer(http.Dir("."+staticDir))))

	r.Use(middleware.Logging)

	authURL := env.Get("AUTH_URL", "https://auth.obiebank.banfico.com")
	baseURL := env.Get("BASE_URL", "https://gw-dev.obiebank.banfico.com/obie-aisp/v3.1/aisp")
	redirectURI := env.Get("REDIRECT_URI", "https://openbankinghacka.herokuapp.com/callback")
	client := http.DefaultClient

	obiebankGateway := gateways.NewObiebank(authURL, baseURL, redirectURI, client)

	user := env.Get("CLIENT_ID", "PSDBR-NCA-AISP01")
	pass := env.Get("PASSWORD", "senha123")
	privateKey := getPrivateKey()

	authService := services.NewAccountAuth(obiebankGateway, user, pass, privateKey)

	authHandler := handlers.NewAuthHandler(authService)

	r.HandleFunc("/auth", authHandler.Auth).Methods("GET")
	r.HandleFunc("/webhook", authHandler.Callback).Methods("GET")

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	port := env.Get("PORT", "5001")
	log.Printf(`%s listening on port: %s `, env.Get("APP", "openbankinghacka"), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func getPrivateKey() string {
	return env.Get("PRIVATE_KEY", `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCf/crE4dptc21t
9IQ6q/QS/CNEVk6DezEp7p3UHpDV2Q8w/IdjDI4rEgGT+c93ijfq6rz50SUU6zZL
CDbEj/7h/YYx7x1nlEUBmz14Q04IML5bucxhRXsr+cfBynAlbFeyrRW6iQf4ynvV
rjBvyyfSL3DSggFNMZzJ3kR5nf6mLYeu3xjd9mbqVan9QJe0PcozX7FEO8ihYBAz
d/zXLYJOM1sO9xZotyjR5psgGLh+Xz3cO2gyHtzKPlY2I8FjgxoGMJMlBN2sNA7Z
IXDJWnZWyxR3Jd4RYSiQV/VLBOL/avXTOqt2eA03pD8d/X/hhJiAd14jQ0AYtGeJ
U9SaMgDNAgMBAAECggEAM4b2bjBbO9LoFHiiuY613gNsjE7LnJbpf5rFJBLwDJ+K
v/kk1WauxvpWncTf550RY0xUrpSIP9N9Oe8cTHQf38LaHGzpaHmdO+Y2huwOTp78
P+h4BX/uKnyYtDYxpivdlsMd82S8t4jnFyuxl9+zJIN476NFLgpjd7RpE04qPHR8
BsjvukOWvNRVvOGzcMU8ffoVBN4wNmvzcDwn1JQ3Lyq+tw2qhSxzCL6BfsiYh9j/
W0+rMX/OQIJz5dIbpZ4f2FIw+yUIvWloyebbwiF15J762z6g7LufdBgr4bmM4ELg
/Z1F9TxaVu+uq3mVxys5eW1YJe6/Nmu9Rgm+oDCU1QKBgQDmDj2kra4BP3TwUHdf
Dh/ixslBtah0ZVLFvEhVjbShDQsHwApbhbyJJRDTkY1HBczCDo15RDVbSCach6SZ
YJuQEVl1OQfywuk7By6omI/KrCM4ov2v2bf+JbuqxjGCHgTtB1P/RcLx836rhnWa
Js14dvhFu0Ck7obkjj2qemmkQwKBgQCyCMezlTCHIze6w47wZ58k3Y8u0I2RIyyM
VYvObfFL2UzuY1Upwb8ttjQ4/inv3uU8L02HSapkKkKn6OiJZVTZW4JVcV3pfKEx
fHIKSsGrue1f3vUtaZy4KvCDTnvmLUpqEAAZecE0Nh0XWyB/UVp9mhjFVWfzmuPh
1GuN9MJ9rwKBgQCzNzjSRvKcykBgzW6QwEIaud0isU9PjXdTzv9Slpe2NqD3IqVu
8toSxKs9BdBXGa+PJSMU6wvd1nEt04VobpgBPWLBLPKCLVDfyRKSCHdL3Zl6j46t
JSBufhqaSNdck+ImfGT1IfVh4tw05wRKWBwM0jFKsTsEwUSYXC6x1bbiXQKBgECs
799NU1PEd3phkIvFGQtLcbiQCt2u6YARk7hqOD5VspzneQiyWcFBb7dEnfeGAcDb
bk63dC7vK0fUVKWVKj3MAI0JohQwMl7H1qXmgnTgFlu9o1PcChLdhoItANWdnmrp
ZR/cG1PcVLUnZaba5wS59kW5wQm+OwrPIENxpzYBAoGALoS0chN6GrgggST13N7P
C1Rdd9KTzDF6uls03N8I/lPaKc73ExtJPE7AnO27qVWc6IlEZnrbWlfQamjQTCGP
PXko2ncCpOGNn0RrGSxwdLuXQHyq8ey8YVx5ZtGMPe5CsEozEX+amxrMqLZDD4jm
JmWDL1i3WVciBrkscCgti1c=
-----END PRIVATE KEY-----`)
}
