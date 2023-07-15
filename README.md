# DESAFIO API ITENS #


## 1 - FRAGMENTAÇÃO DE TRÁFEGO

Para conseguir concluir o primeiro desafio, me dediquei a terminar parte do curso de Golang que estava cursando e colocar meus conhecimentos em prática e nesse processo algumas adaptações foram efetuadas para atender a necessidade do desafio proposto. 

1.1 - API ITENS 
- Contruir uma aplicação em Golang que retornasse uma lista de itens
- Adaptação para a necessidade do desafio
- Implementação de Logs 
- String de Banco de dados para RDS
- Criação da Imagem Docker 
- Criação dos arquivos manifestos kubernetes

1.2 - Provisionamento Cluster EKS
- Criação e adaptação dos arquivos terraform
- Utilização de máquinas spot
- Instalação de Add-Ons ( Kube-Proxy, CoreDNS, CNI-Plugin, EBS CSI Driver)
- Provisionamento do kubecost
- Provisionamento do Aplication LoadBalancing
- Provisionamento do TargetGroup e regras listners ( MLA/MLB/MLC/MLM )

1.3 - Criação do Banco de dados RDS
- Criação das tabelas 


No primeiro momento já tinha intuito de utilizar AWS como cloud provider e API Internet Gateway para fazer o roteamento de tráfego com base no header, após algumas pesquisas identifiquei que seria necessário utilizar uma função lambda para ler o Header e efetuar o forwarding da requisição para o ambiente desejado. 


Por este motivo resolvi colocar a regra do header no ALB, para evitar custos do Lambda, mas teriamos outras opções que poderiam ser exploradas como nginx ou utilizar service mesh istio ou similar. 

Foi provisionado um targetGroup para representar cada ambiente, essa etapa me livrou de configurar ZonaDNS e Ingress de cada aplicação, em um ambiente real seria encaminhado para outro loadbalancer ou endpoint. 

Já é possivel efetuar requisições para o ambiente da API Itens que desejar como mostra na figura:

## 2 - LIMITAÇÃO DE TRÁFEGO 

Apesar do Api Internet Gateway não funcionar muito bem para validar as regras no header, utilizei API Internet Gateway para expor endpoint público e criar limitação de tráfego com base no desafio. 

Foi alcançado o objetivo de limitar as requisições quando receber 1000 requisições por minuto ou 16 requisições por segundo. 

## 3 - MONITORAMENTO

Para o monitoramento ativo do cluster foi provisionado o NewRelic como principal foco de observabilidade, é possivel verificar métricas de recursos de infraestrutura, traces, logs e auto instrumentação com OpenTelemetry. 

Criado um dashboard de exemplo do que podemos entender como necessário para monitorar nossos ambientes, além do monitoramento triviais de Recursos (Memória/CPU/REDE) podemos visualizar métricas integradas com CloudWatch e Prometheus. 

A utilização do NewRelic facilitará ao identificar padrões e alertas que serão criados automaticamente podendo ser notificados em canais de atendimento. 

Alertas personalizados também são possíveis de se criar, podemos pensar no aumento do Error Rate, tanto do lado do API Gateway ou no NewRelic para buscar um monitoramento ativo.

Maiores integrações com newRelic podem ser efetuadas. 

## PLUS KUBECOST

Foi implementado o kubecost para auxiliar no controle dos custos do ambiente, ele trará insigths interessantes para economia e eficiência do ambiente da API de Itens. 

### Considerações Finais

Este challenge foi realmente desafiador, trouxe conteúdo denso de vários aspectos envolvendo infraestrutura, desenvolvimento e operações, com certeza temos vários pontos a melhorar principalmente visando IAC. 


