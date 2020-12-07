# openbanking
Hackathon Repository Backend Platform Code  -  https://www.openbankinghacka.com/

## Pode acessar a API em https://openbankinghacka.herokuapp.com/

### Resgatar URL para logar no Obiebank
```
https://openbankinghacka.herokuapp.com/auth
```
* pegue a URL e abra

### O callback recebe as informações do usuario
```
https://openbankinghacka.herokuapp.com/callback
```
* sera redirecionad para a pagina que pega o Access Token

### Webhook é quem pega o codigo temporario para retornar o Token

```
https://openbankinghacka.herokuapp.com/webhook
```
* e assim devolvemos a pagina de sucesso ou erro

## Para rodar local

### Pré-requisitos

* [Golang](https://github.com/golang/go)

### Como rodar localmente?**

Baixe o repositório, entre no diretório e rode o comando:

```
go run cmd/restapi/main.go
```
Depois acesse a url
```
http://localhost:5001/
```



