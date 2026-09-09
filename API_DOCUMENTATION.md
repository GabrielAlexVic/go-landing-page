# 🚀 Documentação da API - Analytics & Landing Page Metrics Service

Esta documentação descreve todos os endpoints RESTful fornecidos pelo backend em Go (`go-landing-page`), incluindo autenticação, rotas públicas de telemetria para visitantes e rotas protegidas por JWT para o painel administrativo.

---

## 📌 Informações Gerais

* **URL Base (Desenvolvimento):** `http://localhost:8080/api`
* **Formato de Dados:** `JSON` (Headers `Content-Type: application/json` e `Accept: application/json`)
* **Autenticação:** As rotas administrativas requerem o envio do token JWT no cabeçalho HTTP:
  ```http
  Authorization: Bearer <SEU_TOKEN_JWT>
  ```

---

## 🔐 1. Autenticação e Utilitários

### `POST /auth/login`
Autentica o usuário administrador e retorna um token JWT válido por 24 horas.

* **Acesso:** Público

#### Requisição (`Body`):
```json
{
  "email": "admin@gvbyte.com",
  "password": "senha_segura"
}
```

#### Resposta (`200 OK`):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### Respostas de Erro:
* `500 Internal Server Error`: `{"message": "Email ou senha incorretos"}`

---

### `POST /hash-password`
Utilitário para gerar hash bcrypt de senhas para inserção manual de administradores no banco de dados.

* **Acesso:** Público

#### Requisição (`Body`):
```json
{
  "password": "senha_segura"
}
```

#### Resposta (`200 OK`):
```json
{
  "password_hash": "$2a$10$7...hash_gerado..."
}
```

---

## 🌐 2. Rotas Públicas de Telemetria (Landing Page)

Estas rotas são chamadas pelo frontend do site para capturar o comportamento dos visitantes sem necessidade de autenticação.

### `POST /page-views`
Registra a visualização de uma página, incluindo dados de sessão, origem (UTMs) e dispositivo.

* **Acesso:** Público

#### Requisição (`Body`):
```json
{
  "session_id": "sess_xyz12345",
  "visitor_hash": "v_abc67890",
  "page_path": "/",
  "referrer": "https://instagram.com",
  "utm_source": "instagram",
  "utm_medium": "bio",
  "utm_campaign": "lancamento_setembro",
  "utm_term": "software",
  "utm_content": "link_bio",
  "device_type": "mobile",
  "browser": "Chrome",
  "os": "Android",
  "duration_seconds": 15
}
```

#### Resposta (`200 OK`):
```json
{
  "message": "page view created successfully"
}
```

---

### `POST /events`
Registra interações específicas do usuário na página (cliques em botões, profundidade de scroll, etc.).

* **Acesso:** Público

#### Requisição (`Body`):
```json
{
  "session_id": "sess_xyz12345",
  "event_type": "click_cta",
  "event_label": "btn_hero_agendar",
  "metadata": {
    "section": "hero",
    "button_id": "primary_cta"
  }
}
```

#### Resposta (`200 OK`):
```json
{
  "message": "event created successfully"
}
```

---

### `POST /leads`
Captura o preenchimento do formulário de contato/conversão da landing page.

* **Acesso:** Público

#### Requisição (`Body`):
```json
{
  "session_id": "sess_xyz12345",
  "name": "João Silva",
  "email": "joao@empresa.com",
  "company": "Empresa XPTO",
  "phone": "(11) 99999-9999",
  "description": "Gostaria de solicitar um orçamento para desenvolvimento de software sob medida.",
  "utm_source": "google",
  "utm_campaign": "search_ads"
}
```

#### Resposta (`200 OK`):
```json
{
  "message": "lead created successfully"
}
```

---

## 🔒 3. Rotas Protegidas por JWT (Painel Admin)

Todas as rotas abaixo **exigem** o header `Authorization: Bearer <TOKEN_JWT>`.

---

### `GET /metrics/summary`
Retorna o resumo estatístico consolidado do dashboard executivo.

