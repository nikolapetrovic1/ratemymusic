# RateMyMusic
RateMyMusic je aplikacija za diskusiju, ocenjivanje i pisanje recenzija o muzicarima, pesmama i albumima.

# Funkcionalnosti
Glavna funkcionalnost ove aplikacije je pregled i ocenjivanje muzicara, pesama i albuma.
Mogucnost pregleda novih i popularnih muzicara, albuma i pesama. Korisnici takodje mogu da diskutuju tudje recenzije preko komentara.

### Neregistrovani korisnik
- Pregleda i pretrage muzicara, pesama i albuma.
- Registracija i logovanje na sistem

### Registrovani korisnik
- Pregleda i pretrage muzicara, pesama i albuma.
- Ocenjivanje muzicara, albuma i pesama.
- Pisanje recenzija za muzicara, albuma i pesama.
- Pisanje komentara na druge recenezije
- Kreiranja kolekcija pesama, albuma i muzicara
- Pregled profila i profila drugih registrovanih korisnika
- Mogućnost prijave neprikladnog komentara.


### Moderatori
- CRUD operacije nad muzicarima, albumima i pesmama.
- CRUD operacije nad korisnickim ocenama i recenzijama

### Administrator
- CRUD operacije nad registrovanim korisnicima
- Izvestaji o korisnicima
- Mogućnost banovanja korisnika čiji su komentari prijavljeni kao neprikladni.

# Arhitektura sistema
* API gateway - Tehnologija: Go
* Korisnički mirkoservis - Za rukuvanje korisnicima - Go, PostgreSQL 
* Mikroservis za rukovanje muzicarima repozitorijumima - Go, PostgreSQL 
* Mikroservis za rukovanje pesmama repozitorijumima - Go, PostgreSQL 
* Mikroservis za rukovanje albumima repozitorijumima - Go,PostgreSQL 
* Mikroservis za rukovanje recenzijama i ocenama - Rust,PostgreSQL 
* Mikroservis za rukovanje komentarima - Go,PostgreSQL 
* Mikroservis za komunikaciju sa streaming plaformama (Soundcloud ili Spotify) za slusanje muzike - Go ili Python
* Veb interfejs - TypeScript, Angular

## Prosirenje za diplomski rad:
- Slusanje pesma i albuma preko postojecih streaming platformi.


