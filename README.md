# RateMyMusic
RateMyMusic je aplikacija za diskusiju, ocenjivanje i pisanje recenzija o muzicarima, pesmama i albumima.

# Funkcionalnosti

Glavna funkcionalnost ove aplikacije je pregled i ocenjivanje muzicara, pesama i albuma.

### Neregistrovani korisnik
- Pregleda i pretrage muzicara, pesama i albuma.
- Registracija i logovanje u sistem

### Registrovani korisnik
- Funkcionalnosti neregistrovanog korisnika
- Ocenjivanje muzicara, albuma i pesama.
- Pisanje recenzija za muzicara, albuma i pesama.
- Kreiranja kolekcija pesama, albuma i muzicara


### Moderatori
- CRUD operacije nad muzicarima, albumima i pesmama.
- CRUD operacije nad korisnickim ocenama i recenzijama


### Administrator
- CRUD operacije nad registrovanim korisnicima
- 
# Arhitektura sistema
* Korisnički mirkoservis - Go
* Mikroservis za rukovanje repozitorijumima - Go
* Email mikroservis - Go ili Python
* Veb interfejs - TypeScript, Angular

Moguca prosirenja u vidu rule-based sistema ili machine learning systema za predlaganje novih muzicara, albuma i pesmama registrovanim korisnicima na osnovu njihovih prethodnih aktivnosti
