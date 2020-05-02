# Desafio HPA - Bruno Leal

[x] - Criar deployment com nome go-hpa  
[x] - Criar service com nome go-hpa  
[x] - Criar nomeDaImage com nome go-hpa  
[x] - Processo de testes sendo executado ao realizar Pull Request(Trigger-GCP)  
[x] - Processo De HPA - deployment executando(*image: lealbruno/codeedu-desafio-go-hpa*)  


## Docker HUB
Imagem abaixo referente ao exercicio inicial de disponibilizar o laravel
- **[DockerHUB](https://hub.docker.com/repository/docker/lealbruno/codeedu-desafio-go-hpa)** 
- docker push lealbruno/codeedu-desafio-go-hpa


## Colocar aqui as imagens geradas


## Informações Bruno

Primeiro criar uma imagem com o dockerfile.

>docker build -t nome_image .
>docker build -t desafio-k8s-03-go .

Depois criar um container com a imagem gerada no passo 1.
 
>docker run -it --name nome_container -p 8000:8000 nome_da_imagem
>docker run -it --name desafio -p 8000:8000 desafio-k8s-03-go

docker rm $(docker ps -a -q) -f
docker rmi $(docker image ls -a -q)

**Vale ressaltar que alterei o processo de build com a imagem optimezed, lembrar dessa alteração,
meu cloudbuild fez referencia ao container de test, pois nele a versao do go n é optimezed..consigo rodar os testes, foi por isso. Mas a imagem do DockerHUB é a optmized, e a mesma utilizada pelo K8s.
Acredio que no proximo desafio irei aprender a integrar melhor essas etapas.
