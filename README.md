# ligapro-cli

ligapro-cli es un programa de terminal para estar al tanto de la serie A de Ecuador, también llamada LigaPro.
Los datos actualizados son obtenidos de [espn.com](https://www.espn.com/soccer/league/_/name/ecu.1).  

* Para el Web Scraping se uso el framewok [colly](https://go-colly.org/) 🌐
* Para la interfaz se terminal se uso el framewok [lipgloss](https://github.com/charmbracelet/lipgloss) 💻

## Instalación 
Con go instalado, usar el comando:
```bash
go install github.com/Luiggy102/ligapro-cli@latest
```

## Uso
Mostar las opciones con `-h`
```bash
ligapro-cli -h
```
Mostar la tabla actual de posiciones con `-tabla`
```bash
ligapro-cli -tabla
```
![alt text](/media/table.png)
