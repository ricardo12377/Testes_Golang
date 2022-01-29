# Testes_Golang
Realizando testes com Golang.

#Para testar todos os pacotes em um diretório.
  go test ./...

#Resultado ded testes com mais descrições.
  go test -v

#Para rodar testes de modo parelelo às outras.
  inserir t.Parallel dentro da função, recomendo usar essa função em mais de uma para fazer sentido.

#Para ver a porcentagem de código coberto por testes.
  go test --cover
  
#Criar arquivo com detalhes do teste.
  go test --coverfile resultado.txt
  
#Para ler o arquivo do teste.
  go tool cover --func=resultado.txt
  
#Informações do teste, com detalhamento das linhas em HTML.
  go tool cover --html=resultado.txt
  