* **Acesso:** Protegido (Admin)
* **Header Obrigatório:** `Authorization: Bearer <TOKEN_JWT>`

#### Resposta (`200 OK`):
```json
{
  "total_page_views": 1540,
  "unique_visitors": 820,
  "total_leads": 41,
  "total_events": 2300,
  "conversion_rate": 5.0,
  "top_utm_sources": [
    {
      "name": "instagram",
      "count": 450
    },
    {
      "name": "google",
      "count": 210
    },
    {
      "name": "Direto / Nenhum",
      "count": 160
    }
  ],
  "top_devices": [
    {
      "name": "mobile",
      "count": 600
    },
    {
      "name": "desktop",
      "count": 220
    }
  ]
}
```

---

### `GET /page-views`
Lista os acessos registrados com suporte a paginação.

* **Acesso:** Protegido (Admin)
* **Header Obrigatório:** `Authorization: Bearer <TOKEN_JWT>`

#### Parâmetros (`Query` / `Body`):
| Parâmetro | Tipo | Descrição | Padrão |
| :--- | :--- | :--- | :--- |
| `cursor` | `integer` | ID limite para paginação (`WHERE id < cursor`) | `0` |
| `limit` | `integer` | Quantidade máxima de registros a retornar | `50` |

#### Resposta (`200 OK`):
```json
[
  {
    "id": 105,
    "session_id": "sess_xyz12345",
    "visitor_hash": "v_abc67890",
    "page_path": "/",
    "referrer": "https://instagram.com",
    "utm_source": "instagram",
    "utm_medium": "bio",
    "utm_campaign": "lancamento_setembro",
    "utm_term": "",
    "utm_content": "",
    "device_type": "mobile",
    "browser": "Chrome",
    "os": "Android",
    "duration_seconds": 15,
    "created_at": "2026-09-09T17:30:00Z"
  }
]
```

---

### `GET /events`
Lista os eventos de interações registrados.

* **Acesso:** Protegido (Admin)
* **Header Obrigatório:** `Authorization: Bearer <TOKEN_JWT>`

#### Parâmetros (`Query` / `Body`):
| Parâmetro | Tipo | Descrição | Padrão |
| :--- | :--- | :--- | :--- |
| `cursor` | `integer` | ID limite para paginação | `0` |
| `limit` | `integer` | Quantidade máxima de registros | `50` |

#### Resposta (`200 OK`):
```json
[
  {
    "id": 89,
    "session_id": "sess_xyz12345",
    "event_type": "click_cta",
    "event_label": "btn_hero_agendar",
    "metadata": {
      "section": "hero"
    },
    "created_at": "2026-09-09T17:30:05Z"
  }
]
```

---

### `GET /leads`
Lista todos os leads capturados.

* **Acesso:** Protegido (Admin)
* **Header Obrigatório:** `Authorization: Bearer <TOKEN_JWT>`

#### Parâmetros (`Query` / `Body`):
| Parâmetro | Tipo | Descrição | Padrão |
| :--- | :--- | :--- | :--- |
| `cursor` | `integer` | ID limite para paginação | `0` |
| `limit` | `integer` | Quantidade máxima de registros | `50` |

#### Resposta (`200 OK`):
```json
[
  {
    "id": 12,
    "session_id": "sess_xyz12345",
    "name": "João Silva",
    "email": "joao@empresa.com",
    "company": "Empresa XPTO",
    "phone": "(11) 99999-9999",
    "description": "Gostaria de solicitar um orçamento para desenvolvimento de software sob medida.",
    "utm_source": "google",
    "utm_campaign": "search_ads",
    "created_at": "2026-09-09T17:30:10Z"
  }
]
```

---

## 🛠️ Status HTTP Comuns

| Código | Descrição |
| :--- | :--- |
| `200 OK` | Requisição processada com sucesso. |
| `400 Bad Request` | Parâmetros de entrada inválidos ou malformados. |
| `401 Unauthorized` | Token JWT ausente, inválido ou expirado. |
| `500 Internal Server Error` | Erro interno do servidor ou falha no banco de dados. |
