# ligapro-cli

![alt text](./media/inicio.png) 

ligapro-cli es un programa de terminal para estar al tanto de la serie A de Ecuador, también llamada LigaPro.  
Los datos actualizados son obtenidos de:
  
- [espn.com](https://www.espn.com/soccer/league/_/name/ecu.1).  
- [marca.com](https://www.marca.com/futbol/ecuador/primera-etapa/calendario.html)
- [tycsports.com](https://www.tycsports.com/estadisticas/futbol-de-ecuador/tabla-de-goleadores.html)
  
* Para el Web Scraping se uso el framewok [colly](https://go-colly.org/) 🌐
* Para la interfaz se terminal se uso el framewok [lipgloss](https://github.com/charmbracelet/lipgloss) 💻

## Instalación 
Con go instalado, usar el comando:
```bash
go install github.com/Luiggy102/ligapro-cli@latest
```
### Uso
Mostar las opciones con `-h`
```bash
ligapro-cli -h
```
![alt text](./media/ayuda.png) 

Mostar la tabla actual de posiciones con `-tabla`. También se puede buscar por años. Ej: `ligapro-cli -tabla -año 2023 -etapa 2`
```bash
ligapro-cli -tabla
```
![alt text](./media/table.png) 

Mostar los últimos resultados de la presente fecha con `-resultados`
```bash
ligapro-cli -resultados
```
![alt text](./media/resultados.png) 

Mostrar la tabla actual de goleadores con `-goleadores`
```bash
ligapro-cli -goleadores
```
![alt text](./media/goleadores.png) 

Buscar los partidos según su fecha con `-fecha <número-del-partido>`
```bash
ligapro-cli -fecha 15
```
![alt text](./media/fecha.png) 
