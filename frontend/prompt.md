# Angular 19+ Best Practices for this project:

- Use **Signals** for state management (`signal`, `computed`, `effect`).
- Use **`inject()`** function for Dependency Injection instead of constructors.
- Use **Standalone Components** (no NgModules).
- Use **New Control Flow** (`@if`, `@for`, `@switch`).
- Use **`input()`, `output()`, `model()`** (Signal-based inputs).
- Prefer **Zoneless change detection** if possible.
- Backend is written in **Go (Gin framework)** - maintain type safety between Go structs and TS interfaces.
