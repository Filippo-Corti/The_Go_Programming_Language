package main

import (
	"fmt"
	"io"
)

type TimeStamp struct {
	anno, mese, giorno, ora, minuto int
}

func (ts *TimeStamp) String() string {
	return fmt.Sprintf("misura delle ore %d, minuti %d del giorno %d del mese %d dell'anno %d", ts.ora, ts.minuto, ts.giorno, ts.mese, ts.anno)
}

type Misurazione struct {
	timestamp TimeStamp
	temperatura float64
	umidita float64
}
//Es: 1.0° misura delle ore 4, minuti 22 del giorno 11 del mese 12 dell'anno 2022
func (m *Misurazione) StringTemperatura() string {
	return fmt.Sprintf("%.1f° %s", m.temperatura, m.timestamp.String())
}

func (m *Misurazione) StringUmidita() string {
	return fmt.Sprintf("%.1f%% %s", m.umidita, m.timestamp.String())
}

func main() {
	minTemp := Misurazione{TimeStamp{}, 1000, 0}
	maxTemp := Misurazione{TimeStamp{}, -1000, 0}
	minHumid := Misurazione{TimeStamp{}, 0, 1000}
	maxHumid := Misurazione{TimeStamp{}, 0, -1000}
	var anno, mese, giorno, ora, minuto int
	var temperatura, umidita float64
	for {
		_, err := fmt.Scanf("%d-%d-%d/%d:%d,%f,%f", &anno, &mese, &giorno, &ora, &minuto, &temperatura, &umidita)
		if err == io.EOF {
			break
		}
		if temperatura > maxTemp.temperatura {
			maxTemp = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		} else if temperatura < minTemp.temperatura {
			minTemp = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		}
		if umidita > maxHumid.umidita {
			maxHumid = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		} else if umidita < minHumid.umidita {
			minHumid = Misurazione{timestamp: TimeStamp{anno: anno, mese: mese, giorno: giorno, ora: ora, minuto: minuto}, temperatura: temperatura, umidita: umidita}
		}
	}
	fmt.Println("minTemp:", minTemp.StringTemperatura())
	fmt.Println("maxTemp:", maxTemp.StringTemperatura())
	fmt.Println("minHumid:", minHumid.StringUmidita())
	fmt.Println("maxHumid:", maxHumid.StringUmidita())

}