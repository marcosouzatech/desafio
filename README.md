
# DESAFIO API ITENS #

## 1 - FRAGMENTAÇÃO DE TRÁFEGO
 
Para conseguir concluir o primeiro desafio, me dediquei a terminar parte do curso de Golang que estava cursando e colocar meus conhecimentos em prática e nesse processo algumas adaptações foram efetuadas para atender a necessidade do desafio proposto.

**1.1 - API ITENS**

- Contruir uma aplicação em Golang que retornasse uma lista de itens
- Adaptação para a necessidade do desafio
- Implementação de Logs
- String de Banco de dados para RDS
- Criação da Imagem Docker
- Criação dos arquivos manifestos kubernetes
  
**1.2 - Provisionamento Cluster EKS**
- Criação e adaptação dos arquivos terraform
- Utilização de máquinas spot
- Instalação de Add-Ons ( Kube-Proxy, CoreDNS, CNI-Plugin, EBS CSI Driver)
- Provisionamento do kubecost
- Provisionamento do Aplication LoadBalancing
- Provisionamento do TargetGroup e regras listners ( MLA/MLB/MLC/MLM )
  
**1.3 - Criação do Banco de dados RDS**
- Criação das tabelas

> Para fins de demonstração foi criado apenas 1 RDS contendo as 4 tabelas porém essa solução não necessáriamente precisa representar um ambiente real, dependendo da utilização é interessante fragmentar a base de dados para evitar problemas de performance. 

No primeiro momento já tinha intuito de utilizar AWS como cloud provider e API Internet Gateway para fazer o roteamento de tráfego com base no header, após algumas pesquisas identifiquei que seria necessário utilizar uma função lambda para ler o Header e efetuar o forwarding da requisição para o ambiente desejado.

Por este motivo foi decidido colocar a a lógica da fragmentação `site_id` de origem no header no `ApplicationLoadBalancer, e evitar custos do Lambda, mas como alternativas poderíamos explorar outro Gateway como Kong ou KrakenD, um proxy reverso: nginx ou service mesh istio ou algo similar. 

> Foi provisionado um targetGroup para representar cada ambiente, essa etapa me livrou de configurar ZonaDNS e Ingress de cada aplicação, em um ambiente real seria encaminhado para outro loadbalancer ou endpoint.

Para representar a solução foi desenhado o diagrama de arquitetura:

![DIAGRAMA DE ARQUITETURA ](https://github.com/marcosouzatech/desafio/blob/main/img/diagrama.png)


Podemos efetuar requisições para o ambiente da API Itens:
```
curl -H "site_id":"MLB" https://0p4ko8ndke.execute-api.us-east-1.amazonaws.com/api/itens
```
Obs: todos os ambientes estão ativos, e podem ser acessados pelo endpoint público seguindo seu site_id especifico.

Resposta da Chamada:
```
[{"id":1,"Product":"Camiseta","nome":"Brasil 12","Categoria":"Brasil 12","CriadoEm":"2023-07-13T23:44:55Z"},{"id":3,"Product":"Camisa","nome":"Brasil 13","Categoria":"Brasil","CriadoEm":"2023-07-13T23:45:11Z"},{"id":4,"Product":"Camisa1","nome":"Brasil 113","Categoria":"Brasil1","CriadoEm":"2023-07-13T23:45:19Z"},{"id":5,"Product":"Camisa11","nome":"Brasil 1113","Categoria":"Brasil11","CriadoEm":"2023-07-13T23:45:24Z"},{"id":6,"Product":"Camisa111","nome":"Brasil 11113","Categoria":"Brasil111","CriadoEm":"2023-07-13T23:45:31Z"},{"id":7,"Product":"Camisa1111","nome":"Brasil 111113","Categoria":"Brasi1l111","CriadoEm":"2023-07-13T23:45:42Z"}]

```

## 2 - LIMITAÇÃO DE TRÁFEGO

Apesar do Api Internet Gateway não funcionar muito bem para validar as regras no header, utilizei API Internet Gateway para expor endpoint público e criar limitação de tráfego com base no desafio.
  
A regra limita o tráfego em 15 requisições por segundo e assim consegue limitar o tráfego a 1000 requisições por minuto conforme descrito no desafio. 

Ao executar a chamada na api podemos verificar ao atingir o limite recebemos 429 TOO MANY REQUESTS.
  
![TESTE DE CARGA POSTMAN](https://github.com/marcosouzatech/desafio/blob/main/img/teste_postman.png)

![TESTE DE CARGA HEY - EXECUÇÃO 1MINUTO](https://github.com/marcosouzatech/desafio/blob/main/img/teste_hey.png)
  
## 3 - MONITORAMENTO

Para o monitoramento ativo do cluster foi provisionado o NewRelic como principal foco de observabilidade, é possivel verificar métricas de recursos de infraestrutura, traces, logs e auto instrumentação com OpenTelemetry. 

Criado um dashboard de exemplo do que podemos entender como necessário para monitorar nossos ambientes, além do monitoramento triviais de Recursos (Memória/CPU/REDE) podemos visualizar métricas integradas com CloudWatch e Prometheus. 

A utilização do NewRelic facilitará ao identificar padrões e alertas que serão criados automaticamente podendo ser notificados em canais de atendimento. 

Alertas personalizados também são possíveis de se criar, podemos pensar no aumento do Error Rate, tanto do lado do API Gateway ou no NewRelic para buscar um monitoramento ativo.

Maiores integrações com newRelic podem ser efetuadas. 

Como alternativa ao newRelic, podemos utilizar o Dynatrace, Datadog, Elastic Stack entre outras soluções openSource. 

[DASHBOARD - ARGENTINA - VERSÃO PDF ](https://github.com/marcosouzatech/desafio/blob/main/img/dashboard_argentina.pdf)

## CUSTOS 

Foi implementado o kubecost para auxiliar no controle dos custos do ambiente, ele trará insigths interessantes para economia e eficiência do ambiente da API de Itens. 
  
![KUBECOST OVERVIEW](https://github.com/marcosouzatech/desafio/blob/main/img/kubecost.png)

## GIT ACTIONS

Foi implementado um workflow de exemplo para efetuar build da imagem e publish no repositório dockerhub, não foi adicionado o processo de deploy pois nosso cluster não está exposto url públicas e as não foram provisionados runners self-hosted. 

### Considerações Finais
  
Este challenge foi realmente desafiador, trouxe conteúdo denso de vários aspectos envolvendo infraestrutura, desenvolvimento e operações, com certeza temos vários pontos a melhorar principalmente visando IAC, acredito que consegui demonstrar vários aspectos importantes que foram me solicitados. 

> Não foi levado em conta nenhum modelo padrão de versionamento/gitflow


#### REFERÊNCIAS

- Hey- Load to a web application.
https://github.com/rakyll/hey
- Base API - Golang
https://devbook.com.br/curso-golang/
- APM - NEW RELIC
https://newrelic.com/pt
- Kubecost - Monitor and manage Kubernetes spend
https://github.com/kubecost/cost-analyzer-helm-chart
- EVIDÊNCIAS DO PRJETO.
https://github.com/marcosouzatech/desafio/blob/main/img/

##### CONTATOS
- Linkedin: (https://www.linkedin.com/in/marcosouzatech/)
- Youtube: https://www.youtube.com/@maistalkmenosshow
- Email: marcos.souza@luvtech.com.br