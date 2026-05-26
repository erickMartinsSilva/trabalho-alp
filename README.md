# trabalho-alp

Trabalho para a disciplina de **Arquitetura de Linguagens de Programação** — 26.1 · CEFET/RJ

Sistema distribuído de gerenciamento de usuários composto por três camadas: uma **API REST** replicada em três instâncias, um **Load Balancer** que distribui as requisições entre elas, e uma **CLI interativa** para o usuário final.

---

## Arquitetura

```
                                            ┌──────────┐
                                        ┌──▶│  API 1   │──┐
                                        │   │  :8080   │  │
     ┌─────────┐      ┌──────────────┐  │   └──────────┘  │  ┌────────────┐
     │   CLI   │─────▶│ Load Balancer│  │   ┌──────────┐  │  │   SQLite   │
     │ (Go)    │      │   :5080      │──┼──▶│  API 2   │──├─▶│  (volume   │
     └─────────┘      └──────────────┘  │   │  :8081   │  │  │compartil.) │
                                        │   └──────────┘  │  └────────────┘
                                        │   ┌──────────┐  │
                                        └──▶│  API 3   │──┘
                                            │  :8082   │
                                            └──────────┘
```

- **Load Balancer** — Proxy em Go que distribui as requisições entre as três instâncias da API via Round-Robin.
- **API** — servidor HTTP em Go que expõe um CRUD de usuários, persistindo os dados em SQLite.
- **CLI** — interface interativa em modo texto escrito em Go que se comunica com o Load Balancer.

---

## Pré-requisitos

| Ferramenta | Versão mínima | Uso |
|---|---|---|
| [Go](https://go.dev/dl/) | 1.18 | Executar a interface |
| [Docker](https://docs.docker.com/get-docker/) | 20+ | Containerizar API e Load Balancer |
| [Docker Compose](https://docs.docker.com/compose/) | v2 | Orquestrar os serviços de back-end |

---

## Estrutura do projeto

```
trabalho-alp/
├── api/                  # Código-fonte da API REST
│   ├── src/
│   │   ├── controllers/  # Handlers HTTP
│   │   ├── db/           # Inicialização e acesso ao SQLite
│   │   ├── models/       # Structs de domínio
│   │   ├── routes/       # Definição de rotas
│   │   ├── services/     # Regras de negócio
│   │   └── utils/        # Utilitários
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── go.mod
├── loadbalancer/         # Load Balancer HTTP round-robin
│   ├── src/
│   ├── Dockerfile
│   ├── docker-compose.yml
│   └── go.mod
├── cli/                  # Interface de linha de comando
│   └── src/
│       ├── main.go
│       ├── commands.go   # Lógica dos comandos do menu
│       └── client.go     # Cliente HTTP para o Load Balancer
├── run.sh                # Script de inicialização completa
└── README.md
```

---

## Como executar

### Opção 1 — Script automático (recomendado)

Na raiz do projeto, execute:

```bash
chmod +x run.sh
./run.sh
```

O script irá:
1. Subir o Load Balancer e as três instâncias da API via Docker Compose.
2. Iniciar a CLI interativa em seguida.

---

### Opção 2 — Manual passo a passo

**1. Subir o back-end (API + Load Balancer)**

```bash
cd loadbalancer
docker compose up -d --build
```

Isso inicia:
- `alp-api-1` na porta `8080`
- `alp-api-2` na porta `8081`
- `alp-api-3` na porta `8082`
- `alp-loadbalancer` na porta `5080`

**2. Iniciar a CLI**

```bash
cd cli/src
go run .
```

A CLI se conecta por padrão em `http://localhost:5080`. Para apontar para outro endereço, defina a variável de ambiente `LB_HOST`:

```bash
LB_HOST=http://outro-host:5080 go run .
```

---

## Dependências

### API (`api/go.mod`)
| Pacote | Versão | Finalidade |
|---|---|---|
| `github.com/mattn/go-sqlite3` | v1.14.44 | Driver SQLite via CGo |

### Load Balancer e CLI
Utilizam apenas a biblioteca padrão do Go, não possuindo dependências externas.

---

## Parando os serviços

```bash
cd loadbalancer
docker compose down
```

Para remover também o volume de dados do SQLite:

```bash
docker compose down -v
```
