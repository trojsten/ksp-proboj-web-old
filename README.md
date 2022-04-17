# Proboj Web

Koordinačný web určený pre [Proboj Runner](https://github.com/trojsten/ksp-proboj-runner).

## Ako rozbehať?

Potrebuješ Go, Node.JS a pnpm.

```shell
pnpm install
pnpm run build
go run . 
```

Treba si prekopírovať `config.example.json` do `config.json` a upraviť si ho podľa
svojich predstáv:

- `runner_command` - cesta k súboru runnera (`main.py`)
- `runner_debug` - ak je true, stdout a stderr runnera sa zobrazuje na stdout/stderr webu
- `data_folder` - cesta k priečinku, v ktorom sa budú ukladať hry
- `upload_folder` - cesta k priečinku, v ktorom sa budú ukladať nahrané súbory
- `database` - cesta k databáze
- `server_command` - cesta k súboru Proboj servera
- `player_timeout` - koľko sekúnd majú hráči na odpovedanie serveru
- `games_ahead` - koľko hier popredu (voči observeru) má byť pripravených
- `make_command` - cesta k binárke programu `make`
- `run_games` - ak je true, web bude spúšťať hry
- `presenter_ip` - IP adresa, na ktorej sa môže prehrávať automatický observer

## Príprava na Proboj

Asi si chceš vymeniť obrázky `static/bg.jpg` a `static/logo.png`.
Vyrob si priečinok `observer`, v ktorom sú všetky súbory observera. Môže to byť aj symlink.
Observer chce byť `index.html`. Observer dostáva optional GET parametre `file` a `autoplay`:

- `file` - URL adresa k observer logu
- `autoplay` - hodnota 1, ak je observer súčasť autoplayu. Observer by mal vtedy automaticky
  spustiť playback a po skončení presmerovať na `/autoplay/`.

Hráči budú nahrávať `.tar.gz` archív so zdrojákmi a `Makefile`. Web spúšťa target `player`
v Makefile. Je výhodné, aby tento Makefile obsahoval aj target, ktorý vyrobí `.tar.gz`.

Chceš ešte napísať dokumentáciu tvojej hry do `templates/docs.gohtml`.

Hráči na inom operačnom systéme budú potrebovať niečo, čo vie vyrábať `.tar.gz` - napríklad
[7-Zip](https://www.7-zip.org/) to dokáže.

## Výroba klientov

Na vytvorenie používateľov treba pridať riadky do tabuľky `players`. Heslá nie sú hashované
(na naše účely to postačuje). Do SQLite databázy sa dá dostať cli utilitou `sqlite3 data.db`:

```sql
INSERT INTO players VALUES (null, "nazov_skupiny", "heslo", 0); 
```

Názov skupiny by mal byť alfanumerický. Družinka sa vie prihlásiť do rozhrania.

## Mapy

Proboj Web podporuje tzv. mapy. Mapa sa používa pri generovaní hier, reprezentuje `args`
parameter pre Proboj Runner.

## Autoplay

Na počítači, ktorý prehráva hry stačí otvoriť stránku `/autoplay/`, automaticky sa bude
spúšťať observer. (Observer musí byť na toto pripravený, viď *Príprava na Proboj*)
