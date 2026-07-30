# Desafio Técnico DevOps - Korp

Repositório com a solução do desafio técnico DevOps/Infraestrutura.

## Visão Geral

A estrutura sobe os seguintes serviços via Docker Compose e Ansible:

* API Go (`http-server-projeto-korp`): Roda na porta 8080, entregando o JSON em UTC no endpoint /projeto-korp e as métricas em /metrics.
* Nginx: Proxy reverso na porta 80 que encaminha o tráfego para a API.
* Prometheus: Coleta as métricas de requisições e disponibilidade (`up`).
* Grafana: Exibe o dashboard de monitoramento na porta 3000.
* Ansible: Playbook para subir a stack e validar o status do ambiente.


## Como Executar

### Pré-requisitos
* Linux ou WSL2
* Docker e Docker Compose
* Ansible

### Subindo o ambiente

Para provisionar a stack e rodar os testes de verificação automatizados:

```bash
sudo ansible-playbook playbook.yml

```

## Endpoints

* Aplicação (via Nginx): http://localhost/projeto-korp
* Métricas da API: http://localhost:8080/metrics
* Prometheus: http://localhost:9090
* Grafana: http://localhost:3000