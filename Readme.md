## Informações Bruno

Primeiro criar uma imagem com o dockerfile.

>docker build -t nome_image .
>docker build -t desafio-k8s-03-go .

Depois criar um container com a imagem gerada no passo 1.
 
>docker run -it --name nome_container -p 8000:8000 nome_da_imagem
>docker run -it --name desafio -p 8000:8000 desafio-k8s-03-go

docker rm $(docker ps -a -q) -f
docker rmi $(docker image ls -a -q)
