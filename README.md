# Smart Shoper App

## Opis Projekta
Aplikacija Smart Shopper je zasnovana za pomoč uporabnikom pri prihranku časa in denarja, saj omogoča iskanje najcenejših cen za izdelke z njihovega nakupovalnega seznama v različnih trgovinah. Uporabniki lahko postopoma dodajajo izdelke, si ogledajo vse razpoložljive možnosti za vsak izdelek, filtrirajo rezultate po določenih trgovinah in shranjujejo nakupovalne sezname za kasnejšo uporabo. Aplikacija temelji na microservice arhitekturi, ki zagotavlja skalabilnost, modularnost in učinkovito delovanje.

---

## Pregled Mikrostoritev

### 1. **User Service**
**Opis**:  
Ta storitev skrbi za avtentikacijo uporabnikov, upravljanje nakupovalnih seznamov in uporabniške nastavitve.

**Funkcionalnosti**:
- Registracija novih uporabnikov.
- Avtentikacija uporabnikov.
- Upravljanje uporabniških profilov.
- Ustvarjanje novih nakupovalnih seznamov.
- Pridobivanje obstoječih seznamov.
- Posodabljanje ali brisanje seznamov.

**Database**: PostgreSQL  
**Communication Protocol**: REST API  

---

### 2. **Price Aggregation Service**
**Opis**:  
Ta storitev je odgovorna za iskanje najnižjih cen za izdelke, izračun celotnih stroškov in zagotavljanje možnosti optimizacije.

**Funkcionalnosti**:
- Iskanje najnižje cene za vsak izdelek (možna uporaba več trgovin).
- Normalizacija imen izdelkov in cen med različnimi trgovinami.
- Izračun najcenejše trgovine za vse izdelke.
- Omogočanje filtriranja rezultatov po določenih trgovinah.
- Prikaz vseh razpoložljivih trgovin/cen za določen izdelek.

**Database**: MongoDB
**Communication Protocol**: gRPC  

---

### 3. **Product Service**
**Opis**:  
Ta storitev skrbi za pridobivanje cen izdelkov iz različnih trgovin. Uporablja sporočilni posrednik za distribucijo nalog pridobivanja podatkov.

**Funkcionalnosti**:
- Sprejemanje posodobitev cen iz trgovin.
- Objavljanje novih cen.
- Obveščanje servisa za optimizacijo cen o spremembah.

**Database**: None (Stateless)  
**Communication Protocol**: Message Broker (AMQP)  

---