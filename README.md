# MealApp - Angular & Golang Fullstack Project

Projekt stworzony w celu odświeżenia i praktycznego zastosowania frameworka **Angular** oraz integracji go z wydajnym backendem napisanym w języku **Go**.

## 🚀 O projekcie

Aplikacja MealApp to kompletny system typu Fullstack. Frontend zarządza interfejsem użytkownika i komunikacją z API, natomiast backend odpowiada za logikę biznesową oraz trwałość danych.

**Kluczowe cechy:**
* **Architektura:** Rozdzielenie warstwy prezentacji (Angular) od logiki serwerowej (Go).
* **Baza danych:** Wykorzystanie SQLite (przepisane w Go), co eliminuje konieczność instalacji zewnętrznych silników baz danych.
* **Autoinicjalizacja:** Serwis przy starcie sam tworzy niezbędną strukturę bazy danych.

---

## 🛠️ Instrukcja uruchomienia

Postępuj zgodnie z poniższymi krokami, aby uruchomić projekt lokalnie.

### 1. Backend (Golang)
Serwis korzysta z bazy danych `LocalTesBase.db`, która jest automatycznie inicjowana przy starcie.

1.  Wejdź do katalogu backendu:
    ```bash
    cd backend
    ```
2.  Uruchom serwer (pobierze on automatycznie wymagane zależności):
    ```bash
    go run main.go
    ```
    *Opcjonalnie możesz skompilować projekt:* `go build main.go` i uruchomić plik wykonywalny.

### 2. Frontend (Angular)
Upewnij się, że masz zainstalowane środowisko Node.js oraz Angular CLI.

1.  Wejdź do katalogu frontendu:
    ```bash
    cd frontend
    ```
2.  Zainstaluj paczki NPM:
    ```bash
    npm install
    ```
3.  Uruchom aplikację:
    ```bash
    ng serve
    ```

---

## 🌐 Dostęp do aplikacji

Gdy oba serwisy są uruchomione, aplikacja dostępna jest pod adresem:
👉 **[http://localhost:4200/](http://localhost:4200/)**

---

## 💻 Technologie użyte w projekcie

* **Frontend:** Angular (TypeScript)
* **Backend:** Golang
* **Baza danych:** PostgreSql

---
*Projekt przygotowany przez [Adam111p](https://github.com/Adam111p)*
