package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var ingressosDisponiveis = 100
var evento = "Manowar True Metal Concert"

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/reserva", reservaHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := `
	<!DOCTYPE html>
	<html>
	<head><title>{{.Evento}}</title></head>
	<body style="background-color:black; color:yellow; text-align:center;">
		<h1>{{.Evento}}</h1>
		<form action="/reserva" method="POST">
			<p>Are you True Metal? <input type="checkbox" name="truemetal"></p>
			<p>Are you a student? <input type="checkbox" name="student"></p>
			<p>How many tickets? <input type="number" name="ingressos" min="1" max="{{.Disponiveis}}"></p>
			<p>Payment method:
				<select name="pagamento">
					<option value="cash">Cash</option>
					<option value="credit">Credit Card</option>
				</select>
			</p>
			<input type="submit" value="Book Now">
		</form>
	</body>
	</html>
	`
	t := template.Must(template.New("home").Parse(tmpl))
	t.Execute(w, struct {
		Evento      string
		Disponiveis int
	}{
		Evento:      evento,
		Disponiveis: ingressosDisponiveis,
	})
}

func reservaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	trueMetal := r.FormValue("truemetal") == "on"
	student := r.FormValue("student") == "on"
	ingressosStr := r.FormValue("ingressos")
	pagamento := r.FormValue("pagamento")

	ingressos, err := strconv.Atoi(ingressosStr)
	if err != nil || ingressos < 1 || ingressos > ingressosDisponiveis {
		http.Error(w, "Número de ingressos inválido", http.StatusBadRequest)
		return
	}

	if !trueMetal {
		fmt.Fprintln(w, `<h1 style="color:red;">You are not True Metal! Get out!</h1>`)
		return
	}

	ingressosDisponiveis -= ingressos
	total := calcularPreco(ingressos, student)

	fmt.Fprintf(w, `
		<h2>Booking Confirmed!</h2>
		<p>Tickets: %d</p>
		<p>Student: %v</p>
		<p>Payment Method: %s</p>
		<p>Total: $%.2f</p>
		<p>Hail and Kill! 🤘</p>
	`, ingressos, student, pagamento, total)
}

func calcularPreco(qtd int, estudante bool) float64 {
	preco := 1000.0
	if estudante {
		preco = 500.0
	}
	return preco * float64(qtd)
}
