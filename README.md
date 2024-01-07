# Webapplikation zur Verwaltung von Aufgaben

Für die Anwendung der erlernten Theorie in die Praxis soll im Fach Web-Technologie 1 eine Applikation zur Verwaltung von Aufgaben umgesetzt werden. In einer Webapplikation sollen verschiedene Aufgaben erfasst werden können. Für jede dieser Aufgaben wird ein Titel, eine Beschreibung, ein Fälligkeitsdatum, der Aufgabentyp sowie eine oder mehrere zugewiesene Personen erfasst.
Für die zugewiesenen Personen sowie die verschiedenen Aufgabentypen können in einem zusätzlichen Administrationsbereich die entsprechenden Einträge erfasst werden.

## Starten der Applikation

Docker muss auf dem System installiert sein. Das Projekt kann mit dem Befehl `docker compose up -d` gestartet werden. Beim ersten Start werden beide Images erstellt. Dieser Prozess kann einige Sekunden in Anspruch nehmen.

Nach dem Start ist die Applikation unter http://localhost:5173/ erreichbar.

### Anpassen der Ports

Der Frontend Service ist standardmässig unter dem Port `:5173` erreichbar. Dieser Port kann in der Compose konfiguration angepasst werden.

Der Backend Service ist standardmässig unter dem Port `:8080` erreichbar. Sollte dieser Port angepasst werden müssen, so muss...
- Der Port des Backend Services angepasst werden.
- Das Build-arg des Frontend services `BACKEND_URL` angepasst werden mit der gültigen URL.
- Der Service "frontend" neu gebuildet werden.

## Importieren der Daten

Beiliegend ist eine Datei `export.json`. Diese Datei kann im Administrationsbereich der Applikation importiert werden. Beim Importieren werden alle vorhandenen Daten überschreieben.

Das wars, Lars.
