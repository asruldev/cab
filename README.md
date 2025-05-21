```
cmd/
└── server/
    ├── main.go           // ⛩️ Entry point, start server
    ├── router.go         // 🔀 Routing antar modul
internal/
└── auth/                 // 🔐 Modul auth (mandiri)
    ├── domain/           // 🧠 Model: struct + interface
    │   └── model.go
    ├── usecase/          // 💡 ViewModel: logika bisnis
    │   └── usecase.go
    ├── delivery/
    │   └── http/         // 🖥 View/Controller: HTTP Handler
    │       └── handler.go
    ├── repository/       // 🗄 Data: memory/DB/API
    │   └── memory_repo.go
    └── dto/              // 📦 Optional: Request/Response object
        └── dto.go
pkg/
├── config/               // ⚙️ App config & env
├── middleware/           // 🛡 JWT, logging, etc.
├── logger/               // 📋 Logrus/zap setup
└── utils/                // 🛠 Helper umum (hash, convert)
```